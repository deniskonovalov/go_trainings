package documentstore

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
