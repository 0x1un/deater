package conf

import (
	"fmt"
	"testing"
)

func TestReadAgentConfig(t *testing.T) {
	conf := ReadAgentConfig("../conf/duter_agentd.yaml")
	fmt.Println(conf)
}
