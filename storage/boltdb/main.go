package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"sync"
	"unsafe"

	"github.com/boltdb/bolt"
)

func main() {

	syncMap()
	return

	a := [5]int{1, 2, 3, 4, 5}

	// create slice from array
	t := a[:3:3]
	fmt.Println(t, cap(t))
	testPage()
	return
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	er := db.Update(func(t *bolt.Tx) error {
		bk, er := t.CreateBucketIfNotExists([]byte("test"))
		if er != nil {
			return er
		}

		er = bk.Put([]byte("foo"), []byte(""))
		if er != nil {
			return er
		}

		return nil
	})
	if er != nil {
		log.Fatal(er)
	}

	er = db.View(func(t *bolt.Tx) error {
		bk := t.Bucket([]byte("test"))
		if bk == nil {
			return errors.New("bucket not found")
		}

		val := bk.Get([]byte("foo"))
		log.Println("val ", string(val))

		val = bk.Get([]byte("bar"))
		log.Println("val ", string(val))

		return nil
	})
	if er != nil {
		log.Fatal(er)
	}

}

//
type page struct {
	id  int32
	len int32
	ctx int
	// txt string
}

func testPage() {
	// readPage()
	// return
	p := &page{
		id:  23,
		len: 32,
		// txt: "wokao",
		ctx: 998,
	}

	sz := unsafe.Sizeof(*p)
	// var buf = make([]byte, sz)
	// *((*page)(unsafe.Pointer(&buf[0]))) = *p

	var buf []byte
	pbuf := (*slice)(unsafe.Pointer(&buf))
	pbuf.cap = int(sz)
	pbuf.len = int(sz)
	pbuf.data = unsafe.Pointer(p)

	// pp := (*(*[16]byte)(unsafe.Pointer(p)))[:]

	for i := range buf[:1:2] {
		fmt.Printf("%x sz:%d\n", buf[i], sz)
	}

	if er := ioutil.WriteFile("test.bin", buf, 0644); er != nil {
		panic(er)
	}
}

type slice struct {
	data unsafe.Pointer
	len  int
	cap  int
}

func readPage() {
	d, er := ioutil.ReadFile("test.bin")
	if er != nil {
		panic(er)
	}

	p := (*slice)(unsafe.Pointer(&d))

	pp := (*page)(unsafe.Pointer(&d[0]))

	fmt.Println(p, len(d), pp)
}

var sm sync.Map

func syncMap() {
	sm.Store("test", "go")

	sm.Store("test", "go1")

	sm.Store("test2", "go1")

	sm.Load("test")
}
