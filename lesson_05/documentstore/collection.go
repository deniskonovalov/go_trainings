package documentstore

type Collection struct {
	config CollectionConfig
	items  map[string]Document
}

type CollectionConfig struct {
	PrimaryKey string
}

func (s *Collection) Put(doc Document) error {
	key, exists := doc.Fields[s.config.PrimaryKey]

	if !exists || key.Type != DocumentFieldTypeString {
		return ErrDocumentPrimaryKeyIsMissing
	}

	if err := doc.validate(); err != nil {
		return err
	}

	s.items[key.Value.(string)] = doc

	return nil
}

func (s *Collection) Get(key string) (*Document, error) {
	doc, exists := s.items[key]

	if !exists {
		return nil, ErrDocumentNotFound
	}

	return &doc, nil
}

func (s *Collection) Delete(key string) error {
	if _, exists := s.items[key]; !exists {

		return ErrDocumentNotFound
	}

	delete(s.items, key)

	return nil
}

func (s *Collection) List() []Document {
	documentList := make([]Document, 0, len(s.items))

	for _, doc := range s.items {
		documentList = append(documentList, doc)
	}

	return documentList
}
