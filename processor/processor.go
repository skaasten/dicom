package processor

import (
	"bytes"
	"fmt"

	"github.com/suyashkumar/dicom"
)

type Tag uint16

// Returns the specified tags for a given dicom
func Tags(contents []byte, tags []Tag) ([]Tag, error) {
	reader := bytes.NewReader(contents)
	dicomFile, err := dicom.ParseUntilEOF(reader, nil)
	if err != nil {
		fmt.Println("Error parsing DICOM file:", err)
		return []Tag{}, err
	}

	// Print some information about the DICOM file
	fmt.Println("DICOM file details:", dicomFile.Elements)

	// You can access DICOM tags by their names or numbers
	// For example, to get the Patient's Name:
	//patientName := dicomFile.Elements[0x00100010].Value().(string)
	//fmt.Println("Patient's Name:", patientName)

	return []Tag{}, nil
}
