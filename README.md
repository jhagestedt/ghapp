# `ghapp --help`

Lightweight `CLI` to create GitHub installation tokens for GitHub Apps.

```
NAME:
   ghapp - GitHub App CLI

USAGE:
   ghapp [global options] command [command options] [arguments...]

COMMANDS:
   token, t  Create a GitHub App installation token
   help, h   Shows a list of commands or help for one command
```

## 🚀 Install

The packaged binaries of `ghapp` can be found on the [releases](https://github.com/jhagestedt/ghapp/releases) of this repository.

#### 🐧 Linux

```bash
curl -L "https://github.com/jhagestedt/ghapp/releases/latest/download/ghapp_linux_amd64" \
-o ./ghapp && chmod +x ./ghapp
```

#### 🍏 MacOS

```bash
curl -L "https://github.com/jhagestedt/ghapp/releases/latest/download/ghapp_darwin_amd64" \
-o ./ghapp && chmod +x ./ghapp
```

## 🧑‍💻 Usage

A GitHub App installation token can be created by the GitHub App id, the installation id and the private key like described in the [docs](https://docs.github.com/en/developers/apps/building-github-apps/authenticating-with-github-apps).

To create an installation token with `ghapp` the private key can be injected via file or as environment variable.

The GitHub App id and installation id can be passed as flags or environment variables.

```bash
# load private key from file
ghapp token --id 123 --install-id 123456 --private-key-file .ghapp-private-key.pem

# load private key from environment variable
export GHAPP_PRIVATE_KEY=$(cat .ghapp-private-key.pem)
ghapp token --id 123 --install-id 123456
```

A usage on GitHub Actions could look like the following.

```yaml
- id: setup-ghapp
  run: |
    curl -L "https://github.com/jhagestedt/ghapp/releases/latest/download/ghapp_linux_amd64" -o ./usr/local/bin/ghapp
    chmod +x ./usr/local/bin/ghapp
- id: ghapp-install-token
  run: |
    GHAPP_INSTALL_TOKEN=$(ghapp token)
    echo "::set-output name=ghapp_install_token::${GHAPP_INSTALL_TOKEN}"
  env:
    GHAPP_ID: 123
    GHAPP_INSTALL_ID: 123456
    GHAPP_PRIVATE_KEY: ${{ secrets.GHAPP_PRIVATE_KEY }}
```

## 🐳 Docker

There is also a docker container published at [Dockerhub](https://hub.docker.com/repository/docker/jhagestedt/ghapp) that contains the binary that can be used with `docker run` to not install it locally.

```bash
export GHAPP_PRIVATE_KEY=$(cat .ghapp-private-key.pem)
docker run --rm -e GHAPP_PRIVATE_KEY jhagestedt/ghapp token --id 123 --install-id 123456
```
