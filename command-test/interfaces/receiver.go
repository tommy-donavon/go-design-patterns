package interfaces

type IReceiver interface {
	Up()
	Down()
	Left()
	Right()
	NextColor()
	PreviousColor()
}
