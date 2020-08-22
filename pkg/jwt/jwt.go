package jwt

import (
	"encoding/base64"
	"errors"
	"fmt"
	"strings"
)

func Decode(token string) (string, error) {
	s := strings.Split(token, ".")

	if len(s) != 3 {
		return "", errors.New("jwt decode: invalid JWT structure")
	}

	header, err := base64.RawStdEncoding.DecodeString(s[0])
	if err != nil {
		return "", errors.New("jwt decode: base64 decode header failed")
	}

	payload, err := base64.RawStdEncoding.DecodeString(s[1])
	if err != nil {
		return "", errors.New("jwt decode: base64 decode payload failed")
	}

	signature := s[2]

	return fmt.Sprintf(`{"header":%s,"payload":%s,"signature":"%s"}`, header, payload, signature), nil
}
