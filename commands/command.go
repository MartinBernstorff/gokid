package commands

import (
	"fmt"
)

type Command struct {
	Assumptions []NamedCallable
	Action      NamedCallable
	Revert      NamedCallable
}

type NamedCallable struct {
	Name     string
	Callable func() error
}

func NewFailCommand() Command {
	return Command{
		Assumptions: []NamedCallable{},
		Action: NamedCallable{
			Name: "Fail",
			Callable: func() error {
				return fmt.Errorf("fail")
			},
		},
		Revert: NamedCallable{},
	}
}
