package notif

import (
	"fmt"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/third/mgr/v1/template/notif"

	usedfor "github.com/NpoolPlatform/message/npool/notif/mgr/v1/notif"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func validate(info *npool.NotifTemplateReq) error {
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
	case usedfor.EventType_WithdrawalRequest:
	case usedfor.EventType_WithdrawalCompleted:
	case usedfor.EventType_DepositReceived:
	case usedfor.EventType_KYCApproved:
	case usedfor.EventType_KYCRejected:
	default:
		return fmt.Errorf("EventType is invalid")
	}

	return nil
}
