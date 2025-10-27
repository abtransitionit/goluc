# `LoadConfig`

* **Create** an instance **Viper**.
* **Define Global Path:** Determine the path for the global configuration file (`conf.yaml`):
    * Checks the **`GOLUC_CONFIG`** environment variable.
    * If unset, defaults to **`~/wkspc/.config/goluc/workflow/conf.yaml`**.
* **Load Global Config:** 
    * Load the global configuration file **if it exists**.
    * **But** if it contains an **invalid format**, an **error is returned**.
    * If **not found**, the code **skips loading** and continues **without error**.
* **Define Local Path:** Set the local configuration file path to **`conf.yaml`** (relative to the current working directory).
* **Load Local Config:** 
    * Load the local configuration file **if it exists**.
    * If found, **merged** it with the global, **overriding** any conflicting values.
    * If **not found**, the code **skips loading** and continues without error.
* **Return:** Return the final, merged `*viper.Viper` instance.