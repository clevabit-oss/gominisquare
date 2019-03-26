package sandbox

import (
	"gopkg.in/olebedev/go-duktape.v3"
	"runtime"
	"sync"
	"time"
)

type eventLoop struct {
	ctx      *duktape.Context
	m        sync.Mutex
	canceled bool
	tasks    []EventLoopTask
}

func newEventLoop(ctx *duktape.Context) *eventLoop {
	return &eventLoop{
		ctx: ctx,
		m:   sync.Mutex{},
	}
}

func (el *eventLoop) push(task EventLoopTask) bool {
	el.m.Lock()
	defer el.m.Unlock()

	if el.canceled {
		return false
	}

	el.tasks = append(el.tasks, task)
	return true
}

func (el *eventLoop) pop() EventLoopTask {
	el.m.Lock()
	defer el.m.Unlock()

	if len(el.tasks) == 0 {
		return nil
	}

	task, tasks := el.tasks[0], el.tasks[1:]
	el.tasks = tasks

	return task
}

func (el *eventLoop) run() {
	for {
		task := el.pop()
		if task != nil {
			task(el.ctx)
		} else {
			duktape.DukDebugger().Cooperate(el.ctx)
		}

		if el.canceled {
			return
		}

		runtime.Gosched()
		time.Sleep(time.Millisecond)
		runtime.Gosched()
	}
}

func (el *eventLoop) cancel() {
	el.m.Lock()
	defer el.m.Unlock()
	el.canceled = true
}
