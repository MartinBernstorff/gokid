package versioncontrol

import "gokid/shell"

type Stasher interface {
	Save() error
	Pop() error
}

// Stash handles git stash operations
type Stash struct {
	shell shell.Shell
}

func NewStash(s shell.Shell) *Stash {
	return &Stash{
		shell: s,
	}
}

func (s *Stash) Save() error {
	_, err := s.shell.Run("git stash")
	return err
}

func (s *Stash) Pop() error {
	_, err := s.shell.Run("git stash pop")
	return err
}
