package zbx

import (
	"github.com/0x1un/go-zabbix"
)

type Zbx struct {
	session *zabbix.Session
}

func NewZBXConnector(url, username, password string) (*Zbx, error) {
	ses, err := zabbix.NewSession(url, username, password)
	if err != nil {
		return nil, err
	}
	return &Zbx{session: ses}, nil
}
