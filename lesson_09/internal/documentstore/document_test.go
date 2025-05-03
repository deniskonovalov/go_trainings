package documentstore

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_MarshalDocument_Returns_Document(t *testing.T) {
	input := struct {
		ID     string `json:"id"`
		Name   string `json:"name"`
		Age    int    `json:"age"`
		Active bool   `json:"active"`
	}{
		ID:     "123",
		Name:   "John Marshall",
		Age:    30,
		Active: true,
	}

	doc, err := MarshalDocument(input)

	assert.Nil(t, err)
	assert.NotNil(t, doc)

	assert.Equal(t, "123", doc.Fields["id"].Value)
	assert.Equal(t, "John Marshall", doc.Fields["name"].Value)
	assert.Equal(t, 30.0, doc.Fields["age"].Value)
	assert.Equal(t, true, doc.Fields["active"].Value)

	assert.Equal(t, DocumentFieldTypeString, doc.Fields["id"].Type)
	assert.Equal(t, DocumentFieldTypeString, doc.Fields["name"].Type)
	assert.Equal(t, DocumentFieldTypeNumber, doc.Fields["age"].Type)
	assert.Equal(t, DocumentFieldTypeBool, doc.Fields["active"].Type)
}

func Test_MarshalDocument_EmptyInput_Returns_Empty_Document(t *testing.T) {
	input := struct{}{}

	doc, err := MarshalDocument(input)

	assert.Nil(t, err)
	assert.NotNil(t, doc)
	assert.Empty(t, doc.Fields)
}

func Test_MarshalDocument_InvalidInput_Returns_Error(t *testing.T) {
	var input chan int

	doc, err := MarshalDocument(input)

	assert.Nil(t, doc)
	assert.Error(t, err)
}

func Test_UnmarshalDocument(t *testing.T) {
	doc := &Document{
		Fields: map[string]DocumentField{
			"id": {
				Type:  DocumentFieldTypeString,
				Value: "123",
			},
			"name": {
				Type:  DocumentFieldTypeString,
				Value: "John Marshall",
			},
			"age": {
				Type:  DocumentFieldTypeNumber,
				Value: 30.0,
			},
			"active": {
				Type:  DocumentFieldTypeBool,
				Value: true,
			},
		},
	}

	output := struct {
		ID     string `json:"id"`
		Name   string `json:"name"`
		Age    int    `json:"age"`
		Active bool   `json:"active"`
	}{}

	err := UnmarshalDocument(doc, &output)

	assert.Nil(t, err)
	assert.Equal(t, "123", output.ID)
	assert.Equal(t, "John Marshall", output.Name)
	assert.Equal(t, 30, output.Age)
	assert.Equal(t, true, output.Active)
}

func Test_UnmarshalDocument_EmptyDocument(t *testing.T) {
	doc := &Document{
		Fields: map[string]DocumentField{},
	}

	output := struct {
		ID   string
		Name string
	}{}

	err := UnmarshalDocument(doc, &output)

	assert.Nil(t, err)
	assert.Empty(t, output.ID)
	assert.Empty(t, output.Name)
}

func Test_UnmarshalDocument_InvalidFieldValue(t *testing.T) {
	doc := &Document{
		Fields: map[string]DocumentField{
			"id": {
				Type:  DocumentFieldTypeString,
				Value: true,
			},
		},
	}

	output := struct {
		ID string `json:"id"`
	}{}

	err := UnmarshalDocument(doc, &output)

	assert.Error(t, err)
}
