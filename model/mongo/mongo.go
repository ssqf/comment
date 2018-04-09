package mongo

import (
	"document/model"

	mgo "gopkg.in/mgo.v2"
)

const (
	dbName   = "commentDB"
	userColl = "user"
	comment  = "comment"
	arcticle = "arcticle"
)

// Mongodb ...
type Mongodb struct {
	dbSession          *mgo.Session
	db                 *mgo.Database
	UserCollection     *mgo.Collection
	CommentCollection  *mgo.Collection
	ArcticleCollection *mgo.Collection
}

// NewMongoDB 创建一个mongoDB数据库连接
func NewMongoDB() (model.Operater, error) {
	session, err := mgo.Dial("119.27.177.240:27017")
	if err != nil {
		return Mongodb{}, err
	}

	session.SetMode(mgo.Monotonic, true)

	db := session.DB(dbName)
	UserCollection := db.C(userColl)
	CommentCollection := db.C(comment)
	ArcticleCollection := db.C(arcticle)

	mongodb := Mongodb{session, db, UserCollection, CommentCollection, ArcticleCollection}
	return mongodb, nil
}

// CloseMongoDB 关闭db
func CloseMongoDB(db Mongodb) {
	db.dbSession.Close()
}

// GetContent 获取内容
func (db Mongodb) GetContent(id string) string {
	return ""
}

func (db Mongodb) GetUserID(id string) string {
	return ""
}

func (db Mongodb) IncreaseLike(id string) {

}

func (db Mongodb) DecreaseLike(id string) {

}

func (db Mongodb) IncreaseSick(id string) {

}

func (db Mongodb) DecreaseSick(id string) {

}

func (db Mongodb) GetActicle(id string) model.Acticle {
	return model.Acticle{}
}

func (db Mongodb) AddActicle(act model.Acticle) {

}

func (db Mongodb) GetActicleByTopic(topic string) []model.Acticle {
	return []model.Acticle{}
}

func (db Mongodb) GetActicleByKeyword(keyword string) []model.Acticle {
	return []model.Acticle{}
}

func (db Mongodb) GetActicleMostLike() []model.Acticle {
	return []model.Acticle{}
}

func (db Mongodb) GetActicleMostComment() []model.Acticle {
	return []model.Acticle{}
}

func (db Mongodb) GetActicleMostView() []model.Acticle {
	return []model.Acticle{}
}

func (db Mongodb) GetActicleNewest() []model.Acticle {
	return []model.Acticle{}
}

func (db Mongodb) AddComment(model.Comment) error {
	return nil
}

func (db Mongodb) DelComment(id string) error {
	return nil
}

func (db Mongodb) GetComment(id string) model.Comment {
	return model.Comment{}
}

func (db Mongodb) AddUser(user model.User) error {
	return nil
}

func (db Mongodb) DelUser(id string) error {
	return nil
}

func (db Mongodb) ModifyUser(user model.User) error {
	return nil
}

func (db Mongodb) ModifyPasswd(id, passwd string) error {
	return nil
}

func (db Mongodb) ModifyRight(id string, level uint) error {
	return nil
}

func (db Mongodb) IsAvailableNickName(name string) bool {
	return false
}

func (db Mongodb) IsAvailableEmail(email string) bool {
	return false
}

func (db Mongodb) IsAvailableTel(tel string) bool {
	return false
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
