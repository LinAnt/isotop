package relay

type relay struct {
	pin  int
	open bool
}

func NewRelay(pin int) relay {
	r := relay{}
	r.pin = pin
	r.open = false
	return r
}

func (r *relay) Open() {
	r.open = true
}
func (r *relay) Close() {
	r.open = false
}

func (r *relay) IsOpen() bool {
	return r.open
}
