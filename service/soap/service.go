package soap

import (
	"context"
	"encoding/xml"
)

// Resource interface
type Resource interface {
	GetAllUserInfo(ctx context.Context) (GetAllUserInfoResponse, error)
	GetAttendanceLog(ctx context.Context, pin ...int) (AttendanceLogResponse, error)
	GetDeviceAddress() string
}

// Service struct
type Service struct {
	resource Resource
}

// GetAllUserInfoResponse struct
type GetAllUserInfoResponse struct {
	XMLName xml.Name   `xml:"GetAllUserInfoResponse"`
	Users   []UserInfo `xml:"Row"`
	Total   int
}

type GetUserInfoResponse struct {
	XMLName xml.Name `xml:"GetUserInfo"`
	User    UserInfo `xml:"Row"`
}

// UserInfo struct
type UserInfo struct {
	XMLName   xml.Name `xml:"Row"`
	Card      string   `xml:"Card"`
	Group     string   `xml:"Group"`
	Name      string   `xml:"Name"`
	PIN       int      `xml:"PIN"`
	PIN2      int      `xml:"PIN2"`
	Privilege string   `xml:"Privilege"`
	TZ1       string   `xml:"TZ1"`
	TZ2       string   `xml:"TZ2"`
	TZ3       string   `xml:"TZ3"`
}

// AttendanceLogResponse struct
type AttendanceLogResponse struct {
	XMLName    xml.Name `xml:"GetAttLogResponse"`
	Attendance []struct {
		XMLName  xml.Name `xml:"Row"`
		DateTime string   `xml:"DateTime"`
		PIN      int      `xml:"PIN"`
		Status   int      `xml:"Status"`
		Verified int      `xml:"Verified"`
		WorkCode int      `xml:"WorkCode"`
	} `xml:"Row"`
}

// New service
func New(res Resource) *Service {
	return &Service{
		resource: res,
	}
}

// GetDeviceAddress get current service device address
func (s *Service) GetDeviceAddress() string {
	return s.resource.GetDeviceAddress()
}

// GetAllUserInfo get all user
func (s *Service) GetAllUserInfo(ctx context.Context) (GetAllUserInfoResponse, error) {
	return s.resource.GetAllUserInfo(ctx)
}

// GetAttendanceLog info
func (s *Service) GetAttendanceLog(ctx context.Context, pin ...int) (AttendanceLogResponse, error) {
	return s.resource.GetAttendanceLog(ctx, pin...)
}
