package constant

import (
	"context"
	"fmt"
	"github.com/NpoolPlatform/third-manager/pkg/db/ent"
	"github.com/NpoolPlatform/third-manager/pkg/testinit"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/NpoolPlatform/message/npool"
	"github.com/NpoolPlatform/message/npool/third/mgr/v1/contact"
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

var entAppContact = ent.AppContact{
	ID:          uuid.New(),
	AppID:       uuid.New(),
	UsedFor:     usedfor.UsedFor_Signin.String(),
	Sender:      uuid.NewString(),
	Account:     uuid.NewString(),
	AccountType: uuid.NewString(),
}

var (
	id         = entAppContact.ID.String()
	appID      = entAppContact.AppID.String()
	usedFor    = usedfor.UsedFor_Signin
	appContact = contact.ContactReq{
		ID:          &id,
		AppID:       &appID,
		UsedFor:     &usedFor,
		Account:     &entAppContact.Account,
		AccountType: &entAppContact.AccountType,
		Sender:      &entAppContact.Sender,
	}
)

var info *ent.AppContact

func rowToObject(row *ent.AppContact) *ent.AppContact {
	return &ent.AppContact{
		ID:          row.ID,
		CreatedAt:   row.CreatedAt,
		AppID:       row.AppID,
		UsedFor:     row.UsedFor,
		Sender:      row.Sender,
		Account:     row.Account,
		AccountType: row.AccountType,
	}
}

func create(t *testing.T) {
	var err error
	info, err = Create(context.Background(), &appContact)
	if assert.Nil(t, err) {
		if assert.NotEqual(t, info.ID, uuid.UUID{}.String()) {
			entAppContact.ID = info.ID
			entAppContact.CreatedAt = info.CreatedAt
		}
		assert.Equal(t, rowToObject(info), &entAppContact)
	}
}

func createBulk(t *testing.T) {
	entApp := []ent.AppContact{
		{
			ID:          uuid.New(),
			AppID:       uuid.New(),
			UsedFor:     usedfor.UsedFor_Signin.String(),
			Sender:      uuid.NewString(),
			Account:     uuid.NewString(),
			AccountType: uuid.NewString(),
		},
		{
			ID:          uuid.New(),
			AppID:       uuid.New(),
			UsedFor:     usedfor.UsedFor_Signin.String(),
			Sender:      uuid.NewString(),
			Account:     uuid.NewString(),
			AccountType: uuid.NewString(),
		},
	}

	apps := []*contact.ContactReq{}
	for key := range entApp {
		id := entApp[key].ID.String()
		appID = entAppContact.AppID.String()
		usedFor = usedfor.UsedFor_Signin
		apps = append(apps, &contact.ContactReq{
			ID:          &id,
			AppID:       &appID,
			UsedFor:     &usedFor,
			Account:     &entApp[key].Account,
			AccountType: &entApp[key].AccountType,
			Sender:      &entApp[key].Sender,
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
	info, err = Update(context.Background(), &appContact)
	if assert.Nil(t, err) {
		assert.Equal(t, rowToObject(info), &entAppContact)
	}
}

func row(t *testing.T) {
	var err error
	info, err = Row(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, rowToObject(info), &entAppContact)
	}
}

func rows(t *testing.T) {
	infos, total, err := Rows(context.Background(),
		&contact.Conds{
			ID: &npool.StringVal{
				Value: info.ID.String(),
				Op:    cruder.EQ,
			},
		}, 0, 0)
	if assert.Nil(t, err) {
		assert.Equal(t, total, 1)
		assert.Equal(t, rowToObject(infos[0]), &entAppContact)
	}
}

func rowOnly(t *testing.T) {
	var err error
	info, err = RowOnly(context.Background(),
		&contact.Conds{
			ID: &npool.StringVal{
				Value: info.ID.String(),
				Op:    cruder.EQ,
			},
		})
	if assert.Nil(t, err) {
		assert.Equal(t, rowToObject(info), &entAppContact)
	}
}

func count(t *testing.T) {
	count, err := Count(context.Background(),
		&contact.Conds{
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
		&contact.Conds{
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
		assert.Equal(t, rowToObject(info), &entAppContact)
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
