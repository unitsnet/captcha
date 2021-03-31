package store

import (
	"bytes"
	"context"
	"log"
	"os"
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
)

const (
	addr = "localhost:16379"
	db   = 15
)

var (
	logger = log.New(os.Stderr, "", log.LstdFlags)
)

func TestRedisSetGet(t *testing.T) {
	ctx := context.Background()
	s := NewRedisStore(&redis.Options{
		Addr:     addr,
		DB:       db,
		Password: "units668",
	}, 100*time.Second, logger)
	id := "captcha id"
	d := []byte("123456")
	s.Set(ctx, id, d)
	d2 := s.Get(ctx, id, false)

	if d2 == nil || !bytes.Equal(d, d2) {
		t.Errorf("saved %v, getDigits returned got %v", d, d2)
	}
}

func TestRedisGetClear(t *testing.T) {
	ctx := context.Background()
	s := NewRedisStore(&redis.Options{
		Addr:     addr,
		DB:       db,
		Password: "units668",
	}, 100*time.Second, logger)
	id := "captcha id"
	d := []byte("1234566")
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

func TestRedisGC(t *testing.T) {
	ctx := context.Background()
	s := NewRedisStore(&redis.Options{
		Addr:     addr,
		DB:       db,
		Password: "units668",
	}, time.Millisecond*10, logger)
	id := "captcha id"
	d := []byte("1234567")
	s.Set(ctx, id, d)

	time.Sleep(time.Millisecond * 200)
	d2 := s.Get(ctx, id, false)

	if d2 != nil {
		t.Errorf("gc didn't clear (%q=%v)", id, d2)
	}
}
