package config

type Config struct {
	Version      int              `yaml:"version"`
	Rules        map[string]Rule  `yaml:"rules"`
	Inputs       Inputs           `yaml:"inputs"`
	Environments Environments     `yaml:"environments"`
	Suppress     []Suppression    `yaml:"suppress"`
}

type Rule struct {
	Enabled  bool   `yaml:"enabled"`
	Severity string `yaml:"severity"`
}

type Inputs struct {
	Definitions []string `yaml:"definitions"`
	Invocations []string `yaml:"invocations"`
	Wiring      []string `yaml:"wiring"`
}

type Environments struct {
	Default string                       `yaml:"default"`
	Named   map[string]EnvironmentConfig `yaml:",inline"`
}

type EnvironmentConfig struct {
	Match []string `yaml:"match"`
}

type Suppression struct {
	Rule        string `yaml:"rule"`
	Fingerprint string `yaml:"fingerprint"`
	Reason      string `yaml:"reason"`
	Expires     string `yaml:"expires"`
}
