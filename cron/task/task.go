package task

import (
	"github.com/iwanbk/cron"
	"github.com/reyhanfahlevi/soap-absence/config"
	crontask "github.com/reyhanfahlevi/soap-absence/cron"
	"github.com/reyhanfahlevi/soap-absence/cron/task/absence"
)

type Task struct {
	cron       *cron.Cron
	SoapSVCs   map[string]crontask.SoapService
	AbsenceSVC crontask.AbsenceService
}

func (t *Task) Run() {
	scheduler := config.Get().Scheduler

	absence.Init(t.SoapSVCs, t.AbsenceSVC)

	t.cron = register(scheduler)
	t.cron.Start()
}

func register(schedule config.Scheduler) *cron.Cron {
	c := cron.New()

	c.AddFunc(schedule.SyncUserData, absence.HandlerSyncDeviceUserInfo)
	c.AddFunc(schedule.PullAttendanceLog, absence.HandlerPullAttendanceLog)
	return c
}
