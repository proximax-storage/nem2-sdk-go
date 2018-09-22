package sdk

import (
	j "encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"golang.org/x/net/websocket"
	"io"
	"strings"
)

type servicews struct {
	client *ClientWs
}

// Catapult Websocket Client configuration
type ClientWs struct {
	client    *websocket.Conn
	Uid       string
	config    *Config
	common    servicews // Reuse a single struct instead of allocating one for each service on the heap.
	Subscribe *SubscribeService
}

type SubscribeMsg struct {
	UID       string `json:"uid"`
	Subscribe string `json:"subscribe"`
}

func (c *ClientWs) changeURLPort() {
	split := strings.Split(c.config.BaseURL.Host, ":")
	host, port := split[0], "3000"
	c.config.BaseURL.Host = strings.Join([]string{host, port}, ":")
}

func NewClientWs(websocketClient *websocket.Conn, conf *Config) *ClientWs {
	if websocketClient == nil {
		//panic("ws cannot be nil")
	}

	c := &ClientWs{client: websocketClient, config: conf}
	c.common.client = c
	c.Subscribe = (*SubscribeService)(&c.common)

	return c
}

func (c *ClientWs) BuildSubscribe(destination string) (*SubscribeMsg, error) {
	var b SubscribeMsg
	b.UID = c.Uid
	b.Subscribe = destination
	return &b, nil
}

func (c *ClientWs) WsConnect() error {
	c.config.BaseURL.Scheme = "ws"
	c.config.BaseURL.Path = "/ws"
	c.changeURLPort()
	conn, err := websocket.Dial(c.config.BaseURL.String(), "", "http://localhost")
	if err != nil {
		panic(err)
		return err
	}
	c.client = conn

	var msg []byte
	if err = websocket.Message.Receive(c.client, &msg); err != nil {
		return err
	}

	Parser, err := msgParser(msg)
	if err != nil {
		fmt.Println(err)
	}

	if err != nil {
		return err
	}
	c.Uid = Parser.UID

	return nil
}

func (c *ClientWs) Subs(msg *SubscribeMsg, out chan []byte) (chan []byte, error) {
	if err := websocket.JSON.Send(c.client, msg); err != nil {
		return nil, err
	}

	var e error
	go func() {
		var resp []byte

		for {
			if err := websocket.Message.Receive(c.client, &resp); err == io.EOF {
				err = c.WsConnect()
				if err != nil {
					return
				}
				if err = websocket.JSON.Send(c.client, msg); err != nil {
					return
				}
				continue
			} else if err != nil {
				e = errors.Wrap(err, "Error occurred while trying to receive message")
			}

			SubName, _ := restParser(resp)

			if msg.Subscribe == SubName && SubName == "block" {
				out <- resp
				continue
			} else if SubName == strings.Split(msg.Subscribe, "/")[0] {
				out <- resp
			}
		}
	}()
	return out, e
}

func msgParser(msg []byte) (*SubscribeMsg, error) {
	var message SubscribeMsg
	err := json.Unmarshal(msg, &message)
	if err != nil {
		return nil, err
	}
	return &message, nil
}

func restParser(data []byte) (string, error) {
	var raw []j.RawMessage
	err := json.Unmarshal([]byte(fmt.Sprintf("[%v]", string(data))), &raw)
	if err != nil {
		return "", err
	}

	for _, r := range raw {
		var obj map[string]interface{}
		err := json.Unmarshal(r, &obj)
		if err != nil {
			return "", err
		}

		var subscribe string
		if _, ok := obj["block"]; ok {
			subscribe = "block"
		} else if _, ok := obj["status"]; ok {
			subscribe = "status"
		} else if v, ok := obj["meta"]; ok {
			channelName := v.(map[string]interface{})
			subscribe = fmt.Sprintf("%v", channelName["channelName"])
		}
		return subscribe, nil
	}
	return "", nil
}
