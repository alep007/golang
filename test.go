package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

func main() {
	file, err := ioutil.ReadFile("text.txt") //read text file
	if err != nil {
		fmt.Print(err)
	}
	text := strings.ToLower(string(file))     //convert text to lower
	text = strings.Replace(text, "'", "", -1) //replace symbol "'" with "" to avoid unwanted words

	fields := strings.FieldsFunc(text, func(r rune) bool { //separate each words
		return !('a' <= r && r <= 'z')
	})

	words := make(map[string]int)
	for _, field := range fields { //start counting words
		words[field]++
	}

	er := ioutil.WriteFile("output.txt", []byte(createKeyValuePairs(words)), 0755) //create textfile where result is displayed
	if er != nil {
		fmt.Printf("Unable to write file: %v", er)
	}
}

func createKeyValuePairs(m map[string]int) string {
	b := new(bytes.Buffer)
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		fmt.Fprintf(b, "%s=\"%d\"\n", k, m[k])
	}
	return b.String()
}
