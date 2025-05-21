package e2e

import (
	"bufio"
	"encoding/base64"
	"encoding/json"
	"lesson_12/cmd/server/handlers"
	"lesson_12/cmd/server/responses"
	"lesson_12/internal/commands"
	ds "lesson_12/internal/documentstore"
	"net"
	"regexp"
	"strings"
	"testing"
)

// StartTestServer starts test server
func StartTestServer(t *testing.T, address string) (shutdown func()) {
	store := ds.NewStore()

	listener, err := net.Listen("tcp", address)
	if err != nil {
		t.Fatalf("Не вдалося запустити тестовий сервер: %v", err)
	}

	// Channel for stop
	stopChan := make(chan struct{})

	go func() {
		for {
			select {
			case <-stopChan:
				return
			default:
				conn, err := listener.Accept()
				if err != nil {
					return
				}

				handler := handlers.NewHandler(store, conn, nil)
				go handleTestConnection(handler)
			}
		}
	}()

	// Function for server stop
	return func() {
		close(stopChan)
		listener.Close()
	}
}

// handleTestConnection handles connections, copied from cmd/server/main.go
func handleTestConnection(h *handlers.Handler) {
	defer h.Conn.Close()

	scanner := bufio.NewScanner(h.Conn)
	writer := bufio.NewWriter(h.Conn)

	for scanner.Scan() {
		msg := scanner.Text()

		decodedMsg, err := base64.StdEncoding.DecodeString(msg)
		if err != nil {
			continue
		}

		decoded := string(decodedMsg)
		elems := strings.Split(decoded, "{")
		command := strings.Trim(elems[0], " ")

		re := regexp.MustCompile(`\{.*\}`)
		inputJson := re.FindString(decoded)

		var res responses.ServerResponse

		switch command {
		case commands.CreateCollectionCommand:
			res = h.HandleCreateCollection(inputJson)
		case commands.SelectCollectionCommand:
			res = h.HandleSelectCollection(inputJson)
		case commands.ListCollectionsCommand:
			res = h.HandleListCollections()
		case commands.DeleteCollectionCommand:
			res = h.HandleDeleteCollection(inputJson)
		case commands.PutDocumentCommand:
			res = h.HandlePutDocument(inputJson)
		case commands.GetDocumentCommand:
			res = h.HandleGetDocument(inputJson)
		case commands.DeleteDocumentCommand:
			res = h.HandleDeleteDocument(inputJson)
		case commands.ListDocumentsCommand:
			res = h.HandleListDocuments()
		default:
			res = responses.NewServerResponse(responses.StatusFailed, nil, ds.ErrUnknownCommand.Error())
		}

		resultToJson, err := json.Marshal(res)
		if err != nil {
			continue
		}

		resp := base64.StdEncoding.EncodeToString(resultToJson) + "\n"

		_, _ = writer.WriteString(resp)
		_ = writer.Flush()
	}
}
