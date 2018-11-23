// Copyright 2018 ProximaX Limited. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package sdk

import (
	"golang.org/x/net/websocket"
	"strings"
)

var ChanSubscribe struct {
	Block              *SubscribeBlock
	ConfirmedAdded     *SubscribeTransaction
	UnconfirmedAdded   *SubscribeTransaction
	UnconfirmedRemoved *SubscribeHash
	Status             *SubscribeStatus
	PartialAdded       *SubscribeTransaction
	PartialRemoved     *SubscribePartialRemoved
	Cosignature        *SubscribeSigner
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
func (c *SubscribeService) Block() (*SubscribeBlock, error) {
	subBlock := new(SubscribeBlock)
	ChanSubscribe.Block = subBlock
	subBlock.Ch = make(chan *BlockInfo)
	subscribe, err := c.newSubscribe(pathBlock)
	subBlock.subscribe = subscribe
	//subscribe.Ch = subBlock.Ch
	return subBlock, err
}

// ConfirmedAdded notifies when a transaction related to an
// address is included in a block.
// The message contains the transaction.
func (c *SubscribeService) ConfirmedAdded(add string) (*SubscribeTransaction, error) {
	subTransaction := new(SubscribeTransaction)
	ChanSubscribe.ConfirmedAdded = subTransaction
	subTransaction.Ch = make(chan Transaction)
	subscribe, err := c.newSubscribe(pathConfirmedAdded + "/" + add)
	subTransaction.subscribe = subscribe
	//subscribe.Ch = subTransaction.Ch
	return subTransaction, err
}

// UnconfirmedAdded notifies when a transaction related to an
// address is in unconfirmed state and waiting to be included in a block.
// The message contains the transaction.
func (c *SubscribeService) UnconfirmedAdded(add string) (*SubscribeTransaction, error) {
	subTransaction := new(SubscribeTransaction)
	ChanSubscribe.UnconfirmedAdded = subTransaction
	subTransaction.Ch = make(chan Transaction)
	subscribe, err := c.newSubscribe(pathUnconfirmedAdded + "/" + add)
	subTransaction.subscribe = subscribe
	return subTransaction, err
}

// UnconfirmedRemoved notifies when a transaction related to an
// address was in unconfirmed state but not anymore.
// The message contains the transaction hash.
func (c *SubscribeService) UnconfirmedRemoved(add string) (*SubscribeHash, error) {
	subHash := new(SubscribeHash)
	ChanSubscribe.UnconfirmedRemoved = subHash
	subHash.Ch = make(chan *HashInfo)
	subscribe, err := c.newSubscribe(pathUnconfirmedRemoved + "/" + add)
	subHash.subscribe = subscribe
	return subHash, err
}

// Status notifies when a transaction related to an address rises an error.
// The message contains the error message and the transaction hash.
func (c *SubscribeService) Status(add string) (*SubscribeStatus, error) {
	subStatus := new(SubscribeStatus)
	ChanSubscribe.Status = subStatus
	subStatus.Ch = make(chan *StatusInfo)
	subscribe, err := c.newSubscribe(pathStatus + "/" + add)
	subStatus.subscribe = subscribe
	subStatus.subscribe.Ch = subStatus.Ch
	return subStatus, err
}

// PartialAdded notifies when an aggregate bonded transaction related to an
// address is in partial state and waiting to have all required cosigners.
// The message contains a transaction.
func (c *SubscribeService) PartialAdded(add string) (*SubscribeTransaction, error) {
	subTransaction := new(SubscribeTransaction)
	ChanSubscribe.PartialAdded = subTransaction
	subTransaction.Ch = make(chan Transaction)
	subscribe, err := c.newSubscribe(pathPartialAdded + "/" + add)
	subTransaction.subscribe = subscribe
	return subTransaction, err
}

// PartialRemoved notifies when a transaction related to an
// address was in partial state but not anymore.
// The message contains the transaction hash.
func (c *SubscribeService) PartialRemoved(add string) (*SubscribePartialRemoved, error) {
	subPartialRemoved := new(SubscribePartialRemoved)
	ChanSubscribe.PartialRemoved = subPartialRemoved
	subPartialRemoved.Ch = make(chan *PartialRemovedInfo)
	subscribe, err := c.newSubscribe(pathPartialRemoved + "/" + add)
	subPartialRemoved.subscribe = subscribe
	return ChanSubscribe.PartialRemoved, err
}

// Cosignature notifies when a cosignature signed transaction related to an
// address is added to an aggregate bonded transaction with partial state.
// The message contains the cosignature signed transaction.
func (c *SubscribeService) Cosignature(add string) (*SubscribeSigner, error) {
	subCosignature := new(SubscribeSigner)
	ChanSubscribe.Cosignature = subCosignature
	subCosignature.Ch = make(chan *SignerInfo)
	subscribe, err := c.newSubscribe(pathCosignature + "/" + add)
	subCosignature.subscribe = subscribe
	return ChanSubscribe.Cosignature, err
}

// Unsubscribe terminates the specified subscription.
// It does not have any specific param.
func (c *subscribe) unsubscribe() error {
	if err := websocket.JSON.Send(c.conn, sendJson{
		Uid:       c.Uid,
		Subscribe: c.Subscribe,
	}); err != nil {
		return err
	}

	if strings.Split(c.Subscribe, "/")[0] == "status" {
		chType := c.Ch.(chan *StatusInfo)
		close(chType)
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
