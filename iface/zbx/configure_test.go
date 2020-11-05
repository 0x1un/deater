package zbx

import (
	"os"
	"testing"
)

var (
	zbxUrl  = os.Getenv("ZBX_URL")
	zbxUser = os.Getenv("ZBX_USER")
	zbxPass = os.Getenv("ZBX_PASS")
)

func TestZbx_ExportAnyHosts(t *testing.T) {
	zbx, err := NewZBXConnector(zbxUrl, zbxUser, zbxPass)
	if err != nil {
		t.Fatal(err)
	}
	config, err := zbx.ExportAnyHosts("", "")
	if err != nil {
		t.Fatal(err)
	}
	if err := config.Store(); err != nil {
		t.Fatal(err)
	}
}
