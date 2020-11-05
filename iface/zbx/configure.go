package zbx

import (
	"errors"
	"github.com/0x1un/go-zabbix"
	"io/ioutil"
	"os"
)

const (
	defaultExportFileName string = "zbx_any_host_exported"
	defaultExportFileType string = "json"
)

type configureFile struct {
	filename string
	fileType string
	fileData []byte
}

func (c *configureFile) Store() error {
	if len(c.fileData) == 0 {
		return errors.New("数据为空")
	}
	return ioutil.WriteFile(c.filename+"."+c.fileType, c.fileData, os.ModePerm)
}

func (c *configureFile) Data() []byte {
	return c.fileData
}

// ExportSelectedHosts 导出已选择的主机
// selectedID hosts id list
// formatType xml json
func (z *Zbx) ExportSelectedHosts(selectedID []string, formatType string) ([]byte, error) {
	param := zabbix.ConfigurationParamsRequest{
		Format: formatType,
		Options: zabbix.ConfiguraOption{
			Hosts: selectedID,
		},
	}
	respData, err := z.session.ConfiguraExport(param)
	if err != nil {
		return nil, err
	}
	return []byte(respData), nil
}

func (z *Zbx) ExportAnyHosts(filename string, formatType string) (*configureFile, error) {
	hosts, err := z.getAnyHosts()
	if err != nil {
		return nil, err
	}
	if formatType == "" {
		formatType = defaultExportFileType
	}
	bytes, err := z.ExportSelectedHosts(hosts, formatType)
	if err != nil {
		return nil, err
	}
	if filename == "" {
		filename = defaultExportFileName
	}
	return &configureFile{
		filename: filename,
		fileType: formatType,
		fileData: bytes,
	}, nil
}
