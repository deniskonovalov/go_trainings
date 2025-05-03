package documentstore

import (
	"bytes"
	"encoding/gob"
	"github.com/tidwall/btree"
	"os"
)

type Store struct {
	Collections map[string]*Collection
}

// dumpStore is a DTO used to serialize Store.
type dumpStore struct {
	Collections map[string]dumpCollection
}

func NewStore() *Store {
	return &Store{Collections: map[string]*Collection{}}
}

func (s *Store) CreateCollection(name string, cfg *CollectionConfig) (*Collection, error) {
	if _, exists := s.Collections[name]; exists {
		return nil, ErrCollectionAlreadyExists
	}

	collection := &Collection{
		config:  *cfg,
		items:   map[string]*Document{},
		indexes: map[string]*btree.Map[string, string]{},
	}

	s.Collections[name] = collection

	return collection, nil
}

func (s *Store) GetCollection(name string) (*Collection, error) {
	if collection, exists := s.Collections[name]; exists {
		return collection, nil
	}

	return nil, ErrCollectionNotFound
}

func (s *Store) DeleteCollection(name string) error {
	if _, exists := s.Collections[name]; !exists {
		return ErrCollectionNotFound
	}

	delete(s.Collections, name)

	return nil
}

// toDump returns a DTO representation of Store for serialization.
func (s *Store) toDump() dumpStore {
	result := dumpStore{
		Collections: make(map[string]dumpCollection),
	}

	for name, col := range s.Collections {
		result.Collections[name] = col.toDump()
	}

	return result
}

func (s *Store) Dump() ([]byte, error) {
	dumped := s.toDump()

	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(dumped)

	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func NewStoreFromDump(dump []byte) (*Store, error) {
	var dumpStore dumpStore

	buf := bytes.NewBuffer(dump)
	dec := gob.NewDecoder(buf)

	err := dec.Decode(&dumpStore)
	if err != nil {
		return nil, err
	}

	store := &Store{
		Collections: make(map[string]*Collection),
	}

	for name, dumpedCol := range dumpStore.Collections {
		store.Collections[name] = &Collection{
			config:  dumpedCol.Config,
			items:   dumpedCol.Items,
			indexes: make(map[string]*btree.Map[string, string]),
		}

		for field, index := range dumpedCol.Indexes {
			tree := &btree.Map[string, string]{}
			for key, value := range index {
				tree.Set(key, value)
			}
			store.Collections[name].indexes[field] = tree
		}
	}

	return store, nil
}

func (s *Store) DumpToFile(filename string) error {
	data, err := s.Dump()
	if err != nil {
		return err
	}

	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

func NewStoreFromFile(filename string) (*Store, error) {
	data, err := os.ReadFile(filename)

	if err != nil {
		return nil, err
	}

	return NewStoreFromDump(data)
}
