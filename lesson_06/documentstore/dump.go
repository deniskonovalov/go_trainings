package documentstore

import (
	"bytes"
	"encoding/gob"
	"os"
)

// dumpCollection is a DTO used to serialize Collection,
// because original Collection has unexported fields,
type dumpCollection struct {
	Config CollectionConfig
	Items  map[string]Document
}

// toDump returns a DTO representation of Collection for serialization.
func (c *Collection) toDump() dumpCollection {
	return dumpCollection{
		Config: c.config,
		Items:  c.items,
	}
}

// dumpStore is a DTO used to serialize Store.
type dumpStore struct {
	Collections map[string]dumpCollection
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
			config: dumpedCol.Config,
			items:  dumpedCol.Items,
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
