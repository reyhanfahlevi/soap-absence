//+build wireinject

package container

import (
	"github.com/google/wire"
	"github.com/reyhanfahlevi/soap-absence/config"
	absenceresource "github.com/reyhanfahlevi/soap-absence/resource/absence"
	absenceservice "github.com/reyhanfahlevi/soap-absence/service/absence"
	"github.com/tokopedia/affiliate/pkg/httpclient"
	"github.com/tokopedia/affiliate/pkg/safesql"
)

var ResourceProvider = wire.NewSet(absenceresource.New, wire.Bind(new(absenceservice.Resource), new(absenceresource.Resource)))

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
func InitializeAbsenceService(address ...string) (*absenceservice.Service, error) {
	wire.Build(HttpClientProvider, MasterDBProvider, SlaveDBProvider, ResourceProvider, absenceservice.New)
	return &absenceservice.Service{}, nil
}
