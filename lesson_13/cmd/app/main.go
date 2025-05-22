package main

import (
	"fmt"
	"lesson_13/internal/documentstore"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

var randPool = sync.Pool{
	New: func() any {
		seed := time.Now().UnixNano()
		return rand.New(rand.NewSource(seed))
	},
}

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
		return
	}

	if err := c.CreateIndex("name"); err != nil {
		fmt.Println("Index was not created")
		return
	}

	var wg sync.WaitGroup
	numWorkers := 1000

	worker := func(id int) {
		defer wg.Done()

		r := randPool.Get().(*rand.Rand)
		defer randPool.Put(r)

		docID := strconv.Itoa(id)
		docName := fmt.Sprintf("User%03d", r.Intn(numWorkers))

		doc := documentstore.Document{
			Fields: map[string]documentstore.DocumentField{
				"id": {
					Type:  documentstore.DocumentFieldTypeString,
					Value: docID,
				},
				"name": {
					Type:  documentstore.DocumentFieldTypeString,
					Value: docName,
				},
			},
		}

		_ = c.Put(doc)

		if id%10 == 0 {
			_, _ = c.Get(docID)
		}

		if id%25 == 0 {
			_ = c.Delete(docID)
		}

		if id%20 == 0 {
			minVal := fmt.Sprintf("Name%03d", r.Intn(500))
			maxVal := fmt.Sprintf("Name%03d", 500+r.Intn(500))

			qp := documentstore.QueryParams{
				Desc:     false,
				MinValue: &minVal,
				MaxValue: &maxVal,
			}
			_, _ = c.Query("name", qp)
		}
	}

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(i)
	}

	wg.Wait()
	fmt.Println("All goroutines completed.")
}
