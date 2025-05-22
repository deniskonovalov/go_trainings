package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"net"
	"os"
)

func main() {
	us := bufio.NewScanner(os.Stdin)
	uw := bufio.NewWriter(os.Stdout)

	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	defer func() {
		err := conn.Close()
		if err != nil {
			panic(err)
		}
	}()

	cw := bufio.NewWriter(conn)
	cr := bufio.NewReader(conn)

	for us.Scan() {
		msg := us.Text()

		_, err = cw.WriteString(base64.StdEncoding.EncodeToString([]byte(msg)) + "\n")
		if err != nil {
			panic(fmt.Errorf("error writing command: %w", err))
		}

		err = cw.Flush()
		if err != nil {
			panic(fmt.Errorf("error flushing command: %w", err))
		}

		resp, _ := cr.ReadString('\n')

		decodedBytes, err := base64.StdEncoding.DecodeString(resp)
		decoded := string(decodedBytes)

		if err != nil {
			panic(fmt.Errorf("error decoding response: %w", err))
		}

		_, _ = uw.WriteString(decoded + "\n")
		_ = uw.Flush()
	}
}
