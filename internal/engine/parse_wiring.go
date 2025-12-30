package engine

import (
	"os"

	"gopkg.in/yaml.v3"
)

type wiringFile struct {
	Bindings []struct {
		Action      string `yaml:"action"`
		Adapter     string `yaml:"adapter"`
		Environment string `yaml:"environment"`
		Target      string `yaml:"target"`
	} `yaml:"bindings"`
}

func LoadWiring(path string) ([]WiringBinding, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var wf wiringFile
	if err := yaml.Unmarshal(data, &wf); err != nil {
		return nil, err
	}

	var bindings []WiringBinding
	for _, b := range wf.Bindings {
		bindings = append(bindings, WiringBinding{
			Action:      b.Action,
			Adapter:     b.Adapter,
			Environment: b.Environment,
			Target:      b.Target,
		})
	}

	return bindings, nil
}
