package absence

import (
	"context"
	"encoding/xml"
)

// Resource interface
type Resource interface {
	GetAllUserInfo(ctx context.Context) (GetAllUserInfoResponse, error)
	GetAttendanceLog(ctx context.Context, pin ...int) (AttendanceLogResponse, error)
	SaveUserInfo(ctx context.Context, user UserInfo) error
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
		PIN      string   `xml:"PIN"`
		Status   string   `xml:"Status"`
		Verified string   `xml:"Verified"`
		WorkCode string   `xml:"WorkCode"`
	} `xml:"Row"`
}

// New service
func New(res Resource) *Service {
	return &Service{
		resource: res,
	}
}

// GetAllUserInfo get all user
func (s *Service) GetAllUserInfo(ctx context.Context) (GetAllUserInfoResponse, error) {
	return s.resource.GetAllUserInfo(ctx)
}

// GetAttendanceLog info
func (s *Service) GetAttendanceLog(ctx context.Context, pin ...int) (AttendanceLogResponse, error) {
	return s.resource.GetAttendanceLog(ctx, pin...)
}

func (s *Service) SaveUserInfo(ctx context.Context, user UserInfo) error {
	return s.resource.SaveUserInfo(ctx, user)
}
