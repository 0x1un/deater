package impl

import (
	"bytes"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

const (
	writeWait      = 10 * time.Second
	pongWait       = 60 * time.Second
	pingPeriod     = (pongWait * 9) / 10
	maxMessageSize = 400000
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  maxMessageSize,
	WriteBufferSize: 1024,
}

type Client struct {
	conn *websocket.Conn
	send chan []byte
}

func (c *Client) readPump() {
	defer func() {
		_ = c.conn.Close()
	}()
	c.conn.SetReadLimit(maxMessageSize)
	for {
		if err := c.conn.SetReadDeadline(time.Now().Add(pongWait)); err != nil {
			logrus.Errorln(err)
		}
		msgType, msgData, err := c.conn.ReadMessage()
		if err != nil {
			if !websocket.IsCloseError(err, websocket.CloseNormalClosure, websocket.CloseGoingAway) {
				logrus.Errorf("read message error: %s", err.Error())
			}
			break
		}
		switch msgType {
		case websocket.BinaryMessage:
			if len(msgData) != 0 {
				fmt.Println(string(msgData))
			}
		case websocket.TextMessage:
			msgData = bytes.TrimSpace(bytes.Replace(msgData, newline, space, -1))
			msg := string(msgData)
			fmt.Println("recv:", msg)
			c.writePump(websocket.TextMessage, msgData)
		}
	}
}

func (c *Client) writePump(msgType int, msgData []byte) {
	if err := c.conn.WriteMessage(msgType, msgData); err != nil {
		logrus.Errorln("write:", err)
	}
}

func (c *Client) keepalive(timeout time.Duration) {
	lastResp := time.Now()
	c.conn.SetPongHandler(func(msg string) error {
		lastResp = time.Now()
		return nil
	})
	go func() {
		for {
			err := c.conn.WriteMessage(websocket.PingMessage, []byte("keepalive"))
			if err != nil {
				return
			}
			time.Sleep((timeout / 2) * time.Second)
			if time.Now().Sub(lastResp) > timeout {
				_ = c.conn.Close()
				return
			}
		}
	}()
}

func ServeWs(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		logrus.Errorln(err)
		return
	}
	client := &Client{
		conn: conn,
		send: make(chan []byte, 1024),
	}
	go client.readPump()
}
