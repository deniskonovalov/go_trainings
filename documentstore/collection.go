package documentstore

import "errors"

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
		return errors.New("can not put document without primary key field")
	}

	if err := doc.validate(); err != nil {
		return err
	}

	s.items[key.Value.(string)] = doc

	return nil
}

func (s *Collection) Get(key string) (*Document, bool) {
	doc, exists := s.items[key]

	if !exists {
		return nil, false
	}

	return &doc, true
}

func (s *Collection) Delete(key string) bool {
	if _, exists := s.items[key]; !exists {
		return false
	}

	delete(s.items, key)

	return true
}

func (s *Collection) List() []Document {
	documentList := make([]Document, len(s.items))

	for _, doc := range s.items {
		documentList = append(documentList, doc)
	}

	return documentList
}
