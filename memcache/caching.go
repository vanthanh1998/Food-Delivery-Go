package memcache

import (
	"Food-Delivery/common"
	"sync"
	"time"
)

type Caching interface {
	Write(k string, value interface{}) // key , value
	Read(k string) interface{}         // key => return interface
}

type caching struct {
	store  map[string]interface{}
	locker *sync.RWMutex
}

func NewCaching() *caching { // func call run caching
	return &caching{
		store:  make(map[string]interface{}),
		locker: new(sync.RWMutex),
	}
}

func (c *caching) Write(k string, value interface{}) {
	c.locker.Lock()
	defer c.locker.Unlock()
	c.store[k] = value
}

func (c *caching) Read(k string) interface{} {
	c.locker.RLock()
	defer c.locker.RUnlock()
	return c.store[k]
}

// cache range time
func (c *caching) WriteTTL(k string, value interface{}, exp int) { // key, value, exp(time => seconds)
	c.locker.Lock()
	defer c.locker.Unlock()
	c.store[k] = value

	go func() { // hẹn giờ
		defer common.AppRecover()
		<-time.NewTicker(time.Second * time.Duration(exp)).C
		c.Write(k, nil)
	}()
}
