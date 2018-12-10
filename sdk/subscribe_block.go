// Copyright 2018 ProximaX Limited. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package sdk

import (
	"fmt"
	"golang.org/x/net/websocket"
)

var Block *SubscribeBlock

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

// Closes the subscription channel.
func closeChannel(s *subscribe) {
	switch s.Ch.(type) {
	case chan *BlockInfo:
		chType := s.Ch.(chan *BlockInfo)
		close(chType)

	case chan *StatusInfo:
		chType := s.Ch.(chan *StatusInfo)
		delete(statusInfoChannels, s.getAdd())
		close(chType)

	case chan *HashInfo:
		chType := s.Ch.(chan *HashInfo)
		delete(unconfirmedRemovedChannels, s.getAdd())
		close(chType)

	case chan *PartialRemovedInfo:
		chType := s.Ch.(chan *PartialRemovedInfo)
		delete(partialRemovedInfoChannels, s.getAdd())
		close(chType)

	case chan *SignerInfo:
		chType := s.Ch.(chan *SignerInfo)
		delete(signerInfoChannels, s.getAdd())
		close(chType)

	case chan *ErrorInfo:
		chType := s.Ch.(chan *ErrorInfo)
		delete(errChannels, s.getAdd())
		close(chType)

	default:
		chType := s.Ch.(chan Transaction)
		if s.getSubscribe() == "partialAdded" {
			delete(partialAddedChannels, s.getAdd())
		} else if s.getSubscribe() == "unconfirmedAdded" {
			delete(unconfirmedAddedChannels, s.getAdd())
		} else {
			delete(confirmedAddedChannels, s.getAdd())
		}
		close(chType)
	}
}

// Unsubscribe terminates the specified subscription.
// It does not have any specific param.
func (c *subscribe) unsubscribe() error {
	c.conn = connectsWs[c.getAdd()]
	if err := websocket.JSON.Send(c.conn, sendJson{
		Uid:       c.Uid,
		Subscribe: c.Subscribe,
	}); err != nil {
		return err
	}

	closeChannel(c)

	return nil
}

// Generate a new channel and subscribe to the websocket.
// param route A subscription channel route.
// return A pointer Subscribe struct or an error.
func (c *SubscribeService) newSubscribe(route string) (*subscribe, error) {
	subMsg := c.client.buildSubscribe(route)

	err := c.client.subsChannel(subMsg)
	if err != nil {
		return nil, err
	}
	return subMsg, nil
}

func (c *SubscribeService) getClient(add string) *ClientWebsocket {
	if len(connectsWs) == 0 {
		connectsWs[add] = c.client.client
		return c.client
	} else if _, exist := connectsWs[add]; exist {
		return c.client
	} else {
		client, err := NewConnectWs(c.client.config.BaseURL.String(), *c.client.duration)
		if err != nil {
			fmt.Println(err)
		}
		connectsWs[add] = client.client
		return client
	}
}

// Block notifies for every new block.
// The message contains the BlockInfo struct.
func (c *SubscribeService) Block() (*SubscribeBlock, error) {
	subBlock := new(SubscribeBlock)
	Block = subBlock
	subBlock.Ch = make(chan *BlockInfo)
	subscribe, err := c.newSubscribe(pathBlock)
	subBlock.subscribe = subscribe
	subscribe.Ch = subBlock.Ch
	return subBlock, err
}

// ConfirmedAdded notifies when a transaction related to an
// address is included in a block.
// The message contains the transaction.
func (c *SubscribeService) ConfirmedAdded(add string) (*SubscribeTransaction, error) {
	c.client = c.getClient(add)
	subTransaction := new(SubscribeTransaction)
	subTransaction.Ch = make(chan Transaction)
	confirmedAddedChannels[add] = subTransaction.Ch
	subscribe, err := c.newSubscribe(pathConfirmedAdded + "/" + add)
	subTransaction.subscribe = subscribe
	subscribe.Ch = subTransaction.Ch
	return subTransaction, err
}

// UnconfirmedAdded notifies when a transaction related to an
// address is in unconfirmed state and waiting to be included in a block.
// The message contains the transaction.
func (c *SubscribeService) UnconfirmedAdded(add string) (*SubscribeTransaction, error) {
	c.client = c.getClient(add)
	subTransaction := new(SubscribeTransaction)
	subTransaction.Ch = make(chan Transaction)
	unconfirmedAddedChannels[add] = subTransaction.Ch
	subscribe, err := c.newSubscribe(pathUnconfirmedAdded + "/" + add)
	subTransaction.subscribe = subscribe
	subscribe.Ch = unconfirmedAddedChannels[add]
	return subTransaction, err
}

// UnconfirmedRemoved notifies when a transaction related to an
// address was in unconfirmed state but not anymore.
// The message contains the transaction hash.
func (c *SubscribeService) UnconfirmedRemoved(add string) (*SubscribeHash, error) {
	c.client = c.getClient(add)
	subHash := new(SubscribeHash)
	subHash.Ch = make(chan *HashInfo)
	unconfirmedRemovedChannels[add] = subHash.Ch
	subscribe, err := c.newSubscribe(pathUnconfirmedRemoved + "/" + add)
	subHash.subscribe = subscribe
	subscribe.Ch = unconfirmedRemovedChannels[add]
	return subHash, err
}

// Status notifies when a transaction related to an address rises an error.
// The message contains the error message and the transaction hash.
func (c *SubscribeService) Status(add string) (*SubscribeStatus, error) {
	c.client = c.getClient(add)
	subStatus := new(SubscribeStatus)
	subStatus.Ch = make(chan *StatusInfo)
	statusInfoChannels[add] = subStatus.Ch
	subscribe, err := c.newSubscribe(pathStatus + "/" + add)
	subStatus.subscribe = subscribe
	subscribe.Ch = statusInfoChannels[add]
	return subStatus, err
}

// PartialAdded notifies when an aggregate bonded transaction related to an
// address is in partial state and waiting to have all required cosigners.
// The message contains a transaction.
func (c *SubscribeService) PartialAdded(add string) (*SubscribeTransaction, error) {
	c.client = c.getClient(add)
	subTransaction := new(SubscribeTransaction)
	subTransaction.Ch = make(chan Transaction)
	partialAddedChannels[add] = subTransaction.Ch
	subscribe, err := c.newSubscribe(pathPartialAdded + "/" + add)
	subTransaction.subscribe = subscribe
	subscribe.Ch = partialAddedChannels[add]
	return subTransaction, err
}

// PartialRemoved notifies when a transaction related to an
// address was in partial state but not anymore.
// The message contains the transaction hash.
func (c *SubscribeService) PartialRemoved(add string) (*SubscribePartialRemoved, error) {
	c.client = c.getClient(add)
	subPartialRemoved := new(SubscribePartialRemoved)
	subPartialRemoved.Ch = make(chan *PartialRemovedInfo)
	partialRemovedInfoChannels[add] = subPartialRemoved.Ch
	subscribe, err := c.newSubscribe(pathPartialRemoved + "/" + add)
	subPartialRemoved.subscribe = subscribe
	subscribe.Ch = partialRemovedInfoChannels[add]
	return subPartialRemoved, err
}

// Cosignature notifies when a cosignature signed transaction related to an
// address is added to an aggregate bonded transaction with partial state.
// The message contains the cosignature signed transaction.
func (c *SubscribeService) Cosignature(add string) (*SubscribeSigner, error) {
	c.client = c.getClient(add)
	subCosignature := new(SubscribeSigner)
	subCosignature.Ch = make(chan *SignerInfo)
	signerInfoChannels[add] = subCosignature.Ch
	subscribe, err := c.newSubscribe(pathCosignature + "/" + add)
	subCosignature.subscribe = subscribe
	subscribe.Ch = signerInfoChannels[add]
	return subCosignature, err
}

func (c *SubscribeService) Error(add string) *SubscribeError {
	c.client = c.getClient(add)
	subError := new(SubscribeError)
	subError.Ch = make(chan *ErrorInfo)
	errChannels[add] = subError.Ch
	subscribe := new(subscribe)
	subscribe.Subscribe = "error/" + add
	subError.subscribe = subscribe
	subscribe.Ch = errChannels[add]
	return subError
}
