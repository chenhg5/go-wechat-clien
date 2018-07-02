# go-wechat-client

go-wechat go client

## install

```shell
go get 
```

## usage

```go
package main

import (
	wsdk "github.com/chenhg5/go-wechat-client"
	"fmt"
)

func main()  {
	data, err := wsdk.SetAddr("http://localhost:4001/call").SetAcid("2").WxappOauth("123")

	fmt.Println(data)
	fmt.Println(err)
}
```