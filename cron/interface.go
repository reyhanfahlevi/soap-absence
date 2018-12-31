package cron

import (
	"context"

	"github.com/reyhanfahlevi/soap-absence/service/absence"
	"github.com/reyhanfahlevi/soap-absence/service/soap"
)

type AbsenceService interface {
	SaveUserInfo(ctx context.Context, param absence.ParamSaveUserInfo) error
	SaveAttendanceLog(ctx context.Context, param absence.ParamSaveAttendance) error
	GetUserInfoByPin2(ctx context.Context, pin2 int) (absence.User, error)
	GetDevicesAddress(ctx context.Context) ([]string, error)
}

type SoapService interface {
	GetAllUserInfo(ctx context.Context) (soap.GetAllUserInfoResponse, error)
	GetAttendanceLog(ctx context.Context, pin ...int) (soap.AttendanceLogResponse, error)
	GetDeviceAddress() string
}
