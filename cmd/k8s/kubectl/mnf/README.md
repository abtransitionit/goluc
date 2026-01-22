# manage manifest

| Category | Operation | Command | Purpose | Safety level | Cluster-wide impact |
|--------|----------|--------|--------|-------------|--------------------|
| Lifecycle | Apply | `kubectl apply -f $yaml` | Create or update resources declaratively | 🟡 Medium | Depends on manifest (could be cluster-wide) |
| Lifecycle | Delete | `kubectl delete -f $yaml` | Delete exactly the resources defined in the manifest | 🔴 Dangerous | Depends on manifest (could be cluster-wide) |
| Inspection | Client dry-run | `kubectl apply -f $yaml --dry-run=client -o yaml` | See how kubectl interprets and normalizes the manifest | 🟢 Safe | None |
| Inspection | Server dry-run | `kubectl apply -f $yaml --dry-run=server -o yaml` | Validate against the API server, webhooks, and CRDs | 🟢 Safe | None |
| Validation | Client validation | `kubectl apply -f $yaml --dry-run=client` | Validate syntax and basic schema locally | 🟢 Safe | None |
| Validation | Server validation | `kubectl apply -f $yaml --dry-run=server` | Strict validation using cluster admission logic | 🟢 Safe | None |
| Diff | Preview changes | `kubectl diff -f $yaml` | Show what would change compared to the live cluster | 🟢 Safe | None |
| Discovery | List kinds | `kubectl apply -f $yaml --dry-run=client -o yaml \| yq '.items[].kind'` | Identify all resource kinds in the manifest | 🟢 Safe | None |
| Discovery | List objects | `kubectl apply -f $yaml --dry-run=client -o yaml \| yq '.items[] \| [.kind, .metadata.name, (.metadata.namespace // "-")] \| @tsv'` | List kind, name, and namespace | 🟢 Safe | None |
| Discovery | Cluster-scoped | `kubectl apply -f $yaml --dry-run=client -o yaml \| yq '.items[] \| select(.metadata.namespace == null) \| .kind + "/" + .metadata.name'` | Detect cluster-scoped resources | 🟢 Safe | None |
| Observability | Get resources | `kubectl get -f $yaml` | List all resources created by the manifest | 🟢 Safe | None |
| Observability | Describe resources | `kubectl describe -f $yaml` | Show events, conditions, and runtime details | 🟢 Safe | None |
| Ownership | Server-side apply | `kubectl apply -f $yaml --server-side` | Enable field ownership tracking (SSA) | 🟡 Medium | Depends on manifest |
| Ownership | Managed fields | `kubectl get <resource> -o yaml` | Inspect field ownership and conflicts | 🟢 Safe | None |
| Advanced | Replace force | `kubectl replace -f $yaml --force` | Delete and recreate resources (disruptive) | 🔴 Dangerous | Depends on manifest |
| Advanced | Prune | `kubectl apply -f $yaml --prune -l app=my-app` | Remove resources not declared in the manifest | 🔴 Dangerous | Depends on manifest |
