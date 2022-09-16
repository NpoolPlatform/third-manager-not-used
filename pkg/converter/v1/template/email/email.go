package email

import (
	npool "github.com/NpoolPlatform/message/npool/third/mgr/v1/template/email"
	"github.com/NpoolPlatform/third-manager/pkg/db/ent"

	"github.com/NpoolPlatform/message/npool/third/mgr/v1/usedfor"
)

func Ent2Grpc(row *ent.AppEmailTemplate) *npool.EmailTemplate {
	if row == nil {
		return nil
	}

	return &npool.EmailTemplate{
		ID:                row.ID.String(),
		AppID:             row.AppID.String(),
		LangID:            row.LangID.String(),
		UsedFor:           usedfor.UsedFor(usedfor.UsedFor_value[row.UsedFor]),
		Sender:            row.Sender,
		ReplyTos:          row.ReplyTos,
		CCTos:             row.CcTos,
		Subject:           row.Subject,
		Body:              row.Body,
		DefaultToUsername: row.DefaultToUsername,
	}
}

func Ent2GrpcMany(rows []*ent.AppEmailTemplate) []*npool.EmailTemplate {
	infos := []*npool.EmailTemplate{}
	for _, row := range rows {
		infos = append(infos, Ent2Grpc(row))
	}
	return infos
}
