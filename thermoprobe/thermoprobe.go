package thermoprobe

type Thermoprobe interface {
	Read() (float32, error)
}
