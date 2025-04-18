package documentstore

type Collection struct {
	config CollectionConfig
	items  map[string]Document
}

type CollectionConfig struct {
	PrimaryKey string
}

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

func (c *Collection) Put(doc Document) error {
	key, exists := doc.Fields[c.config.PrimaryKey]

	if !exists || key.Type != DocumentFieldTypeString {
		return ErrDocumentPrimaryKeyIsMissing
	}

	if err := doc.validate(); err != nil {
		return err
	}

	c.items[key.Value.(string)] = doc

	return nil
}

func (c *Collection) Get(key string) (*Document, error) {
	doc, exists := c.items[key]

	if !exists {
		return nil, ErrDocumentNotFound
	}

	return &doc, nil
}

func (c *Collection) Delete(key string) error {
	if _, exists := c.items[key]; !exists {

		return ErrDocumentNotFound
	}

	delete(c.items, key)

	return nil
}

func (c *Collection) List() []Document {
	documentList := make([]Document, 0, len(c.items))

	for _, doc := range c.items {
		documentList = append(documentList, doc)
	}

	return documentList
}
