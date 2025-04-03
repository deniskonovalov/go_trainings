package documentstore

import "errors"

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

// Put Updated signature of the function,
// added return of an error if key filed is not exists in Document
func Put(doc Document) (bool, error) {
	key, exists := doc.Fields["key"]

	if !exists || key.Type != DocumentFieldTypeString {
		return false, errors.New("can not put document without key field")
	}

	documentStore[key.Value.(string)] = doc

	return true, nil
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
