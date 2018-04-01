package db

import (
	"fmt"
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Document 文档
type Document struct {
	ID        string
	LikeNum   uint
	SickNum   uint
	UserID    string
	Content   string
	CommentID []string
	Date      string
}

// Comment 评论的结构
type Comment struct {
	Floor     uint
	ActicleID string
}

// Acticle 文章
type Acticle struct {
	Title     string
	Topic     string
	keyWork   []string
	topPicURL []string
	Pageview  uint
}

// User 用户信息
type User struct {
	ID            string
	Nickname      string
	Email         string
	HeadPortraits string
	Bio           string
	Like          string
	Passwd        string
}

var (
	dbName = "MyApp"
)

func init() {
	session, err := mgo.Dial("119.27.177.240:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	user := User{}
	comment := Comment{}
	appDB := session.DB(dbName)
	UserCollection := appDB.C("User")
	CommentCollection := appDB.C("Comment")
	err = UserCollection.Insert(&user)
	if err != nil {
		log.Fatal(err)
	}
	err = CommentCollection.Insert(&comment)
	if err != nil {
		log.Fatal(err)
	}

	user1 := User{}
	err = UserCollection.Find(bson.M{"nickname": "tako"}).One(&user1)
	if err != nil {
		log.Fatal("user:", err)
	}

	comment1 := Comment{}
	err = CommentCollection.Find(bson.M{}).One(&comment1)
	if err != nil {
		log.Fatal("comment:", err)
	}

	fmt.Printf("user:%v\ncomment:%v\n", user1, comment1)
}
