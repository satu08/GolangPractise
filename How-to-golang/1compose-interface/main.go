package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
)

func main() {
	payload := []byte("satyanarayan jadhav")
	err := hashBroadcast(newHashReader(payload))
	if err != nil {
		return
	}
}

type HashReader interface {
	io.Reader
	hash() string
}

type hashReader struct {
	*bytes.Reader
	buf *bytes.Buffer
}

func (h *hashReader) hash() string {
	return hex.EncodeToString(h.buf.Bytes())
}
func newHashReader(b []byte) *hashReader {
	return &hashReader{bytes.NewReader(b), bytes.NewBuffer(b)}
}

func hashBroadcast(r HashReader) error {
	hash := r.hash()
	fmt.Println(hash)
	return broadcast(r)
}

// since we have already used the reader in the above function, so when we pass the reader here, it has nothing to solve this problem, we are using compose interface
func broadcast(r io.Reader) error {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	fmt.Println("string of bytes-", string(b))
	return nil
}
