package tgbot

import (
	"fmt"

	"github.com/olebedev/go-tgbot/models"
)

type Context struct {
	Update  *models.Update
	Keys    map[string]interface{}
	Params  []string
	Session fmt.Stringer

	handlers  []func(*Context) error
	index     uint8
	isSkipped bool
}

// Skip shows that at least next matched handlers should be applied.
func (c *Context) FallThrough() {
	c.isSkipped = true
				
}

// Next executes the pending handlers in the chain inside the calling handler.
func (c *Context) Next() error {
	c.index++
	s := uint8(len(c.handlers))
	var err error
	for ; c.index <= s; c.index++ {
		if c.isSkipped {
			break
		} else {
			err = c.handlers[c.index-1](c)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// Abort stops chain calling.
func (c *Context) Abort() {
	c.index = uint8(len(c.handlers) + 1)
}

// Mostly copied from https://github.com/gin-gonic/gin/blob/develop/context.go

// Set is used to store a new key/value pair exclusivelly for this context.
// It also lazy initializes  c.Keys if it was not used previously.
func (c *Context) Set(key string, value interface{}) {
	if c.Keys == nil {
		c.Keys = make(map[string]interface{})
	}
	c.Keys[key] = value
}

// Get returns the value for the given key, ie: (value, true).
// If the value does not exists it returns (nil, false)
func (c *Context) Get(key string) (value interface{}, exists bool) {
	if c.Keys != nil {
		value, exists = c.Keys[key]
	}
	return
}

// Returns the value for the given key if it exists, otherwise it panics.
func (c *Context) MustGet(key string) interface{} {
	if value, exists := c.Get(key); exists {
		return value
	}
	panic("Key \"" + key + "\" does not exist")
}
