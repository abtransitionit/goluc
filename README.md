# gotplrepo

This repository serves as a standardized template for all future GitHub GO projects within the organization.  

[![LICENSE](https://img.shields.io/badge/license-Apache_2.0-blue.svg)](https://choosealicense.com/licenses/apache-2.0/)

----


# Features  
This project template includes the following components:  


|Component|Description|
|-|-|
|Licensing|Predefined open-source license (Apache 2.0) for legal compliance.|
|Code of Conduct| Ensures a welcoming and inclusive environment for all contributors.|  
|README|Structured documentation template for clear project onboarding.|  



---

# Getting Started  

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

