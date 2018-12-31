package task

import (
	"log"

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
	schedule := config.Get().Scheduler

	absence.Init(t.SoapSVCs, t.AbsenceSVC)

	t.cron = scheduler(schedule)
	t.cron.Start()
}

func scheduler(schedule config.Scheduler) *cron.Cron {
	c := cron.New()

	register(c, schedule.SyncUserData, absence.HandlerSyncDeviceUserInfo, "syncronize user")
	register(c, schedule.PullAttendanceLog, absence.HandlerPullAttendanceLog, "pull attendance log")
	return c
}

func register(c *cron.Cron, schedule string, cmd func(), name ...string) {
	if len(name) > 0 {
		log.Printf("registering %v cron fro %v", name[0], schedule)
	}
	c.AddFunc(schedule, cmd)
}
