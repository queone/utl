package utl

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	goyaml "github.com/goccy/go-yaml"
	"github.com/goccy/go-yaml/lexer"
	"github.com/goccy/go-yaml/token"
	"gopkg.in/yaml.v3"
)

// Trys to read, load, and decode given filePath as some YAML object.
// Returns YAML object and error if any.
func LoadFileYaml(filePath string) (yamlObject interface{}, err error) {
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(fileContent, &yamlObject)
	if err != nil {
		return nil, err
	}
	return yamlObject, nil
}

// Tries to read, load, and recode given filePath as some YAML object as byte slice.
// Returns YAML object as byte slice. and error if any.
func LoadFileYamlBytes(filePath string) (yamlBytes []byte, err error) {
	// Load YAML file into byte slice, including comments
	// Can also JSON file into byte slice!
	yamlBytes, err = os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	// Check YAML formatting compliancy using "github.com/goccy/go-yaml"
	// which provides errors with line numbers
	var yamlObject interface{}
	err = goyaml.Unmarshal(yamlBytes, &yamlObject)
	if err != nil {
		return nil, err
	}
	return yamlBytes, nil // We only care about returning the byte slice
}

// Save given YAML object to given filePath
func SaveFileYaml(yamlObject interface{}, filePath string) {
	yamlData, err := yaml.Marshal(&yamlObject)
	if err != nil {
		panic(err.Error())
	}
	err = os.WriteFile(filePath, yamlData, 0600)
	if err != nil {
		panic(err.Error())
	}
}

// Convert byte slice to YAML interface object
func BytesToYamlObject(yamlBytes []byte) (yamlObject interface{}, err error) {
	buffer := bytes.NewBuffer(yamlBytes)
	decoder := yaml.NewDecoder(buffer)
	err = decoder.Decode(&yamlObject)
	if err != nil {
		return nil, err
	}
	return yamlObject, nil
}

// Convert YAML interface object to byte slice, with option indent spacing
func YamlToBytesIndent(yamlObject interface{}, indent int) (yamlBytes []byte, err error) {
	buffer := &bytes.Buffer{}
	encoder := yaml.NewEncoder(buffer)
	encoder.SetIndent(indent)
	err = encoder.Encode(yamlObject)
	if err != nil {
		return nil, err
	}
	yamlBytes = buffer.Bytes()
	return yamlBytes, nil
}

// With default 2 space indent
func YamlToBytes(yamlObject interface{}) (yamlBytes []byte, err error) {
	indent := 2
	yamlBytes, err = YamlToBytesIndent(yamlObject, indent)
	return yamlBytes, err
}

// Print YAML object
func PrintYaml(yamlObject interface{}) {
	yamlBytes, err := YamlToBytes(yamlObject)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(yamlBytes))
}

// Colorize given token. Internal helper function.
func colorizeString(tk *token.Token, src string) string {
	str := Whi(src)
	switch tk.Type {
	case token.MappingKeyType:
		str = Blu(src)
	case token.StringType, token.SingleQuoteType, token.DoubleQuoteType:
		prev := tk.PreviousType()
		next := tk.NextType()
		if next == token.MappingValueType {
			str = Blu(src)
		} else if prev == token.AnchorType || prev == token.AliasType {
			str = Yel(src)
		} else {
			str = Gre(src)
		}
	case token.IntegerType, token.FloatType, token.BoolType:
		str = Mag(src)
	case token.AnchorType, token.AliasType:
		str = Yel(src)
	case token.CommentType:
		str = Whi(src)
	}
	return str
}

// Print YAML object (that don't usually include comments) in color
func PrintYamlColor(yamlObject interface{}) {
	yamlBytes, err := YamlToBytes(yamlObject)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	PrintYamlBytesColor(yamlBytes)
}

// Print YAML bytes in color, includes comments
// Also prints JSON byte slice in color, but use json.go:PrintJsonBytesColor() alias instead.
// Caller must ensure yamlBytes is proper YAML/JSON
func PrintYamlBytesColor(yamlBytes []byte) {
	tokens := lexer.Tokenize(string(yamlBytes))
	if len(tokens) == 0 {
		return
	}
	printOut := []string{}
	//lineNumber := tokens[0].Position.Line
	for _, tk := range tokens {
		lines := strings.Split(tk.Origin, "\n")
		header := ""
		// if allowLineNumber {
		// 	header = fmt.Sprintf("%2d  ", lineNumber)
		// }
		if len(lines) == 1 {
			line := colorizeString(tk, lines[0])
			if len(printOut) == 0 {
				printOut = append(printOut, header+line)
				//lineNumber++
			} else {
				text := printOut[len(printOut)-1]
				printOut[len(printOut)-1] = text + line
			}
		} else {
			header := ""
			for idx, src := range lines {
				// if allowLineNumber {
				// 	header = fmt.Sprintf("%2d  ", lineNumber)
				// }
				line := colorizeString(tk, src)
				if idx == 0 {
					if len(printOut) == 0 {
						printOut = append(printOut, header+line)
						//lineNumber++
					} else {
						text := printOut[len(printOut)-1]
						printOut[len(printOut)-1] = text + line
					}
				} else {
					printOut = append(printOut, fmt.Sprintf("%s%s", header, line))
					//lineNumber++
				}
			}
		}
	}
	fmt.Println(strings.Join(printOut, "\n"))
}
