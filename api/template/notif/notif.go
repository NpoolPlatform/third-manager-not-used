//nolint:nolintlint,dupl
package notif

import (
	"context"

	usedfor "github.com/NpoolPlatform/message/npool/notif/mgr/v1/notif"

	converter "github.com/NpoolPlatform/third-manager/pkg/converter/v1/template/notif"
	crud "github.com/NpoolPlatform/third-manager/pkg/crud/v1/template/notif"
	constant "github.com/NpoolPlatform/third-manager/pkg/message/const"
	tracer "github.com/NpoolPlatform/third-manager/pkg/tracer/template/notif"

	commontracer "github.com/NpoolPlatform/third-manager/pkg/tracer"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/third/mgr/v1/template/notif"

	"github.com/google/uuid"
)

func (s *Server) CreateNotifTemplate(
	ctx context.Context,
	in *npool.CreateNotifTemplateRequest,
) (
	*npool.CreateNotifTemplateResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateNotifTemplate")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.Trace(span, in.GetInfo())

	err = validate(in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("CreateNotifTemplate", "error", err)
		return &npool.CreateNotifTemplateResponse{}, err
	}

	span = commontracer.TraceInvoker(span, "app", "crud", "Create")

	info, err := crud.Create(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("CreateNotifTemplate", "error", err)
		return &npool.CreateNotifTemplateResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateNotifTemplateResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) CreateNotifTemplates(
	ctx context.Context,
	in *npool.CreateNotifTemplatesRequest,
) (
	*npool.CreateNotifTemplatesResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateNotifTemplates")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if len(in.GetInfos()) == 0 {
		logger.Sugar().Errorw("CreateNotifTemplates", "error", "Infos is empty")
		return &npool.CreateNotifTemplatesResponse{}, status.Error(codes.InvalidArgument, "Infos is empty")
	}

	span = tracer.TraceMany(span, in.GetInfos())
	span = commontracer.TraceInvoker(span, "app", "crud", "CreateBulk")

	rows, err := crud.CreateBulk(ctx, in.GetInfos())
	if err != nil {
		logger.Sugar().Errorw("CreateNotifTemplates", "error", err)
		return &npool.CreateNotifTemplatesResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateNotifTemplatesResponse{
		Infos: converter.Ent2GrpcMany(rows),
	}, nil
}

//nolint:gocyclo
func (s *Server) UpdateNotifTemplate(
	ctx context.Context,
	in *npool.UpdateNotifTemplateRequest,
) (
	*npool.UpdateNotifTemplateResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "UpdateNotifTemplate")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.Trace(span, in.GetInfo())

	if _, err := uuid.Parse(in.GetInfo().GetID()); err != nil {
		logger.Sugar().Errorw("UpdateNotifTemplate", "ID", in.GetInfo().GetID(), "error", err)
		return &npool.UpdateNotifTemplateResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	if _, err := uuid.Parse(in.GetInfo().GetLangID()); err != nil && in.Info.LangID != nil {
		logger.Sugar().Errorw("UpdateSMSTemplate", "LangID", in.GetInfo().GetLangID(), "error", err)
		return &npool.UpdateNotifTemplateResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	if in.GetInfo().UsedFor != nil {
		switch in.GetInfo().GetUsedFor() {
		case usedfor.EventType_WithdrawalRequest:
		case usedfor.EventType_WithdrawalCompleted:
		case usedfor.EventType_DepositReceived:
		case usedfor.EventType_KYCApproved:
		case usedfor.EventType_KYCRejected:
		default:
			logger.Sugar().Errorw("UpdateNotifTemplate", "ID", in.GetInfo().GetID(), "error", "EventType is invalid")
			return &npool.UpdateNotifTemplateResponse{}, status.Error(codes.Internal, "EventType is invalid")
		}
	}

	if in.GetInfo().Title != nil && in.GetInfo().GetTitle() == "" {
		logger.Sugar().Errorw("UpdateNotifTemplate", "ID", in.GetInfo().GetID(), "error", "Title is empty")
		return &npool.UpdateNotifTemplateResponse{}, status.Error(codes.Internal, "Title is empty")
	}

	if in.GetInfo().Content != nil && in.GetInfo().GetContent() == "" {
		logger.Sugar().Errorw("UpdateNotifTemplate", "ID", in.GetInfo().GetID(), "error", "Content is empty")
		return &npool.UpdateNotifTemplateResponse{}, status.Error(codes.Internal, "Content is empty")
	}

	span = commontracer.TraceInvoker(span, "app", "crud", "Update")

	info, err := crud.Update(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("UpdateNotifTemplate", "ID", in.GetInfo().GetID(), "error", err)
		return &npool.UpdateNotifTemplateResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateNotifTemplateResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetNotifTemplate(
	ctx context.Context,
	in *npool.GetNotifTemplateRequest,
) (
	*npool.GetNotifTemplateResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetNotifTemplate")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceID(span, in.GetID())

	id, err := uuid.Parse(in.GetID())
	if err != nil {
		logger.Sugar().Errorw("GetNotifTemplate", "ID", in.GetID(), "error", err)
		return &npool.GetNotifTemplateResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "app", "crud", "Row")

	info, err := crud.Row(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("GetNotifTemplate", "ID", in.GetID(), "error", err)
		return &npool.GetNotifTemplateResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetNotifTemplateResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetNotifTemplateOnly(
	ctx context.Context,
	in *npool.GetNotifTemplateOnlyRequest,
) (
	*npool.GetNotifTemplateOnlyResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetNotifTemplateOnly")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "app", "crud", "RowOnly")

	info, err := crud.RowOnly(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("GetNotifTemplateOnly", "error", err)
		return &npool.GetNotifTemplateOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetNotifTemplateOnlyResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetNotifTemplates(
	ctx context.Context,
	in *npool.GetNotifTemplatesRequest,
) (
	*npool.GetNotifTemplatesResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetNotifTemplates")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceOffsetLimit(span, int(in.GetOffset()), int(in.GetLimit()))
	span = commontracer.TraceInvoker(span, "app", "crud", "Rows")

	rows, total, err := crud.Rows(ctx, in.GetConds(), int(in.GetOffset()), int(in.GetLimit()))
	if err != nil {
		logger.Sugar().Errorw("GetNotifTemplates", "error", err)
		return &npool.GetNotifTemplatesResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetNotifTemplatesResponse{
		Infos: converter.Ent2GrpcMany(rows),
		Total: uint32(total),
	}, nil
}

func (s *Server) ExistNotifTemplate(
	ctx context.Context,
	in *npool.ExistNotifTemplateRequest,
) (
	*npool.ExistNotifTemplateResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistNotifTemplate")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceID(span, in.GetID())

	id, err := uuid.Parse(in.GetID())
	if err != nil {
		logger.Sugar().Errorw("ExistNotifTemplate", "ID", in.GetID(), "error", err)
		return &npool.ExistNotifTemplateResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "app", "crud", "Exist")

	exist, err := crud.Exist(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("ExistNotifTemplate", "ID", in.GetID(), "error", err)
		return &npool.ExistNotifTemplateResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistNotifTemplateResponse{
		Info: exist,
	}, nil
}

func (s *Server) ExistNotifTemplateConds(
	ctx context.Context,
	in *npool.ExistNotifTemplateCondsRequest,
) (
	*npool.ExistNotifTemplateCondsResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistNotifTemplateConds")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "app", "crud", "ExistNotifTemplateConds")

	exist, err := crud.ExistConds(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("ExistNotifTemplateConds", "error", err)
		return &npool.ExistNotifTemplateCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistNotifTemplateCondsResponse{
		Info: exist,
	}, nil
}

func (s *Server) DeleteNotifTemplate(
	ctx context.Context,
	in *npool.DeleteNotifTemplateRequest,
) (
	*npool.DeleteNotifTemplateResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "DeleteNotifTemplate")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceID(span, in.GetID())

	id, err := uuid.Parse(in.GetID())
	if err != nil {
		logger.Sugar().Errorw("DeleteNotifTemplate", "ID", in.GetID(), "error", err)
		return &npool.DeleteNotifTemplateResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "app", "crud", "Delete")

	info, err := crud.Delete(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("DeleteNotifTemplate", "ID", in.GetID(), "error", err)
		return &npool.DeleteNotifTemplateResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteNotifTemplateResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}
