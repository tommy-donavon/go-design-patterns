package interfaces

type IIterator interface {
	HasNext() bool
	GetNext() IComponent
}
