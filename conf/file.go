package conf

import (
	"io"
	"os"

	"gopkg.in/dealancer/validate.v2"
	"gopkg.in/yaml.v2"

	"go.sancus.dev/core/errors"
)

//
// YAML
//
func LoadReader(f io.Reader, c interface{}) error {
	if b, err := io.ReadAll(f); err != nil {
		// read error
		return errors.Wrap(err, "ReadAll")
	} else if len(b) == 0 {
		// empty file
		return nil
	} else if err := yaml.Unmarshal(b, c); err != nil {
		// failed to decode
		return errors.Wrap(err, "Unmarshal")
	} else if err := validate.Validate(c); err != nil {
		// failed to validate
		return errors.Wrap(err, "Validate")
	} else {
		// ready
		return nil
	}
}

func LoadFile(filename string, c interface{}) error {

	if file, err := os.Open(filename); err != nil {
		return err
	} else if err := LoadReader(file, c); err != nil {
		return errors.Wrap(err, "Load: %q", filename)
	} else {
		return nil
	}
}

func WriteTo(f io.Writer, c interface{}) (int, error) {
	b, err := yaml.Marshal(c)
	if err != nil {
		// encoding error
		return 0, err
	}

	return f.Write(b)
}
