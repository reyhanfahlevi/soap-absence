package scheduler

import (
	"context"
	"log"

	"os"
	"os/signal"
	"syscall"

	"github.com/reyhanfahlevi/soap-absence/config"
	"github.com/reyhanfahlevi/soap-absence/container"
	"github.com/reyhanfahlevi/soap-absence/cron"
	"github.com/reyhanfahlevi/soap-absence/cron/task"
)

func Main() error {
	err := config.Init()
	if err != nil {
		return err
	}

	conf := config.Get()
	if len(conf.App.DeviceIPList) == 0 {
		log.Fatal("Empty Device")
	}

	absenSvc, err := container.InitializeAbsenceService()
	if err != nil {
		return err
	}

	// fetching address
	address, err := absenSvc.GetDevicesAddress(context.Background())
	if err != nil {
		return err
	}

	soapSvcDevices := make(map[string]cron.SoapService)
	for _, ip := range address {
		soapSvc, err := container.InitializeSoapService(ip)
		if err != nil {
			return err
		}

		soapSvcDevices[ip] = soapSvc
	}

	cronTask := task.Task{
		AbsenceSVC: absenSvc,
		SoapSVCs:   soapSvcDevices,
	}

	cronTask.Run()

	term := make(chan os.Signal, 1)
	signal.Notify(term, os.Interrupt, syscall.SIGTERM)
	select {
	case <-term:
		log.Println("Signal terminate detected")
	}

	log.Println("👋")
	return nil
}
