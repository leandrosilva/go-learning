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

func newPipeClient() *godis.PipeClient {
	return godis.NewPipeClient("", 0, "")
}

func key(uuid string) string {
	return "DISLOCK:LOCK:" + uuid
}

func tryAcquire(uuid string, client string) bool {
	var acquired = false

	var redis = newPipeClient()
	var key = key(uuid)

	redis.Watch(key)

	// what about this lock?
	redis.Get(key)
	var replyGet = redis.Exec()[0].Elem.String()

	if replyGet != "" {
		if replyGet == client {
			// already has the lock
			acquired = true
		} else {
			// other client has the lock
			acquired = false
		}
	} else {
		// try to acquire the lock
		redis.Set(key, client)
		var replySet = redis.Exec()[0].Elem.String()

		if replySet == "OK" {
			// it was successful to acquire the lock
			acquired = true
		}
	}

	redis.Unwatch()

	return acquired
}

func release(uuid string, client string) {
	var redis = newClient()
	var key = key(uuid)

	redis.Del(key)
}
