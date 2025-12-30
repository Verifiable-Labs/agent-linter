package engine

func DefaultRules() []Rule {
	return []Rule{
		RuleAL001DuplicateActionNames{},
	}
}
