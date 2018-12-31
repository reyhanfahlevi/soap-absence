package absence

import (
	"context"
	"log"

	"sync"

	"github.com/reyhanfahlevi/soap-absence/container"
	"github.com/reyhanfahlevi/soap-absence/cron"
	"github.com/reyhanfahlevi/soap-absence/service/absence"
)

var soapSvc map[string]cron.SoapService
var absenceSvc cron.AbsenceService

var (
	syncUserMutex     sync.Mutex
	syncPullAttMutext sync.Mutex
)

// Init will initialize
func Init(soap map[string]cron.SoapService, absence cron.AbsenceService) {
	soapSvc = soap
	absenceSvc = absence
}

func HandlerPullAttendanceLog() {
	var (
		totalDevice = 0
	)
	ctx := context.Background()

	syncPullAttMutext.Lock()

	for _, v := range soapSvc {
		att, err := v.GetAttendanceLog(ctx)
		if err != nil {
			log.Println(err)
			continue
		}

		for _, a := range att.Attendance {
			user, err := absenceSvc.GetUserInfoByPin2(ctx, a.PIN)
			if err != nil {
				log.Println(err)
				continue
			}

			saveData := absence.ParamSaveAttendance{
				AttendanceLog: absence.AttendanceLog{
					UserID:        user.ID,
					TapTime:       a.DateTime,
					WorkCode:      a.WorkCode,
					Verified:      a.Verified,
					Status:        a.Status,
					DeviceAddress: v.GetDeviceAddress(),
				},
			}

			err = absenceSvc.SaveAttendanceLog(ctx, saveData)
			if err != nil {
				log.Println(err)
				continue
			}
		}

		totalDevice++
	}

	syncPullAttMutext.Unlock()

	log.Printf("-- Total %v Devices --", totalDevice)
}

func HandlerSyncDeviceUserInfo() {
	var (
		totalDevice  = 0
		totalSuccess = 0
	)
	ctx := context.Background()

	syncUserMutex.Lock()

	log.Println("syncronizing user")

	for _, v := range soapSvc {
		users, err := v.GetAllUserInfo(ctx)
		if err != nil {
			log.Println(err)
			continue
		}

		for _, u := range users.Users {
			param := absence.ParamSaveUserInfo{
				User: absence.User{
					Name: u.Name,
					Pin1: u.PIN,
					Pin2: u.PIN2,
				},
			}

			err := absenceSvc.SaveUserInfo(ctx, param)
			if err != nil {
				log.Println(err)
				continue
			}

			totalSuccess++
		}

		totalDevice++

	}

	syncUserMutex.Unlock()

	log.Printf("-- Total Success: %v From %v Devices --", totalSuccess, totalDevice)
}

func HandlerSyncDevices() {
	var (
		countNewDevice = 0
	)

	ctx := context.Background()
	devices, err := absenceSvc.GetAllMachineAddress(ctx)
	if err != nil {
		log.Println(err)
		return
	}

	for _, d := range devices {
		if _, ok := soapSvc[d]; ok {
			continue
		}

		soapSvc[d], err = container.InitializeSoapService(d)
		if err != nil {
			log.Println(err)
			return
		}

		countNewDevice++
	}

	log.Printf("-- Total New Device = %v, Device Total = %v", countNewDevice, len(soapSvc))
}
