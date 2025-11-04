# Howto 
## Add a workflow
- create a new folder
- update the current file `ep.go`
  - add the **import** (related to that new folder)
  - add the line in function `init()`
- into the new folder
  - add a `YAML` file that denotes the workflow [configuration file](#)
  - add a `YAML` file that denotes the workflow [file](#) itself
  - copy/past the `ep.go` file from any workflow
  - create a file (`xxx.go`) whith the same name as the folder
  - update 
    - the package name in the file `ep.go` 
    - the content of these 2 files (cf. any workflow)
    - **BP**: when possible `Phase.Name = Phase.Fn`

### Rules
- `phase:name` **must be** uniq in a workflow file

## Manage a workflow
Use the CLI from `goluc wkf ...`

## Terminology
ebay (**e**phemeral **Ba**stion **y**et)
  - A hardened secured server 
  - Access internal infrastructure via SSH (or RDP).