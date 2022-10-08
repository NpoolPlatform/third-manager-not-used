//nolint:nolintlint
package migrator

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/NpoolPlatform/third-manager/pkg/db/ent"

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

	if len(appContacts) == 0 {
		logger.Sugar().Infow("Migrate appContacts", "Start", "...")

		defer func() {
			logger.Sugar().Infow("Migrate appContacts", "Done", "...", "error", err)
		}()

		appContactInfos, err := gwCli.
			AppContact.
			Query().
			All(ctx)
		if err != nil {
			return err
		}

		err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
			bulk := make([]*ent.ContactCreate, len(appContactInfos))
			for i, info := range appContactInfos {
				usedFor := getUsedFor(appContactInfos[i].UsedFor)
				accountType := signmethod.SignMethodType_Email
				if appContactInfos[i].AccountType == "mobile" {
					accountType = signmethod.SignMethodType_Mobile
				}

				bulk[i] = tx.Contact.
					Create().
					SetID(info.ID).
					SetAppID(info.AppID).
					SetUsedFor(usedFor.String()).
					SetSender(info.Sender).
					SetAccount(info.Account).
					SetAccountType(accountType.String()).
					SetCreatedAt(info.CreateAt).
					SetUpdatedAt(info.UpdateAt)
			}
			_, err = tx.Contact.CreateBulk(bulk...).Save(_ctx)
			return err
		})

		if err != nil {
			return err
		}
	}

	if len(appSMSTemplates) == 0 {
		logger.Sugar().Infow("Migrate appSMSTemplates", "Start", "...")

		defer func() {
			logger.Sugar().Infow("Migrate appSMSTemplates", "Done", "...", "error", err)
		}()

		appSMSTemplateInfos, err := gwCli.
			AppSMSTemplate.
			Query().
			All(ctx)
		if err != nil {
			return err
		}

		err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
			bulk := make([]*ent.SMSTemplateCreate, len(appSMSTemplateInfos))
			for i, info := range appSMSTemplateInfos {
				usedFor := getUsedFor(appSMSTemplateInfos[i].UsedFor)

				bulk[i] = tx.SMSTemplate.
					Create().
					SetID(info.ID).
					SetAppID(info.AppID).
					SetLangID(info.LangID).
					SetUsedFor(usedFor.String()).
					SetSubject(info.Subject).
					SetMessage(info.Message).
					SetCreatedAt(info.CreateAt).
					SetUpdatedAt(info.UpdateAt)
			}
			_, err = tx.SMSTemplate.CreateBulk(bulk...).Save(_ctx)
			return err
		})

		if err != nil {
			return err
		}
	}

	if len(appEmailTemplates) == 0 {
		logger.Sugar().Infow("Migrate appEmailTemplates", "Start", "...")

		defer func() {
			logger.Sugar().Infow("Migrate appSMSTemplates", "Done", "...", "error", err)
		}()

		appEmailTemplateInfos, err := gwCli.
			AppEmailTemplate.
			Query().
			All(ctx)
		if err != nil {
			return err
		}
		err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
			bulk := make([]*ent.EmailTemplateCreate, len(appEmailTemplateInfos))
			for i, info := range appEmailTemplateInfos {
				usedFor := getUsedFor(appEmailTemplateInfos[i].UsedFor)

				bulk[i] = tx.EmailTemplate.
					Create().
					SetID(info.ID).
					SetAppID(info.AppID).
					SetLangID(info.LangID).
					SetDefaultToUsername(info.DefaultToUsername).
					SetUsedFor(usedFor.String()).
					SetSender(info.Sender).
					SetReplyTos(info.ReplyTos).
					SetCcTos(info.CcTos).
					SetSubject(info.Subject).
					SetBody(info.Body).
					SetCreatedAt(info.CreateAt).
					SetUpdatedAt(info.UpdateAt)
			}
			_, err = tx.EmailTemplate.CreateBulk(bulk...).Save(_ctx)
			return err
		})

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
