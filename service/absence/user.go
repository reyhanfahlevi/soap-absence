package absence

import "context"

type ParamSaveUserInfo struct {
	User
}

type ParamSaveAttendance struct {
	AttendanceLog
}

type User struct {
	ID    int64  `db:"id"`
	Email string `db:"email"`
	Pin1  int    `db:"pin"`
	Pin2  int    `db:"pin2"`
	Name  string `db:"name"`
}

type AttendanceLog struct {
	UserID        int64  `db:"user_id"`
	TapTime       string `db:"tap_time"`
	Status        int    `db:"status"`
	Verified      int    `db:"verified"`
	WorkCode      int    `db:"work_code"`
	DeviceAddress string `db:"device_address"`
}

func (s *Service) SaveUserInfo(ctx context.Context, user ParamSaveUserInfo) error {
	return s.resource.SaveUserInfo(ctx, user)
}

func (s *Service) SaveAttendanceLog(ctx context.Context, param ParamSaveAttendance) error {
	return s.resource.SaveAttendanceLog(ctx, param)
}

func (s *Service) GetUserInfoByPin2(ctx context.Context, pin2 int) (User, error) {
	return s.resource.GetUserInfoByPin2(ctx, pin2)
}
