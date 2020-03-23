// hello_server.go
package hello_controller

import (
	"fmt"
	"golang.org/x/net/context"
	"gostudy/grpc/go_server/proto/hello"
)

type HelloController struct{}

func (h *HelloController) SayHello(ctx context.Context, in *hello.HelloRequest) (*hello.HelloResponse, error) {
	return &hello.HelloResponse{Message: fmt.Sprintf("%s", in.Name)}, nil
}

func (h *HelloController) LotsOfReplies(in *hello.HelloRequest, stream hello.Hello_LotsOfRepliesServer) error {
	for i := 0; i < 10; i++ {
		stream.Send(&hello.HelloResponse{Message: fmt.Sprintf("%s %s %d", in.Name, "Reply", i)})
	}
	return nil
}

func (h *HelloController) GetList(ctx context.Context, in *hello.StudentRequest) (*hello.StudentResponse, error) {
	return &hello.StudentResponse{Name: fmt.Sprintf("%s", in.Name)}, nil
}
