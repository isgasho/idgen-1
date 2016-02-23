#idgen#
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

    "github.com/zhwq001/idgen"
 )


func main() {
    runtime.GOMAXPROCS(8)
    gen := idgen.NewIdGener('d')
    id     := gen.Gen()
    fmt.Printf("%d\n", id)
    //ID里面取时间
    fmt.Println(idgen.GetTimeFromId(id))
}
```

###Benchmark Detail###
```
$ go test -bench="."
PASS
BenchmarkGenDecodeEncode-4   2000000          1495 ns/op
BenchmarkGen-4               2000000          1499 ns/op
ok      github.com/zhwq001/idgen    7.867s

```
