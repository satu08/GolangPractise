package main

// func main() {

// 	decodestr, enc := ValidateAndFixUtf8([]byte("å701"))
// 	fmt.Printf("Decoded string in main: %s\n", decodestr)

// 	encodedBack, _, err := EncodeToOriginalEncoding(decodestr, enc)
// 	if err != nil {
// 		log.Fatalf("Encoding back to %s failed: %v", enc, err)
// 	}

// 	fmt.Printf("Successfully re-encoded to %s: %v\n", enc, string(encodedBack))

// }

// func EncodeToOriginalEncoding(utf8Str string, encName string) ([]byte, int, error) {
// 	enc, err := GetEncodingByName(encName)
// 	if err != nil {
// 		return nil, 0, err
// 	}
// 	return transform.Bytes(enc.NewEncoder(), []byte(utf8Str))
// }

// func GetEncodingByName(name string) (encoding.Encoding, error) {
// 	switch name {
// 	case "Shift_JIS":
// 		return japanese.ShiftJIS, nil
// 	case "EUC-JP":
// 		return japanese.EUCJP, nil
// 	case "ISO-2022-JP":
// 		return japanese.ISO2022JP, nil
// 	case "UTF-8":
// 		return nil, nil
// 	case "UTF-8BOM":
// 		return unicode.UTF8BOM, nil
// 	case "UTF-16LE":
// 		return unicode.UTF16(unicode.LittleEndian, unicode.IgnoreBOM), nil
// 	case "UTF-16BE":
// 		return unicode.UTF16(unicode.BigEndian, unicode.IgnoreBOM), nil
// 	case "UTF-16LEBOM":
// 		return unicode.UTF16(unicode.LittleEndian, unicode.UseBOM), nil
// 	case "UTF-16BEBOM":
// 		return unicode.UTF16(unicode.BigEndian, unicode.UseBOM), nil
// 	default:
// 		return nil, errors.New("unsupported encoding: " + name)
// 	}
// }

// func ValidateAndFixUtf8(data []byte) (string, string) {
// 	validBytes := replaceReplacementChar(string(data))
// 	enc, err := jpdec.DetectEncoding([]byte(validBytes))
// 	if err != nil {
// 		fmt.Printf("Error detecting encoding: %s %v\n", string(validBytes), err)
// 		return "", ""
// 	}
// 	fmt.Printf("Detected encoding: %s\n", enc)
// 	decodedStr, err := jpdec.Decode([]byte(validBytes))
// 	if err != nil {
// 		fmt.Printf("Error decoding string: %v %v\n", string(validBytes), err)
// 		return "", ""
// 	}
// 	fmt.Printf("Decoded string: %s\n", decodedStr)
// 	isvalidUtf8 := utf8.ValidString(decodedStr)
// 	if !isvalidUtf8 {
// 		return "", ""
// 	} else {
// 		fmt.Printf("Decoded string is valid UTF-8: %s\n", decodedStr)
// 	}
// 	return decodedStr, enc
// }

// func replaceReplacementChar(s string) string {
// 	return strings.ReplaceAll(s, string('\uFFFD'), "?")
// }

// func main() {
// 	usr, _ := user.Current()
// 	fmt.Println(usr.Name)
// 	fmt.Println(usr.HomeDir)
// 	fmt.Println(filepath.Join(usr.HomeDir, "config"))

// }
