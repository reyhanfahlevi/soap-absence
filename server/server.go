package server

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/reyhanfahlevi/soap-absence/api/http"
	"github.com/reyhanfahlevi/soap-absence/config"
	"github.com/reyhanfahlevi/soap-absence/container"
)

func Main() error {
	err := config.Init()
	if err != nil {
		return err
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
