package absence

import (
	"context"

	"github.com/pkg/errors"
	"github.com/reyhanfahlevi/soap-absence/service/absence"
	"github.com/tokopedia/affiliate/pkg/safesql"
)

// Resource struct
type Resource struct {
	master safesql.MasterDB
	db     safesql.SlaveDB
}

// New resource
func New(masterDB safesql.MasterDB, slaveDB safesql.SlaveDB) *Resource {
	return &Resource{
		db:     slaveDB,
		master: masterDB,
	}
}

// SaveUserInfo save user into db
func (r *Resource) SaveUserInfo(ctx context.Context, user absence.ParamSaveUserInfo) error {
	q := `INSERT INTO userinfo (
		pin,
		pin2,
		name,
		email,
		create_time,
		update_time
	) VALUES (?, ?, ?, ?, now(), now()) ON DUPLICATE KEY UPDATE id=id, update_time = now() `

	_, err := r.master.ExecContext(ctx, q, user.Pin1, user.Pin2, user.Name, user.Email)
	return err
}

// SavDevice save absence device
func (r *Resource) SaveDevice(ctx context.Context, param absence.ParamSaveDevice) error {
	q := `INSERT INTO device_info 
			(
				address,
				name,
				detail,
				create_time,
				active
			) 
			VALUES (
				?, ?, ?, now(), ?
			) ON DUPLICATE KEY UPDATE name=?, detail=?, update_time = now(), active = ?`

	_, err := r.master.ExecContext(ctx, q, param.Address, param.Name, param.Detail, param.Active, param.Name, param.Detail, param.Active)
	return errors.Wrap(err, "failed exec")
}

// SaveAttendanceLog save attendance log
func (r *Resource) SaveAttendanceLog(ctx context.Context, param absence.ParamSaveAttendance) error {
	q := `INSERT INTO att_log
			(
				user_id,
				tap_time,
				status,
				verified,
				work_code,
				device_address
			) VALUES ( ?, ?, ?, ?, ?, ? ) ON DUPLICATE KEY UPDATE device_address = ? `

	_, err := r.master.ExecContext(ctx, q, param.UserID, param.TapTime, param.Status, param.Verified, param.WorkCode, param.DeviceAddress, param.DeviceAddress)
	return errors.Wrapf(err, "failed exec", q)
}

// GetUserInfoByPin2 get user info by soap pin2
func (r *Resource) GetUserInfoByPin2(ctx context.Context, pin2 int) (absence.User, error) {
	var (
		user absence.User
	)
	q := `SELECT
			id,
			name,
			email,
			pin2
		FROM
			userinfo
		WHERE
			pin2 = ?
			`

	err := r.db.GetContext(ctx, &user, q, pin2)
	return user, err
}

// GetAllMachineAddress get all registered device address
func (r *Resource) GetAllMachineAddress(ctx context.Context) ([]string, error) {
	var (
		address []string
	)

	q := `SELECT address FROM device_info WHERE active = true`

	err := r.db.SelectContext(ctx, &address, q)
	return address, err
}
