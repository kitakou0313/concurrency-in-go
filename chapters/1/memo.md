# 1章 並行処理入門

## 並行処理で起きる問題
### 競合状態
2つ以上の操作が正しい順番で行われなければならないところで順序が保障されていなかったときに発生する問題．

#### データ競合
- Read, Writeの実行順序が保障されていなかった時に起きる問題
- 同一のコードを実行しても，結果が定まらない

```
package main

import "fmt"

func main() {
	var data int

	go func() {
		data++ // Write
	}()

	if data == 0 { // Read 1
		fmt.Println("The val is %v") // Read 2
	}
}

```

- 上記のサンプルであれば3通りの結果が考えられる
    - WriteがRead1, Read2の処理のどこで実行されるかによる
    - 例
        - なにも表示されない
        - 1と表示
        - 0と表示
