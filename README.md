#idgen#
- 64bit的id发生器，单实例按时间永远递增，每秒生成2^16~2^17个id
- 最多支持256个实例，多实例同一秒内不能保证按时间递增，但总体是按时间递增的。
```
/***************************
*Genter a int64
*|reserved 1bit|timestamp 32bit|instanceid 8bit|bid 6bit|increment 17bit|
****************************/
```
##Examples##

```go
package main

import (
    "fmt"

    "github.com/zhwq001/idgen"
 )


func main() {
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
