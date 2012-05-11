package lock

import (
	"github.com/simonz05/godis"
	"time"
)

type Lock struct {
	UUID     string
	Client   string
	Acquired bool
}

func New(uuid string, client string) *Lock {
	return &Lock{UUID: uuid, Client: client, Acquired: false}
}

func (l *Lock) TryAcquire() bool {
	if l.Acquired {
		return true
	}

	l.Acquired = tryAcquire(l.UUID, l.Client)
	return l.Acquired
}

func (l *Lock) Acquire() bool {
	if l.Acquired {
		return true
	}

	for {
		if acquired := tryAcquire(l.UUID, l.Client); acquired {
			l.Acquired = true
			break
		}

		// wait a moment and try again
		time.Sleep(1 * 1e9)
	}

	return l.Acquired
}

func (l *Lock) Release() {
	if !l.Acquired {
		return
	}

	l.Acquired = false

	release(l.UUID, l.Client)
}

func newClient() *godis.Client {
	return godis.New("", 0, "")
}

func key(uuid string) string {
	return "DISLOCK:LOCK:" + uuid
}

func tryAcquire(uuid string, client string) bool {
	var acquired = false

	var redis = newClient()
	var key = key(uuid)

	if value, _ := redis.Get(key); value.String() == "" {
		// nobody has this lock, so try to acquire it
		acquired, _ = redis.Setnx(key, client)
	} else {
		if value.String() == client {
			// already has the lock
			acquired = true
		} else {
			// other client has the lock
			acquired = false
		}
	}

	return acquired
}

func release(uuid string, client string) {
	var redis = newClient()
	var key = key(uuid)

	if value, _ := redis.Get(key); value.String() == client {
		redis.Del(key)
	}
}
