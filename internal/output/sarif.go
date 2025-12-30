package output

import (
	"path/filepath"
	"sort"

	"github.com/verifiable-labs/agent-linter/internal/engine"
)

type sarifLog struct {
	Version string     `json:"version"`
	Schema  string     `json:"$schema"`
	Runs    []sarifRun `json:"runs"`
}

type sarifRun struct {
	Tool    sarifTool     `json:"tool"`
	Results []sarifResult `json:"results,omitempty"`
}

type sarifTool struct {
	Driver sarifDriver `json:"driver"`
}

type sarifDriver struct {
	Name           string      `json:"name"`
	InformationURI string      `json:"informationUri,omitempty"`
	Version        string      `json:"version,omitempty"`
	Rules          []sarifRule `json:"rules,omitempty"`
}

type sarifRule struct {
	ID               string `json:"id"`
	Name             string `json:"name,omitempty"`
	ShortDescription struct {
		Text string `json:"text"`
	} `json:"shortDescription,omitempty"`
}

type sarifResult struct {
	RuleID    string           `json:"ruleId"`
	Level     string           `json:"level"`
	Message   sarifMessage     `json:"message"`
	Locations []sarifLocation  `json:"locations"`
}

type sarifMessage struct {
	Text string `json:"text"`
}

type sarifLocation struct {
	PhysicalLocation sarifPhysicalLocation `json:"physicalLocation"`
}

type sarifPhysicalLocation struct {
	ArtifactLocation sarifArtifactLocation `json:"artifactLocation"`
}

type sarifArtifactLocation struct {
	URI string `json:"uri"`
}

func ToSarif(findings []engine.Finding, version string) sarifLog {
	rulesMap := make(map[string]sarifRule)
	var results []sarifResult

	for _, f := range findings {
		// GitHub Code Scanning requires at least one location
		if f.File == "" {
			continue
		}

		level := "warning"
		if f.Severity == engine.SeverityError {
			level = "error"
		}

		uri := filepath.ToSlash(f.File)

		res := sarifResult{
			RuleID:  f.RuleID,
			Level:   level,
			Message: sarifMessage{Text: f.Message},
			Locations: []sarifLocation{
				{
					PhysicalLocation: sarifPhysicalLocation{
						ArtifactLocation: sarifArtifactLocation{
							URI: uri,
						},
					},
				},
			},
		}

		results = append(results, res)

		if _, ok := rulesMap[f.RuleID]; !ok {
			var rule sarifRule
			rule.ID = f.RuleID
			rule.Name = f.RuleID
			rule.ShortDescription.Text = "agent-linter rule " + f.RuleID
			rulesMap[f.RuleID] = rule
		}
	}

	var sarifRules []sarifRule
	for _, r := range rulesMap {
		sarifRules = append(sarifRules, r)
	}

	sort.Slice(sarifRules, func(i, j int) bool {
		return sarifRules[i].ID < sarifRules[j].ID
	})

	sort.Slice(results, func(i, j int) bool {
		if results[i].RuleID == results[j].RuleID {
			return results[i].Locations[0].PhysicalLocation.ArtifactLocation.URI <
				results[j].Locations[0].PhysicalLocation.ArtifactLocation.URI
		}
		return results[i].RuleID < results[j].RuleID
	})

	return sarifLog{
		Version: "2.1.0",
		Schema:  "https://json.schemastore.org/sarif-2.1.0.json",
		Runs: []sarifRun{
			{
				Tool: sarifTool{
					Driver: sarifDriver{
						Name:           "agent-linter",
						InformationURI: "https://verifiablelabs.com",
						Version:        version,
						Rules:          sarifRules,
					},
				},
				Results: results,
			},
		},
	}
}
