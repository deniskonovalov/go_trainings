package documentstore

type Store struct {
	Collections map[string]*Collection
}

func NewStore() *Store {
	return &Store{Collections: map[string]*Collection{}}
}

func (s *Store) CreateCollection(name string, cfg *CollectionConfig) (*Collection, error) {
	if _, exists := s.Collections[name]; exists {
		return nil, ErrCollectionAlreadyExists
	}

	collection := &Collection{config: *cfg, items: map[string]Document{}}
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
