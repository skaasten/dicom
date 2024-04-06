package processor_test

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/skaasten/dicom/processor"
)

const (
	PatientNameGroup  = 0x0010
	PatientAgeElement = 0x1010
)

func TestTags(t *testing.T) {
	content, err := ioutil.ReadFile("../testdata/1.dcm")
	if err != nil {
		fmt.Println("Error opening DICOM file:", err)
		return
	}

	tags := []processor.Tag{
		{PatientNameGroup, PatientAgeElement},
	}
	attrs, err := processor.HeaderAttrs(content, tags)
	if err != nil {
		t.Error(err)
	}

	if len(attrs) != len(tags) {
		t.Errorf("Expected %d tags, but got %d", len(tags), len(attrs))
	}
	expected := "[038Y]"
	got := attrs[0].Value
	if got != expected {
		t.Errorf("got = %s; want %s", got, expected)
	}
}
