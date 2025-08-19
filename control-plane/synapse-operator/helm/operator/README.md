# Synapse Operator Helm Chart

This Helm chart deploys the Synapse Operator on a Kubernetes cluster.

## Prerequisites

- Kubernetes 1.19+
- Helm 3.0+

## Installation

```bash
# Add the repository (once published)
helm repo add synapse https://charts.synapse.io
helm repo update

# Install the chart
helm install synapse-operator synapse/synapse-operator

# Or install from local files
helm install synapse-operator ./helm/synapse-operator
```

## Configuration

The following table lists the configurable parameters of the Synapse Operator chart and their default values:

| Parameter | Description | Default |
|-----------|-------------|---------|
| `replicaCount` | Number of replicas | `1` |
| `image.registry` | Image registry | `ghcr.io` |
| `image.repository` | Image repository | `mdaops/synapse-operator` |
| `image.tag` | Image tag | `v0.0.1` |
| `image.pullPolicy` | Image pull policy | `IfNotPresent` |
| `serviceAccount.create` | Create service account | `true` |
| `rbac.create` | Create RBAC resources | `true` |
| `leaderElection.enabled` | Enable leader election | `true` |
| `metrics.enabled` | Enable metrics | `true` |
| `metrics.port` | Metrics port | `8080` |
| `healthProbe.port` | Health probe port | `8081` |
| `resources.limits.cpu` | CPU limit | `500m` |
| `resources.limits.memory` | Memory limit | `128Mi` |
| `resources.requests.cpu` | CPU request | `10m` |
| `resources.requests.memory` | Memory request | `64Mi` |

## Custom Resource Definitions

The chart installs the following CRDs:
- `setups.synapse.io`

## Uninstalling

```bash
helm uninstall synapse-operator
```

Note: CRDs are not automatically removed. Remove them manually if needed:

```bash
kubectl delete crd setups.synapse.io
```