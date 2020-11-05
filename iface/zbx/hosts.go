package zbx

import "github.com/0x1un/go-zabbix"

func (z *Zbx) getAnyHosts() ([]string, error) {
	param := zabbix.HostGetParams{}
	param.OutputFields = append([]string{}, "hostid")
	hosts, err := z.session.GetHosts(param)
	if err != nil {
		return nil, err
	}
	ids := make([]string, 0)
	for _, host := range hosts {
		if host.HostID == "" {
			continue
		}
		ids = append(ids, host.HostID)
	}
	return ids, nil
}
