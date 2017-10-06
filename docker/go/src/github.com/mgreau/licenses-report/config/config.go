package config

import (
	"github.com/mgreau/licenses-report/types"

	"gopkg.in/yaml.v2"
)

// Parse the licenses-report.yml file
func Parse(data []byte) (*types.Params, error) {
	c := new(types.Params)
	err := yaml.Unmarshal(data, c)
	return c, err
}
