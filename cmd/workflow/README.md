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
# reference
