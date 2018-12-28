package validator

import (
	"net/url"

	"github.com/pkg/errors"
)

var (
	ErrProtocolNotSupported = errors.New("protocol scheme not supported")
)

func ValidateURL(urls string) (bool, error) {
	_, err := url.ParseRequestURI(urls)
	if err != nil {
		return false, ErrProtocolNotSupported
	}

	return true, nil
}
