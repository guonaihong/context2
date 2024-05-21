package context2

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestDetach(t *testing.T) {
	ctx := context.WithValue(context.TODO(), "key", "value")
	ctx, cancel := context.WithCancel(ctx)
	// 模拟父ctx被取消
	cancel()

	ctx2 := Detach(ctx)
	select {
	case <-ctx2.Done():
		t.Error("ctx2 should not be done")
	case <-time.After(time.Second / 100):
		fmt.Printf("ctx2.value:%s\n", ctx2.Value("key"))
	}

}
