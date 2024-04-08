package service

import (
	"github.com/google/uuid"
	"github.com/skaasten/dicom/processor"
)

type Repo interface {
	Add([]byte) uuid.UUID
	Get(key uuid.UUID) ([]byte, bool)
}

type Processor interface {
	HeaderAttrs(contents []byte, tags []processor.Tag) ([]processor.HeaderAttribute, error)
	AsPng(contents []byte) ([][]byte, error)
}

type Dicom struct {
	Repo
}

func New(repo Repo) *Dicom {
	return &Dicom{
		repo,
	}
}
