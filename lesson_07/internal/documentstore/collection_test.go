package documentstore

import (
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

func getTestCollection() Collection {
	return Collection{
		items: map[string]Document{},
		config: CollectionConfig{
			PrimaryKey: "id",
		},
	}
}

func Test_Get_Returns_Document(t *testing.T) {
	doc := getTestDocument("100")

	coll := Collection{
		items: map[string]Document{
			"100": doc,
		},
		config: CollectionConfig{
			PrimaryKey: "id",
		},
	}

	document, err := coll.Get("100")

	assert.Nil(t, err)
	assert.Equal(t, &doc, document)
	assert.Equal(t, doc.Fields["id"].Value, document.Fields["id"].Value)
}

func Test_Get_Returns_ErrDocumentNotFound(t *testing.T) {

	coll := Collection{
		items: map[string]Document{
			"122": getTestDocument("122"),
		},
		config: CollectionConfig{
			PrimaryKey: "id",
		},
	}

	_, err := coll.Get("100")

	if assert.Error(t, err) {
		assert.Equal(t, ErrDocumentNotFound, err)
	}
}

func Test_Put_Adds_Document(t *testing.T) {
	coll := getTestCollection()

	doc := getTestDocument("100")

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
	coll := getTestCollection()

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
	coll := getTestCollection()
	coll.items["100"] = getTestDocument("100")

	err := coll.Delete("100")

	assert.Nil(t, err)
	assert.Empty(t, coll.items)
}

func Test_Delete_Returns_ErrDocumentNotFound(t *testing.T) {
	coll := getTestCollection()
	coll.items["100"] = getTestDocument("100")

	err := coll.Delete("111")

	if assert.Error(t, err) {
		assert.Equal(t, ErrDocumentNotFound, err)
	}
	assert.NotEmpty(t, coll.items)
	assert.Contains(t, coll.items, "100")
}

func Test_List_Returns_All_Documents(t *testing.T) {
	coll := getTestCollection()
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
	coll := getTestCollection()

	l := coll.List()

	assert.IsType(t, []Document{}, l)
	assert.Empty(t, l)
}
