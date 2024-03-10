package forward

import (
	"bufio"
	"bytes"
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"strconv"
	"strings"
)

func Handle_request(connection *net.Conn) error {
	// read request
	var buffer []byte = make([]byte, 4096)
	read_bytes, err := (*connection).Read(buffer)
	if err != nil {
		return err
	} else if read_bytes < 0 {
		return fmt.Errorf("Failed to read repsonse")
	}

	body := string(buffer)
	http_version, err := read_http_version(&body)
	if err != nil {
		return err
	}

	// forward request with correct http version and method
	var data []byte
	var req_err error
	if http_version >= 2 {
		// make http 2 request
		data, req_err = http1_request(buffer)
		if req_err != nil {
			return req_err
		}

	} else {
		// Make http 1 request
	}

	// send response from server to client

	return (*connection).Close()
}

func read_http_version(data *string) (float64, error) {

	var start_index int = strings.Index(*data, "HTTP/")
	if start_index < 0 {
		return -1.0, fmt.Errorf("Failed to find http version")
	}
	start_index += 5 // cut out HTTP/
	var end_index int = strings.Index(*data, "\r\n")
	if start_index >= end_index {
		return -1.0, fmt.Errorf("Failed to find http version")
	}
	http_version_str := (*data)[start_index:end_index]
	return strconv.ParseFloat(http_version_str, 64)
}
func http1_request(data []byte) ([]byte, error) {
	reader := bytes.NewReader(data)
	bufio_reader := bufio.NewReader(reader)
	req, err := http.ReadRequest(bufio_reader)
	if err != nil {
		return []byte{}, err
	}
	conn, err := tls.Dial("tcp", req.URL.Host)
	if err != nil {
		return []byte{}, err
	}
	_, err = conn.Write(data)
	if err != nil {
		return []byte{}, err
	}
	buffer := make([]byte, 4096)
	_, err = conn.Read(buffer)
	return buffer, err

}
