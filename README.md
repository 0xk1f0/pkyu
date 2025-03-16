# pkyu

> [!NOTE]  
> Very early development stage. Nothing's really implemented yet.

Podman Kubernetes YAML Utility

## Motivation

Since I moved most of my stuff from `docker` to `podman`, I always found `podman kube play` a bit cumbersome to use. As I decided to learn Go recently and needed a simple learner project anyway, this little wrapper around `podman` was born to fit my needs.

## Features

- More familiar syntax
- Additional QoL flags like `--pull`
- More to be added...

## Building

```bash
git clone https://github.com/0xk1f0/pkyu.git
cd pkyu/
go build main.go
```
