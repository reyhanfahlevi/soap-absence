package absence

import (
	"context"
	"time"
)

// Resource interface
type Resource interface {
	SaveUserInfo(ctx context.Context, user ParamSaveUserInfo) error
	SaveDevice(ctx context.Context, param ParamSaveDevice) error
	SaveAttendanceLog(ctx context.Context, param ParamSaveAttendance) error
	GetUserInfoByPin2(ctx context.Context, pin2 int) (User, error)
	GetAllMachineAddress(ctx context.Context) ([]string, error)
	GetUserAttendanceLogByID(ctx context.Context, userID int64, minDate, maxDate time.Time) ([]UserAttendanceResource, error)
	GetAllUserAttendanceLog(ctx context.Context, minDate, maxDate time.Time) ([]UserAttendanceResource, error)
}

// Service struct
type Service struct {
	resource Resource
}

// New service
func New(res Resource) *Service {
	return &Service{
		resource: res,
	}
}
