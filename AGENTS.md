# Agent Guidelines for Synapse Developer Platform

## Build/Test Commands
This is a Crossplane-first Kubernetes developer platform. Currently in initial directory structure phase with no build tools configured yet.
- No build commands available (add when Go modules, Helm charts, or other tooling is implemented)
- No test commands available (add when test frameworks are implemented)
- No lint commands available (add when linting tools are configured)

## Architecture Overview
- **Crossplane-First**: All infrastructure provisioning through Crossplane compositions
- **Domain-Driven Structure**: Services grouped by business capability (applications/, control-plane/, governance/, identity/)
- **Kubernetes Native**: Infrastructure resources are Kubernetes resources
- **Single Sync Operator**: Watches multiple K8s resource types and syncs to database

## Code Style Guidelines
Since this is an early-stage project with only directory structure and README:
- Follow domain-driven directory organization as shown in README.md
- Use kebab-case for directory names (control-plane/, sync-operator/)
- Place .gitkeep files in empty directories to maintain structure
- When adding code, follow language-specific best practices (Go for CLI/operators, YAML for Crossplane resources)
- Use clear, descriptive names that reflect business capabilities
- Maintain separation of concerns between domains

## Development Principles
- Declarative infrastructure through Crossplane compositions
- Team ownership of complete business domains
- Clear boundaries with well-defined interfaces between domains
- CLI-first approach for developer experience