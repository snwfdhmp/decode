//Package decode provides shortcuts for simple decoding funcs
package decode

import (
	"encoding/json"
	// "path/filepath"

	"github.com/sirupsen/logrus"
	"github.com/spf13/afero"
	"gopkg.in/yaml.v2"
)

var (
	fs = afero.NewOsFs()
)

//UnmarshalFunc is a type representing a typical Unmarshal func
//(see package encoding/json for example)
type UnmarshalFunc func([]byte, interface{}) error

//Parse parses file located at 'path', relative to current working directory
//using the given UnmarshalFunc, into given recipient
func Parse(u UnmarshalFunc, path string, to interface{}) error {
	logrus.Debugf("Parsing content of file %s as json", path)

	// abs, err := filepath.Abs(path)
	// if err != nil {
	// 	return err
	// }
	content, err := afero.ReadFile(fs, path)
	if err != nil {
		return err
	}
	logrus.Debugf("content: '%s'", content)

	return u(content, to)
}

//JSON runs Parse with encoding/json UnmarshalFunc
func JSON(path string, to interface{}) error {
	return Parse(json.Unmarshal, path, to)
}

//YAML runs Parse with gopkg.in/yaml.v2 UnmarshalFunc
func YAML(path string, to interface{}) error {
	return Parse(yaml.Unmarshal, path, to)
}
