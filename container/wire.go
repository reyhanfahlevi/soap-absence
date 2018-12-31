//+build wireinject

package container

import (
	"github.com/google/wire"
	"github.com/jmoiron/sqlx"
	"github.com/reyhanfahlevi/soap-absence/config"
	absenceresource "github.com/reyhanfahlevi/soap-absence/resource/absence"
	soapres "github.com/reyhanfahlevi/soap-absence/resource/soap"
	absenceservice "github.com/reyhanfahlevi/soap-absence/service/absence"
	soapsvc "github.com/reyhanfahlevi/soap-absence/service/soap"
	"github.com/tokopedia/affiliate/pkg/httpclient"
	"github.com/tokopedia/affiliate/pkg/log"
)

var AbsenceResourceProvider = wire.NewSet(absenceresource.New, wire.Bind(new(absenceservice.Resource), new(absenceresource.Resource)))
var SoapResourceProvider = wire.NewSet(soapres.New, wire.Bind(new(soapsvc.Resource), new(soapres.Resource)))

// DBProvider db provider
func DBProvider() (*sqlx.DB, error) {
	log.Println("connecting db")
	dsn := config.Get().DB.Master
	db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	log.Println("connection established")
	return db, nil
}

// HttpClientProvider http client provider
func HttpClientProvider() *httpclient.Client {
	return httpclient.NewClient()
}

// InitializeAbsenceService init absence service
func InitializeAbsenceService() (*absenceservice.Service, error) {
	wire.Build(DBProvider, AbsenceResourceProvider, absenceservice.New)
	return &absenceservice.Service{}, nil
}

func InitializeSoapService(address string) (*soapsvc.Service, error) {
	wire.Build(HttpClientProvider, SoapResourceProvider, soapsvc.New)
	return &soapsvc.Service{}, nil
}
