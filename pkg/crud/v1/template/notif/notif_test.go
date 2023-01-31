package notif

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
	"github.com/NpoolPlatform/message/npool/third/mgr/v1/template/notif"

	usedfor "github.com/NpoolPlatform/message/npool/notif/mgr/v1/notif"

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

var entNotifTemplate = ent.NotifTemplate{
	ID:      uuid.New(),
	AppID:   uuid.New(),
	LangID:  uuid.New(),
	UsedFor: usedfor.EventType_KYCApproved.String(),
	Title:   uuid.NewString(),
	Content: uuid.NewString(),
}

var (
	id            = entNotifTemplate.ID.String()
	appID         = entNotifTemplate.AppID.String()
	langID        = entNotifTemplate.LangID.String()
	usedFor       = usedfor.EventType_KYCApproved
	NotifTemplate = notif.NotifTemplateReq{
		ID:      &id,
		AppID:   &appID,
		LangID:  &langID,
		UsedFor: &usedFor,
		Title:   &entNotifTemplate.Title,
		Content: &entNotifTemplate.Content,
	}
)

var info *ent.NotifTemplate

func rowToObject(row *ent.NotifTemplate) *ent.NotifTemplate {
	return &ent.NotifTemplate{
		ID:        row.ID,
		CreatedAt: row.CreatedAt,
		UpdatedAt: row.UpdatedAt,
		AppID:     row.AppID,
		LangID:    row.LangID,
		UsedFor:   row.UsedFor,
		Title:     row.Title,
		Content:   row.Content,
	}
}

func create(t *testing.T) {
	var err error
	info, err = Create(context.Background(), &NotifTemplate)
	if assert.Nil(t, err) {
		if assert.NotEqual(t, info.ID, uuid.UUID{}.String()) {
			entNotifTemplate.ID = info.ID
			entNotifTemplate.CreatedAt = info.CreatedAt
		}
		entNotifTemplate.UpdatedAt = info.UpdatedAt
		entNotifTemplate.CreatedAt = info.CreatedAt
		assert.Equal(t, rowToObject(info), &entNotifTemplate)
	}
}

func createBulk(t *testing.T) {
	entApp := []ent.NotifTemplate{
		{
			ID:      uuid.New(),
			AppID:   uuid.New(),
			LangID:  uuid.New(),
			UsedFor: usedfor.EventType_KYCApproved.String(),
			Title:   uuid.NewString(),
			Content: uuid.NewString(),
		},
		{
			ID:      uuid.New(),
			AppID:   uuid.New(),
			LangID:  uuid.New(),
			UsedFor: usedfor.EventType_KYCApproved.String(),
			Title:   uuid.NewString(),
			Content: uuid.NewString(),
		},
	}

	apps := []*notif.NotifTemplateReq{}
	for key := range entApp {
		id := entApp[key].ID.String()
		appID = entApp[key].AppID.String()
		langID := entApp[key].LangID.String()
		usedFor = usedfor.EventType_KYCApproved
		apps = append(apps, &notif.NotifTemplateReq{
			ID:      &id,
			AppID:   &appID,
			LangID:  &langID,
			UsedFor: &usedFor,
			Title:   &entApp[key].Title,
			Content: &entApp[key].Content,
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
	info, err = Update(context.Background(), &NotifTemplate)
	if assert.Nil(t, err) {
		entNotifTemplate.UpdatedAt = info.UpdatedAt
		entNotifTemplate.CreatedAt = info.CreatedAt
		assert.Equal(t, rowToObject(info), &entNotifTemplate)
	}
}

func row(t *testing.T) {
	var err error
	info, err = Row(context.Background(), info.ID)
	if assert.Nil(t, err) {
		entNotifTemplate.UpdatedAt = info.UpdatedAt
		entNotifTemplate.CreatedAt = info.CreatedAt
		assert.Equal(t, rowToObject(info), &entNotifTemplate)
	}
}

func rows(t *testing.T) {
	infos, total, err := Rows(context.Background(),
		&notif.Conds{
			ID: &npool.StringVal{
				Value: info.ID.String(),
				Op:    cruder.EQ,
			},
		}, 0, 0)
	if assert.Nil(t, err) {
		assert.Equal(t, total, 1)
		entNotifTemplate.UpdatedAt = infos[0].UpdatedAt
		entNotifTemplate.CreatedAt = infos[0].CreatedAt
		assert.Equal(t, rowToObject(infos[0]), &entNotifTemplate)
	}
}

func rowOnly(t *testing.T) {
	var err error
	info, err = RowOnly(context.Background(),
		&notif.Conds{
			ID: &npool.StringVal{
				Value: info.ID.String(),
				Op:    cruder.EQ,
			},
		})
	if assert.Nil(t, err) {
		entNotifTemplate.UpdatedAt = info.UpdatedAt
		entNotifTemplate.CreatedAt = info.CreatedAt
		assert.Equal(t, rowToObject(info), &entNotifTemplate)
	}
}

func count(t *testing.T) {
	count, err := Count(context.Background(),
		&notif.Conds{
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
		&notif.Conds{
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
		assert.Equal(t, rowToObject(info), &entNotifTemplate)
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
