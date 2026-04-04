package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"strings"
	"time"
	"unicode/utf8"

	pb "grpc/grpcfile"

	"github.com/juneysec/jpdec"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewStringStreamServiceClient(conn)
	stream, err := client.SendStrings(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	shiftJISBuf := new(bytes.Buffer)
	writer := transform.NewWriter(shiftJISBuf, japanese.ShiftJIS.NewEncoder())
	writer.Write([]byte("こんにちはこんにちは"))
	writer.Close()

	//corruptedData := shiftJISBuf.Bytes()
	decodedStr := utf8DecodeString(string("こんにちは�"))
	//decodedStr1 := utf8DecodeString("ä")
	// Send strings
	stringsToSend := []string{decodedStr}
	for _, str := range stringsToSend {
		if err := stream.Send(&pb.StringMessage{Data: str}); err != nil {
			log.Fatal(err)
		}
		time.Sleep(500 * time.Millisecond) // simulate streaming
	}

	resp, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Server response: %s", resp.Status)
}

func utf8DecodeString(s string) string {
	fmt.Println("Original byte array:", []byte(s))
	fmt.Println("Original string:", s)
	validBytes := replaceReplacementChar(s)
	fmt.Println("Valid byte array:", []byte(validBytes))
	fmt.Println("Valid string:", validBytes)
	enc, err := jpdec.DetectEncoding([]byte(validBytes))
	if err != nil {
		fmt.Println("Error detecting encoding:", err)
	}
	fmt.Println("Detected encoding:", enc)

	decodedStr, err := jpdec.Decode([]byte(validBytes))
	if err != nil {
		fmt.Println("Error decoding:", err)
	}
	fmt.Println("Decoded string:", decodedStr)

	yes := utf8.ValidString(decodedStr)
	if !yes {
		fmt.Println("Decoded string is not valid UTF-8", decodedStr)
	} else {
		fmt.Println("Decoded string is valid UTF-8", decodedStr)
	}
	return decodedStr
}

func replaceReplacementChar(s string) string {
	return strings.ReplaceAll(s, string('\uFFFD'), "?")
}
