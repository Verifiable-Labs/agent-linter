package engine

import (
	"bufio"
	"encoding/json"
	"os"
)

type invocationLine struct {
	Action string         `json:"action"`
	Args   map[string]any `json:"args"`
}

func LoadActionInvocations(path string) ([]ActionInvocation, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var invocations []ActionInvocation
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		var line invocationLine
		if err := json.Unmarshal(scanner.Bytes(), &line); err != nil {
			return nil, err
		}

		invocations = append(invocations, ActionInvocation{
			Action: line.Action,
			Args:   line.Args,
			Source: path,
		})
	}

	return invocations, scanner.Err()
}
