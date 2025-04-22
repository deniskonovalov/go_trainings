package documentstore

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func getTestDocument(id string) Document {
	return Document{
		Fields: map[string]DocumentField{
			"id": {
				Type:  DocumentFieldTypeString,
				Value: id,
			},
		},
	}
}

func getTestCollection() *Collection {
	return &Collection{
		items:  map[string]Document{},
		config: CollectionConfig{},
	}
}

type CollectionBuilder struct {
	collection *Collection
	config     CollectionConfig
	documents  map[string]Document
}

func NewCollectionBuilder(c *Collection) *CollectionBuilder {
	return &CollectionBuilder{
		collection: c,
		config:     c.config,
		documents:  c.items,
	}
}

func (b *CollectionBuilder) WithConfig(c CollectionConfig) *CollectionBuilder {
	b.config = c
	return b
}

func (b *CollectionBuilder) AddDocument(d Document) *CollectionBuilder {
	key, exists := d.Fields[b.config.PrimaryKey]
	if !exists || key.Type != DocumentFieldTypeString {
		fmt.Println("document has no primary key")
		return b
	}

	b.documents[key.Value.(string)] = d

	return b
}

func (b *CollectionBuilder) Build() *Collection {
	b.collection.items = b.documents
	b.collection.config = b.config
	return b.collection
}

func Test_Get_Returns_Document(t *testing.T) {
	config := CollectionConfig{
		PrimaryKey: "id",
	}

	doc := getTestDocument("100")

	collBuilder := NewCollectionBuilder(getTestCollection())
	collBuilder.WithConfig(config)
	collBuilder.AddDocument(doc)
	coll := collBuilder.Build()

	document, err := coll.Get("100")

	assert.Nil(t, err)
	assert.Equal(t, &doc, document)
	assert.Equal(t, doc.Fields["id"].Value, document.Fields["id"].Value)
}

func Test_Get_Returns_ErrDocumentNotFound(t *testing.T) {

	config := CollectionConfig{
		PrimaryKey: "id",
	}

	collBuilder := NewCollectionBuilder(getTestCollection())
	collBuilder.WithConfig(config)
	collBuilder.AddDocument(getTestDocument("122"))
	coll := collBuilder.Build()

	_, err := coll.Get("100")

	if assert.Error(t, err) {
		assert.Equal(t, ErrDocumentNotFound, err)
	}
}

func Test_Put_Adds_Document(t *testing.T) {
	config := CollectionConfig{
		PrimaryKey: "id",
	}

	doc := getTestDocument("100")

	collBuilder := NewCollectionBuilder(getTestCollection())
	collBuilder.WithConfig(config)
	coll := collBuilder.Build()

	err := coll.Put(doc)

	assert.Nil(t, err)
	assert.Contains(t, coll.items, "100")
	assert.Equal(t, doc, coll.items["100"])
	assert.Equal(t, doc.Fields["id"].Value, coll.items["100"].Fields["id"].Value)
}

func Test_Put_Returns_Document_ErrDocumentPrimaryKeyIsMissing(t *testing.T) {
	coll := getTestCollection()

	doc := Document{
		Fields: map[string]DocumentField{
			"test_field_name": {
				Type:  DocumentFieldTypeString,
				Value: "test_field_value",
			},
		},
	}

	err := coll.Put(doc)

	if assert.Error(t, err) {
		assert.Equal(t, ErrDocumentPrimaryKeyIsMissing, err)
	}
	assert.Empty(t, coll.items)
}

func Test_Put_Returns_ValidationError(t *testing.T) {
	config := CollectionConfig{
		PrimaryKey: "id",
	}

	collBuilder := NewCollectionBuilder(getTestCollection())
	collBuilder.WithConfig(config)
	coll := collBuilder.Build()

	doc := Document{
		Fields: map[string]DocumentField{
			"id": {
				Type:  DocumentFieldTypeString,
				Value: "100",
			},
			"test_field_name": {
				Type:  DocumentFieldTypeNumber,
				Value: "test_field_value",
			},
		},
	}

	err := coll.Put(doc)

	if assert.Error(t, err) {
		assert.Equal(t, ErrInvalidTypeNumber, err)
	}
	assert.Empty(t, coll.items)
}

func Test_Delete_Removes_Document(t *testing.T) {
	config := CollectionConfig{
		PrimaryKey: "id",
	}

	collBuilder := NewCollectionBuilder(getTestCollection())
	collBuilder.WithConfig(config)
	coll := collBuilder.Build()

	coll.items["100"] = getTestDocument("100")

	err := coll.Delete("100")

	assert.Nil(t, err)
	assert.Empty(t, coll.items)
}

func Test_Delete_Returns_ErrDocumentNotFound(t *testing.T) {
	config := CollectionConfig{
		PrimaryKey: "id",
	}

	collBuilder := NewCollectionBuilder(getTestCollection())
	collBuilder.WithConfig(config)
	coll := collBuilder.Build()

	coll.items["100"] = getTestDocument("100")

	err := coll.Delete("111")

	if assert.Error(t, err) {
		assert.Equal(t, ErrDocumentNotFound, err)
	}
	assert.NotEmpty(t, coll.items)
	assert.Contains(t, coll.items, "100")
}

func Test_List_Returns_All_Documents(t *testing.T) {
	config := CollectionConfig{
		PrimaryKey: "id",
	}

	collBuilder := NewCollectionBuilder(getTestCollection())
	collBuilder.WithConfig(config)
	coll := collBuilder.Build()

	coll.items["100"] = getTestDocument("100")
	coll.items["101"] = getTestDocument("101")
	coll.items["102"] = getTestDocument("102")

	l := coll.List()

	assert.IsType(t, []Document{}, l)
	assert.Len(t, l, 3)
	assert.Contains(t, l, coll.items["100"])
	assert.Contains(t, l, coll.items["101"])
	assert.Contains(t, l, coll.items["102"])
}

func Test_List_Returns_Empty_Slice(t *testing.T) {
	config := CollectionConfig{
		PrimaryKey: "id",
	}

	collBuilder := NewCollectionBuilder(getTestCollection())
	collBuilder.WithConfig(config)
	coll := collBuilder.Build()

	l := coll.List()

	assert.IsType(t, []Document{}, l)
	assert.Empty(t, l)
}
