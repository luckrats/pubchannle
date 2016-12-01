package pubchannle

func NewPublishChannle() *PublishChannle {
	pLink := &PublishChannle{pTail: NewLinkElement()}
	return pLink
}

type PublishChannle struct {
	pTail *linkElement
}

func (this *PublishChannle) Write(data interface{}) {
	this.pTail = this.pTail.Add(data)
}

func (this *PublishChannle) Close() {
	this.pTail.Close()
}

func (this *PublishChannle) NewSubscribChannle() *SubscribChannle {
	return &SubscribChannle{pTail: this.pTail}

}

type SubscribChannle struct {
	pTail *linkElement
}

func (this *SubscribChannle) Read() (data interface{}, hasMore bool) {
	if this.pTail != nil {
		if !this.pTail.Empty() {
			data = this.pTail.data
			this.pTail = this.pTail.Next()
			return data, true
		}
	}
	return nil, false
}

func (this *SubscribChannle) WaitNotify() <-chan struct{} {
	if this.pTail != nil {
		return this.pTail.arrived
	}
	retv := make(chan struct{})
	close(retv)
	return retv
}
