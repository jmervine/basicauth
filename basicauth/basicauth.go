package basicauth

import (
	"encoding/base64"
	"fmt"
)

type BasicAuth struct {
	Username, Password string
}

func NewBasicAuth(u, p string) *BasicAuth {
	ba := new(BasicAuth)
	ba.Username = u
	ba.Password = p
	return ba
}

func (ba *BasicAuth) String() string {
	return base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", ba.Username, ba.Password)))
}

func (ba *BasicAuth) Header() string {
	return fmt.Sprintf("Authorization: Basic %s", ba.String())
}
