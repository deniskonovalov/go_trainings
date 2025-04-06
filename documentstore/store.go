package documentstore

type Store struct {
	Collections map[string]*Collection
}

func NewStore() *Store {
	return &Store{Collections: map[string]*Collection{}}
}

func (s *Store) CreateCollection(name string, cfg *CollectionConfig) (bool, *Collection) {
	if _, exists := s.Collections[name]; exists {
		return false, nil
	}

	collection := &Collection{config: *cfg, items: map[string]Document{}}
	s.Collections[name] = collection

	return true, collection
}

func (s *Store) GetCollection(name string) (*Collection, bool) {
	if collection, exists := s.Collections[name]; exists {
		return collection, true
	}

	return nil, false
}

func (s *Store) DeleteCollection(name string) bool {
	if _, exists := s.Collections[name]; !exists {
		return false
	}

	delete(s.Collections, name)

	return true
}
