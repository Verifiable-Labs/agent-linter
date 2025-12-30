# 🤖 agent-linter

[![Go Version](https://img.shields.io/badge/go-1.22%2B-00ADD8?logo=go&logoColor=white)](https://go.dev/)
[![License](https://img.shields.io/badge/license-Apache%202.0-blue.svg)](LICENSE)
[![Release](https://img.shields.io/github/v/release/verifiable-labs/agent-linter)](https://github.com/verifiable-labs/agent-linter/releases)
[![CI](https://img.shields.io/github/actions/workflow/status/verifiable-labs/agent-linter/ci.yml?label=CI)](https://github.com/verifiable-labs/agent-linter/actions)

**Deterministic static analysis for executable AI agent behavior**

agent-linter validates what AI agents are allowed to execute before they ever run.  
It focuses exclusively on executable actions and wiring, not prompts, reasoning chains, or model behavior.

Designed for environments where predictability, auditability, and determinism are non negotiable.

---

## ✨ Why agent-linter exists

As AI agents gain the ability to perform real world actions, the primary risk shifts from intelligence quality to execution correctness.

Most failures occur due to:

- Ambiguous or underspecified actions  
- Unsafe defaults  
- Missing or optional parameters  
- Incorrect production wiring  
- Undocumented execution surfaces  

agent-linter applies deterministic, mechanical rules to these surfaces, similar to how traditional linters protect software systems.

---

## 🧠 What agent-linter provides

agent-linter is purpose built to deliver a **clear and bounded execution contract** for AI agents.

It provides:

- Static analysis of agent execution surfaces  
- Deterministic and explainable results  
- Fast, offline operation suitable for CI  
- Enforceable guarantees before runtime  
- Stable rule identifiers for long term governance  

This scope ensures findings are reproducible, auditable, and automation friendly.

---

## 🏗️ Architecture overview

agent-linter operates before execution, validating the full agent execution surface.

    ┌─────────────────────┐
    │  Agent Source Code  │
    │                     │
    │  • Action schemas   │
    │  • Invocations      │
    │  • Wiring config    │
    └─────────┬───────────┘
              │
              ▼
    ┌─────────────────────┐
    │     agent-linter    │
    │                     │
    │  • Rule engine      │
    │  • Deterministic    │
    │  • Offline          │
    └─────────┬───────────┘
              │
              ▼
    ┌──────────────────────────┐
    │ Findings and guarantees  │
    │                          │
    │ • CI pass or fail        │
    │ • SARIF for GitHub       │
    │ • Audit ready output     │
    └──────────────────────────┘

---

## 📦 Installation

### From release binaries

Download the appropriate binary from GitHub Releases and place it on your PATH.

### From source

    go install github.com/verifiable-labs/agent-linter/cmd/agent-linter@latest

---

## 🚀 Quickstart

    agent-linter lint .

### Machine readable output

    agent-linter lint . --format json

### GitHub Code Scanning (SARIF)

    agent-linter lint . --format sarif > agent-linter.sarif

---

## 📥 Supported inputs

agent-linter analyzes three concrete artifacts that together define the executable surface of an agent system:

1. Action definitions  
2. Action invocations  
3. Execution wiring  

These inputs describe everything an agent is permitted to execute.

---

## 📏 Rules

    agent-linter rules

Rule documentation lives under:

    docs/rules

Rules are identified by stable codes such as AL001 and remain consistent across releases.

---

## 🔕 Suppressions

Example configuration:

    suppress:
      - rule: AL201
        fingerprint: "3cf79e3651b54a3cadc6ee70c5e5f93de11896e693f72a3c357bb6baf06cd235"
        reason: "Production wiring allowed for controlled integration tests"
        expires: "2026-06-01"

Suppressions are rule specific, fingerprint specific, and optionally time bounded.  
Expired suppressions automatically stop applying.

---

## 🔁 CI usage

GitHub Actions example:

    - name: Generate SARIF
      run: agent-linter lint . --format sarif > agent-linter.sarif || true

    - name: Upload SARIF
      uses: github/codeql-action/upload-sarif@v3
      with:
        sarif_file: agent-linter.sarif

---

## 🛒 GitHub Marketplace alignment

agent-linter integrates cleanly as:

- A standalone CLI  
- A GitHub Action wrapper  
- A Code Scanning provider via SARIF  

---

## 🔗 Relationship to Verifiable Labs

agent-linter is part of the Verifiable Labs open source ecosystem.

It composes naturally with:

- verifiable-agent-runtime  
- agent-simulator  
- verifiable-agent-debugger  

Each tool operates at a distinct phase of the agent lifecycle.

---

## 📜 License

Apache License 2.0

---

## Project Status 🚀

Current version: v0.1.0  
Status: Stable initial release

Contributions, bug reports, and feedback are warmly welcomed! 🙌

---

Made with ❤️ by Verifiable Labs