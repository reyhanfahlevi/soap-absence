package soap

import (
	"github.com/tokopedia/affiliate/pkg/httpclient"
)

// Resource struct
type Resource struct {
	client  *httpclient.Client
	address string
}

// New resource
func New(client *httpclient.Client, address string) *Resource {
	return &Resource{
		client:  client,
		address: address,
	}
}

// GetDeviceAddress get resource device address
func (r *Resource) GetDeviceAddress() string {
	return r.address
}
