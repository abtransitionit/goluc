# `uidmap`
## What is `uidmap`
- a file located at `/proc/self/uid_map`
- allows 
	- to restrict **unprivileged uers** to access to the entire system like a `root` user
  - **unprivileged users** to be a **virtual root user** within their own namespace
- means that a process with a UID of 0 inside the container doesn't have root privileges on the host system.
- the uiimap file:
```ìni
0 0      4294967295
```

- the uiimap file:
```ìni
0 100000 65536
```
  - tells the Linux kernel how to translate the user IDs. It says:
		- `0`: is The starting UID inside the container.
		- `100000`: is The starting UID on the host system.
		- `65536`: is The number of UIDs to map.
		- Container's UID 0 maps to Host's UID 100000.
		- Container's UID 1 maps to Host's UID 100001.
		- ...
		- Container's UID 65535 maps to Host's UID 165535.


## manage

The most fundamental way to manage UID mappings is by editing two key files as the root user:

* **/etc/subuid**: This file specifies the **subordinate UIDs** that a user is allowed to use. Each line defines a range of UIDs that can be mapped inside a user namespace. The format is `username:start_uid:uid_count`.
* **/etc/subgid**: This file does the same thing for **subordinate GIDs** (group IDs).

For example, a line like `myuser:100000:65536` in `/etc/subuid` would grant the user `myuser` the ability to use a range of 65,536 UIDs, starting from `100000`, in any user namespace they create.


## tools

While you can manually edit those files, there are dedicated commands that do it for you safely, ensuring the ranges don't overlap:

* **`usermod`**: 
	- Can be used to add or remove subordinate UID and GID ranges for a user. 
	- Example: `usermod --add-subuids 100000-165535 myuser` would add the specified UID range to the user.
* **`newuidmap`** and **`newgidmap`**: 
	- These are low-level tools. they are ussually called by programs like container runtimes to configure a new namespace.
	- Set the actual UID and GID mappings for a new user namespace. 
	- part of the package `uidmap`


### 3. Container Management CLI

Most of the time, you'll manage UID mappings indirectly through container-specific command-line flags. Container runtimes like Podman and Docker leverage the `subuid` and `subgid` files to implement their "rootless" mode.

* **Podman**: You can use the `--userns` flag with Podman to automatically handle UID mapping. For example, `--userns=keep-id` will map your host user ID to the same ID inside the container, while `--userns=auto` automatically generates a UID mapping based on the ranges available in your `/etc/subuid` file. You can also use `--uidmap` for more explicit control over the mapping.
* **Docker**: Docker has a feature called "User Namespace Remapping" that you can enable by editing the `daemon.json` configuration file or using the `--userns-remap` flag. This feature maps the `root` user inside a container to an unprivileged user on the host, using the ranges defined in `/etc/subuid` and `/etc/subgid`.