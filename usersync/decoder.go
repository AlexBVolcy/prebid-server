package usersync

import (
	"encoding/base64"
	"encoding/json"
)

type Decoder interface {
	// Decode takes an encoded string and decodes it into a cookie
	Decode(v string) *Cookie
}

type DecodeV1 struct{}

func (d DecodeV1) Decode(encodedValue string) *Cookie {
	jsonValue, err := base64.URLEncoding.DecodeString(encodedValue)
	if err != nil {
		return NewCookie()
	}

	var cookie Cookie
	if err = json.Unmarshal(jsonValue, &cookie); err != nil {
		return NewCookie()
	}

	return &cookie
}