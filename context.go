package tgbot

import "github.com/olebedev/go-tgbot/models"

type Context struct {
	Path    string
	Update  *models.Update
	Keys    map[string]interface{}
	Capture []string

	Text string
	Chat *models.Chat
	From *models.User // ?

	handlers  []func(*Context) error
	index     uint8
	isSkipped bool
}

// FallThrough shows that the next matched handlers should be applied.
func (c *Context) FallThrough() error {
	c.isSkipped = true
	return nil
}

// Next executes the pending handlers in the chain inside the calling handler.
func (c *Context) Next() error {
	c.index++
	var err error
	for ; c.index <= uint8(len(c.handlers)); c.index++ {
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
