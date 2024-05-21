# context2

对context的取消进行分离。比如典型的场景是即想继承rpc服务里面的ctx的trace信息，又想单开个go程发起新的请求。
这时候你可以使用context2分离ctx的影响，同时获取对应的信息

```go
import (
    "github.com/guonaihong/context2"
)
func main() {
    ctx := context.WithValue(context.TODO(), "key", "value")
    ctx, cancel := context.WithCancel(ctx)
    // 模拟父ctx被取消
    cancel()

    ctx2 := Detach(ctx)
    select {
    case <-ctx2.Done():
    // 这里不会执行，父context取消不会影响这里
    case <-time.After(time.Second / 100):
    // 这里会执行， 已经与父context分离了
    fmt.Printf("ctx2.value:%s\n", ctx2.Value("key"))
    }
}
```
