// test/e2e/collection_test.go
package e2e

import (
	"lesson_12/cmd/server/responses"
	"lesson_12/internal/commands"
	"testing"
	"time"
)

func TestCommands(t *testing.T) {
	// Start test server
	serverAddr := "localhost:9090"
	stopServer := StartTestServer(t, serverAddr)
	defer stopServer()

	// Waiting to server start
	time.Sleep(100 * time.Millisecond)

	// Create test client
	client := NewTestClient(t, serverAddr)
	defer client.Close()

	// Test 1: Create collection
	collectionName := "test_collection"
	createResp := client.SendCommand(t, commands.CreateCollectionCommand,
		`{"name": "`+collectionName+`", "primary_key": "id"}`)

	if createResp.Status != responses.StatusOk {
		t.Errorf("create collection: Expected status OK, recieved %s, error %v", createResp.Status, createResp.Error)
	}

	// Test 2: Check collection list
	listResp := client.SendCommand(t, commands.ListCollectionsCommand, "")
	if listResp.Status != responses.StatusOk {
		t.Errorf("list collections: Expected status OK, recieved %s, error %v", listResp.Status, listResp.Error)
	}

	// Check collection is present in list
	collectionFound := false
	if collections, ok := listResp.Data.(map[string]any); ok {
		for name, _ := range collections {
			if name == collectionName {
				collectionFound = true
				break
			}
		}
	}

	if !collectionFound {
		t.Errorf("Collection %s was not found in list", collectionName)
	}

	// Test 3: Select collection and add document
	selectResp := client.SendCommand(t, commands.SelectCollectionCommand, `{"name":"test_collection"}`)
	if selectResp.Status != responses.StatusOk {
		t.Errorf("select collection: Expected status OK, recieved %s, error %v", selectResp.Status, selectResp.Error)
	}

	documentID := "1"
	documentName := "first doc"
	putDocResp := client.SendCommand(t, commands.PutDocumentCommand, `{"id":"`+documentID+`", "name":"`+documentName+`"}`)

	if putDocResp.Status != responses.StatusOk {
		t.Errorf("put document: Expected status OK, recieved %s, error %v", putDocResp.Status, putDocResp.Error)
	}

	assertDocumentResponse(t, putDocResp, commands.PutDocumentCommand, documentID, documentName)

	// Test 4: List document should return list with document
	listDocResp := client.SendCommand(t, commands.ListDocumentsCommand, "")
	if listDocResp.Status != responses.StatusOk {
		t.Errorf("list documents: Expected status OK, recieved %s, error %v", listDocResp.Status, listDocResp.Error)
	}

	found, foundName := findDocumentInList(t, listDocResp.Data, documentID)
	if !found {
		t.Errorf("list documents: Document with id=%s not found in list", documentID)
	} else if foundName != documentName {
		t.Errorf("list documents: Expected name=%s, got %s", documentName, foundName)
	}

	// Test 5: Get document test
	getDocResp := client.SendCommand(t, commands.GetDocumentCommand, `{"id":"`+documentID+`"}`)
	if getDocResp.Status != responses.StatusOk {
		t.Errorf("get document: Expected status OK, recieved %s, error %v", getDocResp.Status, getDocResp.Error)
	}

	assertDocumentResponse(t, getDocResp, commands.GetDocumentCommand, documentID, documentName)

	// Test 6: Delete document test
	delDocResp := client.SendCommand(t, commands.DeleteDocumentCommand, `{"id":"`+documentID+`"}`)
	if delDocResp.Status != responses.StatusOk {
		t.Errorf("get document: Expected status OK, recieved %s, error %v", delDocResp.Status, delDocResp.Error)
	}

	if delDocResp.Data != nil {
		t.Errorf("delete document: Expected no data, recieved %v", delDocResp.Data)
	}

	// List document should return empty map
	listDocResp = client.SendCommand(t, commands.ListDocumentsCommand, "")
	if listDocResp.Status != responses.StatusOk {
		t.Errorf("list documents: Expected status OK, recieved %s, error %v", listDocResp.Status, listDocResp.Error)
	}

	found, _ = findDocumentInList(t, listDocResp.Data, documentID)
	if found {
		t.Errorf("list documents after delete: Document with id=%s should not be in list after deletion", documentID)
	}
}

func assertDocumentResponse(t *testing.T, rsp responses.ServerResponse, cmd string, docID string, docName string) {
	if rsp.Data == nil {
		t.Errorf("%q: Expected document in response Data, got nil", cmd)
	} else {
		docMap, ok := rsp.Data.(map[string]interface{})
		if !ok {
			t.Fatalf("%q: Expected map in Data, got %T", cmd, rsp.Data)
		}

		// Get doc Fields
		fields, ok := docMap["Fields"].(map[string]interface{})
		if !ok {
			t.Fatalf("%q: Expected Fields map, got %T", cmd, docMap["Fields"])
		}

		// Check field id
		idField, ok := fields["id"].(map[string]interface{})
		if !ok {
			t.Fatalf("%q: Expected id field map, got %T", cmd, fields["id"])
		}

		if idValue, ok := idField["Value"].(string); !ok || idValue != docID {
			t.Errorf("%q: Expected id Value=%s, got %v", cmd, docID, idField["Value"])
		}

		// Check field name
		nameField, ok := fields["name"].(map[string]interface{})
		if !ok {
			t.Fatalf("%q: Expected name field map, got %T", cmd, fields["name"])
		}

		if nameValue, ok := nameField["Value"].(string); !ok || nameValue != docName {
			t.Errorf("%q: Expected name Value=%s, got %v", cmd, docName, nameField["Value"])
		}
	}
}

func findDocumentInList(t *testing.T, docs interface{}, docID string) (found bool, docName string) {
	docsList, ok := docs.([]interface{})
	if !ok {
		t.Logf("findDocumentInList: Expected slice, got %T", docs)
		return false, ""
	}

	for _, doc := range docsList {
		docMap, ok := doc.(map[string]interface{})
		if !ok {
			continue
		}

		fields, ok := docMap["Fields"].(map[string]interface{})
		if !ok {
			continue
		}

		idField, ok := fields["id"].(map[string]interface{})
		if !ok {
			continue
		}

		if idValue, ok := idField["Value"].(string); ok && idValue == docID {

			nameField, ok := fields["name"].(map[string]interface{})
			if !ok {
				return true, ""
			}
			nameValue, _ := nameField["Value"].(string)
			return true, nameValue
		}
	}

	return false, ""
}
