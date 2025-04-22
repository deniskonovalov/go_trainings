package documentstore

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func getTestConfig() CollectionConfig {
	return CollectionConfig{
		PrimaryKey: "id",
	}
}

func Test_NewStore(t *testing.T) {
	s := NewStore()

	assert.NotNil(t, s)
	assert.IsType(t, &Store{}, s)
	assert.IsType(t, map[string]*Collection{}, s.Collections)
	assert.Empty(t, s.Collections)
}

func Test_CreateCollection(t *testing.T) {
	s := NewStore()

	cnf := getTestConfig()

	coll, err := s.CreateCollection("test", &cnf)

	assert.Nil(t, err)
	assert.IsType(t, &Collection{}, coll)
	assert.IsType(t, map[string]Document{}, coll.items)
	assert.Empty(t, coll.items)
	assert.Equal(t, cnf, coll.config)
}

func Test_CreateCollection_ErrCollectionAlreadyExists(t *testing.T) {
	s := NewStore()

	cnf := getTestConfig()

	s.Collections["test"] = &Collection{
		items:  map[string]Document{},
		config: cnf,
	}

	coll, err := s.CreateCollection("test", &cnf)

	assert.NotNil(t, err)
	assert.Nil(t, coll)
	assert.Equal(t, ErrCollectionAlreadyExists, err)
}

func Test_GetCollection(t *testing.T) {
	s := NewStore()

	s.Collections["test"] = &Collection{
		items:  map[string]Document{},
		config: getTestConfig(),
	}

	coll, err := s.GetCollection("test")
	assert.Nil(t, err)
	assert.IsType(t, &Collection{}, coll)
	assert.Equal(t, s.Collections["test"], coll)
}

func Test_GetCollection_ErrCollectionNotFound(t *testing.T) {
	s := NewStore()

	coll, err := s.GetCollection("test")

	assert.NotNil(t, err)
	assert.Nil(t, coll)
	assert.Equal(t, ErrCollectionNotFound, err)
}

func Test_DeleteCollection(t *testing.T) {
	s := NewStore()
	s.Collections["test"] = &Collection{
		items:  map[string]Document{},
		config: getTestConfig(),
	}

	err := s.DeleteCollection("test")

	assert.Nil(t, err)
	assert.NotContains(t, s.Collections, "test")
}

func Test_DeleteCollection_ErrCollectionNotFound(t *testing.T) {
	s := NewStore()
	s.Collections["test"] = &Collection{
		items:  map[string]Document{},
		config: getTestConfig(),
	}

	err := s.DeleteCollection("test2")
	assert.NotNil(t, err)
	assert.Contains(t, s.Collections, "test")
	assert.Equal(t, ErrCollectionNotFound, err)
}

func Test_Dump_Empty_Store(t *testing.T) {
	s := NewStore()

	data, err := s.Dump()

	assert.Nil(t, err)
	assert.NotNil(t, data)
	assert.NotEmpty(t, data)
	assert.IsType(t, []byte{}, data)
}

func Test_Dump_Store_With_Collections(t *testing.T) {
	s := NewStore()

	cfg := getTestConfig()
	collection, err := s.CreateCollection("test_collection", &cfg)
	assert.Nil(t, err)

	doc := Document{
		Fields: map[string]DocumentField{
			"id": {
				Type:  DocumentFieldTypeString,
				Value: "123",
			},
			"name": {
				Type:  DocumentFieldTypeString,
				Value: "Test Document",
			},
		},
	}

	err = collection.Put(doc)
	assert.Nil(t, err)

	data, err := s.Dump()

	assert.Nil(t, err)
	assert.NotNil(t, data)
	assert.NotEmpty(t, data)
}

func Test_NewStoreFromDump_Empty_Store(t *testing.T) {
	s := NewStore()
	data, err := s.Dump()

	assert.Nil(t, err)
	assert.NotNil(t, data)
	assert.NotEmpty(t, data)

	restoredStore, err := NewStoreFromDump(data)

	assert.Nil(t, err)
	assert.NotNil(t, restoredStore)
	assert.Equal(t, 0, len(restoredStore.Collections))
}

func Test_NewStoreFromDump_Store_With_Collections(t *testing.T) {
	s := NewStore()
	cfg := getTestConfig()

	collection, err := s.CreateCollection("test_collection", &cfg)
	assert.Nil(t, err)

	doc := Document{
		Fields: map[string]DocumentField{
			"id": {
				Type:  DocumentFieldTypeString,
				Value: "123",
			},
			"name": {
				Type:  DocumentFieldTypeString,
				Value: "Test Document",
			},
		},
	}
	err = collection.Put(doc)
	assert.Nil(t, err)

	dump, err := s.Dump()
	assert.Nil(t, err)

	restoredStore, err := NewStoreFromDump(dump)
	assert.Nil(t, err)
	assert.NotNil(t, restoredStore)

	assert.Equal(t, 1, len(restoredStore.Collections))
	restoredCollection, exists := restoredStore.Collections["test_collection"]
	assert.True(t, exists)
	assert.NotNil(t, restoredCollection)

	restoredDoc, err := restoredCollection.Get("123")
	assert.Nil(t, err)
	assert.NotNil(t, restoredDoc)
	assert.Equal(t, "123", restoredDoc.Fields["id"].Value)
	assert.Equal(t, "Test Document", restoredDoc.Fields["name"].Value)
}

func Test_NewStoreFromDump_Invalid_Data(t *testing.T) {
	invalidData := []byte("invalid serialized data")
	restoredStore, err := NewStoreFromDump(invalidData)

	assert.NotNil(t, err)
	assert.Nil(t, restoredStore)
}

func Test_DumpToFile(t *testing.T) {
	s := NewStore()
	cfg := getTestConfig()

	collection, err := s.CreateCollection("test_collection", &cfg)
	assert.Nil(t, err)

	doc := Document{
		Fields: map[string]DocumentField{
			"id": {
				Type:  DocumentFieldTypeString,
				Value: "123",
			},
			"name": {
				Type:  DocumentFieldTypeString,
				Value: "Test Document",
			},
		},
	}

	err = collection.Put(doc)
	assert.Nil(t, err)

	tmpFile, err := os.CreateTemp("", "store_dump.dump")
	assert.Nil(t, err)

	defer func(name string) {
		err := os.Remove(name)
		if err != nil {

		}
	}(tmpFile.Name())

	err = s.DumpToFile(tmpFile.Name())
	assert.Nil(t, err)

	_, err = os.Stat(tmpFile.Name())
	assert.Nil(t, err)
}

func Test_NewStoreFromFile(t *testing.T) {
	s := NewStore()
	cfg := getTestConfig()

	collection, err := s.CreateCollection("test_collection", &cfg)
	assert.Nil(t, err)

	doc := Document{
		Fields: map[string]DocumentField{
			"id": {
				Type:  DocumentFieldTypeString,
				Value: "123",
			},
			"name": {
				Type:  DocumentFieldTypeString,
				Value: "Test Document",
			},
		},
	}
	err = collection.Put(doc)
	assert.Nil(t, err)

	tmpFile, err := os.CreateTemp("", "store_dump.dump")
	assert.Nil(t, err)

	defer func(name string) {
		err := os.Remove(name)
		if err != nil {

		}
	}(tmpFile.Name())

	err = s.DumpToFile(tmpFile.Name())
	assert.Nil(t, err)

	restoredStore, err := NewStoreFromFile(tmpFile.Name())
	assert.Nil(t, err)
	assert.NotNil(t, restoredStore)

	assert.Equal(t, 1, len(restoredStore.Collections))
	restoredCollection, exists := restoredStore.Collections["test_collection"]
	assert.True(t, exists)
	assert.NotNil(t, restoredCollection)

	restoredDoc, err := restoredCollection.Get("123")
	assert.Nil(t, err)
	assert.Equal(t, "123", restoredDoc.Fields["id"].Value)
	assert.Equal(t, "Test Document", restoredDoc.Fields["name"].Value)
}

func Test_NewStoreFromFile_FileNotFound(t *testing.T) {
	restoredStore, err := NewStoreFromFile("non_existing_file.dump")

	assert.NotNil(t, err)
	assert.Nil(t, restoredStore)
}
