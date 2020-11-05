package agimpl

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"net/url"
	"strconv"
)

type Agent struct {
	Dialer *websocket.Conn
}

func NewAgent(addr string, port int, path string) *Agent {
	diaURL := url.URL{
		Scheme: "ws",
		Host:   addr + ":" + strconv.Itoa(port),
		Path:   path,
	}
	conn, _, err := websocket.DefaultDialer.Dial(diaURL.String(), nil)
	if err != nil {
		logrus.Fatal(err)
	}
	return &Agent{
		Dialer: conn,
	}
}

func (a *Agent) ReadAndMessageHandle() {
	for {
		msgType, msgData, err := a.Dialer.ReadMessage()
		if err != nil {
			if !websocket.IsCloseError(err, websocket.CloseGoingAway, websocket.CloseNormalClosure) {
				logrus.Errorln(err)
			}
			break
		}
		switch msgType {
		case websocket.TextMessage:
			fmt.Println(string(msgData))
		}
	}
}

func (a *Agent) WriteCloseMessage() error {
	return a.Dialer.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "closed by client"))

}

func (a *Agent) WriteTextMessage(text string) error {
	return a.Dialer.WriteMessage(websocket.TextMessage, []byte(text))

}

func (a *Agent) WriteBinaryMessage(data []byte) error {
	return a.Dialer.WriteMessage(websocket.BinaryMessage, data)
}
