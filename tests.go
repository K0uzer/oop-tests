package tests

import (
	"errors"
	"fmt"
	"strings"
	"testing"
)

type fakeStorage struct {
    storage map[string][]byte
}

type CacheService struct {
    store fakeStorage
}

type errorStub struct {
	err error
}

func NewCacheService(store fakeStorage) *CacheService {
    return &CacheService{store: store}
}

func (c *CacheService) GetOrDefault(key string, def []byte) []byte {
	b, e := c.store.Get(key)
	if e != nil {
		return def
	}

	return b
}

func (s *errorStub) Get(key string) ([]byte, error) { 
	return nil, s.err 
}
func (s *errorStub) Put(key string) error { 
	return s.err 
}
func (s *errorStub) Delete(key string) error { 
	return s.err 
}

func TestGetOrDefault_SorageError(t *testing.T) {
	store := &errorStub{err: fmt.Errorf("connecttion refused")}
	scv := NewCacheService(store)

	result := scv.GetOrDefault("key", []byte("fallback"))
	if string(result) != "fallback" {
		t.Errorf("got %q, want %q", result, "fallback")
	}

}

func TestGetOrDefault(t *testing.T) {
	store := newFakeStorage()
	store.storage["getting"] = []byte("hellow")

	svc := NewCacheService(*store)
	
	t.Run("existing key", func(t *testing.T) {
		result := svc.GetOrDefault("greeting", []byte("default"))
		if string(result) != "hello" {
			t.Errorf("got %q, want %q", result, "hellow")
		} 
	})

	t.Run("missing key", func(t *testing.T) {
		result := svc.GetOrDefault("unknown", []byte("default"))
		if string(result) != "default" {
			t.Errorf("got %q, want %q", result, "default")
		}
	})

}



func newFakeStorage() *fakeStorage {
	return &fakeStorage{storage: make(map[string][]byte)}
}

func (s *fakeStorage) Get(key string) ([]byte, error) {
	v, ok := s.storage[key];
	if !ok {
		return nil, errors.New("Такого ключа нет")
	}
	return v, nil
}

func (s *fakeStorage) Put(key string, value []byte) error {
	s.storage[key] = value
	return nil
}

func (s *fakeStorage) Delete(key string) error {
	if _, ok := s.storage[key]; !ok {
		return errors.New("Такого ключа нет")
	}
	fmt.Println("Ключ удален")
	return nil
}

// бенмарк

func benchK1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := ""
		for j := 0; j < 1000; j++ {
			s += "b"
		}
	}
}

func benchK2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var sb strings.Builder
		for j := 0; j < 1000; j++ {
			sb.WriteString("b")
		}
		_ = sb.String()
	}
} 


func Tests() {}