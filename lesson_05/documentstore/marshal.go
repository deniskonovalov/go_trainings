package documentstore

import (
	"encoding/json"
)

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
