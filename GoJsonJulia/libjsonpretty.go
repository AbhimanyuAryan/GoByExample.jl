package main

import (
	"C"
	"encoding/json"
	"fmt"
	"log"
)

//export jsonpretty
func jsonpretty(rawSrc *C.char) *C.char {
	data := new(map[string]interface{})
	err := json.Unmarshal([]byte(C.GoString(rawSrc)), &data)
	if err != nil {
		log.Printf("%s", err)
		return C.CString("")
	}
	src, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		log.Printf("%s", err)
		return C.CString("")
	}
	txt := fmt.Sprintf("%s", src)
	return C.CString(txt)
}

func main() {}
