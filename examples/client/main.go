package client

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/DanteSu/miniblog/internal/pkg/log"
	pb "github.com/DanteSu/miniblog/pkg/proto/miniblog/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

var (
	addr  = flag.String("addr", "localhost:9090", "The address to connect to.")
	limit = flag.Int64("limit", 10, "Limit to list users.")
)

func main() {
	flag.Parse()
	//与服务器建立连接
	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalw("Did not connect", "err", err)
	}
	defer conn.Close()
	c := pb.NewMiniBlogClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	//请求listuser接口
	r, err := c.ListUser(ctx, &pb.ListUserRequest{Offset: 0, Limit: *limit})
	if err != nil {
		log.Fatalw("Could not greet: %v", err)
	}

	//打印结果
	fmt.Println("TOtalCount:", r.TotalCount)
	for _, u := range r.Users {
		d, _ := json.Marshal(u)
		fmt.Println(string(d))
	}

}
