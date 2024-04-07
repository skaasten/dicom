package repository

import (
	"reflect"
	"testing"

	"github.com/google/uuid"
)

func TestAdd(t *testing.T) {
	repo := New()
	fileBuf := []byte{0x00, 0x01, 0x02}
	key := repo.Add(fileBuf)
	content, ok := repo.files[key]
	if !ok {
		t.Error("file buffer was not added to the store")
	}
	if !reflect.DeepEqual(content, fileBuf) {
		t.Error("file buffer does not match the original file buffer")
	}
}

func TestGet_ValidFile(t *testing.T) {
	repo := New()
	fileBuf := []byte{0x00, 0x01, 0x02}
	key := uuid.New()
	repo.files[key] = fileBuf
	content, ok := repo.Get(key)
	if !ok {
		t.Error("file buffer was not found")
	}
	if !reflect.DeepEqual(content, fileBuf) {
		t.Error("file buffer does not match the original file buffer")
	}
}

func TestGet_InvalidFile(t *testing.T) {
	repo := New()
	key := uuid.New()
	content, ok := repo.Get(key)
	if ok != false {
		t.Error("expected the file to not be found")
	}
	if content != nil {
		t.Error("expected content to be nil")
	}
}
