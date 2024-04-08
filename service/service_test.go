package service_test

import (
	"io/ioutil"
	"reflect"
	"testing"

	"github.com/google/uuid"
	"github.com/skaasten/dicom/processor"
	"github.com/skaasten/dicom/repository"
	"github.com/skaasten/dicom/service"
)

func TestAddAndGetImage(t *testing.T) {
	repo := repository.New()
	srv := service.New(repo)

	filePath := "../testdata/5.dcm"
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		t.Fatalf("error opening dicom file: %v", err)
	}
	key := srv.Add(content)
	if key == uuid.Nil {
		t.Errorf("key should not be empty")
	}

	retrieved, ok := srv.Get(key)
	if !ok {
		t.Errorf("file should be returned")
	}
	if !reflect.DeepEqual(retrieved, content) {
		t.Errorf("file contents should match original")
	}

	attrs, err := srv.HeaderAttributes(key, []processor.Tag{
		processor.Tag{0x0010, 0x1010},
	})
	if err != nil {
		t.Errorf("error getting attrs: %v", err)
	}
	expectedValue := "[052Y]"
	if attrs[0].Value != expectedValue {
		t.Errorf("expected: %s, got: %s", expectedValue, attrs[0].Value)
	}
}
