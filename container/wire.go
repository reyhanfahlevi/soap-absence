//+build wireinject

package container

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	absenceresource "github.com/reyhanfahlevi/soap-absence/resource/absence"
	absenceservice "github.com/reyhanfahlevi/soap-absence/service/absence"
	"github.com/tokopedia/affiliate/pkg/httpclient"
	"github.com/tokopedia/affiliate/pkg/safesql"
)

var ResourceProvider = wire.NewSet(absenceresource.New, wire.Bind(new(absenceservice.Resource), new(absenceresource.Resource)))

func MasterDBProvider() (safesql.MasterDB, error) {
	return safesql.OpenMasterDB("mysql", "root:root@/attendance")
}

func SlaveDBProvider() (safesql.SlaveDB, error) {
	return safesql.OpenSlaveDB("mysql", "root:root@/attendance")
}

func HttpClientProvider() *httpclient.Client {
	return httpclient.NewClient()
}

func InitializeService(address ...string) (*absenceservice.Service, error) {
	wire.Build(HttpClientProvider, MasterDBProvider, SlaveDBProvider, ResourceProvider, absenceservice.New)
	return &absenceservice.Service{}, nil
}
