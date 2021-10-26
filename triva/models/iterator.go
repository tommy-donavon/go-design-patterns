package models

import (
	"triva/interfaces"
)

type leafIterator struct {
}

type compositeIterator struct {
	composite       *Section
	index           int
	currentIterator interfaces.IIterator
}

func newCompositeIterator(composite *Section) (interfaces.IIterator, error) {
	it, _ := composite.children[0].GetIterator()
	return &compositeIterator{composite: composite, index: 0, currentIterator: it}, nil
}

func (c *leafIterator) HasNext() bool {
	return false
}

func (c *leafIterator) GetNext() interfaces.IComponent {
	return nil
}
func (c *compositeIterator) HasNext() bool {
	return c.index < len(c.composite.children)
}
func (c *compositeIterator) GetNext() interfaces.IComponent {
	if c.currentIterator.HasNext() {
		leaf := c.currentIterator.GetNext()
		if !c.currentIterator.HasNext() {
			c.index++
			if c.HasNext() {
				c.currentIterator, _ = c.composite.children[c.index].GetIterator()
			}
		}
		return leaf

	} else {
		leaf := c.composite.children[c.index]
		c.index++
		if c.HasNext() {
			c.currentIterator, _ = c.composite.children[c.index].GetIterator()
		}
		return leaf
	}
}
