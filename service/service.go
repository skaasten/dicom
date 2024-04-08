package service

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/skaasten/dicom/processor"
)

type Repo interface {
	Add([]byte) uuid.UUID
	Get(key uuid.UUID) ([]byte, bool)
}

type Dicom struct {
	Repo
}

func New(repo Repo) *Dicom {
	return &Dicom{
		repo,
	}
}

func (d *Dicom) HeaderAttributes(key uuid.UUID, tags []processor.Tag) ([]processor.HeaderAttribute, error) {
	file, ok := d.Get(key)
	if !ok {
		return nil, fmt.Errorf("failed to get file")
	}
	attrs, err := processor.HeaderAttrs(file, tags)
	if err != nil {
		return nil, fmt.Errorf("error getting header attrs: %w", err)
	}
	return attrs, nil
}
