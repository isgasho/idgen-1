#idgenter#
- 64bit的id发生器，单实例按时间永远递增，每秒最多支持生成524,287个id
- 最多支持256个实例，多实例同一秒内不能保证按时间递增，但总体是按时间递增的。
```
/***************************
*Genter a int64
*|reserved 1bit|version 4bit|timestamp 32bit|instanceid 8bit|increment 19bit|
****************************/
```
##Examples##

```go
package main

import (
    "runtime"
    "fmt"

    "idgenter"
 )


func main() {
    runtime.GOMAXPROCS(8)
    genter := idgenter.NewIdGener('d')
    id     := genter.Gen()
    fmt.Printf("%d\n", id)
    //ID里面取时间
    fmt.Println(idgenter.GetTimeFromId(id))
}
```

###Benchmark Detail###
```
$ go test -bench="."
PASS
BenchmarkGen    10000000           237 ns/op
ok      idgenter    2.620s

```
