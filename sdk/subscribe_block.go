package sdk

import "golang.org/x/net/websocket"

type SubscribeService servicews

const (
	pathBlock              = "block"
	pathconfirmedAdded     = "confirmedAdded"
	pathunconfirmedAdded   = "unconfirmedAdded"
	pathunconfirmedRemoved = "unconfirmedRemoved"
	pathstatus             = "status"
	pathpartialAdded       = "partialAdded"
	pathpartialRemoved     = "partialRemoved"
	pathcosignature        = "cosignature"
)

func (c *SubscribeService) Block() (*Subscribe, error) {
	subMsg := c.client.buildSubscribe(pathBlock)
	err := c.client.subsChannel(subMsg)
	if err != nil {
		return nil, err
	}
	return subMsg, nil
}

func (c *SubscribeService) ConfirmedAdded(add string) (*Subscribe, error) {
	subMsg := c.client.buildSubscribe(pathconfirmedAdded + "/" + add)

	err := c.client.subsChannel(subMsg)
	if err != nil {
		return nil, err
	}
	return subMsg, nil
}

func (c *SubscribeService) UnconfirmedAdded(add string) (*Subscribe, error) {
	subMsg := c.client.buildSubscribe(pathunconfirmedAdded + "/" + add)

	err := c.client.subsChannel(subMsg)
	if err != nil {
		return nil, err
	}
	return subMsg, nil
}

func (c *SubscribeService) UnconfirmedRemoved(add string) (*Subscribe, error) {
	subMsg := c.client.buildSubscribe(pathunconfirmedRemoved + "/" + add)

	err := c.client.subsChannel(subMsg)
	if err != nil {
		return nil, err
	}
	return subMsg, nil
}

func (c *SubscribeService) Status(add string) (*Subscribe, error) {
	subMsg := c.client.buildSubscribe(pathstatus + "/" + add)

	err := c.client.subsChannel(subMsg)
	if err != nil {
		return nil, err
	}
	return subMsg, nil
}

func (c *SubscribeService) PartialAdded(add string) (*Subscribe, error) {
	subMsg := c.client.buildSubscribe(pathpartialAdded + "/" + add)

	err := c.client.subsChannel(subMsg)
	if err != nil {
		return nil, err
	}
	return subMsg, nil
}

func (c *SubscribeService) PartialRemoved(add string) (*Subscribe, error) {
	subMsg := c.client.buildSubscribe(pathpartialRemoved + "/" + add)

	err := c.client.subsChannel(subMsg)
	if err != nil {
		return nil, err
	}
	return subMsg, nil
}

func (c *SubscribeService) Cosignature(add string) (*Subscribe, error) {
	subMsg := c.client.buildSubscribe(pathcosignature + "/" + add)

	err := c.client.subsChannel(subMsg)
	if err != nil {
		return nil, err
	}
	return subMsg, nil
}

func (c *Subscribe) Unsubscribe() error {
	if err := websocket.JSON.Send(c.conn, struct {
		UID         string `json:"uid"`
		Unsubscribe string `json:"unsubscribe"`
	}{
		UID:         c.UID,
		Unsubscribe: c.Subscribe,
	}); err != nil {
		return err
	}
	return nil
}
