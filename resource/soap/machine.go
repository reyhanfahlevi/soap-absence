package soap

import (
	"context"

	"encoding/xml"
	"fmt"
	"io/ioutil"

	"github.com/reyhanfahlevi/soap-absence/service/soap"
)

// GetAllUserInfo get all user info
func (r *Resource) GetAllUserInfo(ctx context.Context) (soap.GetAllUserInfoResponse, error) {
	var (
		result soap.GetAllUserInfoResponse
		body   = `<GetAllUserInfo><ArgComKey xsi:type=":xsd:integer">0</ArgComKey></GetAllUserInfo>`
	)

	headers := map[string][]string{
		"Content-Type": {"text/xml"},
	}

	resp, err := r.client.Post(ctx, fmt.Sprintf("%s/iWsService", r.address), headers, body)
	if err != nil {
		return result, err
	}
	defer resp.Body.Close()

	// read response body
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return result, err
	}

	err = xml.Unmarshal(b, &result)
	if err != nil {
		return result, err
	}

	result.Total = len(result.Users)
	return result, nil
}

// GetAttendanceLog get attendance log
func (r *Resource) GetAttendanceLog(ctx context.Context, pin ...int) (soap.AttendanceLogResponse, error) {
	var (
		result soap.AttendanceLogResponse
		body   = `<GetAttLog><ArgComKey xsi:type="xsd:integer">0</ArgComKey><Arg><PIN xsi:type="xsd:integer">%v</PIN></Arg></GetAttLog>`
	)

	if len(pin) > 0 {
		body = fmt.Sprintf(body, pin[0])
	} else {
		body = fmt.Sprintf(body, "All")
	}

	headers := map[string][]string{
		"Content-Type": {"text/xml"},
	}

	resp, err := r.client.Post(ctx, fmt.Sprintf("%s/iWsService", r.address), headers, body)
	if err != nil {
		return result, err
	}
	defer resp.Body.Close()

	// read response body
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return result, err
	}

	err = xml.Unmarshal(b, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}

// GetUserInfo get user info
func (r *Resource) GetUserInfo(ctx context.Context, pin int) (soap.GetUserInfoResponse, error) {
	var (
		result soap.GetUserInfoResponse
		body   = `<GetUserInfo><ArgComKey xsi:type=":xsd:integer">0</ArgComKey>
					<Arg>
						<PIN Xsi:type="xsd:integer">%v</PIN>
					</Arg>
				</GetUserInfo>`
	)

	body = fmt.Sprintf(body, pin)

	headers := map[string][]string{
		"Content-Type": {"text/xml"},
	}

	resp, err := r.client.Post(ctx, fmt.Sprintf("%s/iWsService", r.address), headers, body)
	if err != nil {
		return result, err
	}
	defer resp.Body.Close()

	// read response body
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return result, err
	}

	err = xml.Unmarshal(b, &result)
	if err != nil {
		return result, err
	}

	return result, nil
}
