package grcp

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func GrpcErrorToFiberError(err error) *fiber.Error {
	s, _ := status.FromError(err)
	code := s.Code()

	var httpStatus int
	switch code {
	case codes.OK:
		httpStatus = fiber.StatusOK
	case codes.InvalidArgument:
		httpStatus = fiber.StatusBadRequest
	case codes.NotFound:
		httpStatus = fiber.StatusNotFound
	case codes.AlreadyExists:
		httpStatus = fiber.StatusConflict
	case codes.PermissionDenied:
		httpStatus = fiber.StatusForbidden
	case codes.Unauthenticated:
		httpStatus = fiber.StatusUnauthorized
	case codes.ResourceExhausted:
		httpStatus = fiber.StatusTooManyRequests
	case codes.FailedPrecondition, codes.Aborted:
		httpStatus = fiber.StatusPreconditionFailed
	case codes.OutOfRange:
		httpStatus = fiber.StatusRequestedRangeNotSatisfiable
	case codes.Internal:
		httpStatus = fiber.StatusInternalServerError
	case codes.Unavailable:
		httpStatus = fiber.StatusServiceUnavailable
	case codes.DataLoss:
		httpStatus = fiber.StatusPartialContent
	case codes.Unimplemented:
		httpStatus = fiber.StatusNotImplemented
	default:
		httpStatus = fiber.StatusInternalServerError
	}

	return fiber.NewError(httpStatus, fmt.Sprintf("%s: %s", s.Code().String(), s.Message()))
}
