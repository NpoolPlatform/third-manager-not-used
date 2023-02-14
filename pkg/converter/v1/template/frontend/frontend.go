package frontend

import (
	npool "github.com/NpoolPlatform/message/npool/third/mgr/v1/template/frontend"
	"github.com/NpoolPlatform/third-manager/pkg/db/ent"

	usedfor "github.com/NpoolPlatform/message/npool/third/mgr/v1/usedfor"
)

func Ent2Grpc(row *ent.FrontendTemplate) *npool.FrontendTemplate {
	if row == nil {
		return nil
	}

	return &npool.FrontendTemplate{
		ID:        row.ID.String(),
		AppID:     row.AppID.String(),
		LangID:    row.LangID.String(),
		UsedFor:   usedfor.UsedFor(usedfor.UsedFor_value[row.UsedFor]),
		Title:     row.Title,
		Content:   row.Content,
		Sender:    row.Sender,
		CreatedAt: row.CreatedAt,
		UpdatedAt: row.UpdatedAt,
	}
}

func Ent2GrpcMany(rows []*ent.FrontendTemplate) []*npool.FrontendTemplate {
	infos := []*npool.FrontendTemplate{}
	for _, row := range rows {
		infos = append(infos, Ent2Grpc(row))
	}
	return infos
}
