package pubchannle

func NewLinkElement() *linkElement {
	return &linkElement{arrived: make(chan struct{}), data: nil, pNext: nil}
}

type linkElement struct {
	arrived chan struct{}
	data    interface{}
	pNext   *linkElement
}

func (this *linkElement) Add(data interface{}) *linkElement {
	this.data = data
	this.pNext = &linkElement{arrived: make(chan struct{})}
	close(this.arrived)
	return this.pNext
}

func (this *linkElement) Next() *linkElement {
	if this.pNext != nil {
		return this.pNext
	}
	return this
}

func (this *linkElement) Close() {
	this.data = this.arrived
	close(this.arrived)
}

func (this *linkElement) Empty() bool {
	return this.data == this.arrived
}
