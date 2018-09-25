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
	client        *websocket.Conn
	Uid           string
	config        *Config
	common        servicews // Reuse a single struct instead of allocating one for each service on the heap.
	Subscribe     *SubscribeService
	subscriptions map[string]chan<- []byte
}

type Subscribe struct {
	UID       string `json:"uid"`
	Subscribe string `json:"subscribe"`
	ChIn      chan []byte
	conn      *websocket.Conn
}

func (c *ClientWs) changeURLPort() {
	split := strings.Split(c.config.BaseURL.Host, ":")
	host, port := split[0], "3000"
	c.config.BaseURL.Host = strings.Join([]string{host, port}, ":")
}

func NewConnectWs(conf *Config) (*ClientWs, error) {
	c := &ClientWs{config: conf}
	c.common.client = c
	c.Subscribe = (*SubscribeService)(&c.common)
	c.subscriptions = make(map[string]chan<- []byte)

	err := c.wsconnect()
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (c *ClientWs) BuildSubscribe(destination string) *Subscribe {
	b := new(Subscribe)
	b.ChIn = make(chan []byte)
	subName := strings.Split(destination, "/")[0]
	c.subscriptions[subName] = b.ChIn
	b.UID = c.Uid
	b.Subscribe = destination
	b.conn = c.client
	return b
}

func (c *ClientWs) wsconnect() error {
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

	imsg, err := msgparser(msg)
	if err != nil {
		fmt.Println(err)
	}

	if err != nil {
		return err
	}
	c.Uid = imsg.UID

	return nil
}

func (c *ClientWs) SubsChannel(msg *Subscribe) error {
	if err := websocket.JSON.Send(c.client, struct {
		UID       string `json:"uid"`
		Subscribe string `json:"subscribe"`
	}{
		UID:       msg.UID,
		Subscribe: msg.Subscribe,
	}); err != nil {
		return err
	}

	var e error
	go func() {
		var resp []byte

		for {
			if err := websocket.Message.Receive(c.client, &resp); err == io.EOF {
				err = c.wsconnect()
				if err != nil {
					return
				}
				if err = websocket.JSON.Send(c.client, struct {
					UID       string `json:"uid"`
					Subscribe string `json:"subscribe"`
				}{
					UID:       msg.UID,
					Subscribe: msg.Subscribe,
				}); err != nil {
					return
				}
				continue
			} else if err != nil {
				e = errors.Wrap(err, "Error occurred while trying to receive message")
			}

			subName, _ := restparser(resp)
			c.subscriptions[subName] <- resp
		}
	}()
	return e
}

func msgparser(msg []byte) (*Subscribe, error) {
	var message Subscribe
	err := json.Unmarshal(msg, &message)
	if err != nil {
		return nil, err
	}
	return &message, nil
}

func restparser(data []byte) (string, error) {
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
