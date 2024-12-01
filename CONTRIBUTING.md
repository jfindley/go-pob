# Contributing to Path of Building

## Table of contents
1. [Reporting bugs](#reporting-bugs)
2. [Requesting features](#requesting-features)
3. [Contributing code](#contributing-code)
4. [Setting up a development installation](#setting-up-a-development-installation)
5. [Setting up a development environment](#setting-up-a-development-environment)
6. [Testing](#testing)
7. [Linting](#linting)

## Reporting bugs

### Before creating an issue:
* Check that the bug hasn't been reported in an existing issue. View similar issues to the left of the submit button.
* Make sure you are running the latest version of the program. Click "Check for Update" in the bottom left corner.
* If you've found an issue with offence or defence calculations, make sure you check the breakdown for that calculation in the Calcs tab to see how it is being performed, as this may help you find the cause.

### When creating an issue:
* Select the "Bug Report" issue template and fill out all fields.
* Please provide detailed instructions on how to reproduce the bug, if possible.
* Provide a build share code for a build that is affected by the bug, if possible.
  In the "Import/Export Build" tab, click "Generate", then "Share" and add the link to your post.

Build share codes allow us to reproduce bugs much more quickly.

## Requesting features
Feature requests are always welcome. Note that not all requests will receive an immediate response.

### Before submitting a feature request:
* Check that the feature hasn't already been requested. Look at all issues with titles that might be related to the feature.
* Make sure you are running the latest version of the program, as the feature may already have been added. Click "Check for Update" in the bottom left corner.

### When submitting a feature request:
* Select the "Feature Request" issue template and fill out all fields.
* Be specific! The more details, the better.
* Small requests are fine, even if it's just adding support for a minor modifier on a rarely-used unique.

## Contributing code

### Before submitting a pull request:
* Familiarise yourself with the code base [here](docs/rundown.md) to get you started.
* There is a [Discord](https://discordapp.com/) server for **active development** on the fork and members are happy to answer your questions there.
  If you are interested in joining, send a private message to any of **Cinnabarit#1341**, **LocalIdentity#9871**, **Yamin#5575** and we'll send you an invitation.

### When submitting a pull request:
* **Pull requests must be created against the `dev` branch**, as all changes to the code are staged there before merging to `main`.
* Make sure that the changes have been thoroughly tested!

## Setting up a development installation
Note: This tutorial assumes that you are already familiar with Git.

Clone the repository using this command:
```shell
git clone -b dev https://github.com/Vilsol/go-pob.git
```

### Windows

If you are on Windows, you will semi-manually need to get all the dependencies needed for development.

The suggested way to achieve this is via WinGet (except for `golangci-lint`, it has an outdated version, please use manual step below).

```shell
winget install --id=Task.Task  -e
winget install --id=Schniz.fnm  -e
winget install --id=GoLang.Go -v "1.23.3" -e
```

If you do not have WinGet, you will need to follow the manual install instructions for each:

* Task: https://taskfile.dev/installation/
* fnm: https://github.com/Schniz/fnm?tab=readme-ov-file#installation
* Go: https://go.dev/doc/install
* golangci-lint https://golangci-lint.run/welcome/install/

After installing dependencies, you have two more steps

1. For fnm: make sure that you add it to your shell https://github.com/Schniz/fnm?tab=readme-ov-file#powershell and then run `fnm use` in this project directory
2. For node: execute `corepack enable` to be able to use `pnpm`

### Linux

If you are on Linux, you have an option to have all the dependencies automatically maintained with `devbox`.
Just follow the installation instructions here: https://www.jetify.com/docs/devbox/installing_devbox/.
If possible, you should also setup `direnv` so you don't need to run `devbox shell` every time: https://direnv.net/

If unable, then have a look at the manual Windows instructions, and install the same tools.

## Setting up a development environment

Note: This tutorial assumes that you are already familiar with the development tool of your choice.

If you want to use a text editor, [Visual Studio Code](https://code.visualstudio.com/) (proprietary) is recommended.
All you need are the [Go](https://marketplace.visualstudio.com/items?itemName=golang.go) and [Svelte](https://marketplace.visualstudio.com/items?itemName=svelte.svelte-vscode) extensions.

If you want to use an IDE instead, [GoLand](https://www.jetbrains.com/go/) (for backend) and [WebStorm](https://www.jetbrains.com/webstorm/) (for frontend) are recommended.

### Backend (Go)

To build the WASM binary, you can use the following command:

```shell
task build-go
```

To re-generate any new typings that have been exposed from Go, you can use this:
```shell
task generate
```

If you want those to run continuously while developing, you can use:
```shell
task dev-go
```

### Frontend (Svelte)

The frontend resides in the `./frontend`

You can start the dev server via:
```shell
task dev-frontend
```

You can build it via:
```shell
task build-frontend
```

## Testing

PoB uses the builtin Go testing platform together with [testza](https://github.com/MarvinJWendt/testza) framework.

All tests are run twice, first time in native Go, second time through WASM to ensure that nothing breaks on either deployment.

### Running Native Tests

Simply executing `task test` should run all tests.

### Running WASM Tests (you probably don't need this)

First ensure that you have the appropriate NodeJS version installed. (current version can be seen in [devbox.json](devbox.json))

Then you should be able to execute all tests using `./.github/wasm_test.sh` script.

## Linting

If any of these linters fail, the CI build will not pass.

### Backend (Go)

The backend is linted using [`golangci-lint`](https://golangci-lint.run/usage/install/).

You can execute it via `task lint-go` and format with `task format-go`.

### Frontend (Svelte)

The frontend is linted using `prettier` and `eslint`.

You can execute those by using `task lint-frontend` and format with `task format-frontend`.

## Setting up a PoB reference environment

You want to clone the `dev` branch of this repo https://github.com/Vilsol/PathOfBuilding (important, as it's pinned to specific commit)

After that, simply running `devbox run install` will setup emmylua and then `devbox run launch` will run PoB with a debugger.

Then add the following config in your local clone under `.vscode/launch.js` and run the debugger via VSCode UI:
```json
{
    "version": "0.2.0",
    "configurations": [
        {
            "type": "emmylua_new",
            "request": "launch",
            "name": "Debug",
            "host": "localhost",
            "port": 9966,
            "ext": [
                ".lua",
                ".lua.txt",
                ".lua.bytes"
            ],
            "ideConnectDebugger": true
        }
    ]
}
```

## Writing code

In general, follow Go's best practices.

If you are using regex, make sure that you precompile it instead of creating a new one every time. (as Go's regex implementation runs in [linear](https://github.com/golang/go/blob/1176052bb40378272cfbe83d873b65fcc2ed8502/src/regexp/regexp.go#L15-L19) time)

### Converting Code from PoB

[CONVERSION_EXAMPLES.md](CONVERSION_EXAMPLES.md)