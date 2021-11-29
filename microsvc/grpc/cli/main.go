package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"step/microsvc/grpc/proto"
	"step/misc/randtools"
	"strconv"
	"sync"
	"time"

	"google.golang.org/grpc"
)

type client struct {
	proto.MonsterClient
	proto.HelloClient
}

func (c *client) run(cmd string) {
	switch cmd {
	case "one":
		c.one()
	case "hello":
		c.hello()
	case "pull":
		c.pull()

	case "push":
		c.push()

	case "chat":
		c.chat()

	case "mchat":
		c.chat()

	default:
		c.one()
	}
}

func (c *client) one() {
	id := randtools.Range(1, 100)
	res, er := c.Translate(context.Background(), &proto.Query{
		Id: int32(id),
	})

	log.Println(res, er)
}

func (c *client) hello() {
	res, er := c.Say(context.Background(), &proto.Query{})
	log.Println(res, er)
}

func (c *client) pull() {
	// id := randtools.Range(1, 100)
	stream, er := c.PullMetrics(context.Background(), &proto.Query{})
	if er != nil {
		panic(er)
	}

	for {
		msg, er := stream.Recv()
		if er != nil {
			log.Println(er)
			break
		}
		log.Println(msg)

	}
	c.one()
}

func (c *client) push() {
	st, er := c.PushMetrics(context.Background())
	if er != nil {
		panic(er)
	}

	for i := 0; i < 10; i++ {
		er := st.Send(&proto.Metric{
			Ts:  time.Now().Unix(),
			Mem: int32(i + 1),
		})
		if er != nil {
			panic(er)
		}
	}
	_, er = st.CloseAndRecv()
	if er != nil {
		panic(er)
	}

}

func (c *client) mchat() {
	var wg sync.WaitGroup
	for i := 1; i < 10; i++ {
		wg.Add(1)
		go func() {
			c.chat()
			wg.Done()
		}()
	}
	wg.Wait()
}

func (c *client) chat() {
	st, er := c.Chat(context.Background())
	if er != nil {
		panic(er)
	}
	myid := randtools.Range(1, 100)
	er = st.Send(&proto.Msg{
		From: int32(myid),
	})
	if er != nil {
		panic(er)
	}

	hello, er := st.Recv()
	if er != nil {
		panic(er)
	}
	log.Println("hello", hello)

	dc := make(chan struct{})
	defer close(dc)
	go func() {
		for {
			select {
			case <-dc:
				return
			default:
			}

			msg, er := st.Recv()
			if er != nil {
				fmt.Println(er)
				continue
			}

			log.Printf("recv %-4d -> %-4d: %s\n", msg.From, msg.To, msg.Content)
		}
	}()

	for {
		seed := randtools.Range(1, 9999)
		var to = seed
		if len(hello.Onlines) != 0 {
			to = int(hello.Onlines[seed%len(hello.Onlines)])
		}

		if seed%3 == 0 {
			to = 0
		}

		er := st.Send(&proto.Msg{
			From:    int32(myid),
			To:      int32(to),
			Content: strconv.Itoa(seed),
		})
		if er != nil {
			break
		}

		time.Sleep(time.Second)
	}
}

var c *client

func main() {
	addr := "0.0.0.0:9999"
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("connect server error: %v", err)
	}
	defer conn.Close()

	c = &client{MonsterClient: proto.NewMonsterClient(conn),
		HelloClient: proto.NewHelloClient(conn)}

	cmd := ""
	flag.StringVar(&cmd, "cmd", "one", "specify cmd")
	flag.Parse()
	c.run(cmd)
}
