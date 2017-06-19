# Notify
[![Build Status](https://secure.travis-ci.org/tadaken3/notify.png)](https://secure.travis-ci.org/tadaken3/notify)


Notifyはチャットワーク、LINE、Slackに通知を行うためのシンプルなCLIツールです。

## Description

Notifyは、主要なチャットツールに通知を行うためのシンプルなCLIツールです。対応しているチャットサービスはチャットワーク、LINEです。サブコマンドで、チャットツールを指定します。利用するにあたって、事前にアクセストークンを取得してください。

## Demo

![demo](https://github.com/tadaken3/notify/blob/master/demo.gif)

## Features

- サブコマンドで、チャットサービスを指定して、メッセージとトークンを渡すだけで、通知を行えます。
- トークンは設定ファイルから読み込むこともできます。

## Exsamle

```bash
$ notify line -m "sent message to line"
$ notify chtwk -m "sent message to chatwork"
$ notify slack -m "sent message to slack"
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
  slack       notify to slact
  version     Print the version number of notify

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

[MIT](https://github.com/tadaken3/notify/blob/master/LICENSE)

## Author

[tadaken3](https://github.com/tadaken3)
