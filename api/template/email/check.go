package email

import (
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/third/mgr/v1/template/email"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func validate(info *npool.EmailTemplateReq) error {
	if info.AppID == nil {
		logger.Sugar().Errorw("validate", "AppID", info.AppID)
		return status.Error(codes.InvalidArgument, "AppID is empty")
	}

	if _, err := uuid.Parse(info.GetAppID()); err != nil {
		logger.Sugar().Errorw("validate", "AppID", info.GetAppID(), "error", err)
		return status.Error(codes.InvalidArgument, "AppID is invalid")
	}
	if info.LangID == nil {
		logger.Sugar().Errorw("validate", "LangID", info.LangID)
		return status.Error(codes.InvalidArgument, "LangID is empty")
	}

	if _, err := uuid.Parse(info.GetLangID()); err != nil {
		logger.Sugar().Errorw("validate", "LangID", info.GetLangID(), "error", err)
		return status.Error(codes.InvalidArgument, "LangID is invalid")
	}

	if info.DefaultToUsername == nil || info.GetDefaultToUsername() == "" {
		logger.Sugar().Errorw("validate", "DefaultToUsername", info.DefaultToUsername, "GetDefaultToUsername", info.GetDefaultToUsername())
		return status.Error(codes.InvalidArgument, "DefaultToUsername is empty")
	}

	if info.UsedFor == nil {
		logger.Sugar().Errorw("validate", "UsedFor", info.UsedFor, "GetUsedFor", info.GetUsedFor())
		return status.Error(codes.InvalidArgument, "UsedFor is empty")
	}

	if info.Sender == nil || info.GetSender() == "" {
		logger.Sugar().Errorw("validate", "Sender", info.Sender, "GetSender", info.GetSender())
		return status.Error(codes.InvalidArgument, "Sender is empty")
	}

	if info.Subject == nil || info.GetSender() == "" {
		logger.Sugar().Errorw("validate", "Subject", info.Sender, "GetSubject", info.GetSender())
		return status.Error(codes.InvalidArgument, "Subject is empty")
	}

	if info.Body == nil || info.GetBody() == "" {
		logger.Sugar().Errorw("validate", "Body", info.Body, "GetBody", info.GetBody())
		return status.Error(codes.InvalidArgument, "Body is empty")
	}

	return nil
}

func Validate(info *npool.EmailTemplateReq) error {
	return validate(info)
}
