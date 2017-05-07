# Notify
[![Build Status](https://secure.travis-ci.org/tadaken3/notify.png)](https://secure.travis-ci.org/tadaken3/notify)

Notify is simple CLI application to nottify any chat serviece.

## Description

Notify is a simple CLI tool for notifying major chat tools. Supported chat service is chat work, LINE. Specify the chat tool with the subcommand. Please obtain an access token before using it.

## Demo



## Features

- You can specify a chat service with a subcommand and notify by just passing a message and a token
- The token can also be read from the configuration file.

## Exsamle

```bash
$ notify line -m "sent message to line"
$ notify chtwk -m "sent message to chatwork"
```

## Usage

```bash
This application is a simple CLI tool.
You can quickly notify to any chat serviece

Usage:
  notify [command]

Available Commands:
  chtwk       notify to chatwork
  help        Help about any command
  line        notify to line

Flags:
      --config string    config file (default is $HOME/.notify.yaml)
  -h, --help             help for notify
  -m, --message string   message (default "This message is from notify")
  -t, --token string     access token

Use "notify [command] --help" for more information about a command.
```

## Requirement

set your credentials in $HOME/.notify.yaml/credentials :

[default]
line_access_token: XXXXXX
chatwork_access_token = XXXXXX
room_id = XXXXXX

## Install

If you have Go development environment:
```bash
$ go get github.com/tadaken3/notify
```

## Contribution

1. Fork ([https://github.com/tadaken3/notify](https://github.com/tadaken3/notify))
2. Create a feature branch
3. Commit your changes
4. Rebase your local changes against the master branch
5. Run test suite with the `go test ./...` command and confirm that it passes
6. Run `gofmt -s`
7. Create new Pull Request

## Licence

[MIT](https://github.com/tcnksm/tool/blob/master/LICENCE)

## Author

[tadaken3](https://github.com/tadaken3)
