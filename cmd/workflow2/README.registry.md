
# 1. User runs the command

You execute:

```
go run . wkf2 kbe view -f
```

Cobra resolves your `view` command and finally runs:

```go
if showFunction {
    registry := phase2.GetFnRegistry()
```

**Action:**
➡️ Your code is asking the `phase2` package for *the registry instance that contains all registered workflow functions.*


# 2. `GetFnRegistry()` returns the global, shared registry

Because you wrote this in `phase2`:

```go
var globalRegistry = &FnRegistry{
    functionMap: make(map[string]PhaseFn),
}

func GetFnRegistry() *FnRegistry {
    return globalRegistry
}
```

**Action:**
➡️ `GetFnRegistry()` returns **a pointer to a single shared global registry** created **at program startup**.

This is the key fix. Before this fix your registry was recreated each time (empty). Now it is a **true singleton**.

---

# 3. BEFORE this moment — the registry was already filled

This is the part people sometimes misunderstand.

## 3a. When your program started, Go loaded all imported packages

This includes your workflow modules such as:

```go
package ssh

import (
    "github.com/abtransitionit/gocore/phase2"
    "github.com/abtransitionit/gotask/mock/node"
)
```

## 3b. Go automatically executed all `init()` functions

This code:

```go
func init() {
    registry := phase2.GetFnRegistry()

    registry.Add("CheckSshConf", node.CheckSshConf)
    registry.Add("CheckSshAccess", node.CheckSshAccess)
}
```

**Action:**
➡️ The functions `CheckSshConf` and `CheckSshAccess` were registered in `globalRegistry` *before the Cobra command runs.*

This happens because:

* `init()` always runs at startup
* `GetFnRegistry()` returns *the* same global instance
* `Add()` stores them in the shared `map[string]PhaseFn`


# 4. So when your view command runs:

```go
registry := phase2.GetFnRegistry()
```

**Action:**
➡️ You are retrieving the **already-populated** global registry.


# 5. Your viewer builds a table from it

You iterate through the registry:

```go
keys := registry.List()
for _, key := range keys {
    fn, _ := registry.Get(key)
    ...
}
```

Because the registry was filled during `init()`, the result now contains:

```
CheckSshConf
CheckSshAccess
```


# 6. Why it was empty before (the bug)

Before you changed to a global singleton, you had something like:

```go
func GetFnRegistry() *FnRegistry {
    return &FnRegistry{functionMap: make(map[string]PhaseFn)}
}
```

This meant:

* Every call to `GetFnRegistry()` created a *new empty registry*
* `ssh.init()` used registry instance A
* `view` used registry instance B ← empty

Therefore your printed table was empty.


# 7. Your fix made everything work

By creating:

```go
var globalRegistry = &FnRegistry{
    functionMap: make(map[string]PhaseFn),
}
```

and returning it:

```go
func GetFnRegistry() *FnRegistry {
    return globalRegistry
}
```

You ensured:

* one registry
* shared by all packages
* filled during `init`
* visible when printing

That’s exactly why your output now shows the correct table:

```
│ 1 │ CheckSshConf   │ gotask/mock │ node.CheckSshConf   │
│ 2 │ CheckSshAccess │ gotask/mock │ node.CheckSshAccess │
```


# Short version for your documentation

Below is a clean block you can paste directly into your docs:


## **How the function registry is initialized (execution flow)**

1. **At startup**, Go imports all packages in the project.

2. Every imported workflow package (e.g., `ssh`) executes its `init()` function.

3. Inside each `init()`, functions are registered using:

   ```go
   phase2.GetFnRegistry().Add("Key", fn)
   ```

4. `GetFnRegistry()` returns **the single global instance**:

   ```go
   var globalRegistry = &FnRegistry{functionMap: make(map[string]PhaseFn)}
   ```

5. Therefore all calls to `Add()` store functions in the same shared map.

6. Later, when the user runs:

   ```
   wkf2 kbe view -f
   ```

   the command calls:

   ```go
   registry := phase2.GetFnRegistry()
   ```

   and retrieves the same global registry that was already populated during startup.

7. The `view` command prints the list of all registered workflow functions.

---

If you want, I can format this into a prettier, ready-to-publish Markdown page for your documentation.
