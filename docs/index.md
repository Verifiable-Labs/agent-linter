# agent-linter rules

This directory contains the complete rule set used by agent-linter to validate executable AI agent behavior.

Each rule represents a deterministic check applied to an agent’s execution surface before runtime.

---

## 🎯 Rule philosophy

Rules in agent-linter follow strict design constraints:

- Deterministic and reproducible
- Purely static analysis
- No inference or probabilistic judgment
- No runtime execution
- Explainable findings with clear remediation guidance

The goal is to make agent behavior predictable, auditable, and enforceable.

---

## 🧱 What rules analyze

Rules operate on three concrete artifacts:

1. Action definitions  
2. Action invocations  
3. Execution wiring  

Together, these describe everything an agent is permitted to execute.

---

## 🆔 Rule identifiers

Each rule has a stable identifier, for example:

- AL001
- AL102
- AL201

Rule IDs are permanent and will not be reused or repurposed.  
This guarantees long term compatibility with CI systems, suppressions, and audits.

---

## 📂 Rule documentation structure

Each rule is documented in its own file:

- One rule per file
- File name matches the rule ID
- Clear description of the violation
- Explanation of risk
- Remediation guidance
- Examples when appropriate

Example:

    docs/rules/AL201.md

---

## 🔕 Suppressions

Rules may be suppressed using fingerprints when justified.

Suppressions must be:

- Rule specific
- Fingerprint specific
- Documented with a reason
- Optionally time bounded

Expired suppressions automatically stop applying.

---

## 🧪 Rule evolution

Rules may evolve to become stricter over time, but they will not change meaning.

Breaking semantic changes require a new rule ID.

Deprecated rules will be documented clearly.

---

## 📚 Adding new rules

When adding a new rule:

- Choose the next available rule ID
- Add comprehensive documentation
- Include deterministic tests
- Ensure findings are precise and actionable

Rules that overlap significantly with existing ones may be rejected.

---

## 🔗 Relationship to the agent lifecycle

Rules apply at a single phase of the lifecycle:

Design time and CI, before execution.

They are intentionally separate from:

- Runtime enforcement
- Simulation
- Debugging

This separation keeps the system simple and auditable.

---

agent-linter rules exist to make executable AI systems behave like well engineered software.
