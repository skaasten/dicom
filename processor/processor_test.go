package processor_test

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/skaasten/dicom/processor"
)

func TestTags(t *testing.T) {
	content, err := ioutil.ReadFile("../testdata/1.dcm")
	if err != nil {
		fmt.Println("Error opening DICOM file:", err)
		return
	}

	tags := []processor.Tag{0x0010, 0x0020, 0x0028}
	retrievedTags, err := processor.Tags(content, tags)
	if err != nil {
		t.Error(err)
	}

	if len(retrievedTags) != len(tags) {
		t.Errorf("Expected %d tags, but got %d", len(tags), len(retrievedTags))
	}

	expectedFirstTag := tags[0]
	if retrievedTags[0] != expectedFirstTag {
		t.Errorf("Expected first tag to be %v, but got %v", expectedFirstTag, retrievedTags[0])
	}
}
