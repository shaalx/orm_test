package main

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

var ORM orm.Ormer

func init() {
	orm.RegisterDriver("mysql", orm.DR_MySQL)
	// orm.RegisterDataBase("local", "mysql", "root:1234@tcp(localhost:3306)/os?charset=utf8")
	orm.RegisterDataBase("default", "mysql", "cdb_outerroot:root1234@tcp(55c354e17de4e.sh.cdb.myqcloud.com:7276)/os?charset=utf8")
	orm.RegisterModel(new(Groups), new(User), new(Grel), new(Message), new(Item), new(Friends), new(Login))
	ORM = orm.NewOrm()
}

func insertRel() {
	group := &Groups{Id: 1, Name: "group1"}
	ORM.Insert(group)
	user1 := &User{Id: 1, Name: "jack1", Password: "a1212"}
	user2 := &User{Id: 2, Name: "jack2", Password: "a1212"}
	grel1 := &Grel{User: user1, Group: group}
	grel2 := &Grel{User: user2, Group: group}
	ORM.Insert(grel1)
	ORM.Insert(grel2)
}

func insertMsg() {
	uin := &User{Id: 1}
	uout := &User{Id: 2}
	msg := &Message{Uin: uin, Uout: uout, Content: "hello,2.", Timestamp: time.Now().Unix()}
	n, err := ORM.Insert(msg)
	fmt.Println(n, err)
}
func queryMsg() {
	var msg []Message
	n, err := ORM.QueryTable((*Message)(nil)).RelatedSel().All(&msg)
	fmt.Println(n, err)
	for _, it := range msg {
		fmt.Println(it.String())
	}
}
func insertUser() {
	user1 := &User{Id: 1, Name: "jack1", Password: "a1212"}
	user2 := &User{Id: 2, Name: "jack2", Password: "a1212"}
	item1 := &Item{Id: 1, Name: "jacak1's item", Uid: user1}
	item2 := &Item{Id: 2, Name: "jacak2's item", Uid: user2}
	ORM.Insert(item1)
	ORM.Insert(item2)
	ORM.Insert(user1)
	ORM.Insert(user2)
}
func insertFriends() {
	user1 := &User{Id: 1, Name: "jack1", Password: "a1212"}
	user2 := &User{Id: 2, Name: "jack2", Password: "a1212"}
	item1 := &Item{Id: 1, Name: "jack1's item", Uid: user1}
	item2 := &Item{Id: 2, Name: "jack2's item", Uid: user2}
	friends1 := &Friends{Uid: user1, Iid: item1, Fuid: user2}
	friends2 := &Friends{Uid: user2, Iid: item2, Fuid: user1}
	n, err := ORM.Insert(friends1)
	fmt.Println(n, err)
	n, err = ORM.Insert(friends2)
	fmt.Println(n, err)
}
func queryFriends() {
	var items []Item
	ORM.QueryTable((*Item)(nil)).Filter("Uid__Id", 1).All(&items)
	ids := make([]int32, 0, len(items))
	for _, it := range items {
		ids = append(ids, it.Id)
	}
	var friends []Friends
	ORM.QueryTable((*Friends)(nil)).Filter("Iid__Id__in", ids).RelatedSel("Fuid").All(&friends)
	fmt.Println(friends)
	for _, it := range friends {
		fmt.Println(it.Fuid.Name)
	}
}
func quickFriends() {
	var friends []Friends
	ORM.QueryTable((*Friends)(nil)).Filter("Uid__Id", 1).RelatedSel("Fuid").All(&friends)
	fmt.Println(friends)
	for _, it := range friends {
		fmt.Printf("%+v\n", it.Fuid)
	}
}
func main() {
	insertUser()
	insertRel()
	insertMsg()
	queryMsg()

	insertFriends()

	queryFriends()
	quickFriends()

}

type User struct {
	Id       int32  `orm:"pk"`
	Name     string `orm:"name"`
	Password string `orm:"password"`
}

type Groups struct {
	Id   int32 `orm:"pk"`
	Name string
}
type Grel struct {
	Id       int32   `orm:"pk"`
	User     *User   `orm:"rel(fk);column(uid);on_delete(cascade)"`
	Group    *Groups `orm:"rel(fk);column(gid);on_delete(cascade)"`
	Nickname string  `orm:"null"`
}
type Message struct {
	Id        int32  `orm:"pk"`
	Uin       *User  `orm:"rel(fk);column(uin);on_delete(do_nothing)"`
	Uout      *User  `orm:"rel(fk);column(uout);on_delete(do_nothing)"`
	Content   string `orm:"column(content)"`
	Timestamp int64
}

func (m *Message) String() string {
	return fmt.Sprintf("%+v ===> %+v :%s [%v]", m.Uin.Name, m.Uout.Name, m.Content, m.Timestamp)
}

type Item struct {
	Id   int32 `orm:"pk"`
	Name string
	Uid  *User `orm:"rel(fk);column(uid);on_delete(do_nothing)"`
}
type Friends struct {
	Id   int32 `orm:"pk"`
	Uid  *User `orm:"rel(fk);column(uid);on_delete(do_nothing)"`
	Iid  *Item `orm:"rel(fk);column(iid);on_delete(do_nothing)"`
	Fuid *User `orm:"rel(fk);column(fuid);on_delete(do_nothing)"`
}
type Login struct {
	Id        int32 `orm:"pk"`
	Uid       *User `orm:"rel(fk);column(uid);on_delete(do_nothing)"`
	Timestamp int64
	Ip        string
}

func checkerr(err error) bool {
	if nil != err {
		fmt.Println(err)
		return true
	}
	return false
}
