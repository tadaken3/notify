
# [現在作業中]Notify

Notifyはチャットワーク、LINE、Slackに通知を行うためのシンプルなCLIツールです。

## ToDo

- REEDMEの英訳
- Slack通知機能の実装
- チャットワーク通知機能の実装
- ビルドする
- ステータスをとる

## Description

Notifyは、主要なチャットツールに通知を行うためのシンプルなCLIツールです。対応しているチャットサービスはチャットワーク、LINE、Slackです。サブコマンドで、チャットツールを指定します。利用するにあたって、事前にアクセストークンを取得してください。

## Demo

## Features

- サブコマンドで、チャットサービスを指定して、メッセージとトークンを渡すだけで、通知を行えます。
- トークンは設定ファイルから読み込むこともできます。

## Usage
line LINE
chtwk チャットワーク
slack Slack
nofify -l 'access token' アクセストークンをセットします。
notify -m 'message' メッセージ本文になります。省略した場合は、デフォルトで設定されているメッセージが送信されます。

## Requirement

- Go

## Install

以下のコマンドをターミナルにペーストしてください。

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
