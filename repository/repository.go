package repository

import (
	"github.com/google/uuid"
)

// ImageStore represents a repository for dicom files
type Repo struct {
	files map[uuid.UUID][]byte
}

func New() *Repo {
	return &Repo{
		files: make(map[uuid.UUID][]byte),
	}
}

func (r *Repo) Get(key uuid.UUID) ([]byte, bool) {
	content, ok := r.files[key]
	return content, ok
}

func (r *Repo) Add(content []byte) uuid.UUID {
	key := uuid.New()
	r.files[key] = content
	return key
}
