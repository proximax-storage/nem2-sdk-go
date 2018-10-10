package sdk

import (
	"bytes"
	j "encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"golang.org/x/net/websocket"
	"io"
	"net/url"
	"strings"
)

type serviceWs struct {
	client *ClientWs
}

var chanSubscribe = struct {
	block  chan BlockInfo
	status chan StatusInfo
}{}

// Catapult Websocket Client configuration
type ClientWs struct {
	client        *websocket.Conn
	Uid           string
	config        *Config
	common        serviceWs // Reuse a single struct instead of allocating one for each service on the heap.
	Subscribe     *SubscribeService
	subscriptions map[string]chan<- interface{}
}

type Subscribe struct {
	Uid       string `json:"uid"`
	Subscribe string `json:"subscribe"`
	ChIn      chan interface{}
	conn      *websocket.Conn
}

type sendJson struct {
	Uid       string `json:"uid"`
	Subscribe string `json:"subscribe"`
}

func (c *ClientWs) changeURLPort() {
	c.config.BaseURL.Scheme = "ws"
	c.config.BaseURL.Path = "/ws"
	split := strings.Split(c.config.BaseURL.Host, ":")
	host, port := split[0], "3000"
	c.config.BaseURL.Host = strings.Join([]string{host, port}, ":")
}

func NewConnectWs(host string) (*ClientWs, error) {
	u, err := url.Parse(host)
	if err != nil {
		return nil, err
	}
	newconf := &Config{BaseURL: u}
	c := &ClientWs{config: newconf}
	c.common.client = c
	c.Subscribe = (*SubscribeService)(&c.common)
	c.subscriptions = make(map[string]chan<- interface{})

	err = c.wsConnect()
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (c *ClientWs) buildSubscribe(destination string) *Subscribe {
	b := new(Subscribe)
	b.ChIn = make(chan interface{})
	subName := strings.Split(destination, "/")[0]
	c.subscriptions[subName] = b.ChIn
	b.Uid = c.Uid
	b.Subscribe = destination
	b.conn = c.client
	return b
}

func (c *ClientWs) wsConnect() error {
	c.changeURLPort()
	conn, err := websocket.Dial(c.config.BaseURL.String(), "", "http://localhost")
	if err != nil {
		return err
	}
	c.client = conn

	var msg []byte
	if err = websocket.Message.Receive(c.client, &msg); err != nil {
		return err
	}

	imsg, err := msgParser(msg)
	if err != nil {
		fmt.Println(err)
	}

	if err != nil {
		return err
	}
	c.Uid = imsg.Uid

	return nil
}

func (c *ClientWs) subsChannel(msg *Subscribe) error {
	if err := websocket.JSON.Send(c.client, sendJson{
		Uid:       msg.Uid,
		Subscribe: msg.Subscribe,
	}); err != nil {
		return err
	}

	var e error
	go func() {
		var resp []byte

		for {
			if err := websocket.Message.Receive(c.client, &resp); err == io.EOF {
				err = c.wsConnect()
				if err != nil {
					fmt.Println(err)
					return
				}
				if err = websocket.JSON.Send(c.client, sendJson{
					Uid:       msg.Uid,
					Subscribe: msg.Subscribe,
				}); err != nil {
					fmt.Println(err)
					return
				}
				continue
			} else if err != nil {
				e = errors.Wrap(err, "Error occurred while trying to receive message")
			}

			subName, _ := restParser(resp)
			e = c.buildType(subName, resp)
		}
	}()
	return e
}

func msgParser(msg []byte) (*Subscribe, error) {
	var message Subscribe
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

	var subscribe string
	for _, r := range raw {
		var obj map[string]interface{}
		err := json.Unmarshal(r, &obj)
		if err != nil {
			return "", err
		}

		if _, ok := obj["block"]; ok {
			subscribe = "block"
		} else if _, ok := obj["status"]; ok {
			subscribe = "status"
		} else if _, ok := obj["signer"]; ok {
			subscribe = "signer"
		} else if v, ok := obj["meta"]; ok {
			channelName := v.(map[string]interface{})
			subscribe = fmt.Sprintf("%v", channelName["channelName"])
		} else {
			subscribe = "none"
		}
	}
	return subscribe, nil
}

func (c *ClientWs) buildType(name string, t []byte) error {
	switch name {
	case "block":
		var b blockInfoDTO
		err := json.Unmarshal(t, &b)
		if err != nil {
			return err
		}
		data, err := b.toStruct()
		if err != nil {
			return err
		}
		c.subscriptions[name] <- data
		return nil

	case "status":
		var data StatusInfo
		err := json.Unmarshal(t, &data)
		if err != nil {
			return err
		}
		c.subscriptions[name] <- data
		return nil

	case "signer":
		var data SignerInfo
		err := json.Unmarshal(t, &data)
		if err != nil {
			return err
		}
		c.subscriptions["cosignature"] <- data
		return nil

	case "unconfirmedRemoved":
		var data SubscribeHash
		err := json.Unmarshal(t, &data)
		if err != nil {
			return err
		}
		c.subscriptions[name] <- data
		return nil

	case "partialRemoved":
		var data SubscribePartialRemoved
		err := json.Unmarshal(t, &data)
		if err != nil {
			return err
		}
		c.subscriptions[name] <- data
		return nil

	default:
		data, err := MapTransaction(bytes.NewBuffer([]byte(t)))
		if err != nil {
			return err
		}
		c.subscriptions[name] <- data
		return nil
	}
}
