// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package container

import (
	"github.com/google/wire"
	"github.com/reyhanfahlevi/soap-absence/config"
	absence2 "github.com/reyhanfahlevi/soap-absence/resource/absence"
	soap2 "github.com/reyhanfahlevi/soap-absence/resource/soap"
	"github.com/reyhanfahlevi/soap-absence/service/absence"
	"github.com/reyhanfahlevi/soap-absence/service/soap"
	"github.com/tokopedia/affiliate/pkg/httpclient"
	"github.com/tokopedia/affiliate/pkg/safesql"
)

// Injectors from wire.go:

func InitializeAbsenceService() (*absence.Service, error) {
	masterDB, err := MasterDBProvider()
	if err != nil {
		return nil, err
	}
	slaveDB, err := SlaveDBProvider()
	if err != nil {
		return nil, err
	}
	resource := absence2.New(masterDB, slaveDB)
	service := absence.New(resource)
	return service, nil
}

func InitializeSoapService(address string) (*soap.Service, error) {
	client := HttpClientProvider()
	resource := soap2.New(client, address)
	service := soap.New(resource)
	return service, nil
}

// wire.go:

var AbsenceResourceProvider = wire.NewSet(absence2.New, wire.Bind(new(absence.Resource), new(absence2.Resource)))

var SoapResourceProvider = wire.NewSet(soap2.New, wire.Bind(new(soap.Resource), new(soap2.Resource)))

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
