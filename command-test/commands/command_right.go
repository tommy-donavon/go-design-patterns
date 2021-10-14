package commands

import "command-ui/interfaces"

type commandRight struct {
	receiver interfaces.IReceiver
}

func NewCommandRight(receiver interfaces.IReceiver) *commandRight {
	return &commandRight{receiver}
}

func (cu *commandRight) Execute() {
	cu.receiver.Right()
}

func (cu *commandRight) Undo() {
	cu.receiver.Left()
}
