package main

import (
	"github.com/0x1un/deater/server/impl"
	"github.com/sirupsen/logrus"
	"net/http"
)

func main() {
	logrus.SetReportCaller(true)
	http.HandleFunc("/ws", impl.ServeWs)
	logrus.Fatal(http.ListenAndServe(":8888", nil))
}
