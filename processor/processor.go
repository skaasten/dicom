package processor

import (
	"bytes"
	"encoding/json"
	"fmt"
	"image/png"

	"github.com/suyashkumar/dicom"
	"github.com/suyashkumar/dicom/pkg/frame"
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

func (t Tag) MarshalJSON() ([]byte, error) {
	// Convert the group and element fields to hexadecimal strings
	groupHex := fmt.Sprintf("%04X", t.Group)
	elementHex := fmt.Sprintf("%04X", t.Element)

	jsonStr := groupHex + ":" + elementHex
	// Marshal the map to JSON
	return json.Marshal(jsonStr)
}

// HeaderAttrs Returns the specified header attributs for a given dicom
func HeaderAttrs(contents []byte, tags []Tag) ([]HeaderAttribute, error) {
	reader := bytes.NewReader(contents)
	data, err := dicom.ParseUntilEOF(reader, nil)
	if err != nil {
		fmt.Println("Error parsing DICOM file:", err)
		return []HeaderAttribute{}, err
	}

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

func AsPng(contents []byte) ([][]byte, error) {
	reader := bytes.NewReader(contents)
	dataset, err := dicom.ParseUntilEOF(reader, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to parse dicom data: %w", err)
	}
	images := [][]byte{}
	for _, elem := range dataset.Elements {
		fmt.Println("elem", elem)
		if elem.Tag == tag.PixelData {
			fmt.Println("found image")
			tagImages, err := writePixelDataElement(elem, "")
			if err != nil {
				return images, fmt.Errorf("failed to convert image: %w", err)
			}
			images = append(images, tagImages...)
		}
	}
	return images, nil
}

func writePixelDataElement(e *dicom.Element, suffix string) ([][]byte, error) {
	imageInfo := e.Value.GetValue().(dicom.PixelDataInfo)
	images := [][]byte{}
	for _, f := range imageInfo.Frames {
		image, err := generateImage(f)
		if err != nil {
			return images, fmt.Errorf("failed to generate image: %w", err)
		}
		images = append(images, image)
	}
	return images, nil
}

func generateImage(fr *frame.Frame) ([]byte, error) {
	i, err := fr.GetImage()
	var buf bytes.Buffer
	if err != nil {
		return nil, fmt.Errorf("failed to get image: %w", err)
	}
	err = png.Encode(&buf, i)
	if err != nil {
		return nil, fmt.Errorf("failed to encode png: %w", err)
	}
	return buf.Bytes(), err
}
