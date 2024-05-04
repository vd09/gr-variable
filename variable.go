package gr_variable

import "github.com/vd09/gr-variable/chan_variable"

func NewGrChannel[T any]() GrChannel[T] {
	return chan_variable.NewCharVar[T]()
}

func NewGrChannelWithLength[T any](length int) GrChannel[T] {
	return chan_variable.NewCharVarWithLength[T](length)
}
