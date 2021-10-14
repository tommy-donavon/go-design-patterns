package interfaces

type ICommand interface {
	Execute()
	Undo()
}
