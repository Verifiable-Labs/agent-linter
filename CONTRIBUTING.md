# Contributing to agent-linter

Thank you for your interest in contributing to agent-linter.  
This project values correctness, determinism, and long term maintainability over rapid feature growth.

Contributions are welcome when they align with these principles.

---

## 🎯 Project principles

Before contributing, please understand the core design goals:

- Deterministic behavior only
- No probabilistic or heuristic based rules
- Stable rule identifiers
- Clear, explainable findings
- Safe for CI and regulated environments

Changes that compromise these principles will not be accepted.

---

## 🧩 Types of contributions

We welcome:

- New lint rules
- Improvements to existing rules
- Documentation enhancements
- Bug fixes
- Performance improvements
- Test coverage additions

Large architectural changes should be discussed before implementation.

---

## 🧪 Rule design guidelines

When adding or modifying rules:

- Rules must be purely static
- Rules must not depend on runtime behavior
- Rules must produce the same output for the same input
- Error messages must be precise and actionable
- Rule IDs must remain stable once released

Each rule must include documentation under `docs/rules`.

---

## 🧪 Testing requirements

All changes must include appropriate tests.

- New rules require positive and negative test cases
- Bug fixes require a regression test
- Tests must be deterministic and reproducible

CI must pass before review.

---

## 📄 Documentation

If your change affects user facing behavior, documentation updates are required.

This includes:

- README updates when applicable
- Rule documentation for new or changed rules
- Clear examples when helpful

---

## 🔁 Development workflow

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Run tests locally
5. Commit with a clear message
6. Open a pull request

Small, focused pull requests are preferred.

---

## 📜 License

By contributing, you agree that your contributions will be licensed under the Apache License 2.0.

---

## 🤝 Code of conduct

All contributors are expected to be respectful and constructive.  
Harassment or abusive behavior will not be tolerated.

Thank you for helping make agent-linter reliable and trustworthy.
