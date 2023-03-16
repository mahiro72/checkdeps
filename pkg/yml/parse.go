package yml

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

type depsYml struct {
	Spec struct {
		Module struct {
			Name string `yaml:"name"`
		}
		Observes []string            `yaml:"observes"`
		Layers   map[string][]string `yaml:"layers"`
	}
}

func Parse(b []byte) (*depsYml, error) {
	var y depsYml
	err := yaml.Unmarshal([]byte(b), &y)
	if err != nil {
		return nil, err
	}

	if err := checkYml(&y); err != nil {
		return nil, err
	}

	return &y, nil
}

func checkYml(d *depsYml) error {
	if d.Spec.Module.Name == "" {
		return fmt.Errorf("error: deps.yml module name not found")
	}

	return nil
}
