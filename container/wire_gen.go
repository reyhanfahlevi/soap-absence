// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package container

import (
	"github.com/google/wire"
	absence2 "github.com/reyhanfahlevi/soap-absence/resource/absence"
	"github.com/reyhanfahlevi/soap-absence/service/absence"
	"github.com/tokopedia/affiliate/pkg/httpclient"
	"github.com/tokopedia/affiliate/pkg/safesql"
)

// Injectors from wire.go:

func InitializeService(address ...string) (*absence.Service, error) {
	client := HttpClientProvider()
	masterDB, err := MasterDBProvider()
	if err != nil {
		return nil, err
	}
	slaveDB, err := SlaveDBProvider()
	if err != nil {
		return nil, err
	}
	resource := absence2.New(client, masterDB, slaveDB, address...)
	service := absence.New(resource)
	return service, nil
}

// wire.go:

var ResourceProvider = wire.NewSet(absence2.New, wire.Bind(new(absence.Resource), new(absence2.Resource)))

func MasterDBProvider() (safesql.MasterDB, error) {
	return safesql.OpenMasterDB("mysql", "root:root@/attendance")
}

func SlaveDBProvider() (safesql.SlaveDB, error) {
	return safesql.OpenSlaveDB("mysql", "root:root@/attendance")
}

func HttpClientProvider() *httpclient.Client {
	return httpclient.NewClient()
}
