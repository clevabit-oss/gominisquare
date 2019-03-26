package sandbox

import "gopkg.in/olebedev/go-duktape.v3"

type EventLoopTask = func(ctx *duktape.Context)
