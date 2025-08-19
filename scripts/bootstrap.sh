#!/bin/bash
set -e

echo "Adding Helm repositories..."
helm repo add crossplane https://charts.crossplane.io/stable
helm repo add kyverno https://kyverno.github.io/kyverno
helm repo update

echo "Installing Crossplane..."
helm install crossplane crossplane/crossplane \
  --namespace crossplane-system \
  --create-namespace \
  --wait

echo "Installing Kyverno..."
helm install kyverno kyverno/kyverno \
  --namespace kyverno-system \
  --create-namespace \
  --wait

echo "Components installed:"
echo "  - Crossplane: crossplane-system namespace"
echo "  - Kyverno: kyverno-system namespace"
