//Package decode provides shortcuts for simple decoding funcs
package decode

import (
	"encoding/json"

	// "path/filepath"

	"github.com/spf13/afero"
	yaml "gopkg.in/yaml.v2"
)

var (
	fs = afero.NewOsFs()
)

// UnmarshalFunc represents an function used for Unmarshal-ing bytes into a struct
// Example: json.Unmarshal from package encoding/json
type UnmarshalFunc func([]byte, interface{}) error

// Parse parses file located at 'path' using the given UnmarshalFunc, into given recipient
// Any standard Marshal func of type :
//     func (b []byte, data interface{}) error
// can be used.
func Parse(u UnmarshalFunc, path string, to interface{}) error {
	content, err := afero.ReadFile(fs, path) //read file
	if err != nil {
		return err
	}

	return u(content, to)
}

//JSON calls Parse with encoding/json Unmarshal func
func JSON(path string, to interface{}) error {
	return Parse(json.Unmarshal, path, to)
}

//YAML calls Parse with gopkg.in/yaml.v2 Unmarshal func
func YAML(path string, to interface{}) error {
	return Parse(yaml.Unmarshal, path, to)
}
