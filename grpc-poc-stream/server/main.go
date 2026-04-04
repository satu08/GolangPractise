// // package main

// // import (
// // 	"fmt"
// // 	"io/ioutil"
// // 	"strings"

// // 	"github.com/saintfish/chardet"
// // 	"golang.org/x/text/encoding"
// // 	"golang.org/x/text/encoding/charmap"
// // 	"golang.org/x/text/encoding/japanese"
// // 	"golang.org/x/text/encoding/korean"
// // 	"golang.org/x/text/encoding/simplifiedchinese"
// // 	"golang.org/x/text/encoding/traditionalchinese"
// // 	"golang.org/x/text/transform"
// // )

// // func detectEncoding(input []byte) (*chardet.Result, error) {
// // 	detector := chardet.NewTextDetector()
// // 	return detector.DetectBest(input)
// // }

// // func getDecoder(charset string) (*encoding.Decoder, error) {
// // 	switch strings.ToLower(charset) {
// // 	case "utf-8":
// // 		return nil, nil
// // 	case "iso-8859-1":
// // 		return charmap.ISO8859_1.NewDecoder(), nil
// // 	case "windows-1252":
// // 		return charmap.Windows1252.NewDecoder(), nil
// // 	case "shift_jis", "sjis":
// // 		return japanese.ShiftJIS.NewDecoder(), nil
// // 	case "euc-jp":
// // 		return japanese.EUCJP.NewDecoder(), nil
// // 	case "iso-2022-jp":
// // 		return japanese.ISO2022JP.NewDecoder(), nil
// // 	case "euc-kr":
// // 		return korean.EUCKR.NewDecoder(), nil
// // 	case "gb18030":
// // 		return simplifiedchinese.GB18030.NewDecoder(), nil
// // 	case "gbk":
// // 		return simplifiedchinese.GBK.NewDecoder(), nil
// // 	case "big5":
// // 		return traditionalchinese.Big5.NewDecoder(), nil
// // 	default:
// // 		return nil, fmt.Errorf("unsupported charset: %s", charset)
// // 	}
// // }

// // func convertToUTF8(input []byte) (string, error) {
// // 	// Step 1: Detect encoding
// // 	result, err := detectEncoding(input)
// // 	if err != nil {
// // 		return "", err
// // 	}
// // 	fmt.Printf("Detected charset: %s (confidence: %d%%)\n", result.Charset, result.Confidence)

// // 	// Step 2: Get appropriate decoder
// // 	decoder, err := getDecoder(result.Charset)
// // 	if err != nil {
// // 		return "", err
// // 	}

// // 	// If UTF-8, return directly
// // 	if decoder == nil {
// // 		return string(input), nil
// // 	}

// // 	// Step 3: Decode
// // 	reader := transform.NewReader(strings.NewReader(string(input)), decoder)
// // 	decoded, err := ioutil.ReadAll(reader)
// // 	if err != nil {
// // 		return "", err
// // 	}
// // 	return string(decoded), nil
// // }

// // func main() {
// // 	// Example input: You can load from a file or use real encoded bytes
// // 	input := []byte("あいう") // Japanese Shift_JIS example

// // 	utf8Text, err := convertToUTF8(input)
// // 	if err != nil {
// // 		fmt.Println("Error:", err)
// // 		return
// // 	}
// // 	fmt.Println("Converted to UTF-8:", utf8Text)
// // }

// package main

// import (
// 	"bytes"
// 	"fmt"
// 	"unicode/utf8"

// 	"github.com/juneysec/jpdec"
// 	"golang.org/x/text/encoding/japanese"
// 	"golang.org/x/text/transform"
// )

// func main() {

// 	shiftJISBuf := new(bytes.Buffer)
// 	writer := transform.NewWriter(shiftJISBuf, japanese.ShiftJIS.NewEncoder())
// 	writer.Write([]byte("さようなら�"))
// 	writer.Close()

// 	corruptedData := shiftJISBuf.Bytes()
// 	//byteArray := []byte("さようなら")
// 	// Example JIS encoded byte slice

// 	fmt.Printf("Original byte array:%X\n", corruptedData)
// 	enc, err := jpdec.DetectEncoding(corruptedData)
// 	if err != nil {
// 		fmt.Println("Error detecting encoding:", err)
// 		return
// 	}
// 	fmt.Println("Detected encoding:", enc)

// 	decodedStr, err := jpdec.Decode(corruptedData)
// 	if err != nil {
// 		fmt.Println("Error decoding:", err)
// 		return
// 	}
// 	fmt.Println("Decoded string:", decodedStr)

// 	yes := utf8.ValidString(decodedStr)
// 	if !yes {
// 		fmt.Println("Decoded string is not valid UTF-8", decodedStr)
// 	} else {
// 		fmt.Println("Decoded string is valid UTF-8", decodedStr)
// 	}

// }

// type MyService_SendStringServer interface {
// 	Send(*MyStringMessage) error
// }
// type MyStringMessage struct {
// 	Value string
// }

// func sendDecodedStringToStream(stream MyService_SendStringServer, decodedStr string) error {
// 	msg := &MyStringMessage{Value: decodedStr}
// 	return stream.Send(msg)
// }

package main

import (
	"fmt"
	"io"
	"log"
	"net"

	"google.golang.org/grpc"
	// Import your generated protobuf package
	pb "grpc/grpcfile"
)

type server struct {
	pb.UnimplementedStringStreamServiceServer
}

func (s *server) SendStrings(stream pb.StringStreamService_SendStringsServer) error {
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.StreamResponse{Status: "All strings received"})
		}
		if err != nil {
			return err
		}
		fmt.Printf("Received: %s\n", msg.Data)
	}
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	pb.RegisterStringStreamServiceServer(s, &server{})

	fmt.Println("gRPC server listening on :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
