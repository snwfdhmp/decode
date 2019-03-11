package decode

import (
	"encoding/json"
	"testing"

	// "github.com/sirupsen/logrus"
	"github.com/spf13/afero"
	// "gopkg.in/yaml.v2"
)

//ExampleType is a simple type for testing purposes
type ExampleType struct {
	Author string `json:"author"`
}

var (
	//ExampleData will hold data for our tests
	ExampleData = ExampleType{
		"snwfdhmp",
	}
	//ExamplePathJSON will hold path for test data file on JSON func
	ExamplePathJSON = "test.json"
	//ExamplePathYAML will hold path for test data file on YAML func
	ExamplePathYAML = "test.yaml"
)

//TestJSON runs tests for JSON func
func TestJSON(t *testing.T) {
	b, err := json.Marshal(ExampleData)
	if err != nil {
		t.Errorf("err: %s", err)
		return
	}
	if err := afero.WriteFile(fs, ExamplePathJSON, b, 0755); err != nil {
		t.Errorf("err: %s", err)
		return
	}

	var result ExampleType
	if err := JSON(ExamplePathJSON, &result); err != nil {
		t.Errorf("err: %s", err)
		return
	}

	if result != ExampleData {
		t.Errorf("result=%s exampleData=%s (different)", result, ExampleData)
		return
	}
}

//TestParse runs tests for Parse func
func TestParse(t *testing.T) {
	b, err := json.Marshal(ExampleData)
	if err != nil {
		t.Errorf("err: %s", err)
		return
	}
	if err := afero.WriteFile(fs, ExamplePathJSON, b, 0755); err != nil {
		t.Errorf("err: %s", err)
		return
	}

	var result ExampleType
	if err := Parse(json.Unmarshal, ExamplePathJSON, &result); err != nil {
		t.Errorf("err: %s", err)
		return
	}

	if result != ExampleData {
		t.Errorf("result=%s exampleData=%s (different)", result, ExampleData)
		return
	}
}

//TestYAML runs tests for YAML func
func TestYAML(t *testing.T) {
	b, err := json.Marshal(ExampleData)
	if err != nil {
		t.Errorf("err: %s", err)
		return
	}
	if err := afero.WriteFile(fs, ExamplePathYAML, b, 0755); err != nil {
		t.Errorf("err: %s", err)
		return
	}

	var result ExampleType
	if err := YAML(ExamplePathYAML, &result); err != nil {
		t.Errorf("err: %s", err)
		return
	}

	if result != ExampleData {
		t.Errorf("result=%s exampleData=%s (different)", result, ExampleData)
		return
	}
}

func init() {
	// logrus.SetLevel(logrus.DebugLevel)
}
