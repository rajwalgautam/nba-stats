# nba-stats

## Setting up local dependencies

This repo requires the following to be installed to run locally -
* Rancher desktop https://docs.rancherdesktop.io/getting-started/installation/
* Kubernetes (through HomeBrew or https://kubernetes.io/releases/download/)

Once the above are installed, run `make local-deps` in order to install local dependencies to the `rancher-desktop` namespace in your local Kubernetes environment. `make clean-local-deps` will clean up all dependencies in the namespace.