<!--
SPDX-FileCopyrightText: 2022 Kalle Fagerberg

SPDX-License-Identifier: CC-BY-4.0
-->

# hns-tree

I found that the kubectl-hns command line was too slow when printing the
tree of namespaces, so I just created a small utility script to render them
quickly.

## Install

Requires Go 1.18 or higher

```sh
go install github.com/jilleJr/hns-tree@latest
```

## Usage

```console
$ hns-tree --help
Usage of hns-tree:
  -h, --help                show this help text
      --kubeconfig string   (optional) absolute path to the kubeconfig file (default "/home/kallefagerberg/.kube/config")
  -o, --output output       output format: tree, json, or yaml (default tree)
```

```console
$ time hns-tree
default
hnc-system
kube-node-lease
kube-public
kube-system
test-root
└── test-child1
│   ├── test-child3
│   ├── test-child4
└── test-child2
    └── test-child5
hns-tree  0.03s user 0.01s system 57% cpu 0.066 total
```

Versus the time it takes for kubectl-hns:

```console
$ time kubectl-hns tree --all-namespaces
default
hnc-system
kube-node-lease
kube-public
kube-system
test-root
├── [s] test-child1
│   ├── [s] test-child3
│   └── [s] test-child4
└── [s] test-child2
    └── [s] test-child5

[s] indicates subnamespaces
kubectl-hns tree --all-namespaces  0.04s user 0.04s system 1% cpu 5.828 total
```
