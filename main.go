package main

import (
	"fmt"
	"learningGo/documentstore"
)

func main() {
	document1 := documentstore.Document{
		Fields: map[string]documentstore.DocumentField{
			"key": {
				Type:  documentstore.DocumentFieldTypeString,
				Value: "key1",
			},
			"boolField": {
				Type:  documentstore.DocumentFieldTypeBool,
				Value: true,
			},
			"numberField": {
				Type:  documentstore.DocumentFieldTypeNumber,
				Value: 123,
			},
		},
	}

	document2 := documentstore.Document{
		Fields: map[string]documentstore.DocumentField{
			"key": {
				Type:  documentstore.DocumentFieldTypeString,
				Value: "key2",
			},
			"arrayField": {
				Type:  documentstore.DocumentFieldTypeArray,
				Value: []int{1, 2, 3},
			},
			"objectField": {
				Type:  documentstore.DocumentFieldTypeObject,
				Value: document1,
			},
		},
	}

	if _, err := documentstore.Put(document1); err != nil {
		fmt.Println(err)
	}

	if _, err := documentstore.Put(document2); err != nil {
		fmt.Println(err)
	}

	documentWithoutKey := documentstore.Document{
		Fields: map[string]documentstore.DocumentField{
			"someField": {
				Type:  documentstore.DocumentFieldTypeString,
				Value: "Some String Value",
			},
		},
	}

	if _, err := documentstore.Put(documentWithoutKey); err != nil {
		fmt.Println(err)
	}

	foundDocument, _ := documentstore.Get("key1")
	fmt.Println(foundDocument)

	getByIncorrectKey, _ := documentstore.Get("incorrectKey")
	fmt.Println(getByIncorrectKey)

	listOfDocuments := documentstore.List()
	fmt.Println(listOfDocuments)

	documentstore.Delete("key1")

	listOfDocuments = documentstore.List()
	fmt.Println(listOfDocuments)
}
