package documentstore

import "encoding/json"

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

func MarshalDocument(input any) (*Document, error) {
	d := Document{
		Fields: make(map[string]DocumentField),
	}
	jsonData, err := json.Marshal(input)

	if err != nil {
		return nil, err
	}

	m := make(map[string]any)
	err = json.Unmarshal(jsonData, &m)

	if err != nil {
		return nil, err
	}

	for k, v := range m {
		d.Fields[k] = DocumentField{
			Type:  getFieldType(v),
			Value: v,
		}
	}

	return &d, err
}

func UnmarshalDocument(doc *Document, output any) error {
	m := make(map[string]any)

	for k, v := range doc.Fields {
		m[k] = v.Value
	}

	jsonData, err := json.Marshal(m)

	if err != nil {
		return err
	}

	err = json.Unmarshal(jsonData, &output)

	if err != nil {
		return err
	}

	return nil
}

func (d Document) validate() error {
	for _, field := range d.Fields {
		switch field.Type {
		case DocumentFieldTypeString:
			if _, ok := field.Value.(string); !ok {
				return ErrInvalidTypeString
			}
		case DocumentFieldTypeNumber:
			if !IsValidNumber(field.Value) {
				return ErrInvalidTypeNumber
			}
		case DocumentFieldTypeBool:
			if _, ok := field.Value.(bool); !ok {
				return ErrInvalidTypeBool
			}
		case DocumentFieldTypeArray:
			if !IsValidSlice(field.Value) {
				return ErrInvalidTypeArray
			}
		case DocumentFieldTypeObject:
			if !IsValidMap(field.Value) {
				return ErrInvalidTypeObject
			}
		}
	}

	return nil
}

func getFieldType(input any) DocumentFieldType {
	switch input.(type) {
	case string:
		return DocumentFieldTypeString
	case float64:
		return DocumentFieldTypeNumber
	case bool:
		return DocumentFieldTypeBool
	default:
		return DocumentFieldTypeObject
	}
}
