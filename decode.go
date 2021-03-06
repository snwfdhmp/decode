//Package decode provides shortcuts for simple decoding funcs
package decode

import (
	"encoding/json"
	"io"
	"os"

	"github.com/go-yaml/yaml"
)

// NewDecoderFunc represents an function used for Unmarshal-ing bytes into a struct
// Example: json.Unmarshal from package encoding/json
type NewDecoderFunc func(io.Reader) Decoder

// Decoder represents an object implementing the Decode func
type Decoder interface {
	//Decode represents the functions used for decoding. See json.Decoder.Decode or yaml.Decoder.Decode
	Decode(interface{}) error
}

//JSON calls Parse with encoding/json Unmarshal func
func JSON(path string, to interface{}) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewDecoder(file).Decode(to)
}

//YAML calls Parse with github.com/go-yaml/yaml Unmarshal func
func YAML(path string, to interface{}) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	return yaml.NewDecoder(file).Decode(to)
}
