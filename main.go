package main

import (
	"fmt"
	"learningGo/documentstore"
)

func main() {
	store := documentstore.NewStore()
	collectionConfig := &documentstore.CollectionConfig{PrimaryKey: "key"}

	isCreated, collection := store.CreateCollection("my_first_collection", collectionConfig)

	if !isCreated {
		fmt.Println("Collection was not created")
		return
	}

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

	if err := collection.Put(document1); err != nil {
		fmt.Println(err)
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
				Value: map[string]any{"SomeFiled": "someValue"},
			},
		},
	}

	if err := collection.Put(document2); err != nil {
		fmt.Println(err)
	}

	allDocuments := collection.List()

	fmt.Println(allDocuments)

	store.CreateCollection("my_second_collection", collectionConfig)

	secondCollection, _ := store.GetCollection("my_second_collection")

	doc3 := documentstore.Document{
		Fields: map[string]documentstore.DocumentField{
			"key": {
				Type:  documentstore.DocumentFieldTypeString,
				Value: "key3",
			},
			"boolField": {
				Type:  documentstore.DocumentFieldTypeBool,
				Value: true,
			},
		},
	}

	if err := secondCollection.Put(doc3); err != nil {
		fmt.Println(err)
	}

	fmt.Println(secondCollection.List())

	secondCollection.Delete("key3")

	fmt.Println(secondCollection.List())
}
