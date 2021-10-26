package models

import (
	"triva/interfaces"
)

type Section struct {
	name     string
	children []interfaces.IComponent
}

func NewSection(name string) *Section {
	return &Section{name: name}
}

func (s *Section) AskQuestion() error {
	it, _ := s.GetIterator()
	for it.HasNext() {
		if item := it.GetNext(); item != nil {
			if err := item.AskQuestion(); err != nil {
				return err
			}
		}
	}
	return nil
}

func (s *Section) GetIterator() (interfaces.IIterator, error) {
	it, err := newCompositeIterator(s)
	if err != nil {
		return nil, err
	}
	return it, nil
}
func (s *Section) Add(c interfaces.IComponent) error {
	s.children = append(s.children, c)
	return nil
}
func (s *Section) Remove(c interfaces.IComponent) error {
	for i, v := range s.children {
		if v == c {
			s.children = append(s.children[:i], s.children[i+1:]...)
		}
	}
	return nil
}

func (s *Section) GetName() string {
	return s.name
}

func (s *Section) Size() int {
	return len(s.children)
}
