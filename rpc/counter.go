package rpc

import (
	"github.com/projecteru2/core/log"
	"github.com/projecteru2/core/types"
	"github.com/projecteru2/core/utils"

	"golang.org/x/net/context"
)

// gRPC上全局的计数器
// 只有在任务数为0的时候才给停止
// 为啥会加在gRPC server上呢?
// 因为一个入口给一个最简单了...

// 增加一个任务, 在任务调用之前要调用一次.
// 否则任务不被追踪, 不保证任务能够正常完成.
func (v *Vibranium) taskAdd(ctx context.Context, name string, verbose bool) context.Context {
	if ctx != nil {
		ctx = context.WithValue(ctx, types.TracingID, utils.RandomString(8))
	}
	if verbose {
		log.Debugf(ctx, "[task] %s added", name)
	}
	v.counter.Add(1)
	v.TaskNum++
	return ctx
}

// 完成一个任务, 在任务执行完之后调用一次.
// 否则计数器用完不会为0, 你也别想退出这个进程了.
func (v *Vibranium) taskDone(ctx context.Context, name string, verbose bool) {
	if verbose {
		log.Debugf(ctx, "[task] %s done", name)
	}
	v.counter.Done()
	v.TaskNum--
}

// Wait for all tasks done
// 会在外面graceful之后调用.
// 不完成不给退出进程.
func (v *Vibranium) Wait() {
	v.counter.Wait()
}
