# hns-tree

I found that the kubectl-hnc command line was too slow when printing the
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
test-root
└── test-child2
│   ├── test-child5
└── test-child1
    └── test-child4
    └── test-child3
kube-system
kube-public
kube-node-lease
hnc-system
default
hns-tree  0.03s user 0.01s system 57% cpu 0.066 total
```
