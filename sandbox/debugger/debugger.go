package debugger

import (
	"fmt"
	"gominisquare/duktape"
	"log"
	"unsafe"
)

type GominiDebugger struct {
	transport *duktape.SocketTransport
	ctx       *duktape.Context
}

func New(address string, ctx *duktape.Context) (*GominiDebugger, error) {
	requestFn := func(ctx *duktape.Context, uData unsafe.Pointer, nValues int) int {
		fmt.Printf("number of arguments received: %d", nValues)
		return 0
	}

	detachedFn := func(ctx *duktape.Context, uData unsafe.Pointer) {
		println("detached transport")
	}

	transport, err := duktape.NewSocketTransport(ctx, "tcp", address, requestFn, detachedFn, nil)
	if err != nil {
		return nil, err
	}

	return &GominiDebugger{
		transport: transport,
		ctx:       ctx,
	}, nil
}

func (d *GominiDebugger) Start() {
	d.transport.Listen(func(err error) {
		log.Printf("error: %s", err.Error())
	})
}

func (d *GominiDebugger) Cooperate() {
	duktape.DukDebugger().Cooperate(d.ctx)
}

func (d *GominiDebugger) Detach() {
	duktape.DukDebugger().Detach(d.ctx)
}

func (d *GominiDebugger) Close() error {
	return d.transport.Close()
}
