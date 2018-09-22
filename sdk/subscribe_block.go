package sdk

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

func (c *SubscribeService) Block() (chan []byte, error) {
	ch := make(chan []byte)
	subMsg, err := c.client.BuildSubscribe(pathBlock)
	if err != nil {
		return nil, err
	}
	ch, err = c.client.Subs(subMsg, ch)
	if err != nil {
		return nil, err
	}
	return ch, nil
}

func (c *SubscribeService) ConfirmedAdded(add string) (chan []byte, error) {
	ch := make(chan []byte)
	subMsg, err := c.client.BuildSubscribe(pathconfirmedAdded + "/" + add)
	if err != nil {
		return nil, err
	}
	ch, err = c.client.Subs(subMsg, ch)
	if err != nil {
		return nil, err
	}
	return ch, nil
}

func (c *SubscribeService) UnConfirmedAdded(add string) (chan []byte, error) {
	ch := make(chan []byte)
	subMsg, err := c.client.BuildSubscribe(pathunconfirmedAdded + "/" + add)
	if err != nil {
		return nil, err
	}
	ch, err = c.client.Subs(subMsg, ch)
	if err != nil {
		return nil, err
	}
	return ch, nil
}

func (c *SubscribeService) UnconfirmedRemoved(add string) (chan []byte, error) {
	ch := make(chan []byte)
	subMsg, err := c.client.BuildSubscribe(pathunconfirmedRemoved + "/" + add)
	if err != nil {
		return nil, err
	}
	ch, err = c.client.Subs(subMsg, ch)
	if err != nil {
		return nil, err
	}
	return ch, nil
}

func (c *SubscribeService) Status(add string) (chan []byte, error) {
	ch := make(chan []byte)
	subMsg, err := c.client.BuildSubscribe(pathstatus + "/" + add)
	if err != nil {
		return nil, err
	}
	ch, err = c.client.Subs(subMsg, ch)
	if err != nil {
		return nil, err
	}
	return ch, nil
}

func (c *SubscribeService) PartialAdded(add string) (chan []byte, error) {
	ch := make(chan []byte)
	subMsg, err := c.client.BuildSubscribe(pathpartialAdded + "/" + add)
	if err != nil {
		return nil, err
	}
	ch, err = c.client.Subs(subMsg, ch)
	if err != nil {
		return nil, err
	}
	return ch, nil
}

func (c *SubscribeService) PartialRemoved(add string) (chan []byte, error) {
	ch := make(chan []byte)
	subMsg, err := c.client.BuildSubscribe(pathpartialRemoved + "/" + add)
	if err != nil {
		return nil, err
	}
	ch, err = c.client.Subs(subMsg, ch)
	if err != nil {
		return nil, err
	}
	return ch, nil
}

func (c *SubscribeService) Cosignature(add string) (chan []byte, error) {
	ch := make(chan []byte)
	subMsg, err := c.client.BuildSubscribe(pathcosignature + "/" + add)
	if err != nil {
		return nil, err
	}
	ch, err = c.client.Subs(subMsg, ch)
	if err != nil {
		return nil, err
	}
	return ch, nil
}

//func (c *SubscribeService) Unsubscribe(add string) (chan []byte, error) {
//	ch := make(chan []byte)
//	subMsg, err := c.client.BuildSubscribe(pathcosignature + "/" + add)
//	if err != nil {
//		return nil, err
//	}
//	ch, err = c.client.Subs(subMsg, ch)
//	if err != nil {
//		return nil, err
//	}
//	return ch, nil
//}
