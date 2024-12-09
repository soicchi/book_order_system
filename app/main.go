package main

import "event_system/cmd"

// cobraを利用しているため、main関数でcmd.Execute()を呼び出す。
// これにより、コマンドライン引数を解析し、実行するコマンドを決定する。
// たとえばAPIサーバーを起動したかったら、以下のように実行する。
// ex) go run main.go server
func main() {
	cmd.Execute()
}
