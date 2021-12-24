package main

import (
	"context"
	"flag"
	"log"
	"misc/mongodb"
	"os"
	"sync"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	c         *mongo.Client
	coll      *mongo.Collection
	count     int = 1e4
	batch         = 100
	itemCount     = 1005
	mongoURI      = "mongodb://127.0.0.1:27017/test"
)

var cmd string

func main() {
	switch cmd {
	case "multi":
		InsertMulti()
	default:
		Insert()
	}
}

func init() {
	ft := flag.NewFlagSet("flag", flag.ContinueOnError)
	ft.StringVar(&cmd, "cmd", "", "-cmd specify cmd")
	ft.IntVar(&count, "count", 1e4, "-count specify player count")
	ft.IntVar(&itemCount, "item", 1e4, "-item specify item count")
	ft.StringVar(&mongoURI, "mongo", mongoURI, "-mongo specify mongo URI")

	ft.Parse(os.Args[1:])
}

func InsertMulti() {
	c = newMongoClient()
	coll = c.Database("test").Collection("itemMutli")
	for i := 0; i < count; i++ {
		p := mongodb.NewPlayerItems(i)
		p.Add(itemCount)
		_, err := coll.InsertOne(context.Background(), p)
		if err != nil {
			log.Println("insert failed: ", err)
		}
	}
}

func Insert() {
	c = newMongoClient()
	coll = c.Database("test").Collection("item")

	ch := make(chan struct{}, 4)
	for i := 0; i < 4; i++ {
		ch <- struct{}{}
	}
	var wg sync.WaitGroup
	itemOnce := func() {
		for i := 0; i < count; i += batch {
			num := batch
			if i+batch > count {
				num = count - i
			}
			i := i
			wg.Add(1)
			go func() {
				<-ch

				InsertOnce(i, num)
				wg.Done()
				ch <- struct{}{}
			}()
		}
	}

	for i := 0; i < itemCount; i++ {
		itemOnce()
	}

	wg.Wait()
}

func InsertOnce(start, n int) {
	var list = make([]interface{}, 0, n)
	for i := 0; i < n; i++ {
		list = append(list, mongodb.NewItem(start+i))
	}

	_, err := coll.InsertMany(context.Background(), list)
	if err != nil {
		log.Println("insert failed:", err)
	}
}

func newMongoClient() *mongo.Client {
	clientOptions := options.Client()

	clientOptions.SetMaxPoolSize(uint64(100))

	client, err := mongo.Connect(context.TODO(), clientOptions.ApplyURI(mongoURI))
	if err != nil {
		log.Fatal(err.Error())
		return nil
	}

	er := client.Ping(context.Background(), &readpref.ReadPref{})
	if er != nil {
		panic(er)
	}

	coll := client.Database("test").Collection("item")
	iname, er := coll.Indexes().CreateOne(context.Background(),
		mongo.IndexModel{
			Keys: bson.D{{"pid", 1}},
		})
	log.Println(iname, er)
	return client
}
