package main

import (
	"fmt"
	"lesson_09/internal/documentstore"
)

func main() {
	store := documentstore.NewStore()
	collectionConfig := &documentstore.CollectionConfig{PrimaryKey: "id"}

	c, err := store.CreateCollection("documents", collectionConfig)

	if err != nil {
		fmt.Println("Collection was not created")
		return
	}

	if err := c.CreateIndex("id"); err != nil {
		fmt.Println("Index was not created")
	}

	d1 := documentstore.Document{
		Fields: map[string]documentstore.DocumentField{
			"id": {
				Type:  documentstore.DocumentFieldTypeString,
				Value: "1",
			},
			"name": {
				Type:  documentstore.DocumentFieldTypeString,
				Value: "Ronald",
			},
			"value": {
				Type:  documentstore.DocumentFieldTypeString,
				Value: "100",
			},
		},
	}

	d2 := documentstore.Document{
		Fields: map[string]documentstore.DocumentField{
			"id": {
				Type:  documentstore.DocumentFieldTypeString,
				Value: "2",
			},
			"name": {
				Type:  documentstore.DocumentFieldTypeString,
				Value: "Max",
			},
			"value": {
				Type:  documentstore.DocumentFieldTypeString,
				Value: "80",
			},
		},
	}

	d3 := documentstore.Document{
		Fields: map[string]documentstore.DocumentField{
			"id": {
				Type:  documentstore.DocumentFieldTypeString,
				Value: "3",
			},
			"name": {
				Type:  documentstore.DocumentFieldTypeString,
				Value: "Donna",
			},
			"value": {
				Type:  documentstore.DocumentFieldTypeString,
				Value: "120",
			},
		},
	}

	d4 := documentstore.Document{
		Fields: map[string]documentstore.DocumentField{
			"id": {
				Type:  documentstore.DocumentFieldTypeString,
				Value: "4",
			},
			"name": {
				Type:  documentstore.DocumentFieldTypeString,
				Value: "Bruce",
			},
			"value": {
				Type:  documentstore.DocumentFieldTypeString,
				Value: "20",
			},
		},
	}

	d1New := documentstore.Document{
		Fields: map[string]documentstore.DocumentField{
			"id": {
				Type:  documentstore.DocumentFieldTypeString,
				Value: "1",
			},
			"name": {
				Type:  documentstore.DocumentFieldTypeString,
				Value: "Aaron",
			},
			"value": {
				Type:  documentstore.DocumentFieldTypeString,
				Value: "100",
			},
		},
	}

	if err := c.Put(d1); err != nil {
		fmt.Println(err)
	}

	if err := c.Put(d2); err != nil {
		fmt.Println(err)
	}

	if err := c.Put(d3); err != nil {
		fmt.Println(err)
	}

	if err := c.Put(d4); err != nil {
		fmt.Println(err)
	}

	if err := c.CreateIndex("value"); err != nil {
		fmt.Println("Index was not created")
	}

	if err := c.CreateIndex("name"); err != nil {
		fmt.Println("Index was not created")
	}

	if err := c.Put(d1New); err != nil {
		fmt.Println(err)
	}

	minVal := d1New.Fields["name"].Value.(string)
	maxVal := d2.Fields["name"].Value.(string)

	qp := documentstore.QueryParams{
		Desc:     false,
		MinValue: &minVal,
		MaxValue: &maxVal,
	}

	list, err := c.Query("name", qp)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%+v\n", list)

	err = store.DumpToFile("documents.dump")
	if err != nil {
		fmt.Println(err)
	}

	newStore, err := documentstore.NewStoreFromFile("documents.dump")
	if err != nil {
		fmt.Println(err)
		return
	}

	newCol, err := newStore.GetCollection("documents")
	if err != nil {
		fmt.Println(err)
		return
	}

	newList, err := newCol.Query("name", qp)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%+v\n", newList)
}
