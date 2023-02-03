package notif

import (
	npool "github.com/NpoolPlatform/message/npool/third/mgr/v1/template/notif"
	"github.com/NpoolPlatform/third-manager/pkg/db/ent"

	usedfor "github.com/NpoolPlatform/message/npool/notif/mgr/v1/notif"
)

func Ent2Grpc(row *ent.NotifTemplate) *npool.NotifTemplate {
	if row == nil {
		return nil
	}

	return &npool.NotifTemplate{
		ID:        row.ID.String(),
		AppID:     row.AppID.String(),
		LangID:    row.LangID.String(),
		UsedFor:   usedfor.EventType(usedfor.EventType_value[row.UsedFor]),
		Title:     row.Title,
		Content:   row.Content,
		Sender:    row.Sender,
		CreatedAt: row.CreatedAt,
		UpdatedAt: row.UpdatedAt,
	}
}

func Ent2GrpcMany(rows []*ent.NotifTemplate) []*npool.NotifTemplate {
	infos := []*npool.NotifTemplate{}
	for _, row := range rows {
		infos = append(infos, Ent2Grpc(row))
	}
	return infos
}
