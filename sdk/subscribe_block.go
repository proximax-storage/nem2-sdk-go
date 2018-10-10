package sdk

import "golang.org/x/net/websocket"

// structure for Subscribe status
type SubscribeHash struct {
	Hash string `json:"hash"`
}

// structure for Subscribe PartialRemoved
type SubscribePartialRemoved struct {
	Meta SubscribeHash `json:"meta"`
}

type SubscribeService serviceWs

// const routers path for methods SubscribeService
const (
	pathBlock              = "block"
	pathConfirmedAdded     = "confirmedAdded"
	pathUnconfirmedAdded   = "unconfirmedAdded"
	pathUnconfirmedRemoved = "unconfirmedRemoved"
	pathStatus             = "status"
	pathPartialAdded       = "partialAdded"
	pathPartialRemoved     = "partialRemoved"
	pathCosignature        = "cosignature"
)

// Block notifies for every new block.
// The message contains the BlockInfo struct.
func (c *SubscribeService) Block() (*Subscribe, error) {
	return c.newChannel(pathBlock)
}

// ConfirmedAdded notifies when a transaction related to an
// address is included in a block.
// The message contains the transaction.
func (c *SubscribeService) ConfirmedAdded(add string) (*Subscribe, error) {
	return c.newChannel(pathConfirmedAdded + "/" + add)
}

// UnconfirmedAdded notifies when a transaction related to an
// address is in unconfirmed state and waiting to be included in a block.
// The message contains the transaction.
func (c *SubscribeService) UnconfirmedAdded(add string) (*Subscribe, error) {
	return c.newChannel(pathUnconfirmedAdded + "/" + add)
}

// UnconfirmedRemoved notifies when a transaction related to an
// address was in unconfirmed state but not anymore.
// The message contains the transaction hash.
func (c *SubscribeService) UnconfirmedRemoved(add string) (*Subscribe, error) {
	return c.newChannel(pathUnconfirmedRemoved + "/" + add)
}

// Status notifies when a transaction related to an address rises an error.
// The message contains the error message and the transaction hash.
func (c *SubscribeService) Status(add string) (*Subscribe, error) {
	return c.newChannel(pathStatus + "/" + add)
}

// PartialAdded notifies when an aggregate bonded transaction related to an
// address is in partial state and waiting to have all required cosigners.
// The message contains a transaction.
func (c *SubscribeService) PartialAdded(add string) (*Subscribe, error) {
	return c.newChannel(pathPartialAdded + "/" + add)
}

// PartialRemoved notifies when a transaction related to an
// address was in partial state but not anymore.
// The message contains the transaction hash.
func (c *SubscribeService) PartialRemoved(add string) (*Subscribe, error) {
	return c.newChannel(pathPartialRemoved + "/" + add)
}

// Cosignature notifies when a cosignature signed transaction related to an
// address is added to an aggregate bonded transaction with partial state.
// The message contains the cosignature signed transaction.
func (c *SubscribeService) Cosignature(add string) (*Subscribe, error) {
	return c.newChannel(pathCosignature + "/" + add)
}

// Unsubscribe terminates the specified subscription.
// It does not have any specific param.
func (c *Subscribe) Unsubscribe() error {
	if err := websocket.JSON.Send(c.conn, sendJson{
		Uid:       c.Uid,
		Subscribe: c.Subscribe,
	}); err != nil {
		return err
	}
	return nil
}

// Generate a new channel and subscribe to the websocket.
// param route A subscription channel route.
// return A pointer Subscribe struct or an error.
func (c *SubscribeService) newChannel(route string) (*Subscribe, error) {
	subMsg := c.client.buildSubscribe(route)

	err := c.client.subsChannel(subMsg)
	if err != nil {
		return nil, err
	}
	return subMsg, nil
}
