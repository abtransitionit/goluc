# Purpose
- define **worflows** 

# The config YAML
This config file holds a a set of variables used by the **worflow**. To find the file :
- `Viper` first checks an environment variable `GOLUC_CONFIG` (that point to that **YAML**)
- If it’s not set → fallback to `~/wkspc/.config/goluc/workflow/conf.yaml`.
- Then it merges any project-local `./conf.yaml` (if present), so local values override global ones.

## Example
```yaml
# global config : ~/wkspc/.config/goluc/workflow/conf.yaml
workflow:
  kindn:
    node:
      all: ["o1u"]
```

```yaml
# local config: ./conf.yaml
workflow:
  kindn:
    node:
      all: ["o1u","o2a"]
```

## Precedence
**variables** in this config can be override by `cobra flags` and **OS** environment variables.
```bash
export GOLUC_WKF_KINDN_EXAMPLE_KEY="env_value"
goluc wkf kindn --example_key="flag_value"
```

|Priority|Type|Example|Comment|
|-|-|-|-|
|1|Command-line flags|`--someflag=value`|**Highest priority**; overrides everything else.|
|2|Environment variables |`GOLUC_KINDN_SOMEFLAG=value`|Overrides files but can be overridden by CLI flags. |
|3|Local config |`./conf.yaml` under `workflow.kindn:` |Specific to the project; overrides user/workflow-specific and global. |
|4|Global project config |`~/.config/goluc/workflow/kindn.yaml` |Overrides global config; user-specific. |
|5|Global config |`~/.config/goluc/workflow/conf.yaml`|**Lowest file-based priority**; shared defaults for all workflows.|
|7|Package defaults|`conf.yaml`|used if nothing else provides a value.|
|6|Default values|`cmd.Flags().String("someflag", "default")` |Fallback default; used if nothing else provides a value.|


So when the CLI runs:

* If a flag is provided, it **always wins**.
* If no flag, env var can override.
* Then local `./conf.yaml` can override global YAML file.
* Then global YAML (custom via harcoded envar `GOLUC_CONFIG` or hardcoded file `~/.config/goluc/workflow/kindn.yaml` ).
* If nothing exists, defaults apply.

flags > env vars > local `conf.yaml` > global YAML ``> global conf.yaml > defaults

## Get variables
```bash
# using viper
v.GetString("example_key")

# using sub
sub.GetString("example_key")

# using struct
workflow.kindn.example_key
```

## Set variables
```bash
workflow.kindn.example_key = "global_value".
```

- flag > env var > project-local YAML (`./conf.yaml`) > global YAML (`~/wkspc/.config/goluc/workflow/conf.yaml`) > defaults.


#  Node and NodeSet
## definition
In the workflow YAML, each phase has a field:

```yaml
Node: all
```

- A **node** represents a VM where a workflow phase will run.
- A **NodeSet** is a **named group of nodes**.


**NodeSet** are define in `conf.yaml` than referenced in `phase.yaml`. For example:

```yaml
# conf.yaml
Node:
  all:
    - worker1
    - worker2
    - dbserver

  frontend:
    - frontend1
    - frontend2
```

Then in your phase:

```yaml
# phase.yaml
Node: all   # will later expand to -> worker1, worker2, dbserver
```


## What is NodeSet resolution

It’s the process of **replacing a NodeSet name with the actual list of nodes it represents**.

### Example
In the code `ResolveNodeSets` :

```go
func (wf *Workflow2) ResolveNodeSets(v *viper.Viper, logger logx.Logger)
```

Currently, only logs something like:

```
NodeSet "all" → [worker1, worker2, dbserver]
```

After full resolution:

* Each `Phase2.Node` that uses a NodeSet name should be replaced (or mapped) to the list of nodes.
* Your workflow engine can then run the phase on all these nodes.

---

### 4️⃣ Why is this important?

Without NodeSet resolution:

* You can’t easily run a phase on multiple nodes.
* Every phase would have to manually list its nodes.
* Parallel or distributed execution becomes tricky.

With NodeSet resolution:

* You define node groups once.
* Phases can simply reference the group.
* Execution engine can iterate over actual nodes automatically.

---

💡 **Example workflow after NodeSet resolution**:

Before:

```yaml
Node: all
```

After resolution:

```go
Phase2.NodeList = []string{"worker1", "worker2", "dbserver"}
```

Then your execution engine can loop over `NodeList` for the phase.

