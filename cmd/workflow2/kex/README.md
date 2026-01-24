
# container Image source of trust


| Requirement              | Why                                     |
| ------------------------ | --------------------------------------- |
| Private registry         | Control what enters the cluster         |
| Immutable tags / digests | No “latest” drift                       |
| Admission enforcement    | Prevent bypass                          |


**purpose**:
- use a private cluster container registry
- deploy it inside the cluster for the `POC`.
- use **admission control policy** to enforce that only images from it can run inside the cluster.

**the registry**:
- the Docker `registry:2` 



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

## phases
**Step 1**

* Deploy registry **`registry:2`** as `ClusterIP`
* TLS enabled
* Basic auth
* Push/pull works

**Step 2**

* Add ValidatingAdmissionPolicy`
* Enforce by only allow `registry.mycorp.local/*`
* Use digests


**Step 3** : Access strategy

* If needed: temporary NodePort
* Or SSH port-forward
* Or CI inside cluster



after that:
* `Pods` pull images **from inside the cluster**
* CI runs **inside the cluster** (or via VPN / SSH)

# 🟡 Step 1
## phase **auth**
create a secret that can be used 
- as a `.htpasswd` file for a website or a private docker registry
- by an ingress controller

**create a hashed secret**
```sh
# define var
lNs="cimreg"
lUserName=$lNs
lOutFile="/tmp/auth"

# create a pwd and save it in the file /tmp/auth (a name recognized by most ingress)
kubectl run htpasswd-gen \
  -n $lNs  \
  --restart=Never --rm -i --tty \
  --image=httpd:2.4-alpine \
  -- sh -c "apk add -q --no-cache apache2-utils && htpasswd -Bbn $lUserName $(head -c 20 /dev/urandom | base64)" > $lOutFile
```
action:
- creates a temporary Kubernetes pod (instanciate from **cim** `httpd:2.4-alpine` )
- Inside that container, installs `apache2-utils`, which contains the htpasswd tool.
- generates a secure, encrypted password entry 
- immediately deletes that pod so there’s no footprint left behind.


```sh
# define var
lNs="cimreg"
lUserName=$lNs
lOutFile="/tmp/auth"

# create a pwd and output it
kubectl run htpasswd-gen \
  --image=httpd:2.4-alpine \
  --restart=Never --rm -i --tty \
  -n $lNs  \
  -- sh -c "apk add -q --no-cache apache2-utils && htpasswd -Bbn $lUserName $(head -c 20 /dev/urandom | base64) > $lOutFile && cat $lOutFile"
```
action:
- same as above
- output something like `cimreg:$2y$05$5E7...`

**create a k8s secret** from this hash secret

```sh
# define var
lNs="cimreg"
lUserName=$lNs
lK8sSecretName="registry-htpasswd"

# create secret from copy/past
kubectl create secret generic $lK8sSecretName \
  -n $lNs \
  --from-literal=auth='<PASTE_HTPASSWD_LINE_HERE>'
```
```sh
# define var
lNs="cimreg"
lUserName=$lNs
lK8sSecretName="ingress-basic-auth-secret"

# create secret from file
kubectl create secret generic $lK8sSecretName --from-file=$lOutFile \
  -n $lNs \
  --from-file=$lOutFile
```


**create and save all at once**
```sh
lNs="cimreg"
lUserName=$lNs
lK8sSecretName="registry-htpasswd"

kubectl run htpasswd-gen \
  --image=httpd:2.4-alpine \
  --restart=Never --rm \
  -n $lNs 
  --quiet=true \
  -- sh -c "apk add -q --no-cache apache2-utils && htpasswd -Bbn $lUserName $(head -c 20 /dev/urandom | base64)" \
| kubectl create secret generic $lK8sSecretName \
  -n cimreg \
  --quiet=true \
  --from-file=auth=/dev/stdin
```
**Bonus example**: use that secret for nginx ingress
```yaml
metadata:
  name: my-app-ingress
  annotations:
    # This tells the Ingress to use basic authentication
    nginx.ingress.kubernetes.io/auth-type: basic
    # This points to the secret you just created
    nginx.ingress.kubernetes.io/auth-secret: ingress-basic-auth-secret
    # This is the message the user sees in the login popup
    nginx.ingress.kubernetes.io/auth-realm: "Authentication Required"
```

# 🟡 Step 2
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



# Terminology
## htpasswd file
**definition**

- a txt file
- a set of username/password i nthe forn `user:HasfedPwd`
- 1 per line
- commonly used for basic Authentification using username/pwd by
  - nginx ingress
  - docker registry `registry:2`
  - website


**Example**: Nginx Ingress

You need the secret key to be `auth` and the secret type should be generic:

```yaml
apiVersion: v1
kind: Secret
metadata:
  name: basic-auth
  namespace: mynamespace
type: Opaque
data:
  auth: <base64>
```

Then in Ingress:

```yaml
nginx.ingress.kubernetes.io/auth-type: basic
nginx.ingress.kubernetes.io/auth-secret: basic-auth
nginx.ingress.kubernetes.io/auth-realm: "Protected"
```

> ⚠️ Nginx expects the secret to contain:

```
auth: <htpasswd content>
```


**Example**: Docker Registry

Your htpasswd file must be mounted as `/auth/htpasswd` (example with registry):

```yaml
volumeMounts:
- name: auth
  mountPath: /auth

volumes:
- name: auth
  secret:
    secretName: basic-auth
```

Then registry config:

```yaml
auth:
  htpasswd:
    realm: basic-realm
    path: /auth/htpasswd
```

> ⚠️ Docker expects a file called `htpasswd` mounted in the registry container.

**Example**: website

Same: mount secret and configure your web server to use it.




# Todo Todo Todo Todo Todo Todo Todo

## What your registry actually needs

A Docker registry with basic auth only needs:

* a **bcrypt hash** (`user:hash`)
* stored in a **Kubernetes Secret**
* mounted read-only into the registry pod

👉 **Nobody ever needs to know the password**, including you.

So generating a random one and throwing it away is not just OK — it’s ideal.



### 1️⃣ Generate htpasswd in a throwaway pod

```sh
kubectl run htpasswd-gen \
  -n registry \
  --quiet --restart=Never --rm -i \
  --image=httpd:2.4-alpine \
  -- sh -c '
    apk add -q --no-cache apache2-utils &&
    htpasswd -Bbn registry-user \
      $(head -c 20 /dev/urandom | base64)
  '
```

This outputs:

```
registry-user:$2y$05$...
```

No files. No password. No traces.

---

### 2️⃣ Pipe directly into a Secret (best practice)

```sh
kubectl run htpasswd-gen \
  -n registry \
  --quiet --restart=Never --rm -i \
  --image=httpd:2.4-alpine \
  -- sh -c '
    apk add -q --no-cache apache2-utils &&
    htpasswd -Bbn registry-user \
      $(head -c 20 /dev/urandom | base64)
  ' \
| kubectl create secret generic registry-auth \
    -n registry \
    --from-file=htpasswd=/dev/stdin \
    --dry-run=client -o yaml \
| kubectl apply -f -
```

✔ idempotent
✔ no password exposure
✔ reproducible automation

---

## How the registry consumes it

### Secret

```yaml
apiVersion: v1
kind: Secret
metadata:
  name: registry-auth
type: Opaque
data:
  htpasswd: <base64>
```

---

### Registry Deployment snippet

```yaml
volumeMounts:
  - name: auth
    mountPath: /auth
    readOnly: true

volumes:
  - name: auth
    secret:
      secretName: registry-auth
```

---

### Registry config

```yaml
env:
  - name: REGISTRY_AUTH
    value: htpasswd
  - name: REGISTRY_AUTH_HTPASSWD_REALM
    value: Registry Realm
  - name: REGISTRY_AUTH_HTPASSWD_PATH
    value: /auth/htpasswd
```

That’s it. Docker registry will just work.

---

## How clients authenticate (inside cluster)

Other pods use a **dockerconfigjson** secret:

```sh
kubectl create secret docker-registry registry-pull \
  -n my-namespace \
  --docker-server=registry.registry.svc.cluster.local:5000 \
  --docker-username=registry-user \
  --docker-password='(the password you *never* stored)'
```

⚠️ That’s why this registry is usually:

* push-only from CI (same pipeline that generated the password)
* pull-only via service accounts you control

If you **don’t want any human or pod to log in manually**, you can instead:

* inject the same generated password into a second secret
* or use token-based auth instead

---

## If this registry is truly internal

Strong recommendation:

* restrict access with **NetworkPolicies**
* don’t expose it via Ingress
* keep auth simple and boring
* rotate the secret by re-running the same command

---

## Bottom line

For a custom **in-cluster private registry**:

* ✅ `htpasswd -Bbn`
* ✅ random password from `/dev/urandom`
* ✅ throwaway pod
* ✅ pipe to secret
* ❌ no human-known password

You’re doing the *right* thing.

If you want next steps, I can help you with:

* CI push credentials flow
* rotating auth without downtime
* switching to token auth later
* mirroring images into this registry cleanly

Just say where you want to go next 🚀
