package mongodb

import (
	"math/rand"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Item struct {
	ID         int32  `bson:"id"` //物品ID
	PID        string `bson:"pid"`
	Count      int32  `bson:"count"`
	ExpireTime int64  `bson:"epTime,omitempty"` //到期时间, 0表示永不过期
	UID        string `bson:"_id"`              //物品实例ID
	GetTime    int64  `bson:"get_time"`
	CreateTime int64  `bson:"create_time"`
	Number     int64  `bson:"number"`
	Instance   int64  `bson:"instance"`
}

func NewItem(i int) *Item {
	now := time.Now().Unix()

	return &Item{
		ID:         int32(i),
		PID:        strconv.Itoa(i),
		ExpireTime: now,
		UID:        primitive.NewObjectID().Hex(),
		GetTime:    now,
		CreateTime: now,
		Number:     rand.Int63(),
		Instance:   rand.Int63(),
		Count:      rand.Int31(),
	}
}

func init() {
	rand.Seed(time.Now().Unix())
}

type PlayerItems struct {
	PID   string `bson:"_id"`
	Items map[string]*DBItem
}

func NewPlayerItems(i int) *PlayerItems {
	return &PlayerItems{
		PID:   strconv.Itoa(i),
		Items: map[string]*DBItem{},
	}
}

type DBItem struct {
	ID         int32  `bson:"id"` //物品ID
	Count      int32  `bson:"count"`
	ExpireTime int64  `bson:"epTime,omitempty"` //到期时间, 0表示永不过期
	UID        string `bson:"-"`                //物品实例ID
	GetTime    int64  `bson:"get_time"`
	CreateTime int64  `bson:"create_time"`
	Number     int64  `bson:"number"`
	Instance   int64  `bson:"instance"`
}

func NewDBItem() *DBItem {
	now := time.Now().Unix()
	return &DBItem{
		ID:         rand.Int31(),
		ExpireTime: now,
		UID:        primitive.NewObjectID().Hex(),
		GetTime:    now,
		CreateTime: now,
		Number:     rand.Int63(),
		Instance:   rand.Int63(),
		Count:      rand.Int31(),
	}
}

func (p *PlayerItems) Add(n int) {
	for i := 0; i < n; i++ {
		it := NewDBItem()
		p.Items[it.UID] = it
	}
}
