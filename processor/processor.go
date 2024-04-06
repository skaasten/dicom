package processor

import (
	"bytes"
	"fmt"

	"github.com/suyashkumar/dicom"
	"github.com/suyashkumar/dicom/pkg/tag"
)

type Tag struct {
	Group   uint16 `json:"group"`
	Element uint16 `json:"element"`
}

type HeaderAttribute struct {
	Tag   Tag    `json:"tag"`
	Value string `json:"value"`
}

// HeaderAttrs Returns the specified header attributs for a given dicom
func HeaderAttrs(contents []byte, tags []Tag) ([]HeaderAttribute, error) {
	reader := bytes.NewReader(contents)
	data, err := dicom.ParseUntilEOF(reader, nil)
	if err != nil {
		fmt.Println("Error parsing DICOM file:", err)
		return []HeaderAttribute{}, err
	}
	//fmt.Println(data.Elements)

	attrs := []HeaderAttribute{}
	for _, t := range tags {
		dicomTag := tag.Tag{t.Group, t.Element}
		elem, err := data.FindElementByTag(dicomTag)
		if err != nil {
			return attrs, err
		}
		attr := HeaderAttribute{
			Tag:   t,
			Value: elem.Value.String(),
		}
		attrs = append(attrs, attr)
	}
	return attrs, nil
}
