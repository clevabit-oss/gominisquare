package sandbox

import "gominisquare/duktape"

type EventLoopTask = func(ctx *duktape.Context)