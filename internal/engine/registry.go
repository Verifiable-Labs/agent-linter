package engine

func DefaultRules() []Rule {
	return []Rule{
		RuleAL001DuplicateActionNames{},
		RuleAL101InvocationMustReferenceKnownAction{},
		RuleAL102InvocationMissingRequiredParams{},
	}
}
