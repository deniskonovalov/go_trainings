package documentstore

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/tidwall/btree"
	"testing"
)

func getTestDocument(id string) *Document {
	return &Document{
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
		items:   map[string]*Document{},
		config:  CollectionConfig{},
		indexes: map[string]*btree.Map[string, string]{},
	}
}

type CollectionBuilder struct {
	collection *Collection
	config     CollectionConfig
	documents  map[string]*Document
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

func (b *CollectionBuilder) AddDocument(d *Document) *CollectionBuilder {
	key, exists := d.Fields[b.config.PrimaryKey]
	if !exists || key.Type != DocumentFieldTypeString {
		fmt.Println("document has no primary Key")
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
	assert.Equal(t, &doc, &document)
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

	assert.Error(t, err)
	assert.Equal(t, ErrDocumentNotFound, err)

}

func Test_Put_Adds_Document(t *testing.T) {
	config := CollectionConfig{
		PrimaryKey: "id",
	}

	doc := getTestDocument("100")

	collBuilder := NewCollectionBuilder(getTestCollection())
	collBuilder.WithConfig(config)
	coll := collBuilder.Build()

	err := coll.Put(*doc)

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
	assert.Contains(t, l, *coll.items["100"])
	assert.Contains(t, l, *coll.items["101"])
	assert.Contains(t, l, *coll.items["102"])
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

func Test_CreateIndex_Creates_Index(t *testing.T) {
	config := CollectionConfig{
		PrimaryKey: "id",
	}

	doc1 := getTestDocument("1")
	doc1.Fields["field1"] = DocumentField{
		Type:  DocumentFieldTypeString,
		Value: "value1",
	}

	doc2 := getTestDocument("2")
	doc2.Fields["field1"] = DocumentField{
		Type:  DocumentFieldTypeString,
		Value: "value2",
	}

	collBuilder := NewCollectionBuilder(getTestCollection())
	collBuilder.WithConfig(config)
	collBuilder.AddDocument(doc1)
	collBuilder.AddDocument(doc2)
	coll := collBuilder.Build()

	err := coll.CreateIndex("field1")

	assert.Nil(t, err)
	assert.Contains(t, coll.indexes, "field1")

	val1, exists1 := coll.indexes["field1"].Get("value1")
	val2, exists2 := coll.indexes["field1"].Get("value2")

	assert.True(t, exists1)
	assert.True(t, exists2)
	assert.Equal(t, "1", val1)
	assert.Equal(t, "2", val2)
}

func Test_CreateIndex_Returns_ErrIndexAlreadyExists(t *testing.T) {
	config := CollectionConfig{
		PrimaryKey: "id",
	}

	collBuilder := NewCollectionBuilder(getTestCollection())
	collBuilder.WithConfig(config)
	coll := collBuilder.Build()

	err := coll.CreateIndex("field1")
	assert.Nil(t, err)

	err = coll.CreateIndex("field1")

	assert.NotNil(t, err)
	assert.Equal(t, ErrIndexAlreadyExists, err)
}

func Test_DeleteIndex_Removes_Index(t *testing.T) {
	config := CollectionConfig{
		PrimaryKey: "id",
	}

	collBuilder := NewCollectionBuilder(getTestCollection())
	collBuilder.WithConfig(config)
	coll := collBuilder.Build()

	err := coll.CreateIndex("field1")
	assert.Nil(t, err)

	err = coll.DeleteIndex("field1")
	assert.Nil(t, err)
	assert.NotContains(t, coll.indexes, "field1")
}

func Test_DeleteIndex_Returns_ErrIndexDoesNotExist(t *testing.T) {
	config := CollectionConfig{
		PrimaryKey: "id",
	}

	collBuilder := NewCollectionBuilder(getTestCollection())
	collBuilder.WithConfig(config)
	coll := collBuilder.Build()

	err := coll.DeleteIndex("nonexistent_field")

	assert.NotNil(t, err)
	assert.Equal(t, ErrIndexDoesNotExist, err)
}

func Test_Query_Returns_Correct_Results(t *testing.T) {
	config := CollectionConfig{
		PrimaryKey: "id",
	}

	doc1 := getTestDocument("1")
	doc1.Fields["field1"] = DocumentField{
		Type:  DocumentFieldTypeString,
		Value: "value1",
	}

	doc2 := getTestDocument("2")
	doc2.Fields["field1"] = DocumentField{
		Type:  DocumentFieldTypeString,
		Value: "value2",
	}

	doc3 := getTestDocument("3")
	doc3.Fields["field1"] = DocumentField{
		Type:  DocumentFieldTypeString,
		Value: "value3",
	}

	collBuilder := NewCollectionBuilder(getTestCollection())
	collBuilder.WithConfig(config)
	collBuilder.AddDocument(doc1)
	collBuilder.AddDocument(doc2)
	collBuilder.AddDocument(doc3)
	coll := collBuilder.Build()

	err := coll.CreateIndex("field1")
	assert.Nil(t, err)

	minVal := "value1"
	maxVal := "value2"

	qp := QueryParams{
		Desc:     false,
		MinValue: &minVal,
		MaxValue: &maxVal,
	}

	result, err := coll.Query("field1", qp)

	fmt.Printf("result: %v\n", result)
	assert.Nil(t, err)
	assert.Len(t, result, 2)
	assert.Contains(t, result, *doc1)
	assert.Contains(t, result, *doc2)
}

func Test_Query_Returns_ErrIndexDoesNotExist(t *testing.T) {
	config := CollectionConfig{
		PrimaryKey: "id",
	}

	collBuilder := NewCollectionBuilder(getTestCollection())
	collBuilder.WithConfig(config)
	coll := collBuilder.Build()

	minVal := "value1"
	maxVal := "value2"

	qp := QueryParams{
		Desc:     false,
		MinValue: &minVal,
		MaxValue: &maxVal,
	}

	result, err := coll.Query("nonexistent_field", qp)

	assert.NotNil(t, err)
	assert.Nil(t, result)
	assert.Equal(t, ErrIndexDoesNotExist, err)
}
