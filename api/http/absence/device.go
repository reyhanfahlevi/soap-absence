package absence

import (
	"net/http"

	"github.com/reyhanfahlevi/soap-absence/api"
	"github.com/reyhanfahlevi/soap-absence/pkg/response"
	"github.com/reyhanfahlevi/soap-absence/service/absence"
)

var (
	absenceSvc api.AbsenceService
)

func Init(abseSvc api.AbsenceService) {
	absenceSvc = abseSvc
}

func HandlerAddNewDevice(w http.ResponseWriter, r *http.Request) {
	var (
		paramDevice absence.ParamSaveDevice
		resp        response.Response
	)
	defer resp.Render(w, r)

	paramDevice.Address = r.FormValue("address")
	paramDevice.Name = r.FormValue("name")
	paramDevice.Detail = r.FormValue("detail")

	if r.FormValue("active") != "false" {
		paramDevice.Active = true
	}

	err := absenceSvc.SaveDevice(r.Context(), paramDevice)
	if err != nil {
		resp.SetError(err, http.StatusUnprocessableEntity)
		return
	}

	resp.Data = struct {
		Success bool
	}{
		true,
	}
	return
}
