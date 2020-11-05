package main

import (
	"fmt"
	conf "github.com/0x1un/deater/agent/agent_conf"
	"github.com/0x1un/deater/agent/agimpl"
	"github.com/0x1un/deater/iface/zbx"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"log"
	"math"
	"os"
	"os/signal"
	"time"
)

var (
	z   *zbx.Zbx
	cfg *viper.Viper
)

func init() {
	cfg = conf.ReadAgentConfig("conf/duter_agentd.yaml")
	var err error
	zbxURL := fmt.Sprintf("http://%s:%s/%s", cfg.GetString("zbx.zbx_api_host"), cfg.GetString("zbx.zbx_api_port"), cfg.GetString("zbx.zbx_api_suffix"))
	z, err = zbx.NewZBXConnector(zbxURL, cfg.GetString("zbx.zbx_api_username"), cfg.GetString("zbx.zbx_api_password"))
	if err != nil {
		logrus.Fatal(err)
	}
	_ = z
}

func main() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	agent := agimpl.NewAgent(cfg.GetString("duter_server_host"), cfg.GetInt("duter_server_port"), "ws")
	defer func() { _ = agent.Dialer.Close() }()
	go agent.ReadAndMessageHandle()
	go func() {
		for i := 0; i < math.MaxUint16; i++ {
			if err := agent.WriteTextMessage(fmt.Sprintf("%d", i)); err != nil {
				logrus.Errorln(err)
			}
			ticker := time.NewTicker(1 * time.Second)
			<-ticker.C
		}
	}()

	for {
		select {
		case <-interrupt:
			log.Println("Ctrl-C...")
			if err := agent.WriteCloseMessage(); err != nil {
				log.Println(err)
				return
			}
			return
		}
	}
}
