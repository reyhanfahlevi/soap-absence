//+build wireinject

package container

import (
	"github.com/google/wire"
	"github.com/reyhanfahlevi/soap-absence/config"
	absenceresource "github.com/reyhanfahlevi/soap-absence/resource/absence"
	soapres "github.com/reyhanfahlevi/soap-absence/resource/soap"
	absenceservice "github.com/reyhanfahlevi/soap-absence/service/absence"
	soapsvc "github.com/reyhanfahlevi/soap-absence/service/soap"
	"github.com/tokopedia/affiliate/pkg/httpclient"
	"github.com/tokopedia/affiliate/pkg/safesql"
)

var AbsenceResourceProvider = wire.NewSet(absenceresource.New, wire.Bind(new(absenceservice.Resource), new(absenceresource.Resource)))
var SoapResourceProvider = wire.NewSet(soapres.New, wire.Bind(new(soapsvc.Resource), new(soapres.Resource)))

// MasterDBProvider master db provider
func MasterDBProvider() (safesql.MasterDB, error) {
	conf := config.Get()
	return safesql.OpenMasterDB("mysql", conf.DB.Master)
}

// SlaveDBProvider slave db provider
func SlaveDBProvider() (safesql.SlaveDB, error) {
	conf := config.Get()
	return safesql.OpenSlaveDB("mysql", conf.DB.Slave)
}

// HttpClientProvider http client provider
func HttpClientProvider() *httpclient.Client {
	return httpclient.NewClient()
}

// InitializeAbsenceService init absence service
func InitializeAbsenceService() (*absenceservice.Service, error) {
	wire.Build(MasterDBProvider, SlaveDBProvider, AbsenceResourceProvider, absenceservice.New)
	return &absenceservice.Service{}, nil
}

func InitializeSoapService(address string) (*soapsvc.Service, error) {
	wire.Build(HttpClientProvider, SoapResourceProvider, soapsvc.New)
	return &soapsvc.Service{}, nil
}
