# Developer notes

## Overall architecture

 - OpenTofu is used to deploy infrastructure. That includes all is necessary in order to launch Kubernetes clusters - modules should conclude producing a kubeconfig file and context
   - `tf` files in `tofu/main/` specify whole testing environments
   - `tf` files in `tofu/modules/` implement components (platform-specific or platform-agnostic)
 - the `dartboard` Golang program runs OpenTofu to create Kubernetes clusters, then Helm/kubectl to deploy and configure software under test (Rancher and/or any other component). It is designed to be idempotent
 - a Mimir-backed Grafana instance in an own cluster displays results and acts as long-term result storage

## Porting OpenTofu files to new platforms

 - create a new `tofu/main` subdirectory copying over `tf` files from `aws`
 - edit `variables.tf` to include any platform-specific information
 - edit `main.tf` to use platform-specific providers, add modules as appropriate
   - platform-specific modules are prefixed with the platform name (eg. `tofu/modules/aws_*`)
   - platform-agnostic modules are not prefixed
   - platform-specific wrappers are normally created for platform-agnostic modules (eg. `aws_k3s` wraps `k3s`)
 - adapt `outputs.tf` - please note the exact structure is expected by `dartboard` - change with care

It is assumed all created clusters will be able to reach one another with the same domain names, from the same network. That network might not be the same network of the machine running OpenTofu.

Created clusters may or may not be directly reachable from the machine running OpenTofu. In the current `aws` implementation, for example, all access goes through an SSH bastion host and tunnels, but that is an implementation detail and may change in future. For new platforms there is no requirement - clusters might be directly reachable with an Internet-accessible FQDN, or be behind a bastion host, Tailscale, Boundary or other mechanism. Structures in `outputs.tf` have been designed to accommodate for all cases, in particular:
 - `local_` variables refer to domain names and ports as used by the machine running OpenTofu,
 - `private_` variables refer to domain names and ports as used by the clusters in their network,
 - values may coincide.

`node_access_commands` are an optional convenience mechanism to allow a user to SSH into a particular node directly.

## Hacks and workarounds

In some situations we want to add code which "uncleanly" works around bugs in other software or limitations of some kind. Those can be discussed in the PR on a case-by-case basis, but they have to be documented with a comment starting with `HACK:`, so that they can be tracked later, eg.:

https://github.com/search?q=repo%3Amoio%2Fscalability-tests+HACK%3A&type=code

If at all possible, also include a condition for the removal of the hack (eg. dependency is updated to a version that fixes a certain issue).
