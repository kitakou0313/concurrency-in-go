package main

func main() {

	// channelから繰り返しの値を送信する
	input := make(chan interface{}, 0)
	output := make(chan string, 0)
	for _, s := range []string{"a", "b", "c"} {
		select {
		case <-input:
			return
		default:
			output <- s
		}
	}

	// 停止シグナルを待つ無限ループ
	for {
		select {
		case <-input:
			return
		default:
		}
	}
}
