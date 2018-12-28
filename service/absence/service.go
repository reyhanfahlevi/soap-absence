package absence

import (
	"context"
)

// Resource interface
type Resource interface {
	SaveUserInfo(ctx context.Context, user ParamSaveUserInfo) error
	SaveDevice(ctx context.Context, param ParamSaveDevice) error
	SaveAttendanceLog(ctx context.Context, param ParamSaveAttendance) error
	GetUserInfoByPin2(ctx context.Context, pin2 int) (User, error)
	GetAllMachineAddress(ctx context.Context) ([]string, error)
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
