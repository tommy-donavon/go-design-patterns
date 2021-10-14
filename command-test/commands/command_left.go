package commands

import "command-ui/interfaces"

type commandLeft struct {
	receiver interfaces.IReceiver
}

func NewCommandLeft(receiver interfaces.IReceiver) *commandLeft {
	return &commandLeft{receiver}
}

func (cu *commandLeft) Execute() {
	cu.receiver.Left()
}

func (cu *commandLeft) Undo() {
	cu.receiver.Right()
}
