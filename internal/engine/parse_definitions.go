package engine

import (
	"encoding/json"
	"fmt"
	"os"
)

type openAITool struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Parameters  struct {
		Properties map[string]struct {
			Type string `json:"type"`
		} `json:"properties"`
		Required []string `json:"required"`
	} `json:"parameters"`
}

func LoadActionDefinitions(path string) ([]ActionDefinition, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var tools []openAITool
	if err := json.Unmarshal(data, &tools); err != nil {
		return nil, fmt.Errorf("invalid action definitions %q: %w", path, err)
	}

	var defs []ActionDefinition
	for _, t := range tools {
		params := make(map[string]Parameter)
		required := make(map[string]bool)
		for _, r := range t.Parameters.Required {
			required[r] = true
		}

		for name, p := range t.Parameters.Properties {
			params[name] = Parameter{
				Type:     p.Type,
				Required: required[name],
			}
		}

		defs = append(defs, ActionDefinition{
			Name:        t.Name,
			Description: t.Description,
			Parameters:  params,
		})
	}

	return defs, nil
}
