# go-cleanarchitecture-tasks
[cleanarchitecture-tasks](https://github.com/cloud-ace/cleanarchitecture-tasks/tree/develop/src/jp/cloudace/tech/clean/demo/tasks)をGoで実装してみたかった。結構手抜きがある。

## 動作方法
```shell script
$ go run main.go
input > [resource] [action]
```

### resource
- help: ヘルプを出力
- tasks: タスク操作

### action
- list [-c]: 一覧表示(-cオプションで完了のみ表示)
- create [args: name, description]: タスクを作成する
- complete: タスクを完了にする

### コマンド例
|処理|コマンド|
|---|---|
|タスク作成|tasks create work meeting|
|タスク一覧取得|tasks list|
|完了タスク一覧取得|tasks list -c|
|タスク完了|tasks complete [TaskID]|

## memo
- `presenters/print_presenter.go`
    - Java: printTasksというメソッドを作ってる。
    - Go: CompletedTaskとTaskが別の型なのでprintTasks実装できず。
