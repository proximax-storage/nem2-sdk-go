package sdk

import "golang.org/x/net/websocket"

type SubscribeHash struct {
	Hash string `json:"hash"`
}

type SubscribeService servicews

// const routers path for methods SubscribeService
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

// Block notifies for every new block.
// The message contains the BlockInfo struct.
func (c *SubscribeService) Block() (*Subscribe, error) {
	subMsg := c.client.buildSubscribe(pathBlock)
	err := c.client.subsChannel(subMsg)
	if err != nil {
		return nil, err
	}
	return subMsg, nil
}

// ConfirmedAdded notifies when a transaction related to an
// address is included in a block.
// The message contains the transaction.
func (c *SubscribeService) ConfirmedAdded(add string) (*Subscribe, error) {
	subMsg := c.client.buildSubscribe(pathconfirmedAdded + "/" + add)

	err := c.client.subsChannel(subMsg)
	if err != nil {
		return nil, err
	}
	return subMsg, nil
}

// UnconfirmedAdded notifies when a transaction related to an
// address is in unconfirmed state and waiting to be included in a block.
// The message contains the transaction.
func (c *SubscribeService) UnconfirmedAdded(add string) (*Subscribe, error) {
	subMsg := c.client.buildSubscribe(pathunconfirmedAdded + "/" + add)

	err := c.client.subsChannel(subMsg)
	if err != nil {
		return nil, err
	}
	return subMsg, nil
}

// UnconfirmedRemoved notifies when a transaction related to an
// address was in unconfirmed state but not anymore.
// The message contains the transaction hash.
func (c *SubscribeService) UnconfirmedRemoved(add string) (*Subscribe, error) {
	subMsg := c.client.buildSubscribe(pathunconfirmedRemoved + "/" + add)

	err := c.client.subsChannel(subMsg)
	if err != nil {
		return nil, err
	}
	return subMsg, nil
}

// Status notifies when a transaction related to an address rises an error.
// The message contains the error message and the transaction hash.
func (c *SubscribeService) Status(add string) (*Subscribe, error) {
	subMsg := c.client.buildSubscribe(pathstatus + "/" + add)

	err := c.client.subsChannel(subMsg)
	if err != nil {
		return nil, err
	}
	return subMsg, nil
}

// PartialAdded notifies when an aggregate bonded transaction related to an
// address is in partial state and waiting to have all required cosigners.
// The message contains a transaction.
func (c *SubscribeService) PartialAdded(add string) (*Subscribe, error) {
	subMsg := c.client.buildSubscribe(pathpartialAdded + "/" + add)

	err := c.client.subsChannel(subMsg)
	if err != nil {
		return nil, err
	}
	return subMsg, nil
}

// PartialRemoved notifies when a transaction related to an
// address was in partial state but not anymore.
// The message contains the transaction hash.
func (c *SubscribeService) PartialRemoved(add string) (*Subscribe, error) {
	subMsg := c.client.buildSubscribe(pathpartialRemoved + "/" + add)

	err := c.client.subsChannel(subMsg)
	if err != nil {
		return nil, err
	}
	return subMsg, nil
}

// Cosignature notifies when a cosignature signed transaction related to an
// address is added to an aggregate bonded transaction with partial state.
// The message contains the cosignature signed transaction.
func (c *SubscribeService) Cosignature(add string) (*Subscribe, error) {
	subMsg := c.client.buildSubscribe(pathcosignature + "/" + add)

	err := c.client.subsChannel(subMsg)
	if err != nil {
		return nil, err
	}
	return subMsg, nil
}

// Unsubscribe terminates the specified subscription.
// It does not have any specific param.
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
