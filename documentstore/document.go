package documentstore

import (
	"errors"
	"learningGo/validation"
)

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
			if !validation.IsValidNumber(field.Value) {
				return errors.New("field " + key + " is not number")
			}
		case DocumentFieldTypeBool:
			if _, ok := field.Value.(bool); !ok {
				return errors.New("field " + key + " is not bool")
			}
		case DocumentFieldTypeArray:
			if !validation.IsValidSlice(field.Value) {
				return errors.New("field " + key + " is not array")
			}
		case DocumentFieldTypeObject:
			if !validation.IsValidMap(field.Value) {
				return errors.New("field " + key + " is not object")
			}
		}
	}

	return nil
}
