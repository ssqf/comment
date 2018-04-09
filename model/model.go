package model

import (
	"document/config"
	"log"
)

// Document 文档
type document struct {
	ID        string
	LikeNum   uint
	SickNum   uint
	UserID    string
	Content   string
	CommentID []string
	Date      string
}

// Acticle 文章
type Acticle struct {
	document
	Title       string
	Topic       string
	keyWork     []string
	topPicURL   []string
	PageviewNum uint
}

// Comment 评论的结构
type Comment struct {
	document
	ActicleID string
}

// User 用户信息
type User struct {
	ID          string
	NickName    string
	Email       string
	HeadPicture string
	Bio         string
	Like        []string
	Passwd      string
	Tel         string
	RightLevel  uint
}

// Docer 对数据的操作
type docInterface interface {
	GetUserID(id string) string
	IncreaseLike(id string)
	DecreaseLike(id string)
	IncreaseSick(id string)
	DecreaseSick(id string)
}

// acticleInterface 操作文章的接口
type acticleInterface interface {
	GetActicle(id string) Acticle
	AddActicle(act Acticle)
	GetActicleByTopic(topic string) []Acticle
	GetActicleByKeyword(keyword string) []Acticle
	GetActicleMostLike() []Acticle
	GetActicleMostComment() []Acticle
	GetActicleMostView() []Acticle
	GetActicleNewest() []Acticle
}

// Operater 所有操作接口
type Operater interface {
	commentInterface
	acticleInterface
	userInterface
}

// CommentInterface 操作评论的接口
type commentInterface interface {
	AddComment(Comment) error
	DelComment(id string) error
	GetComment(id string) Comment
}

// UserInterface 用户操作接口
type userInterface interface {
	AddUser(user User) error
	DelUser(id string) error
	ModifyUser(user User) error
	ModifyPasswd(id, passwd string) error
	ModifyRight(id string, level uint) error
	IsAvailableNickName(name string) bool
	IsAvailableEmail(email string) bool
	IsAvailableTel(tel string) bool
}

var operater Operater

// GetOperater 获取操作
func GetOperater() Operater {
	return operater
}
func init() {
	dbType, err := config.GetString("dbType")
	if err != nil {
		log.Fatalf("init db operater,can't obtain dbType from config file error:%v", err)
	}
	switch dbType {
	case "mongodb":
	case "mysql":
	case "psql":
	case "redis":
	default:
		log.Fatalf("init db operater,%s db don't support", dbType)
	}
}
