# drone-git-sync

[Drone](https://www.drone.io/) plugin to sync current ref to a remote `git` repository.

## Build

Build the binary with the following commands:

```sh
go build
```

## Docker

Build the docker image with the following commands:

```sh
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -a -o release/$GOOS/$GOARCH/drone-git-sync
# docker
docker build --rm -t ghcr.io/virzz/drone-git-sync .
# podman
podman build --rm -t ghcr.io/virzz/drone-git-sync .
```

## Usage

Execute from the working directory:

```sh
docker run --rm \
  -e PLUGIN_SSH_KEY="$(cat "${HOME}/.ssh/id_rsa")" \
  -e PLUGIN_REMOTE=git@github.com:foo/bar.git \
  -e PLUGIN_FORCE=false \
  -v "$(pwd):$(pwd)" \
  -w "$(pwd)" \
  ghcr.io/virzz/drone-git-sync
```

## .drone.yml

```yaml
steps:
  - name: Sync Github
    image: ghcr.io/virzz/drone-git-sync
    settings:
      remote: "ssh://git@github.com:foo/bar.git"
      force: true
      ssh_key:
        from_secret: SYNC_GITHUB_KEY
```

## References

> [drone-git](https://github.com/drone/drone-git)
> [drone-git-push](https://github.com/appleboy/drone-git-push)