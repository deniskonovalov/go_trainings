package main

import (
	"bufio"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"lesson_12/cmd/server/handlers"
	"lesson_12/cmd/server/responses"
	"net"
	"regexp"
	"strings"

	"lesson_12/internal/commands"
	ds "lesson_12/internal/documentstore"
)

var store = ds.NewStore()

func main() {
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := l.Accept()

		handler := handlers.NewHandler(store, conn, nil)

		if err != nil {
			panic(err)
		}

		go handleConnection(handler)
	}
}

func handleConnection(h *handlers.Handler) {

	defer func() {
		err := h.Conn.Close()
		if err != nil {
			panic(err)
		}
	}()

	scanner := bufio.NewScanner(h.Conn)
	writer := bufio.NewWriter(h.Conn)

	for scanner.Scan() {
		msg := scanner.Text()

		decodedMsg, err := base64.StdEncoding.DecodeString(msg)
		if err != nil {
			panic(err)
		}

		decoded := string(decodedMsg)
		elems := strings.Split(decoded, "{")
		command := strings.Trim(elems[0], " ")

		fmt.Printf("command: %s\n", command)

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
			panic(err)
		}

		resp := base64.StdEncoding.EncodeToString(resultToJson) + "\n"

		_, _ = writer.WriteString(resp)
		_ = writer.Flush()
	}
}
