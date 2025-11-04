# Howto 
## Add a workflow
- create a folder
- add a `YAML` file that denotes the workflow [configuration file](#)
- add a `YAML` file that denotes the workflow [file](#) itself
- copy/past the `ep.go` file from any workflow
- create a file (`xxx.go`) whith the same name of the folder
- update the package name in the file `ep.go` 
- update the content of these 2 files (cf. any workflow)
  - **BP**: when possible `Phase.Name = Phase.Fn`

### Rules
- `phase:name` **must be** uniq in a workflow file

## Manage a workflow
Use the CLI from `goluc wkf ...`

