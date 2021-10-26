package interfaces

type IComponent interface {
	AskQuestion() error
	GetIterator() (IIterator, error)
	Add(IComponent) error
	Remove(IComponent) error
}
