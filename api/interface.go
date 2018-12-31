package api

import (
	"context"

	"time"

	"github.com/reyhanfahlevi/soap-absence/service/absence"
	"github.com/reyhanfahlevi/soap-absence/service/soap"
)

type AbsenceService interface {
	SaveUserInfo(ctx context.Context, param absence.ParamSaveUserInfo) error
	SaveDevice(ctx context.Context, param absence.ParamSaveDevice) error
	GetUserAttendanceLogByID(ctx context.Context, userID int64, minDate, maxDate time.Time) (absence.UserAttendanceResponse, error)
	GetAllUserAttendanceLog(ctx context.Context, minDate, maxDate time.Time) (map[int64]absence.UserAttendanceResponse, error)
}

type SoapService interface {
	GetAllUserInfo(ctx context.Context) (soap.GetAllUserInfoResponse, error)
	GetAttendanceLog(ctx context.Context, pin ...int) (soap.AttendanceLogResponse, error)
	GetUserInfo(ctx context.Context, pin int) (soap.GetUserInfoResponse, error)
}
