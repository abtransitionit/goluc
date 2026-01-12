
# Kubernetes Cluster — First-Class TODOs

This document defines 
  - **what remains to be done** once a `KBE` cluster is created
  - organize a set of domain around **first-class Kubernetes concerns**
  - Each domain defines a clear responsibility and scope.

## The state after this workflow is palyed
- Pods can talk to each other
- Nothing can be reached from the **internet** unless 
  - being exposed from **in**side.
  - being reachable **out**side.


# 1️⃣ Domains Overview

| Domain   | Acronym meaning                                 | Name            | Purpose                         | Core Question                                        |
| -------- | ----------------------------------------------- | --------------- | ------------------------------- | ---------------------------------------------------- |
| **kco**  | **K**ubernetes **C**ore **O**perations          | Core Operations | Cluster foundations & lifecycle | *Is the cluster well-defined, stable, and operable?* |
| **kex**  | **K**ubernetes **Ex**tensions                   | Extensions      | Functional capabilities         | *What capabilities does the cluster gain?*           |
| **kobe** | **K**ubernetes **Ob**servability **E**xtensions | Observability   | Visibility & diagnosis          | *Can we see what’s happening?*                       |
| **kse**  | **K**ubernetes **S**ecurity **E**xtensions      | Security        | Trust & enforcement             | *What is allowed, trusted, and protected?*           |

These domains are **orthogonal by design**:

* Extensions add features
* Observability explains behavior
* Security constrains behavior
* Core operations keep everything coherent over time


# 2️⃣ Domain Definitions

## ✅ kco — Kubernetes Core Operations

This domain defines the **baseline contract** of the cluster.
Everything else depends on it.

Scope includes:

* Cluster bootstrap rules
* Node roles, labels, taints, annotations
* Namespace layout and ownership
* Resource quotas and default limits
* Upgrade and versioning policy
* etcd backup and restore strategy
* Operational conventions (naming, environments)

> kco answers: *“What kind of cluster is this?”*


## ✅ kex — Kubernetes Extensions

Extensions add **capabilities** to the cluster without redefining its core.

Characteristics:

* Feature-oriented
* Replaceable
* Explicitly enabled or disabled

Typical examples:

* Ingress controllers
* Gateway API
* Storage backends
* Load-balancer integrations
* Service mesh features
* CSI / CNI add-ons beyond the baseline

📌 *kex* is about **feature enablement**, not policy.


## ✅ kobe — Kubernetes Observability Extensions

This domain ensures the cluster is **observable**, not just running.

Minimal expectations:

* Metrics
* Logs
* Basic event visibility

Typical components:

* `metrics-server`
* `kube-state-metrics`
* Log access (`kubectl logs`)
* Optional but powerful:

  * Cilium Hubble (CLI / UI)
  * Tracing backends

> kobe answers: *“Can we understand failures and behavior?”*


## ✅ kse — Kubernetes Security Extensions

Security is a **first-class domain**, not an afterthought.

Scope includes:

* Network policies (Cilium)
* Default-deny strategies
* Pod Security (standards / admission)
* RBAC and service accounts
* Secrets management
* Admission control and validation
* Image policies and provenance

> kse answers: *“What is allowed to happen, and what is forbidden?”*

# `ingress` in the cluster
# Role of an `ingress`
```sh
# generic
Internet - > [Ingress Controller] - > [Service] - > [Pod]

# nginx
Internet - > [NGINX Pod] - > [kube-proxy / Cilium] - > [Service → Pod]


# cilium
Internet - > [Envoy (Cilium-managed)] - > [Cilium L7 routing] - > [Service → Pod]
```


Ingress = HTTP(S) reverse proxy that **route** traffic based on **host/path**:
- listens on ports (80 / 443)
- terminates TLS

## open source possible ingress for K8s
|name|comment|
|-|-|
|Traefik|popular, simple, **legacy**|
|HAProxy|**legacy**|
|NGINX Ingress Controller|production grade, simple, No CNI-specific|
|Cilium Ingress|production grade, **modern**|

# Test ingress used in the cluster
```sh
kubectl get ingress -A
kubectl get pods -A | grep -i ingress
kubectl get ingressclass
```
# Todo
## dig
- NodePort
- MetalLB
- cloud LB
- NodePort vs MetalLB decision

## Practical Mapping (GitOps-friendly)

```text
cluster/
├── kco/
│   ├── namespaces/
│   ├── quotas/
│   ├── node-labels/
│   └── backups/
├── kex/
│   ├── ingress/
│   ├── storage/
│   └── gateway-api/
├── kobe/
│   ├── metrics/
│   ├── logs/
│   └── tracing/
└── kse/
    ├── network-policy/
    ├── rbac/
    └── pod-security/
```


## 4️⃣ Baseline Validation (Mandatory)

Before adding features, validate **cluster fundamentals**.

### Validate Cilium datapath + CoreDNS

```bash
kubectl run test --image=busybox -- sleep 1h
kubectl exec -it test -- sh
```

Inside the pod:

```sh
nslookup kubernetes.default
ping <another-pod-ip>
```

✅ Confirms:

* Pod scheduling
* DNS resolution
* Pod-to-pod networking

---

## 5️⃣ Define the Cluster Contract (Early & Explicit)

Before expanding the cluster, decide and document:

* **Ingress model**
  Cilium Ingress / NGINX / Gateway API / none

* **Storage model**
  local-path / Longhorn / external / none

* **Load-balancing model**
  NodePort / MetalLB / cloud LB (OVH) / none

* **OS & kernel tolerance**
  Are all nodes first-class or best-effort?

📌 Decisions can evolve, but **undecided clusters rot quickly**.

---

## 6️⃣ Lock Down Networking Early (Cilium Advantage)

Networking policy should be applied **before applications multiply**.

Minimum steps:

* Enable default-deny (namespace or cluster-wide)
* Write at least one real NetworkPolicy
* Explicitly test blocked traffic

Mindset:

> “Nothing talks unless explicitly allowed.”

Retrofitting this later is painful.

---

## 7️⃣ Node Hygiene (Heterogeneous Cluster)

Your cluster runs:

* Ubuntu
* Fedora
* Debian
* Rocky
* Alma

First-class actions:

* Label nodes by OS, role, stability
* Decide where control-plane and critical workloads may land

```bash
kubectl label node vps-xxxx os=ubuntu
```

This avoids subtle kernel / eBPF issues later.

---

## 8️⃣ Backup Before You Need It

At minimum:

* Regular etcd snapshots
* Stored **off-cluster**
* Restore procedure documented

If you do nothing else, do this.

---

## 9️⃣ Source of Truth (GitOps Lite)

Keep it boring and explicit:

* One repository
* YAML or Helm values
* No manual drift

Even this is enough:

```text
cluster/
  cilium/
  core/
  apps/
```

Future-you will thank present-you.

---

## TL;DR — The 5 Non-Negotiables

1. ✅ Validate pod networking and DNS
2. 🧭 Define ingress / storage / LB decisions on paper
3. 🔒 Apply at least one real NetworkPolicy
4. 👀 Install minimal observability
5. 🏷 Label and classify nodes

