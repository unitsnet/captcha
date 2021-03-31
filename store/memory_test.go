package store

import (
	"bytes"
	"context"
	"testing"
	"time"
)

func TestMemorySetGet(t *testing.T) {
	s := NewMemoryStore(time.Second, time.Second*2)
	id := "captcha id"
	d := []byte("123456")
	ctx := context.Background()
	s.Set(ctx, id, d)
	d2 := s.Get(ctx, id, false)

	if d2 == nil || !bytes.Equal(d, d2) {
		t.Errorf("saved %v, getDigits returned got %v", d, d2)
	}
}

func TestMemoryGetClear(t *testing.T) {
	s := NewMemoryStore(time.Second, time.Second*2)
	id := "captcha id"
	d := []byte("123456")
	ctx := context.Background()
	s.Set(ctx, id, d)
	d2 := s.Get(ctx, id, true)
	if d2 == nil || !bytes.Equal(d, d2) {
		t.Errorf("saved %v, getDigitsClear returned got %v", d, d2)
	}

	d2 = s.Get(ctx, id, false)
	if d2 != nil {
		t.Errorf("getDigitClear didn't clear (%q=%v)", id, d2)
	}
}

func TestMemoryGC(t *testing.T) {
	s := NewMemoryStore(time.Millisecond*10, time.Millisecond*100)
	id := "captcha id"
	d := []byte("123456")
	ctx := context.Background()
	s.Set(ctx, id, d)

	time.Sleep(time.Millisecond * 200)
	d2 := s.Get(ctx, id, false)

	if d2 != nil {
		t.Errorf("gc didn't clear (%q=%v)", id, d2)
	}
}
