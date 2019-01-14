// Copyright 2018 ProximaX Limited. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package sdk

import (
	"errors"
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
func (s *subscribe) closeChannel() error {
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

	case chan Transaction:
		chType := s.Ch.(chan Transaction)
		if s.getSubscribe() == "partialAdded" {
			delete(partialAddedChannels, s.getAdd())
		} else if s.getSubscribe() == "unconfirmedAdded" {
			delete(unconfirmedAddedChannels, s.getAdd())
		} else {
			delete(confirmedAddedChannels, s.getAdd())
		}
		close(chType)

	default:
		return errors.New("WRONG TYPE CHANNEL")
	}
	return nil
}

// Unsubscribe terminates the specified subscription.
// It does not have any specific param.
func (c *subscribe) unsubscribe() error {
	c.conn = connectsWs[c.getAdd()].conn
	if err := websocket.JSON.Send(c.conn, sendJson{
		Uid:       c.Uid,
		Subscribe: c.Subscribe,
	}); err != nil {
		return err
	}

	if err := c.closeChannel(); err != nil {
		return err
	}

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

func (c *SubscribeService) getClient(add string) (*ClientWebsocket, error) {
	if len(connectsWs) == 0 {
		obj := uidConn{
			uid:  c.client.Uid,
			conn: c.client.client,
		}
		connectsWs[add] = &obj
		return c.client, nil
	} else if obj, exist := connectsWs[add]; exist {
		c.client.client = obj.conn
		c.client.Uid = obj.uid
		return c.client, nil
	} else {
		client, err := NewConnectWs(c.client.config.BaseURL.String(), *c.client.duration)

		if err != nil {
			return nil, err
		}
		obj := uidConn{
			uid:  client.Uid,
			conn: client.client,
		}
		connectsWs[add] = &obj
		return client, nil
	}
}

// Block notifies for every new block.
// The message contains the BlockInfo struct.
func (c *SubscribeService) Block() (*SubscribeBlock, error) {
	subBlock := new(SubscribeBlock)
	Block = subBlock
	subBlock.Ch = make(chan *BlockInfo)
	subscribe, err := c.newSubscribe(pathBlock)
	if err != nil {
		return nil, err
	}
	subBlock.subscribe = subscribe
	subscribe.Ch = subBlock.Ch
	return subBlock, nil
}

// ConfirmedAdded notifies when a transaction related to an
// address is included in a block.
// The message contains the transaction.
func (c *SubscribeService) ConfirmedAdded(add *Address) (*SubscribeTransaction, error) {
	if client, err := c.getClient(add.Address); err != nil {
		return nil, err
	} else {
		c.client = client
	}
	subTransaction := new(SubscribeTransaction)
	subTransaction.Ch = make(chan Transaction)
	confirmedAddedChannels[add.Address] = subTransaction.Ch
	subscribe, err := c.newSubscribe(pathConfirmedAdded + "/" + add.Address)
	if err != nil {
		return nil, err
	}
	subTransaction.subscribe = subscribe
	subscribe.Ch = subTransaction.Ch
	//fmt.Printf("Address %v Uid %v Socketc [%v] \n", add.Address, c.client.Uid, c.client.client)
	return subTransaction, nil
}

// UnconfirmedAdded notifies when a transaction related to an
// address is in unconfirmed state and waiting to be included in a block.
// The message contains the transaction.
func (c *SubscribeService) UnconfirmedAdded(add *Address) (*SubscribeTransaction, error) {
	if client, err := c.getClient(add.Address); err != nil {
		return nil, err
	} else {
		c.client = client
	}
	subTransaction := new(SubscribeTransaction)
	subTransaction.Ch = make(chan Transaction)
	unconfirmedAddedChannels[add.Address] = subTransaction.Ch
	subscribe, err := c.newSubscribe(pathUnconfirmedAdded + "/" + add.Address)
	if err != nil {
		return nil, err
	}
	subTransaction.subscribe = subscribe
	subscribe.Ch = unconfirmedAddedChannels[add.Address]
	//fmt.Printf("Address %v Uid %v Socketc [%v] \n", add.Address, c.client.Uid, c.client.client)
	return subTransaction, nil
}

// UnconfirmedRemoved notifies when a transaction related to an
// address was in unconfirmed state but not anymore.
// The message contains the transaction hash.
func (c *SubscribeService) UnconfirmedRemoved(add *Address) (*SubscribeHash, error) {
	if client, err := c.getClient(add.Address); err != nil {
		return nil, err
	} else {
		c.client = client
	}
	subHash := new(SubscribeHash)
	subHash.Ch = make(chan *HashInfo)
	unconfirmedRemovedChannels[add.Address] = subHash.Ch
	subscribe, err := c.newSubscribe(pathUnconfirmedRemoved + "/" + add.Address)
	if err != nil {
		return nil, err
	}
	subHash.subscribe = subscribe
	subscribe.Ch = unconfirmedRemovedChannels[add.Address]
	return subHash, nil
}

// Status notifies when a transaction related to an address rises an error.
// The message contains the error message and the transaction hash.
func (c *SubscribeService) Status(add *Address) (*SubscribeStatus, error) {
	if client, err := c.getClient(add.Address); err != nil {
		return nil, err
	} else {
		c.client = client
	}
	subStatus := new(SubscribeStatus)
	subStatus.Ch = make(chan *StatusInfo)
	statusInfoChannels[add.Address] = subStatus.Ch
	subscribe, err := c.newSubscribe(pathStatus + "/" + add.Address)
	if err != nil {
		return nil, err
	}
	subStatus.subscribe = subscribe
	subscribe.Ch = statusInfoChannels[add.Address]
	return subStatus, nil
}

// PartialAdded notifies when an aggregate bonded transaction related to an
// address is in partial state and waiting to have all required cosigners.
// The message contains a transaction.
func (c *SubscribeService) PartialAdded(add *Address) (*SubscribeTransaction, error) {
	if client, err := c.getClient(add.Address); err != nil {
		return nil, err
	} else {
		c.client = client
	}
	subTransaction := new(SubscribeTransaction)
	subTransaction.Ch = make(chan Transaction)
	partialAddedChannels[add.Address] = subTransaction.Ch
	subscribe, err := c.newSubscribe(pathPartialAdded + "/" + add.Address)
	if err != nil {
		return nil, err
	}
	subTransaction.subscribe = subscribe
	subscribe.Ch = partialAddedChannels[add.Address]
	return subTransaction, nil
}

// PartialRemoved notifies when a transaction related to an
// address was in partial state but not anymore.
// The message contains the transaction hash.
func (c *SubscribeService) PartialRemoved(add *Address) (*SubscribePartialRemoved, error) {
	if client, err := c.getClient(add.Address); err != nil {
		return nil, err
	} else {
		c.client = client
	}
	subPartialRemoved := new(SubscribePartialRemoved)
	subPartialRemoved.Ch = make(chan *PartialRemovedInfo)
	partialRemovedInfoChannels[add.Address] = subPartialRemoved.Ch
	subscribe, err := c.newSubscribe(pathPartialRemoved + "/" + add.Address)
	if err != nil {
		return nil, err
	}
	subPartialRemoved.subscribe = subscribe
	subscribe.Ch = partialRemovedInfoChannels[add.Address]
	return subPartialRemoved, nil
}

// Cosignature notifies when a cosignature signed transaction related to an
// address is added to an aggregate bonded transaction with partial state.
// The message contains the cosignature signed transaction.
func (c *SubscribeService) Cosignature(add *Address) (*SubscribeSigner, error) {
	if client, err := c.getClient(add.Address); err != nil {
		return nil, err
	} else {
		c.client = client
	}
	subCosignature := new(SubscribeSigner)
	subCosignature.Ch = make(chan *SignerInfo)
	signerInfoChannels[add.Address] = subCosignature.Ch
	subscribe, err := c.newSubscribe(pathCosignature + "/" + add.Address)
	if err != nil {
		return nil, err
	}
	subCosignature.subscribe = subscribe
	subscribe.Ch = signerInfoChannels[add.Address]
	return subCosignature, nil
}

func (c *SubscribeService) Error(add *Address) (*SubscribeError, error) {
	address := "block"
	if add != nil {
		address = add.Address
	}
	if client, err := c.getClient(address); err != nil {
		return nil, err
	} else {
		c.client = client
	}
	subError := new(SubscribeError)
	subError.Ch = make(chan *ErrorInfo)
	errChannels[address] = subError.Ch
	subscribe := new(subscribe)
	subscribe.Subscribe = "error/" + address
	subError.subscribe = subscribe
	subscribe.Ch = errChannels[address]
	return subError, nil
}
