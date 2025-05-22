package e2e

import (
	"bufio"
	"encoding/base64"
	"encoding/json"
	"lesson_13/cmd/server/responses"
	"net"
	"testing"
)

// TestClient represents the test client for connection with the test server
type TestClient struct {
	conn   net.Conn
	writer *bufio.Writer
	reader *bufio.Reader
}

// NewTestClient creates new test client
func NewTestClient(t *testing.T, address string) *TestClient {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		t.Fatalf("Не вдалося з'єднатися з сервером: %v", err)
	}

	return &TestClient{
		conn:   conn,
		writer: bufio.NewWriter(conn),
		reader: bufio.NewReader(conn),
	}
}

// SendCommand sends command to the server and returns the response
func (c *TestClient) SendCommand(t *testing.T, command string, jsonData string) responses.ServerResponse {

	msg := command
	if jsonData != "" {
		msg += " " + jsonData
	}

	_, err := c.writer.WriteString(base64.StdEncoding.EncodeToString([]byte(msg)) + "\n")
	if err != nil {
		t.Fatalf("Помилка відправки команди: %v", err)
	}

	err = c.writer.Flush()
	if err != nil {
		t.Fatalf("Помилка при промиванні буфера: %v", err)
	}

	resp, err := c.reader.ReadString('\n')
	if err != nil {
		t.Fatalf("Помилка при читанні відповіді: %v", err)
	}

	decodedBytes, err := base64.StdEncoding.DecodeString(resp)
	if err != nil {
		t.Fatalf("Помилка декодування відповіді: %v", err)
	}

	var serverResponse responses.ServerResponse
	err = json.Unmarshal(decodedBytes, &serverResponse)
	if err != nil {
		t.Fatalf("Помилка парсингу відповіді: %v", err)
	}

	return serverResponse
}

func (c *TestClient) Close() {
	if c.conn != nil {
		c.conn.Close()
	}
}
