# Goluc
This repository is a `GO` project that provide a `Linux` CLI named `LUC` (aka. **L**inux **U**nified **C**LI). 


# Badges

[![Dev CI](https://github.com/abtransitionit/gotest/actions/workflows/ci-dev.yaml/badge.svg?branch=dev)](https://github.com/abtransitionit/gotest/actions/workflows/ci-dev.yaml)
[![Main CI](https://github.com/abtransitionit/gotest/actions/workflows/ci-main.yaml/badge.svg?branch=main)](https://github.com/abtransitionit/gotest/actions/workflows/ci-main.yaml)
[![LICENSE](https://img.shields.io/badge/license-Apache_2.0-blue.svg)](https://choosealicense.com/licenses/apache-2.0/)

----




---

# Getting Started  

## Developper
- the project uses codes form multiple GO projects (e.g : `gocore`, `golinux`, `gotask`).

## 1. Create a repository from this template (e.g. `gomine`)
- on `github.com` create `gomine` : an empty git repo without `README` and `.gitcore`
- git clone the tpl repo: `gotplrepo` into `gomine`
```shell
git clone https://github.com/abtransitionit/gotplrepo.git gomine
```
- reset history and init repo
```shell
cd gomine
rm -rf .git
git init -b main  
```
- update `GO` path in the file `go.mod`
```shell
# do update
go mod init github.com/abtransitionit/gomine
# check updtae
cat go.mod
```
- commit the code "initial setup from template"
- update .git/config
```shell
git remote add origin https://github.com/abtransitionit/gomine.git
```
- push the code
```shell
git push -u origin main
```

## 2. Update the README
- update `gotplrepo` to `gomine`
- review each sections and update/add content when needed


---

# Contributing  

We welcome contributions! Before participating, please review:  
- **[Code of Conduct](.github/CODE_OF_CONDUCT.md)** – Our community guidelines.  
- **[Contributing Guide](.github/CONTRIBUTING.md)** – How to submit issues, PRs, and more.  


----


# Release History & Changelog  

Track version updates and changes:  
- **📦 Latest Release**: `vX.X.X` ([GitHub Releases](#))  
- **📄 Full Changelog**: See [CHANGELOG.md](CHANGELOG.md) for detailed version history.  

---

# Todo

```go
var rootSDesc = "LUC  is a user-friendly, auto-documented command-line interface."
var rootLDesc = rootSDesc + ` It simplifies daily tasks for DevOps engineers and developers by providing a unified and consistent CLI experience. LUC can, for example:
	→ Manage containers and container images,
	→ Manage Linux OS packages and repositories using a unified interface — no need to worry about whether it's apt or dnf nor if it's debian, fedora or ubuntu
	→ Manage remote VM objects,
	→ Simplify the creation and management of Kubernetes clusters across virtual machines,
	→ ...and much more.

As a Linux cross-distribution CLI, LUC is also well-suited and ready for full automation and integration into any CI/CD pipelines.`
``
