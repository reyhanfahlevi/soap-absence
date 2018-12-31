package absence

import (
	"net/http"
	"strconv"
	"time"

	"github.com/pkg/errors"
	"github.com/reyhanfahlevi/soap-absence/pkg/response"
	"github.com/reyhanfahlevi/soap-absence/service/absence"
)

func HandlerGetUserAttendance(w http.ResponseWriter, r *http.Request) {
	var (
		resp       response.Response
		res        absence.UserAttendanceResponse
		err        error
		timeLayout = "2006-01-02"
	)
	defer resp.Render(w, r)

	userID, _ := strconv.ParseInt(r.FormValue("user_id"), 10, 64)
	startDate, err := time.Parse(timeLayout, r.FormValue("start_date"))
	if err != nil {
		startDate = time.Now()
	}

	endDate, err := time.Parse(timeLayout, r.FormValue("end_date"))
	if err != nil {
		endDate = time.Now()
	}

	//email := r.FormValue("email")

	if userID != 0 {
		res, err = absenceSvc.GetUserAttendanceLogByID(r.Context(), userID, startDate, endDate)
		if err != nil {
			resp.SetError(err)
			return
		}
	}

	resp.Data = res
	return
}

func HandlerGetAllUserAttendance(w http.ResponseWriter, r *http.Request) {
	var (
		resp       response.Response
		timeLayout = "2006-01-02"
	)
	defer resp.Render(w, r)

	startDate, err := time.Parse(timeLayout, r.FormValue("start_date"))
	if err != nil {
		startDate = time.Now()
	}

	endDate, err := time.Parse(timeLayout, r.FormValue("end_date"))
	if err != nil {
		endDate = time.Now()
	}

	if startDate.After(endDate) {
		resp.SetError(errors.New("invalid date range"))
		return
	}

	res, err := absenceSvc.GetAllUserAttendanceLog(r.Context(), startDate, endDate)
	if err != nil {
		resp.SetError(err)
		return
	}

	resp.Data = res
	return
}
