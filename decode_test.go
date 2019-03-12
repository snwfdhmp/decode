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

	fs = afero.NewOsFs()
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

	fakePath := "/tmp/thisFileDoesNot_/Exists"
	if err := JSON(fakePath, &result); err == nil {
		exists, err := afero.Exists(fs, fakePath)
		if err != nil {
			t.Errorf("could not execute tests properly: could not check if fakePath exists (%s): %s", fakePath, err)
			return
		} else if exists {
			t.Errorf("could not execute tests properly: the used fakePath is an existing file (%s)", fakePath)
			return
		}
		t.Errorf("error not handled when file doesnt exist")
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

	fakePath := "/tmp/thisFileDoesNot_/Exists"
	if err := YAML(fakePath, &result); err == nil {
		exists, err := afero.Exists(fs, fakePath)
		if err != nil {
			t.Errorf("could not execute tests properly: could not check if fakePath exists (%s): %s", fakePath, err)
			return
		} else if exists {
			t.Errorf("could not execute tests properly: the used fakePath is an existing file (%s)", fakePath)
			return
		}
		t.Errorf("error not handled when file doesnt exist")
		return
	}
}

func init() {
	// logrus.SetLevel(logrus.DebugLevel)
}
