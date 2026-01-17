
# container Image source of trust


| Requirement              | Why                                     |
| ------------------------ | --------------------------------------- |
| Private registry         | Control what enters the cluster         |
| Immutable tags / digests | No “latest” drift                       |
| Admission enforcement    | Prevent bypass                          |


**purpose**:
- use a private cluster container registry
- installdeploy it inside the cluster for the `POC`.
- use **admission control policy** to enforce that only images from it can run inside the cluster.

**the registry**:
- the Docker `registry:2` 

👉 **`registry:2` (Docker Distribution)**
Yes, it’s boring. That’s why it’s perfect here.


**deployement**
- option
  * `Deployment` or `StatefulSet`
  * `PersistentVolume` (local-path, Ceph, Longhorn, etc.)
  * TLS (self-signed or real cert)
  * Auth (basic or token)

- constraints
  * **No NodePort**
  * **ClusterIP + Ingress**
  * **TLS required** (even internally)

Example architecture:

```
registry.kube-system.svc.cluster.local
```



**admission policy**


> If the image does not come from *that registry*, it does not run.


```yaml
apiVersion: admissionregistration.k8s.io/v1
kind: ValidatingAdmissionPolicy
metadata:
  name: restrict-image-registries
spec:
  failurePolicy: Fail
  matchConstraints:
    resourceRules:
    - apiGroups: [""]
      apiVersions: ["v1"]
      operations: ["CREATE","UPDATE"]
      resources: ["pods"]
  validations:
  - expression: "object.spec.containers.all(c, c.image.startsWith('registry.mycorp.local/'))"
    message: "Images must come from registry.mycorp.local"
```

**advantages**
- No extra pods
- No webhook latency
- Declarative
- Easy to audit

## Steps

* Deploy 
* Deploy registry **`registry:2`** as `ClusterIP`
* Add **ValidatingAdmissionPolicy**
* Enforce `registry.mycorp.local/*`
* Use digests

after that:
* `Pods` pull images **from inside the cluster**
* CI runs **inside the cluster** (or via VPN / SSH)

## Todo
Once the registry is in place, make images trustworthy

**Short-term**

* **Digest-only deployments**

  ```yaml
  image: registry.mycorp.local/app@sha256:...
  ```
* Disable `latest`

**Midle-term**
* Add TLS + auth hardening
* CI pushes only
* Read-only pull credentials in cluster

**Medium-term**
* **cosign** for image signing
* Verify signatures in admission
* Still works with `registry:2`

**Long-term**

* Replace registry with **Harbor** or **nexus**
* Keep the same admission logic





If you want, next we can:

* Design the **exact registry deployment YAML**
* Write the **full admission policy with namespace exceptions**
* Decide **inside vs outside cluster** based on your OVH topology

This is a very solid direction 👌
