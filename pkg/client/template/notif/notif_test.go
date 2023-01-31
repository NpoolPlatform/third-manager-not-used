package notif

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"google.golang.org/grpc/credentials/insecure"

	"github.com/NpoolPlatform/go-service-framework/pkg/config"

	"bou.ke/monkey"
	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	"google.golang.org/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	val "github.com/NpoolPlatform/message/npool"

	npool "github.com/NpoolPlatform/message/npool/third/mgr/v1/template/notif"
	testinit "github.com/NpoolPlatform/third-manager/pkg/testinit"
	"github.com/google/uuid"

	"github.com/stretchr/testify/assert"

	usedfor "github.com/NpoolPlatform/message/npool/notif/mgr/v1/notif"
)

func init() {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	if err := testinit.Init(); err != nil {
		fmt.Printf("cannot init test stub: %v\n", err)
	}
}

var appDate = npool.NotifTemplate{
	ID:      uuid.NewString(),
	AppID:   uuid.NewString(),
	LangID:  uuid.NewString(),
	UsedFor: usedfor.EventType_KYCApproved,
	Title:   uuid.NewString(),
	Content: uuid.NewString(),
}

var (
	appInfo = npool.NotifTemplateReq{
		ID:      &appDate.ID,
		AppID:   &appDate.AppID,
		LangID:  &appDate.LangID,
		UsedFor: &appDate.UsedFor,
		Title:   &appDate.Title,
		Content: &appDate.Content,
	}
)

var info *npool.NotifTemplate

func createNotifTemplate(t *testing.T) {
	var err error
	info, err = CreateNotifTemplate(context.Background(), &appInfo)
	if assert.Nil(t, err) {
		appDate.UpdatedAt = info.UpdatedAt
		appDate.CreatedAt = info.CreatedAt
		assert.Equal(t, info, &appDate)
	}
}

func createNotifTemplates(t *testing.T) {
	appDates := []npool.NotifTemplate{
		{
			ID:      uuid.NewString(),
			AppID:   uuid.NewString(),
			LangID:  uuid.NewString(),
			UsedFor: usedfor.EventType_KYCApproved,
			Title:   uuid.NewString(),
			Content: uuid.NewString(),
		},
		{
			ID:      uuid.NewString(),
			AppID:   uuid.NewString(),
			LangID:  uuid.NewString(),
			UsedFor: usedfor.EventType_KYCApproved,
			Title:   uuid.NewString(),
			Content: uuid.NewString(),
		},
	}

	apps := []*npool.NotifTemplateReq{}
	for key := range appDates {
		apps = append(apps, &npool.NotifTemplateReq{
			ID:      &appDates[key].ID,
			AppID:   &appDates[key].AppID,
			LangID:  &appDates[key].LangID,
			UsedFor: &appDates[key].UsedFor,
			Title:   &appDates[key].Title,
			Content: &appDates[key].Content,
		})
	}

	infos, err := CreateNotifTemplates(context.Background(), apps)
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 2)
	}
}

func updateNotifTemplate(t *testing.T) {
	var err error
	info, err = UpdateNotifTemplate(context.Background(), &appInfo)
	if assert.Nil(t, err) {
		appDate.UpdatedAt = info.UpdatedAt
		appDate.CreatedAt = info.CreatedAt
		assert.Equal(t, info, &appDate)
	}
}

func getNotifTemplate(t *testing.T) {
	var err error
	info, err = GetNotifTemplate(context.Background(), info.ID)
	if assert.Nil(t, err) {
		appDate.UpdatedAt = info.UpdatedAt
		appDate.CreatedAt = info.CreatedAt
		assert.Equal(t, info, &appDate)
	}
}

func getNotifTemplates(t *testing.T) {
	infos, total, err := GetNotifTemplates(context.Background(),
		&npool.Conds{
			ID: &val.StringVal{
				Value: info.ID,
				Op:    cruder.EQ,
			},
		}, 0, 1)
	if assert.Nil(t, err) {
		assert.Equal(t, total, uint32(1))
		appDate.UpdatedAt = infos[0].UpdatedAt
		appDate.CreatedAt = infos[0].CreatedAt
		assert.Equal(t, infos[0], &appDate)
	}
}

func getNotifTemplateOnly(t *testing.T) {
	var err error
	info, err = GetNotifTemplateOnly(context.Background(),
		&npool.Conds{
			ID: &val.StringVal{
				Value: info.ID,
				Op:    cruder.EQ,
			},
		})
	if assert.Nil(t, err) {
		appDate.UpdatedAt = info.UpdatedAt
		appDate.CreatedAt = info.CreatedAt
		assert.Equal(t, info, &appDate)
	}
}

func existNotifTemplate(t *testing.T) {
	exist, err := ExistNotifTemplate(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, exist, true)
	}
}

func existNotifTemplateConds(t *testing.T) {
	exist, err := ExistNotifTemplateConds(context.Background(),
		&npool.Conds{
			ID: &val.StringVal{
				Value: info.ID,
				Op:    cruder.EQ,
			},
		},
	)
	if assert.Nil(t, err) {
		assert.Equal(t, exist, true)
	}
}

func deleteNotifTemplate(t *testing.T) {
	info, err := DeleteNotifTemplate(context.Background(), info.ID)
	if assert.Nil(t, err) {
		appDate.UpdatedAt = info.UpdatedAt
		appDate.CreatedAt = info.CreatedAt
		assert.Equal(t, info, &appDate)
	}
}

func TestMainOrder(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	gport := config.GetIntValueWithNameSpace("", config.KeyGRPCPort)

	monkey.Patch(grpc2.GetGRPCConn, func(service string, tags ...string) (*grpc.ClientConn, error) {
		return grpc.Dial(fmt.Sprintf("localhost:%v", gport), grpc.WithTransportCredentials(insecure.NewCredentials()))
	})

	t.Run("createNotifTemplate", createNotifTemplate)
	t.Run("createNotifTemplates", createNotifTemplates)
	t.Run("getNotifTemplate", getNotifTemplate)
	t.Run("getNotifTemplates", getNotifTemplates)
	t.Run("getNotifTemplateOnly", getNotifTemplateOnly)
	t.Run("updateNotifTemplate", updateNotifTemplate)
	t.Run("existNotifTemplate", existNotifTemplate)
	t.Run("existNotifTemplateConds", existNotifTemplateConds)
	t.Run("delete", deleteNotifTemplate)
}
