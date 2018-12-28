package server

import (
	"log"

	"github.com/reyhanfahlevi/soap-absence/api/http"
	"github.com/reyhanfahlevi/soap-absence/config"
	"github.com/reyhanfahlevi/soap-absence/container"
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

	server := http.Server{
		AbsenceSvc: absenSvc,
	}

	server.Serve(config.Get().App.ApiPort)
	return nil
}
