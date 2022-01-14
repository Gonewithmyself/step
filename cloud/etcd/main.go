package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/concurrency"
)

var cli *clientv3.Client

func main() {
	// KV()
	// watch()
	// election()
	locker()
}

func KV() {
	// _, er := cli.KV.Put(context.TODO(), "test", "hello")
	// if er != nil {
	// 	panic(er)
	// }

	rsp, er := cli.KV.Get(context.TODO(), "/",
		// clientv3.WithPrefix(),
		clientv3.WithKeysOnly(),
		// clientv3.WithRange("z"),
		// clientv3.WithMinCreateRev(0),
	)
	if er != nil {
		fmt.Println(er)
		return
	}

	for _, kv := range rsp.Kvs {
		fmt.Println(kv)
	}

	fmt.Println(rsp.Count)
}

func watch() {
	runtime.GOMAXPROCS(1)
	ctx, cancel := context.WithCancel(context.Background())

	wch := cli.Watch(ctx, string([]byte{1}),
		clientv3.WithRange(string([]byte{255})),
		// clientv3.w
		// clientv3.WithPrefix(),
	)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		// for {
		// 	select {
		// 	case msg := <-wch:
		// 		fmt.Println("recv", msg)
		// 	case <-ctx.Done():
		// 		return
		// 	}
		// }
		for msg := range wch {
			for _, ev := range msg.Events {
				fmt.Println("recv", ev)
			}
		}

		wg.Done()
	}()

	put := func() {
		defer cancel()
		cli.KV.Put(context.TODO(), "test", "hello")
		tx := cli.Txn(context.TODO())
		tx.Then(
			clientv3.OpPut("/", "xx"),
			clientv3.OpPut("AZ", "xx")).
			Commit()
		time.Sleep(time.Second)
	}
	put()
	wg.Wait()
}

func locker() {
	key := "/lock/2"
	var counter int64
	const max int64 = 100
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		s1, er := concurrency.NewSession(cli)
		if er != nil {
			panic(er)
		}
		lk := concurrency.NewLocker(s1, key)
		// fmt.Println("get lock 1")
		for {
			lk.Lock()
			fmt.Println("get lock 1")
			if counter < max {
				counter++
				lk.Unlock()
			} else {
				lk.Unlock()
				break
			}
		}

	}()

	go func() {
		defer wg.Done()
		s1, er := concurrency.NewSession(cli)
		if er != nil {
			panic(er)
		}
		lk := concurrency.NewLocker(s1, key)
		for {
			lk.Lock()
			fmt.Println("get lock 2")
			if counter < max {
				counter++
				lk.Unlock()
			} else {
				lk.Unlock()
				break
			}
		}
	}()

	wg.Wait()
	fmt.Println(counter)

}

func election() {
	key := "/node"

	s1, er := concurrency.NewSession(cli)
	if er != nil {
		panic(er)
	}
	e1 := concurrency.NewElection(s1, key)

	s2, er := concurrency.NewSession(cli)
	if er != nil {
		panic(er)
	}
	e2 := concurrency.NewElection(s2, key)

	var wg sync.WaitGroup
	ech := make(chan *concurrency.Election, 1)
	doCompaign := func(e *concurrency.Election, name string) {
		time.Sleep(time.Second)
		er := e.Campaign(context.Background(), name)
		if er != nil {
			panic(er)
		}

		ech <- e
		wg.Done()
	}

	wg.Add(2)
	go doCompaign(e1, "e1")
	go doCompaign(e2, "e2")

	wg.Add(1)
	go func() {
		var count = 2
		for e := range ech {
			fmt.Println("become leader", e)
			e.Resign(context.Background())
			count--
			if count == 0 {
				break
			}
		}

		wg.Done()
	}()

	wg.Wait()

}

func init() {
	c, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:23791"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		// handle error!
		fmt.Printf("connect to etcd failed, err:%v\n", err)
		return
	}

	cli = c
}
