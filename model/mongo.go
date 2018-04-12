package model

import (
	"crypto/md5"
	"crypto/sha512"
	"document/config"
	"errors"
	"fmt"
	"log"

	"gopkg.in/mgo.v2/bson"

	mgo "gopkg.in/mgo.v2"
)

const (
	dbName   = "commentDB"
	userColl = "user"
	comment  = "comment"
	arcticle = "arcticle"
)

var (
	errDup = errors.New("nick、email or tel duplicate")
)

// mongodb ...
type mongodb struct {
	dbSession          *mgo.Session
	db                 *mgo.Database
	UserCollection     *mgo.Collection
	CommentCollection  *mgo.Collection
	ArcticleCollection *mgo.Collection
}

// NewMongoDB 创建一个mongoDB数据库连接
func NewMongoDB() (Operater, error) {
	conn, err := config.GetString("mongoConn")
	if err != nil {
		log.Fatalf("Get mongoConn config error:%v", err)
		return mongodb{}, err
	}

	session, err := mgo.Dial(conn)
	if err != nil {
		log.Fatalf("Connect mongodb error:%v", err)
		return mongodb{}, err
	}

	session.SetMode(mgo.Monotonic, true)

	db := session.DB(dbName)

	UserCollection := db.C(userColl)
	index := mgo.Index{
		Key:        []string{"nick"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}
	err = UserCollection.EnsureIndex(index)
	if err != nil {
		log.Fatalf("user ensure nick index error: % v ", err)
	}

	index.Key = []string{"email"}
	err = UserCollection.EnsureIndex(index)
	if err != nil {
		log.Fatalf("user ensure email index error: % v ", err)
	}

	index.Key = []string{"tel"}
	err = UserCollection.EnsureIndex(index)

	if err != nil {
		index.Key = []string{"tel"}
		log.Fatalf("user ensure tel index error: % v ", err)
	}

	CommentCollection := db.C(comment)
	ArcticleCollection := db.C(arcticle)

	mongodb := mongodb{session, db, UserCollection, CommentCollection, ArcticleCollection}
	return mongodb, nil
}

// CloseMongoDB 关闭db
func CloseMongoDB(db mongodb) {
	db.dbSession.Close()
}

// 基础方法
func (db mongodb) GetUserID(id string) string {
	return ""
}

func (db mongodb) IncreaseLike(id string) {

}

func (db mongodb) DecreaseLike(id string) {

}

func (db mongodb) IncreaseSick(id string) {

}

func (db mongodb) DecreaseSick(id string) {

}

// 文章相关方法
func (db mongodb) GetActicle(id string) Acticle {
	return Acticle{}
}

func (db mongodb) AddActicle(act Acticle) {

}

func (db mongodb) GetActicleByTopic(topic string) []Acticle {
	return []Acticle{}
}

func (db mongodb) GetActicleByKeyword(keyword string) []Acticle {
	return []Acticle{}
}

func (db mongodb) GetActicleMostLike() []Acticle {
	return []Acticle{}
}

func (db mongodb) GetActicleMostComment() []Acticle {
	return []Acticle{}
}

func (db mongodb) GetActicleMostView() []Acticle {
	return []Acticle{}
}

func (db mongodb) GetActicleNewest() []Acticle {
	return []Acticle{}
}

func (db mongodb) Search(keyword string) []Acticle {
	return []Acticle{}
}

func (db mongodb) IncreaseActiclPageview(is string) {

}

// 评论的方法
func (db mongodb) AddComment(Comment) error {
	return nil
}

func (db mongodb) DelComment(id string) error {
	return nil
}

func (db mongodb) GetComment(id string) Comment {
	return Comment{}
}

// 用户操作方法
func (db mongodb) AddUser(user User) error {
	userID := bson.NewObjectId()
	user.ID = userID
	user.Passwd = transPasswd(user.Passwd)
	err := db.UserCollection.Insert(&user)
	if mgo.IsDup(err) {
		return errDup
	}
	return err
}

func (db mongodb) DelUserByID(id string) error {
	err := db.UserCollection.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	return err
}

func (db mongodb) GetUserByID(id string) (User, error) {
	user := User{}
	err := db.UserCollection.FindId(bson.ObjectIdHex(id)).One(&user)
	return user, err
}

func (db mongodb) GetUserIDByNick(nick string) (string, error) {
	var id string
	err := db.UserCollection.Find(bson.M{"nick": nick, "_id": 1}).One(&id)
	return id, err
}

func (db mongodb) ModifyUser(user User) error {
	return nil
}

func (db mongodb) ModifyPasswd(id, passwd string) error {
	return nil
}

func (db mongodb) ModifyRight(id string, level uint) error {
	return nil
}

func (db mongodb) IsAvailableNickName(name string) bool {
	return false
}

func (db mongodb) IsAvailableEmail(email string) bool {
	return false
}

func (db mongodb) IsAvailableTel(tel string) bool {
	return false
}

func (db mongodb) IsLogin(nick, passwd string) (string, bool) {
	return "id", false
}

// func init() {
// 	session, err := mgo.Dial("119.27.177.240:27017")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer session.Close()

// 	// Optional. Switch the session to a monotonic behavior.
// 	session.SetMode(mgo.Monotonic, true)
// 	appDB := session.DB(dbName)
// 	UserCollection := appDB.C("User")
// 	CommentCollection := appDB.C("Comment")
// 	err = UserCollection.Insert(&user)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	err = CommentCollection.Insert(&comment)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	user1 := User{}
// 	err = UserCollection.Find(bson.M{"nickname": "tako"}).One(&user1)
// 	if err != nil {
// 		log.Fatal("user:", err)
// 	}

// 	comment1 := Comment{}
// 	err = CommentCollection.Find(bson.M{}).One(&comment1)
// 	if err != nil {
// 		log.Fatal("comment:", err)
// 	}

// 	fmt.Printf("user:%v\ncomment:%v\n", user1, comment1)
// }

func transPasswd(originalPasswd string) string {
	shaPwd := sha512.Sum512([]byte(originalPasswd))
	md5Pwd := md5.Sum(shaPwd[:])
	newPwd := fmt.Sprintf("%x", md5Pwd)
	return newPwd
}
