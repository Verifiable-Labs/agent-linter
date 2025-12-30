package engine

type Inputs struct {
	Definitions []ActionDefinition
	Invocations []ActionInvocation
	Wiring      []WiringBinding
}
