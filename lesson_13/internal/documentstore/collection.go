package documentstore

import (
	"github.com/tidwall/btree"
	"sync"
)

type Collection struct {
	mu      sync.RWMutex
	config  CollectionConfig
	items   map[string]*Document
	indexes map[string]*btree.Map[string, string]
}

type CollectionConfig struct {
	PrimaryKey string
}

// dumpCollection is a DTO used to serialize Collection,
// because original Collection has unexported fields,
type dumpCollection struct {
	Config  CollectionConfig
	Items   map[string]*Document
	Indexes map[string]map[string]string
}

type QueryParams struct {
	Desc     bool
	MinValue *string
	MaxValue *string
}

// toDump returns a DTO representation of Collection for serialization.
func (c *Collection) toDump() dumpCollection {
	dumpIndexes := make(map[string]map[string]string)
	c.mu.RLock()
	defer c.mu.RUnlock()

	for field, tree := range c.indexes {
		dumpIndexes[field] = make(map[string]string)
		tree.Scan(func(key, value string) bool {
			dumpIndexes[field][key] = value
			return true
		})
	}

	return dumpCollection{
		Config:  c.config,
		Items:   c.items,
		Indexes: dumpIndexes,
	}
}

func (c *Collection) Put(doc Document) error {
	k, exists := doc.Fields[c.config.PrimaryKey]
	if !exists || k.Type != DocumentFieldTypeString {

		return ErrDocumentPrimaryKeyIsMissing
	}
	keyString := k.Value.(string)

	c.mu.Lock()

	if existingDoc, exists := c.items[keyString]; exists {
		for field, tree := range c.indexes {
			if docField, exists := existingDoc.Fields[field]; exists {
				tree.Delete(docField.Value.(string))
			}
		}
	}

	if err := doc.validate(); err != nil {
		return err
	}

	c.items[keyString] = &doc

	for field, tree := range c.indexes {
		docField, exists := doc.Fields[field]
		if exists && docField.Type == DocumentFieldTypeString {
			tree.Set(docField.Value.(string), keyString)
		}
	}

	c.mu.Unlock()

	return nil
}

func (c *Collection) Get(key string) (*Document, error) {
	c.mu.RLock()

	doc, exists := c.items[key]

	c.mu.RUnlock()

	if !exists {
		return nil, ErrDocumentNotFound
	}

	return doc, nil
}

func (c *Collection) Delete(key string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	doc, exists := c.items[key]

	if !exists {
		return ErrDocumentNotFound
	}

	for field, tree := range c.indexes {
		docField, exists := doc.Fields[field]

		if !exists {
			continue
		}

		tree.Delete(docField.Value.(string))
	}

	delete(c.items, key)

	return nil
}

func (c *Collection) List() []Document {
	c.mu.RLock()
	defer c.mu.RUnlock()

	documentList := make([]Document, 0, len(c.items))

	for _, doc := range c.items {
		documentList = append(documentList, *doc)
	}

	return documentList
}

func (c *Collection) CreateIndex(fieldName string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if _, exists := c.indexes[fieldName]; exists {
		return ErrIndexAlreadyExists
	}

	tree := btree.Map[string, string]{}

	for _, doc := range c.items {
		val, exists := doc.Fields[fieldName]

		if !exists || val.Type != DocumentFieldTypeString {
			continue
		}

		treeKey := val.Value.(string)
		docKey := doc.Fields[c.config.PrimaryKey].Value.(string)

		tree.Set(treeKey, docKey)
	}

	c.indexes[fieldName] = &tree

	return nil
}

func (c *Collection) DeleteIndex(fieldName string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if _, exists := c.indexes[fieldName]; !exists {
		return ErrIndexDoesNotExist
	}

	delete(c.indexes, fieldName)

	return nil
}

func (c *Collection) Query(fieldName string, params QueryParams) ([]Document, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	tree, exists := c.indexes[fieldName]

	if !exists {
		return nil, ErrIndexDoesNotExist
	}

	minVal := *params.MinValue
	maxVal := *params.MaxValue

	docList := make([]Document, 0)

	if params.Desc {
		tree.Descend(maxVal, func(key string, value string) bool {
			if key < minVal {
				return false
			}
			docList = append(docList, *c.items[value])
			return true
		})
	} else {
		tree.Ascend(minVal, func(key string, value string) bool {
			if key > maxVal {

				return false
			}
			docList = append(docList, *c.items[value])
			return true
		})
	}

	return docList, nil
}

func (c *Collection) GetPrimaryKey() string {
	return c.config.PrimaryKey
}
