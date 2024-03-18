package utl

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

// Reads, load, and decode given filePath as a JSON object text file.
// Returns JSON object and err if any.
func LoadFileJson(filePath string) (jsonObject interface{}, err error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	byteValue, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(byteValue), &jsonObject)
	if err != nil {
		return nil, err
	}
	return jsonObject, nil
}

// Reads, load, and decode given filePath as a gzipped JSON object text file.
// Returns JSON object and err if any.
// The compression helps speed things up for some very large objects.
func LoadFileJsonGzip(filePath string) (jsonObject interface{}, err error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	gzipReader, err := gzip.NewReader(f)
	if err != nil {
		return nil, err
	}
	defer gzipReader.Close()

	byteValue, err := io.ReadAll(gzipReader)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(byteValue), &jsonObject)
	if err != nil {
		return nil, err
	}

	return jsonObject, nil
}

// Save given JSON object as text file
func SaveFileJson(jsonObject interface{}, filePath string) {
	jsonData, err := json.Marshal(jsonObject)
	if err != nil {
		panic(err.Error())
	}
	err = os.WriteFile(filePath, jsonData, 0600)
	if err != nil {
		panic(err.Error())
	}
}

// Save given JSON object as gzipped text file
func SaveFileJsonGzip(jsonObject interface{}, filePath string) {
	jsonData, err := json.Marshal(jsonObject)
	if err != nil {
		panic(err.Error())
	}

	file, err := os.Create(filePath)
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()

	gzipWriter := gzip.NewWriter(file)
	defer gzipWriter.Close()

	_, err = gzipWriter.Write(jsonData)
	if err != nil {
		panic(err.Error())
	}
}

// Prints JSON object, flushing the output buffer
func PrintJson(jsonObject interface{}) {
	pretty, err := Prettify(jsonObject)
	if err != nil {
		fmt.Printf("Prettify() error\n")
	} else {
		fmt.Println(pretty)
	}
	os.Stdout.Sync()
}

// Convert JSON interface object to byte slice, with option to indent spacing
func JsonBytesReindent(jsonBytes []byte, indent int) (jsonBytes2 []byte, err error) {
	var prettyJson bytes.Buffer
	indentStr := strings.Repeat(" ", indent)
	err = json.Indent(&prettyJson, jsonBytes, "", indentStr)
	if err != nil {
		return nil, err
	}
	jsonBytes2 = prettyJson.Bytes()
	return jsonBytes2, nil
}

// Convert JSON interface object to byte slice, with option to indent spacing
func JsonToBytesIndent(jsonObject interface{}, indent int) (jsonBytes []byte, err error) {
	indentStr := strings.Repeat(" ", indent)
	jsonBytes, err = json.MarshalIndent(jsonObject, "", indentStr)
	if err != nil {
		return nil, err
	}
	return jsonBytes, nil
}

// Convert JSON interface object to byte slice, with default 2-space indentation
func JsonToBytes(jsonObject interface{}) (jsonBytes []byte, err error) {
	indent := 2 // With default 2 space indent
	jsonBytes, err = JsonToBytesIndent(jsonObject, indent)
	return jsonBytes, err
}

// Convert JSON byte slice to JSON interface object, with default 2-space indentation
func JsonBytesToJsonObj(jsonBytes []byte) (jsonObject interface{}, err error) {
	err = json.Unmarshal(jsonBytes, &jsonObject)
	return jsonObject, err
}

// NOTE: To be replaced by JsonToBytes()
func Prettify(jsonObject interface{}) (pretty string, err error) {
	j, err := json.MarshalIndent(jsonObject, "", "  ")
	return string(j), err
}

// Print JSON object in color
func PrintJsonColor(jsonObject interface{}) {
	jsonBytes, err := JsonToBytes(jsonObject)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	PrintJsonBytesColor(jsonBytes)
}

// Prints JSON byte slice in color. Just an alias of yaml.go:PrintYamlBytesColor().
func PrintJsonBytesColor(jsonBytes []byte) {
	PrintYamlBytesColor(jsonBytes)
}

// Combines two string-to-string maps, with keys from the second map overwriting those
// from the first if duplicates exist. Returns the merged map.
func MergeMaps(m1, m2 map[string]string) (result map[string]string) {
	result = map[string]string{}
	for k, v := range m1 {
		result[k] = v
	}
	for k, v := range m2 {
		result[k] = v
	}
	return result
}

// Non-recursive merge of first-level attributes in JSON object y onto object x
// If attribute exists in y, it is overwritten
func MergeObjects(x, y map[string]interface{}) (obj map[string]interface{}) {
	obj = x
	for k, v := range x { // Update existing x values with updated y values
		obj[k] = v
		if y[k] != nil {
			obj[k] = y[k]
		}
	}
	for k := range y { // Add new y values to x
		if x[k] == nil {
			obj[k] = y[k]
		}
	}
	return obj
}

// Recursive function returns True if filter string value is anywhere within jsonObject
func StringInJson(jsonObject interface{}, filter string) bool {
	switch value := jsonObject.(type) {
	case string:
		return SubString(value, filter)
	case []interface{}:
		for _, v := range value {
			if StringInJson(v, filter) {
				return true
			}
		}
	case map[string]interface{}:
		for _, v := range value {
			if StringInJson(v, filter) {
				return true
			}
		}
	}
	return false
}
