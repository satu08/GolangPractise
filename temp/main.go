package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"
	"log"
	"strings"
	"unicode/utf16"
	"unicode/utf8"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

// func main() {

// 	hexStr := "1B 24 42 46 7C 4B 5C 1B 28 42"
// 	bytesData, err := hexStringToBytes(hexStr)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}
// 	fmt.Printf("Byte Data: %v\n", bytesData)
// 	var decodedString string
// 	err = stringDecoder(bytesData, &decodedString)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}
// }

// func stringDecoder(byteData []byte, v any) error {
// 	data := make([]byte, 0, len(byteData))
// 	data = append(data, byteData...)
// 	log.Print("stringDecoder: raw data: ", string(data))
// 	log.Printf("stringDecoder: raw data (hex): % X", data)
// 	strp := v.(*string)
// 	*strp = string(data)
// 	log.Printf("stringDecoder: decoded data: %s", *strp)
// 	return nil
// }

// func hexStringToBytes(hexStr string) ([]byte, error) {
// 	// Remove spaces
// 	cleanHex := strings.ReplaceAll(hexStr, " ", "")
// 	// Decode the hex string
// 	return hex.DecodeString(cleanHex)
// }

// func validateJSONLD(thing string, contextFile string) error {
// 	// Load the JSON-LD context
// 	contextData, err := os.ReadFile(contextFile)
// 	log.Printf("contextData: %s", contextData)
// 	if err != nil {
// 		return fmt.Errorf("failed to load context file: %w", err)
// 	}

// 	// Create a JSON-LD processor
// 	_ = ld.NewJsonLdProcessor()
// 	_ = ld.NewJsonLdOptions("")

// 	// Parse the Thing Description
// 	thingReader := bytes.NewReader([]byte(thing))
// 	thingData, err := ld.DocumentFromReader(thingReader)
// 	if err != nil {
// 		return fmt.Errorf("failed to parse Thing Description: %w", err)
// 	}

// 	return nil
// }

// func main() {
// 	thing := `{
//         "@context": [
//             "https://www.w3.org/2019/wot/td/v1",
//             {
//                 "bacv": "https://example.org/bacnet",
//                 "btzf": "http://bt.schema.siemens.io/shared/btzf#"
//             }
//         ],
//         "id": "urn:bacnet-1100_0",
//         "title": "BACnet-Edge-Router-1100"
//     }`

// 	contextFile := "resources/v1-1.json"

// 	err := validateJSONLD(thing, contextFile)
// 	if err != nil {
// 		log.Fatalf("Validation failed: %v", err)
// 	} else {
// 		log.Println("Thing Description is valid!")
// 	}
// }

// Encoding detection stub functions
func isAnsiX34(data []byte) bool {
	for _, b := range data {
		if b > 0x7F {
			return false
		}
	}
	return true
}

func isJisC6226(data []byte) bool {
	// for i := 0; i < len(data)-2; i++ {
	// 	if data[i] == 0x1B && data[i+1] == '$' && (data[i+2] == 'B' || data[i+2] == '@') {
	// 		return true
	// 	}
	// }
	// return false
	for len(data) > 0 {
		r, size := utf8.DecodeRune(data)
		if r == utf8.RuneError && size == 1 {
			// invalid UTF-8 sequence
			return false
		}

		// Check known Japanese character ranges
		switch {
		case r >= 0x3040 && r <= 0x309F: // Hiragana
			return true
		case r >= 0x30A0 && r <= 0x30FF: // Katakana
			return true
		case r >= 0x4E00 && r <= 0x9FFF: // CJK Unified Ideographs (Kanji)
			return true
		case r >= 0x3000 && r <= 0x303F: // Japanese punctuation and symbols
			return true
		}
		data = data[size:]
	}
	return false
}

func isIso88591(data []byte) bool {
	// for _, b := range data {
	// 	if (b >= 0x80 && b <= 0x9F) || b == 0xC0 || b == 0xC1 || b == 0xC2 || b == 0xC3 || b == 0xC4 || b == 0xC5 || b == 0xC6 || b == 0xC7 || b == 0xC8 || b == 0xC9 {
	// 		return true
	// 	}
	// }
	// return false

	for i := range data {
		// If the byte is an invalid UTF-8 start byte (but valid in ISO-8859-1)
		if data[i] >= 0x80 && data[i] <= 0x9F {
			return true // These are control characters in ISO-8859-1, invalid in UTF-8
		}
	}
	return false

}

// func isUcs2(data []byte) bool {
// 	if len(data) >= 2 {
// 		bom := uint16(data[0])<<8 | uint16(data[1])
// 		return bom == 0xFEFF || bom == 0xFFFE
// 	}
// 	return false
// }

func isUcs2(data []byte) bool {
	if len(data)%2 != 0 {
		return false
	}

	// BOM check
	if len(data) >= 2 {
		beBOM := data[0] == 0xFE && data[1] == 0xFF
		leBOM := data[0] == 0xFF && data[1] == 0xFE
		if beBOM {
			return validateUcs2(data[2:], true)
		}
		if leBOM {
			return validateUcs2(data[2:], false)
		}
	}

	// Heuristically try both
	return validateUcs2(data, true) || validateUcs2(data, false)
}

func validateUcs2(data []byte, bigEndian bool) bool {
	for i := 0; i < len(data); i += 2 {
		var cp uint16
		if bigEndian {
			cp = uint16(data[i])<<8 | uint16(data[i+1])
		} else {
			cp = uint16(data[i+1])<<8 | uint16(data[i])
		}

		// UCS-2 supports only code points in BMP (U+0000 to U+FFFF),
		// but must exclude surrogate pairs (U+D800–U+DFFF)
		if cp >= 0xD800 && cp <= 0xDFFF {
			return false
		}
	}

	return true
}

// func isUcs4(data []byte) bool {

// 	if len(data) >= 4 {
// 		beBOM := data[0] == 0x00 && data[1] == 0x00 && data[2] == 0xFE && data[3] == 0xFF
// 		leBOM := data[0] == 0xFF && data[1] == 0xFE && data[2] == 0x00 && data[3] == 0x00
// 		return beBOM || leBOM
// 	}
// 	return false
// }

func isUcs4(data []byte) bool {
	if len(data)%4 != 0 {
		return false
	}

	// Check for BOM first
	if len(data) >= 4 {
		beBOM := data[0] == 0x00 && data[1] == 0x00 && data[2] == 0xFE && data[3] == 0xFF
		leBOM := data[0] == 0xFF && data[1] == 0xFE && data[2] == 0x00 && data[3] == 0x00
		if beBOM {
			return validateUcs4(data[4:], true)
		}
		if leBOM {
			return validateUcs4(data[4:], false)
		}
	}

	// Heuristically try both and accept if either is valid
	return validateUcs4(data, true) || validateUcs4(data, false)
}

func validateUcs4(data []byte, bigEndian bool) bool {
	for i := 0; i < len(data); i += 4 {
		var cp uint32
		if bigEndian {
			cp = uint32(data[i])<<24 | uint32(data[i+1])<<16 | uint32(data[i+2])<<8 | uint32(data[i+3])
		} else {
			cp = uint32(data[i+3])<<24 | uint32(data[i+2])<<16 | uint32(data[i+1])<<8 | uint32(data[i])
		}

		// Valid Unicode code point range: U+0000 to U+10FFFF, excluding surrogates
		if cp > 0x10FFFF || (cp >= 0xD800 && cp <= 0xDFFF) {
			return false
		}
	}
	return true
}

// Main detection function
func detectEncoding(data []byte) string {
	fmt.Println("detectEncoding: raw data:", data)
	log.Printf("detectEncoding: raw data: % X", data)
	switch {
	// case isUcs4(data):
	// 	return "UCS-4"
	// case isUcs2(data):
	// 	return "UCS-2"
	case isJisC6226(data):
		return "JIS_C6226"
	case isIso88591(data):
		return "ISO_8859_1"
	case isAnsiX34(data):
		return "ANSI_X3.4 (ASCII)"

	default:
		return "Unknown"
	}
}

func main() {
	// Example test

	validData := []byte("Hello, World!")
	fmt.Println("Valid data:", detectEncoding(validData))
	validvalue := ansi34toGoString(validData)
	fmt.Println("Valid data:", validvalue)
	validdata5, err := validateAndFixUTF8(string(validData))
	if err != nil {
		log.Fatalf("Error fixing UTF-8: %v", err)
	}
	fmt.Println("Fixed UTF-8 data 1:", validdata5)
	asciiData := []byte("さようなら")
	fmt.Println("encoding of japanese characters:", detectEncoding(asciiData))
	asdata1, err := validateAndFixUTF8(string(asciiData))
	if err != nil {
		log.Fatalf("Error fixing UTF-8: %v", err)
	}
	fmt.Println("Fixed UTF-8 data 1:", asdata1)
	//value := jisToGoString(asciiData)
	fmt.Println("ASCII data:", jisToGoString(asciiData))
	asdata, err := validateAndFixUTF8(jisToGoString(asciiData))
	if err != nil {
		log.Fatalf("Error fixing UTF-8: %v", err)
	}
	fmt.Println("Fixed UTF-8 data:", asdata)
	fmt.Println("Detected encoding:", detectEncoding(asciiData))

	ucs2BE := []byte("ä") // UCS-2 with BOM
	fmt.Printf("iso8859 data in hex: %X\n", ucs2BE)
	value1 := iso8859ToUtf8(ucs2BE)
	fmt.Println("iso8859 data:", value1)
	fmt.Println("Detected encoding:", detectEncoding(ucs2BE))

	ucs2BE2 := []byte{0xC2} // UCS-2 with BOM
	fmt.Printf("iso8859 in hex: %X\n", ucs2BE2)
	fmt.Println("Detected encoding:", detectEncoding(ucs2BE2))
	value13 := iso8859ToUtf8(ucs2BE2)
	fmt.Println("iso8859data:", value13)

	ucs4 := []byte{0x00, 0x00, 0x00, 0x48, // H
		0x00, 0x00, 0x00, 0x65, // e
		0x00, 0x00, 0x00, 0x6C, // l
		0x00, 0x00, 0x00, 0x6C, // l
		0x00, 0x00, 0x00, 0x6F}
	fmt.Println("Detected encoding:", detectEncoding(ucs4))
	// Convert []byte to []uint32
	// var ucs4Uint32 []uint32
	// for i := 0; i < len(ucs4); i += 4 {
	// 	if i+3 < len(ucs4) {
	// 		ucs4Uint32 = append(ucs4Uint32, uint32(ucs4[i])<<24|uint32(ucs4[i+1])<<16|uint32(ucs4[i+2])<<8|uint32(ucs4[i+3]))
	// 	}
	// }
	// value4 := ucs4ToGoString(ucs4Uint32)
	value4 := ansi34toGoString(ucs4)
	fmt.Println("UCS-4 data:", value4)

	ucs2data := []byte{0x00, 0x61, 0x00, 0x62, 0x00, 0x63} // UCS-2 with BOM
	fmt.Println("Detected encoding:", detectEncoding(ucs2data))
	// Convert []byte to []uint16
	// var ucs2Uint16 []uint16
	// for i := 0; i < len(ucs2data); i += 2 {
	// 	if i+1 < len(ucs2data) {
	// 		ucs2Uint16 = append(ucs2Uint16, uint16(ucs2data[i])<<8|uint16(ucs2data[i+1]))
	// 	}
	// }
	// value45 := ucs2ToGoString(ucs2Uint16)
	value45 := ansi34toGoString(ucs2data)
	fmt.Println("UCS-2 data:", value45)

}

func jisToGoString(jisArr []byte) string {
	decoder := japanese.ShiftJIS.NewDecoder()
	bytes, _ := io.ReadAll(transform.NewReader(bytes.NewBuffer(jisArr), decoder))
	if strings.Contains(string(bytes), "\\") || strings.Contains(string(bytes), "\uFFFD") || strings.Contains(string(bytes), "､ｰ､") {
		log.Printf("found faulty bytes during JIS conversion: %s", string(bytes))
	}
	return string(bytes)
}

func validateAndFixUTF8(input string) (string, error) {
	if utf8.ValidString(input) {
		log.Printf("validateAndFixUTF8: input is valid UTF-8: %s", input)
		return input, nil
	}

	decoder := transform.NewReader(bytes.NewReader([]byte(input)), unicode.UTF8.NewDecoder())
	fixedBytes, err := io.ReadAll(decoder)
	if err != nil {
		return "", fmt.Errorf("failed to decode string: %w", err)
	}
	return string(fixedBytes), nil
}

func iso8859ToUtf8(iso8859Buff []byte) string {
	buf := make([]rune, len(iso8859Buff))
	for i, b := range iso8859Buff {
		buf[i] = rune(b)
	}
	return string(buf)
}

func ucs4ToGoString(array []uint32) string {
	var utf8Runes []rune
	for _, u32 := range array {
		utf8Runes = append(utf8Runes, rune(u32))
	}
	utf8String := string(utf8Runes)
	return utf8String
}

func ucs2ToGoString(array []uint16) string {
	u16s := make([]uint16, 1)
	ret := &bytes.Buffer{}
	b8buf := make([]byte, 4)
	for _, u16 := range array {
		u16s[0] = uint16(u16)
		r := utf16.Decode(u16s)
		n := utf8.EncodeRune(b8buf, r[0])
		ret.Write(b8buf[:n])
	}
	return ret.String()
}

func ansi34toGoString(array []byte) string {
	decoder := unicode.UTF8.NewDecoder()
	bytes, err := io.ReadAll(transform.NewReader(bytes.NewReader(array), decoder))
	if err != nil {
		log.Printf("found faulty bytes during ansii34 conversion: %s", string(bytes))
	}
	return string(bytes)
}
