//go:generate protoc -I ./transport --go_out=plugins=grpc:./transport ./transport/communication.proto

package main

import (
	"fmt"
	"golang.org/x/net/context"
	"gominisquare/duktape"
	"gominisquare/sandbox/debugger"
	"gominisquare/transport"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"io/ioutil"
	"log"
	"net"
	"os"
	"path/filepath"
	"time"
)

type server struct {
}

func (s *server) Call(ctx context.Context, request *transport.Request) (*transport.Response, error) {
	panic("implement me")
}

func (s *server) ServerStream(request *transport.Request, stream transport.KernelSyscall_ServerStreamServer) error {
	panic("implement me")
}

func (s *server) ClientStream(stream transport.KernelSyscall_ClientStreamServer) error {
	panic("implement me")
}

func (s *server) Stream(stream transport.KernelSyscall_StreamServer) error {
	panic("implement me")
}

func main() {
	f, _ := os.Stat("/tmp/test")
	if f != nil {
		os.Remove("/tmp/test")
	}

	ctx := duktape.New()
	dbg, err := debugger.New("localhost:9091", ctx)
	if err != nil {
		panic(err)
	}
	dbg.Start()

	go func() {
		fmt.Println("DebuggerTransport ready")
		file, err := filepath.Abs("./test.js")
		if err != nil {
			panic(err)
		}
		println(file)
		f, err := os.Open(file)
		if err != nil {
			panic(err)
		}
		code, err := ioutil.ReadAll(f)
		if err != nil {
			panic(err)
		}
		ctx.PushString(string(code))
		ctx.PushString("test.js")

		if err := ctx.Pcompile(0); err != nil {
			panic(fmt.Sprintf("aaaaaaargh: %v", err.Error()))
		}

		ret := ctx.Pcall(0)
		if ret != duktape.ExecSuccess {
			fmt.Printf("output: %s\n", ctx.SafeToString(-1))
		}

		executions := 1
		for {
			ctx.GetGlobalString("test")
			ret := ctx.Pcall(0)
			if ret != duktape.ExecSuccess {
				fmt.Printf("error message: %s", ctx.SafeToString(-1))
			}
			ctx.Pop()

			dbg.Cooperate()
			time.Sleep(time.Second)

			log.Printf("executions: %d", executions)

			if executions == 10 {
				dbg.Detach()
				dbg.Close()

			}

			executions++
		}
	}()

	port, err := net.Listen("unix", "/tmp/test")
	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
	}

	s := grpc.NewServer()

	transport.RegisterKernelSyscallServer(s, &server{})

	reflection.Register(s)

	if err := s.Serve(port); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	dbg.Close()
}
