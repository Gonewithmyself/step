package main

import (
	"context"
	"fmt"
	"log"
	"runtime"
	"step/microsvc/grpc/proto"
	"time"
)

type monster struct {
	chatroom
}

type chatroom struct {
	conns map[int32]proto.Monster_ChatServer
	hist  []*proto.Msg
}

func (c *chatroom) add(id int32, srv proto.Monster_ChatServer) {
	c.conns[id] = srv
}

func (c *chatroom) del(id int32) {
	delete(c.conns, id)
}

func (c *chatroom) telOnlines(sv proto.Monster_ChatServer) {
	msg := &proto.Msg{}
	for id := range c.conns {
		msg.Onlines = append(msg.Onlines, id)
	}
	sv.Send(msg)
}

func (c *chatroom) route(msg *proto.Msg) {
	if msg.To != 0 {
		to := c.conns[msg.To]
		if to == nil {
			return
		}
		to.Send(msg)
		return
	}

	c.hist = append(c.hist, msg)
	if len(c.hist) > 40 {
		tail := c.hist[20:40]
		c.hist = c.hist[:0]
		c.hist = append(c.hist, tail...)
	}
	for id, conn := range c.conns {
		if id != msg.From {
			conn.Send(msg)
		}

	}
}

func (mst *monster) Translate(ctx context.Context, req *proto.Query) (res *proto.Result, er error) {
	if req.Id < 0 {
		er = fmt.Errorf("wrong id(%v)", req.Id)
		return
	}

	res = &proto.Result{}
	res.Mean = fmt.Sprintf("%d - %v", req.Id, req.Word)
	return
}

func (mst *monster) PullMetrics(req *proto.Query, srv proto.Monster_PullMetricsServer) error {
	n := 10
	for i := 0; i < n; i++ {
		var stats runtime.MemStats
		runtime.ReadMemStats(&stats)
		er := srv.Send(&proto.Metric{
			Mem: int32(stats.HeapInuse),
			Ts:  time.Now().Unix(),
		})
		if er != nil {
			log.Println(er)
		}

		time.Sleep(time.Second)
	}
	return nil
}

func (mst *monster) PushMetrics(srv proto.Monster_PushMetricsServer) error {
	var ok = true
	for i := 0; i < 10; i++ {
		msg, er := srv.Recv()
		if er != nil {
			log.Println(er)
			ok = false
			break
		}

		log.Println(msg)
	}
	if ok {
		srv.SendAndClose(&proto.None{})
	}
	return nil
}

func (mst *monster) Chat(srv proto.Monster_ChatServer) error {
	msg, er := srv.Recv()
	if er != nil {
		log.Println("first msg", er)
		return er
	}
	mst.telOnlines(srv)

	stk := [1024]byte{}
	// runtime.Stack(stk[:], false)
	log.Printf("recv  %s \n", stk)

	mst.add(msg.From, srv)
	defer mst.del(msg.From)

	for i := range mst.hist {
		if er := srv.Send(mst.hist[i]); er != nil {
			log.Println("send history ", er)
			break
		}
	}

	for {
		msg, er := srv.Recv()
		if er != nil {
			break
		}

		mst.route(msg)
	}
	return nil
}

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
}
