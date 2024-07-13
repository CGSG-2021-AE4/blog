package json

import (
	"context"
	"os"

	"github.com/google/uuid"
)

type ContentStore struct {
	dir string
}

func NewContentStore(dir string) *ContentStore {
	return &ContentStore{
		dir: dir,
	}
}

func (cs *ContentStore) Get(ctx context.Context, id uuid.UUID) ([]byte, error) {
	return os.ReadFile(cs.dir + id.String() + ".md")
}

func (cs *ContentStore) Create(ctx context.Context, id uuid.UUID, content []byte) error {
	return os.WriteFile(cs.dir+id.String()+".md", content, 0777)
}

func (cs *ContentStore) Update(ctx context.Context, id uuid.UUID, content []byte) error {
	return os.WriteFile(cs.dir+id.String()+".md", content, 0777)
}

func (cs *ContentStore) Delete(ctx context.Context, id uuid.UUID) error {
	return os.Remove(cs.dir + id.String() + ".md")
}

func (cs *ContentStore) Close() error {
	return nil
}
