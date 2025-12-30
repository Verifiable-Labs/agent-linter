package engine

func DefaultRules() []Rule {
	return []Rule{
		RuleAL001DuplicateActionNames{},
		RuleAL003DisallowPermissiveParamTypes{},
		RuleAL101InvocationMustReferenceKnownAction{},
		RuleAL102InvocationMissingRequiredParams{},
		RuleAL201NoProdByDefault{},
	}
}
