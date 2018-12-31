package absence

import (
	"context"
	"time"

	"github.com/bmizerany/pq"
)

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
	UserID        int64  `db:"user_id,omitempty"`
	TapTime       string `db:"tap_time,omitempty"`
	Status        int    `db:"status,omitempty"`
	Verified      int    `db:"verified,omitempty"`
	WorkCode      int    `db:"work_code,omitempty"`
	DeviceAddress string `db:"device_address,omitempty"`
}

type UserAttendanceResponse struct {
	UserID     int64               `json:"user_id"`
	Email      string              `json:"email"`
	Name       string              `json:"name"`
	TapTime    []time.Time         `json:"tap_time"`
	TapTimeTxt map[string][]string `json:"tap_time_txt"`
}

type UserAttendanceResource struct {
	UserID  int64       `db:"user_id"`
	Name    string      `db:"name"`
	Email   string      `db:"email"`
	TapTime pq.NullTime `db:"tap_time"`
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

func (s *Service) GetUserAttendanceLogByID(ctx context.Context, userID int64, minDate, maxDate time.Time) (UserAttendanceResponse, error) {
	var (
		resp UserAttendanceResponse
	)

	att, err := s.resource.GetUserAttendanceLogByID(ctx, userID, minDate, maxDate)
	if err != nil {
		return resp, err
	}

	if len(att) > 0 {
		resp.UserID = att[0].UserID
		resp.Name = att[0].Name
		resp.Email = att[0].Email
		resp.TapTime = []time.Time{}
	}

	resp.TapTimeTxt = make(map[string][]string)
	timeIter := minDate
	for ok := true; ok; ok = timeIter.Before(maxDate) || timeIter.Equal(maxDate) {
		resp.TapTimeTxt[timeIter.Format("2006-01-02")] = []string{}
		timeIter = timeIter.AddDate(0, 0, 1)
	}

	for _, d := range att {
		if !d.TapTime.Valid {
			continue
		}

		resp.TapTime = append(resp.TapTime, d.TapTime.Time)
		if _, ok := resp.TapTimeTxt[d.TapTime.Time.Format("2006-01-02")]; !ok {
			resp.TapTimeTxt[d.TapTime.Time.Format("2006-01-02")] = []string{}
		}

		resp.TapTimeTxt[d.TapTime.Time.Format("2006-01-02")] = append(resp.TapTimeTxt[d.TapTime.Time.
			Format("2006-01-02")], d.TapTime.Time.Format("15:04"))
	}

	return resp, nil
}

func (s *Service) GetAllUserAttendanceLog(ctx context.Context, minDate, maxDate time.Time) (map[int64]UserAttendanceResponse, error) {
	var (
		resp map[int64]UserAttendanceResponse
	)

	att, err := s.resource.GetAllUserAttendanceLog(ctx, minDate, maxDate)
	if err != nil {
		return resp, err
	}

	resp = make(map[int64]UserAttendanceResponse)
	for _, d := range att {
		// copy from map, to avoid cannot assign to struct field  xxx in map
		tmpField, ok := resp[d.UserID]
		if !ok {
			timeIter := minDate
			defaultTimeTxt := make(map[string][]string)
			for ok := true; ok; ok = timeIter.Before(maxDate) || timeIter.Equal(maxDate) {
				defaultTimeTxt[timeIter.Format("2006-01-02")] = []string{}
				timeIter = timeIter.AddDate(0, 0, 1)
			}

			tmpField = UserAttendanceResponse{
				UserID:     d.UserID,
				TapTimeTxt: defaultTimeTxt,
				Name:       d.Name,
				Email:      d.Email,
				TapTime:    []time.Time{},
			}
		}

		if !d.TapTime.Valid {
			resp[d.UserID] = tmpField
			continue
		}

		tmp := append(tmpField.TapTime, d.TapTime.Time)
		tmpField.TapTime = tmp
		if _, ok := tmpField.TapTimeTxt[d.TapTime.Time.Format("2006-01-02")]; !ok {
			tmpField.TapTimeTxt[d.TapTime.Time.Format("2006-01-02")] = []string{}
		}

		tmpField.TapTimeTxt[d.TapTime.Time.Format("2006-01-02")] = append(tmpField.TapTimeTxt[d.TapTime.Time.
			Format("2006-01-02")], d.TapTime.Time.Format("15:04"))

		resp[d.UserID] = tmpField
	}

	return resp, nil
}
