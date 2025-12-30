package engine

type ActionDefinition struct {
	Name        string
	Description string
	Parameters  map[string]Parameter
}

type Parameter struct {
	Type     string
	Required bool
}

type ActionInvocation struct {
	Action string
	Args   map[string]any
	Source string
}

type WiringBinding struct {
	Action      string
	Adapter     string
	Environment string
	Target      string
}
