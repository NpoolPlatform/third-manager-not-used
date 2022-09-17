//nolint:nolintlint
package migrator

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/NpoolPlatform/message/npool/appuser/mgr/v2/signmethod"
	"github.com/NpoolPlatform/message/npool/third/mgr/v1/usedfor"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/go-service-framework/pkg/config"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	constant "github.com/NpoolPlatform/go-service-framework/pkg/mysql/const"
	"github.com/NpoolPlatform/third-manager/pkg/db"

	gwent "github.com/NpoolPlatform/third-gateway/pkg/db/ent"
	gwconst "github.com/NpoolPlatform/third-gateway/pkg/message/const"

	gwcconst "github.com/NpoolPlatform/third-gateway/pkg/const"

	crudcontact "github.com/NpoolPlatform/third-manager/pkg/crud/v1/contact"
	crudemail "github.com/NpoolPlatform/third-manager/pkg/crud/v1/template/email"
	crudsms "github.com/NpoolPlatform/third-manager/pkg/crud/v1/template/sms"

	contactpb "github.com/NpoolPlatform/message/npool/third/mgr/v1/contact"
	emailpb "github.com/NpoolPlatform/message/npool/third/mgr/v1/template/email"
	smspb "github.com/NpoolPlatform/message/npool/third/mgr/v1/template/sms"
)

func Migrate(ctx context.Context) error {
	return migrationThirdGateway(ctx)
}

const (
	keyUsername = "username"
	keyPassword = "password"
	keyDBName   = "database_name"
	maxOpen     = 10
	maxIdle     = 10
	MaxLife     = 3
)

func dsn(hostname string) (string, error) {
	username := config.GetStringValueWithNameSpace(constant.MysqlServiceName, keyUsername)
	password := config.GetStringValueWithNameSpace(constant.MysqlServiceName, keyPassword)
	dbname := config.GetStringValueWithNameSpace(hostname, keyDBName)

	svc, err := config.PeekService(constant.MysqlServiceName)
	if err != nil {
		logger.Sugar().Warnw("dsb", "error", err)
		return "", err
	}

	return fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true&interpolateParams=true",
		username, password,
		svc.Address,
		svc.Port,
		dbname,
	), nil
}

func open(hostname string) (conn *sql.DB, err error) {
	hdsn, err := dsn(hostname)
	if err != nil {
		return nil, err
	}

	conn, err = sql.Open("mysql", hdsn)
	if err != nil {
		return nil, err
	}

	// https://github.com/go-sql-driver/mysql
	// See "Important settings" section.

	conn.SetConnMaxLifetime(time.Minute * MaxLife)
	conn.SetMaxOpenConns(maxOpen)
	conn.SetMaxIdleConns(maxIdle)

	return conn, nil
}

//nolint
func migrationThirdGateway(ctx context.Context) (err error) {
	cli, err := db.Client()
	if err != nil {
		return err
	}

	appContacts, err := cli.Contact.Query().Limit(1).All(ctx)
	if err != nil {
		return err
	}

	appSMSTemplates, err := cli.SMSTemplate.Query().Limit(1).All(ctx)
	if err != nil {
		return err
	}

	appEmailTemplates, err := cli.EmailTemplate.Query().Limit(1).All(ctx)
	if err != nil {
		return err
	}

	thirdGW, err := open(gwconst.ServiceName)
	if err != nil {
		return err
	}

	defer thirdGW.Close()

	gwCli := gwent.NewClient(gwent.Driver(entsql.OpenDB(dialect.MySQL, thirdGW)))

	defer func() {
		logger.Sugar().Infow("Migrate appContacts", "Done", "...", "error", err)
	}()

	if len(appContacts) == 0 {
		logger.Sugar().Infow("Migrate appContacts", "Start", "...")
		appContactInfos, err := gwCli.
			AppContact.
			Query().
			All(ctx)
		if err != nil {
			return err
		}

		appContactInfosC := []*contactpb.ContactReq{}

		for key := range appContactInfos {
			id := appContactInfos[key].ID.String()
			appID := appContactInfos[key].AppID.String()
			usedFor := getUsedFor(appContactInfos[key].UsedFor)
			accountType := signmethod.SignMethodType_Email
			if appContactInfos[key].AccountType == "mobile" {
				accountType = signmethod.SignMethodType_Mobile
			}
			appContactInfosC = append(appContactInfosC, &contactpb.ContactReq{
				ID:          &id,
				AppID:       &appID,
				UsedFor:     &usedFor,
				Account:     &appContactInfos[key].Account,
				AccountType: &accountType,
				Sender:      &appContactInfos[key].Sender,
			})
		}

		_, err = crudcontact.CreateBulk(ctx, appContactInfosC)
		if err != nil {
			return err
		}
	}

	if len(appSMSTemplates) == 0 {
		logger.Sugar().Infow("Migrate appSMSTemplates", "Start", "...")

		appSMSTemplateInfos, err := gwCli.
			AppSMSTemplate.
			Query().
			All(ctx)
		if err != nil {
			return err
		}

		appSMSTemplateInfosC := []*smspb.SMSTemplateReq{}

		for key := range appSMSTemplateInfos {
			id := appSMSTemplateInfos[key].ID.String()
			appID := appSMSTemplateInfos[key].AppID.String()
			usedFor := getUsedFor(appSMSTemplateInfos[key].UsedFor)
			langID := appSMSTemplateInfos[key].LangID.String()
			appSMSTemplateInfosC = append(appSMSTemplateInfosC, &smspb.SMSTemplateReq{
				ID:      &id,
				AppID:   &appID,
				LangID:  &langID,
				UsedFor: &usedFor,
				Subject: &appSMSTemplateInfos[key].Subject,
				Message: &appSMSTemplateInfos[key].Message,
			})
		}

		_, err = crudsms.CreateBulk(ctx, appSMSTemplateInfosC)
		if err != nil {
			return err
		}
	}

	if len(appEmailTemplates) == 0 {
		logger.Sugar().Infow("Migrate appEmailTemplates", "Start", "...")
		appEmailTemplateInfos, err := gwCli.
			AppEmailTemplate.
			Query().
			All(ctx)
		if err != nil {
			return err
		}

		appEmailTemplateInfosC := []*emailpb.EmailTemplateReq{}

		for key := range appEmailTemplateInfos {
			id := appEmailTemplateInfos[key].ID.String()
			appID := appEmailTemplateInfos[key].AppID.String()
			usedFor := getUsedFor(appEmailTemplateInfos[key].UsedFor)
			langID := appEmailTemplateInfos[key].LangID.String()
			appEmailTemplateInfosC = append(appEmailTemplateInfosC, &emailpb.EmailTemplateReq{
				ID:                &id,
				AppID:             &appID,
				LangID:            &langID,
				UsedFor:           &usedFor,
				Sender:            &appEmailTemplateInfos[key].Sender,
				ReplyTos:          appEmailTemplateInfos[key].ReplyTos,
				CCTos:             appEmailTemplateInfos[key].CcTos,
				Subject:           &appEmailTemplateInfos[key].Subject,
				Body:              &appEmailTemplateInfos[key].Body,
				DefaultToUsername: &appEmailTemplateInfos[key].DefaultToUsername,
			})
		}

		_, err = crudemail.CreateBulk(ctx, appEmailTemplateInfosC)
		if err != nil {
			return err
		}
	}

	return nil
}

func getUsedFor(usedFor string) usedfor.UsedFor {
	switch usedFor {
	case gwcconst.UsedForSignup:
		return usedfor.UsedFor_Signup
	case gwcconst.UsedForSignin:
		return usedfor.UsedFor_Signin
	case gwcconst.UsedForUpdate:
		return usedfor.UsedFor_Update
	case gwcconst.UsedForContact:
		return usedfor.UsedFor_Contact
	case gwcconst.UsedForSetWithdrawAddress:
		return usedfor.UsedFor_SetWithdrawAddress
	case gwcconst.UsedForWithdraw:
		return usedfor.UsedFor_Withdraw
	case gwcconst.UsedForCreateInvitationCode:
		return usedfor.UsedFor_CreateInvitationCode
	case gwcconst.UsedForSetCommission:
		return usedfor.UsedFor_SetCommission
	case gwcconst.UsedForSetTransferTargetUser:
		return usedfor.UsedFor_SetTransferTargetUser
	case gwcconst.UsedForTransfer:
		return usedfor.UsedFor_Transfer
	default:
		return usedfor.UsedFor_DefaultUsedFor
	}
}
