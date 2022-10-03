package sms

import (
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/third/mgr/v1/template/sms"
	"github.com/NpoolPlatform/message/npool/third/mgr/v1/usedfor"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func validate(info *npool.SMSTemplateReq) error { //nolint
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

	if info.UsedFor == nil {
		logger.Sugar().Errorw("validate", "UsedFor", info.UsedFor, "GetUsedFor", info.GetUsedFor())
		return status.Error(codes.InvalidArgument, "UsedFor is empty")
	}

	switch info.GetUsedFor() {
	case usedfor.UsedFor_Signup:
	case usedfor.UsedFor_Signin:
	case usedfor.UsedFor_Update:
	case usedfor.UsedFor_Contact:
	case usedfor.UsedFor_SetWithdrawAddress:
	case usedfor.UsedFor_Withdraw:
	case usedfor.UsedFor_CreateInvitationCode:
	case usedfor.UsedFor_SetCommission:
	case usedfor.UsedFor_SetTransferTargetUser:
	case usedfor.UsedFor_Transfer:
	default:
		return status.Error(codes.InvalidArgument, "Invalid UsedFor")
	}

	if info.Subject == nil || info.GetSubject() == "" {
		logger.Sugar().Errorw("validate", "Subject", info.Subject, "GetSubject", info.GetSubject())
		return status.Error(codes.InvalidArgument, "Subject is empty")
	}

	if info.Message == nil || info.GetMessage() == "" {
		logger.Sugar().Errorw("validate", "Message", info.Message, "GetMessage", info.GetMessage())
		return status.Error(codes.InvalidArgument, "Message is empty")
	}

	return nil
}

func Validate(info *npool.SMSTemplateReq) error {
	return validate(info)
}
