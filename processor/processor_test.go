package processor_test

import (
	"io/ioutil"
	"reflect"
	"testing"

	"github.com/skaasten/dicom/processor"
)

const (
	PatientNameGroup     = 0x0010
	PatientAgeElement    = 0x1010
	PatientWeightElement = 0x1030
)

func TestHeaderAttrs(t *testing.T) {
	testCases := []struct {
		name        string
		filePath    string
		tags        []processor.Tag
		expected    []processor.HeaderAttribute
		expectedErr error
	}{
		{
			name:     "Single Tag",
			filePath: "../testdata/1.dcm",
			tags: []processor.Tag{
				{PatientNameGroup, PatientAgeElement},
			},
			expected: []processor.HeaderAttribute{
				processor.HeaderAttribute{
					processor.Tag{PatientNameGroup, PatientAgeElement},
					"[038Y]",
				},
			},
		},
		{
			name:     "Multiple tags",
			filePath: "../testdata/1.dcm",
			tags: []processor.Tag{
				{PatientNameGroup, PatientAgeElement},
				{PatientNameGroup, PatientWeightElement},
			},
			expected: []processor.HeaderAttribute{
				processor.HeaderAttribute{
					processor.Tag{PatientNameGroup, PatientAgeElement},
					"[038Y]",
				},
				processor.HeaderAttribute{
					processor.Tag{PatientNameGroup, PatientWeightElement},
					"[69.7344]",
				},
			},
		},
		// Add more test cases as needed
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			content, err := ioutil.ReadFile(tc.filePath)
			if err != nil {
				t.Fatalf("Error opening DICOM file: %v", err)
			}

			attrs, err := processor.HeaderAttrs(content, tc.tags)
			if err != tc.expectedErr {
				t.Errorf("Unexpected error: got %v, want %v", err, tc.expectedErr)
			}

			if len(attrs) != len(tc.tags) {
				t.Errorf("Expected %d tags, but got %d", len(tc.tags), len(attrs))
			}
			got := attrs
			if !reflect.DeepEqual(got, tc.expected) {
				t.Errorf("got = %v; want %v", got, tc.expected)
			}
		})
	}
}
