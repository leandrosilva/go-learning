package lock

import (
	"github.com/simonz05/godis"
	"testing"
)

// WARNING: This test suite requires Redis is up and runnig

func reset() {
	var redis = godis.New("", 0, "")
	redis.Flushdb()
}

func TestNew(t *testing.T) {
	var uuid = "nuclear_read"
	var l = New(uuid, "homer_simpson")

	if l.UUID != uuid {
		t.Errorf("lock.New(%v) = %v, want %v", uuid, l.UUID, uuid)
	}

	if l.Acquired {
		t.Errorf("lock.Acquired = %v, want %v", l.Acquired, false)
	}
}

func TestTryAcquire_Sucess(t *testing.T) {
	reset()

	var uuid = "nuclear_read"
	var l = New(uuid, "homer_simpson")

	if sucess := l.TryAcquire(); !sucess {
		t.Errorf("lock.TryAcquire() = %v, want = %v", sucess, true)
	}

	if !l.Acquired {
		t.Errorf("lock.Acquired = %v, want = %v", l.Acquired, true)
	}
}

func TestTryAcquire_Fail(t *testing.T) {
	reset()

	var uuid = "nuclear_read"
	var l = New(uuid, "mr_burns")

	if sucess := l.TryAcquire(); !sucess {
		t.Errorf("lock.TryAcquire() = %v, want = %v", sucess, true)
	}

	if !l.Acquired {
		t.Errorf("lock.Acquired = %v, want = %v", l.Acquired, true)
	}

	var l2 = New(uuid, "homer_simpson")

	if sucess := l2.TryAcquire(); sucess {
		t.Errorf("lock.TryAcquire() = %v, want = %v", sucess, false)
	}

	if l2.Acquired {
		t.Errorf("lock.Acquired = %v, want = %v", l2.Acquired, false)
	}
}

func TestRelease(t *testing.T) {
	reset()

	var uuid = "nuclear_read"
	var l = New(uuid, "homer_simpson")

	if sucess := l.TryAcquire(); !sucess {
		t.Errorf("lock.TryAcquire() = %v, want = %v", sucess, true)
	}

	if !l.Acquired {
		t.Errorf("lock.Acquired = %v, want = %v", l.Acquired, true)
	}

	l.Release()

	if l.Acquired {
		t.Errorf("lock.Acquired = %v, want = %v", l.Acquired, false)
	}
}
