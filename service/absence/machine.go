package absence

import (
	"context"

	"github.com/reyhanfahlevi/soap-absence/pkg/validator"
)

type ParamSaveDevice struct {
	Address string
	Name    string
	Detail  string
	Active  bool
}

// SaveDevice save device info
func (s *Service) SaveDevice(ctx context.Context, param ParamSaveDevice) error {
	_, err := validator.ValidateURL(param.Address)
	if err != nil {
		return err
	}
	return s.resource.SaveDevice(ctx, param)
}

// GetDevicesAddress get all the device address
func (s *Service) GetDevicesAddress(ctx context.Context) ([]string, error) {
	return s.resource.GetAllMachineAddress(ctx)
}
