package email

import (
	"context"
	"fmt"

	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/third-manager/pkg/db/ent"
	"github.com/NpoolPlatform/third-manager/pkg/testinit"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/NpoolPlatform/message/npool"
	"github.com/NpoolPlatform/message/npool/third/mgr/v1/template/email"
	"github.com/NpoolPlatform/message/npool/third/mgr/v1/usedfor"

	"github.com/google/uuid"

	"github.com/stretchr/testify/assert"
)

func init() {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	if err := testinit.Init(); err != nil {
		fmt.Printf("cannot init test stub: %v\n", err)
	}
}

var entAppEmailTemplate = ent.AppEmailTemplate{
	ID:                uuid.New(),
	AppID:             uuid.New(),
	LangID:            uuid.New(),
	DefaultToUsername: uuid.NewString(),
	UsedFor:           usedfor.UsedFor_Signin.String(),
	Sender:            uuid.NewString(),
	ReplyTos:          []string{uuid.NewString()},
	CcTos:             []string{uuid.NewString()},
	Subject:           uuid.NewString(),
	Body:              uuid.NewString(),
}

var (
	id               = entAppEmailTemplate.ID.String()
	appID            = entAppEmailTemplate.AppID.String()
	langID           = entAppEmailTemplate.LangID.String()
	usedFor          = usedfor.UsedFor_Signin
	appEmailTemplate = email.EmailTemplateReq{
		ID:                &id,
		AppID:             &appID,
		LangID:            &langID,
		UsedFor:           &usedFor,
		Sender:            &entAppEmailTemplate.Sender,
		ReplyTos:          entAppEmailTemplate.ReplyTos,
		CCTos:             entAppEmailTemplate.CcTos,
		Subject:           &entAppEmailTemplate.Subject,
		Body:              &entAppEmailTemplate.Body,
		DefaultToUsername: &entAppEmailTemplate.DefaultToUsername,
	}
)

var info *ent.AppEmailTemplate

func rowToObject(row *ent.AppEmailTemplate) *ent.AppEmailTemplate {
	return &ent.AppEmailTemplate{
		ID:                row.ID,
		CreatedAt:         row.CreatedAt,
		AppID:             row.AppID,
		LangID:            row.LangID,
		DefaultToUsername: row.DefaultToUsername,
		UsedFor:           row.UsedFor,
		Sender:            row.Sender,
		ReplyTos:          row.ReplyTos,
		CcTos:             row.CcTos,
		Subject:           row.Subject,
		Body:              row.Body,
	}
}

func create(t *testing.T) {
	var err error
	info, err = Create(context.Background(), &appEmailTemplate)
	if assert.Nil(t, err) {
		if assert.NotEqual(t, info.ID, uuid.UUID{}.String()) {
			entAppEmailTemplate.ID = info.ID
			entAppEmailTemplate.CreatedAt = info.CreatedAt
		}
		assert.Equal(t, rowToObject(info), &entAppEmailTemplate)
	}
}

func createBulk(t *testing.T) {
	entApp := []ent.AppEmailTemplate{
		{
			ID:                uuid.New(),
			AppID:             uuid.New(),
			LangID:            uuid.New(),
			DefaultToUsername: uuid.NewString(),
			UsedFor:           usedfor.UsedFor_Signin.String(),
			Sender:            uuid.NewString(),
			ReplyTos:          []string{uuid.NewString()},
			CcTos:             []string{uuid.NewString()},
			Subject:           uuid.NewString(),
			Body:              uuid.NewString(),
		},
		{
			ID:                uuid.New(),
			AppID:             uuid.New(),
			LangID:            uuid.New(),
			DefaultToUsername: uuid.NewString(),
			UsedFor:           usedfor.UsedFor_Signin.String(),
			Sender:            uuid.NewString(),
			ReplyTos:          []string{uuid.NewString()},
			CcTos:             []string{uuid.NewString()},
			Subject:           uuid.NewString(),
			Body:              uuid.NewString(),
		},
	}

	apps := []*email.EmailTemplateReq{}
	for key := range entApp {
		id := entApp[key].ID.String()
		appID = entApp[key].AppID.String()
		langID := entApp[key].LangID.String()
		usedFor = usedfor.UsedFor_Signin
		apps = append(apps, &email.EmailTemplateReq{
			ID:                &id,
			AppID:             &appID,
			LangID:            &langID,
			UsedFor:           &usedFor,
			Sender:            &entApp[key].Sender,
			ReplyTos:          entApp[key].ReplyTos,
			CCTos:             entApp[key].CcTos,
			Subject:           &entApp[key].Subject,
			Body:              &entApp[key].Body,
			DefaultToUsername: &entApp[key].DefaultToUsername,
		})
	}
	infos, err := CreateBulk(context.Background(), apps)
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 2)
		assert.NotEqual(t, infos[0].ID, uuid.UUID{}.String())
		assert.NotEqual(t, infos[1].ID, uuid.UUID{}.String())
	}
}

func update(t *testing.T) {
	var err error
	info, err = Update(context.Background(), &appEmailTemplate)
	if assert.Nil(t, err) {
		assert.Equal(t, rowToObject(info), &entAppEmailTemplate)
	}
}

func row(t *testing.T) {
	var err error
	info, err = Row(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, rowToObject(info), &entAppEmailTemplate)
	}
}

func rows(t *testing.T) {
	infos, total, err := Rows(context.Background(),
		&email.Conds{
			ID: &npool.StringVal{
				Value: info.ID.String(),
				Op:    cruder.EQ,
			},
		}, 0, 0)
	if assert.Nil(t, err) {
		assert.Equal(t, total, 1)
		assert.Equal(t, rowToObject(infos[0]), &entAppEmailTemplate)
	}
}

func rowOnly(t *testing.T) {
	var err error
	info, err = RowOnly(context.Background(),
		&email.Conds{
			ID: &npool.StringVal{
				Value: info.ID.String(),
				Op:    cruder.EQ,
			},
		})
	if assert.Nil(t, err) {
		assert.Equal(t, rowToObject(info), &entAppEmailTemplate)
	}
}

func count(t *testing.T) {
	count, err := Count(context.Background(),
		&email.Conds{
			ID: &npool.StringVal{
				Value: info.ID.String(),
				Op:    cruder.EQ,
			},
		},
	)
	if assert.Nil(t, err) {
		assert.Equal(t, count, count)
	}
}

func exist(t *testing.T) {
	exist, err := Exist(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, exist, true)
	}
}

func existConds(t *testing.T) {
	exist, err := ExistConds(context.Background(),
		&email.Conds{
			ID: &npool.StringVal{
				Value: info.ID.String(),
				Op:    cruder.EQ,
			},
		},
	)
	if assert.Nil(t, err) {
		assert.Equal(t, exist, true)
	}
}

func deleteT(t *testing.T) {
	info, err := Delete(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, rowToObject(info), &entAppEmailTemplate)
	}
}

func TestMainOrder(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	t.Run("create", create)
	t.Run("createBulk", createBulk)
	t.Run("row", row)
	t.Run("rows", rows)
	t.Run("rowOnly", rowOnly)
	t.Run("update", update)
	t.Run("exist", exist)
	t.Run("existConds", existConds)
	t.Run("delete", deleteT)
	t.Run("count", count)
}
