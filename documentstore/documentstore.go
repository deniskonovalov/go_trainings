package documentstore

import (
	"errors"
)

var documentStore = map[string]Document{}

type DocumentFieldType string

const (
	DocumentFieldTypeString DocumentFieldType = "string"
	DocumentFieldTypeNumber DocumentFieldType = "number"
	DocumentFieldTypeBool   DocumentFieldType = "bool"
	DocumentFieldTypeArray  DocumentFieldType = "array"
	DocumentFieldTypeObject DocumentFieldType = "object"
)

type DocumentField struct {
	Type  DocumentFieldType
	Value interface{}
}

type Document struct {
	Fields map[string]DocumentField
}

func (d Document) validate() error {
	for key, field := range d.Fields {
		switch field.Type {
		case DocumentFieldTypeString:
			if _, ok := field.Value.(string); !ok {
				return errors.New("field " + key + " is not string")
			}
		case DocumentFieldTypeNumber:
			if _, ok := field.Value.(int64); !ok {
				if _, ok := field.Value.(float64); !ok {
					return errors.New("field " + key + " is not number")
				}
			}
		case DocumentFieldTypeBool:
			if _, ok := field.Value.(bool); !ok {
				return errors.New("field " + key + " is not bool")
			}
		case DocumentFieldTypeArray:
			if _, ok := field.Value.([]any); !ok {
				return errors.New("field " + key + " is not array")
			}
		}
	}

	return nil
}

// Put Updated signature of the function,
// added return of an error if key filed is not exists in Document
func Put(doc Document) error {
	key, exists := doc.Fields["key"]

	if !exists || key.Type != DocumentFieldTypeString {
		return errors.New("can not put document without key field")
	}

	if err := doc.validate(); err != nil {
		return err
	}

	documentStore[key.Value.(string)] = doc

	return nil
}

func Get(key string) (*Document, bool) {
	doc, exists := documentStore[key]

	if !exists {
		return nil, false
	}

	return &doc, true
}

func Delete(key string) bool {
	if _, exists := documentStore[key]; !exists {
		return false
	}
	delete(documentStore, key)
	return true
}

func List() []Document {
	var documentList []Document

	for _, doc := range documentStore {
		documentList = append(documentList, doc)
	}

	return documentList
}
