package main

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

var ORM orm.Ormer

func init() {
	orm.RegisterDriver("mysql", orm.DR_MySQL)
	orm.RegisterDataBase("default", "mysql", "root:1234@tcp(localhost:3306)/os?charset=utf8")
	orm.RegisterModel(new(Groups), new(User), new(Grel))
	ORM = orm.NewOrm()
	orm.Debug = true
}

func main() {
	insertUser()
	// insertGrel()
	// deleteGrel()
	// deleteUser()
}
func insertUser() {
	user := User{Name: "jack2", Password: "jack123"}
	ORM.Insert(&user)
}
func deleteUser() {
	ORM.QueryTable((*User)(nil)).Filter("Name", "jack2").Delete()
}
func deleteGrel() {
	ORM.QueryTable((*Grel)(nil)).Filter("Id", 1).Delete()
}
func insertGrel() {
	user := User{Name: "jack2", Password: "jack123"}
	groups := Groups{Id: 5, Name: "gopher5"}

	grel := Grel{Group: &groups, User: &user, Nickname: "jackson"}
	ORM.Insert(&groups)
	ORM.Insert(&user)
	n2, err2 := ORM.Insert(&grel)
	checkerr(err2)
	fmt.Println(n2)
}

type User struct {
	Name     string `orm:"pk"`
	Password string `orm:"password"`
}

type Groups struct {
	Id   int32 `orm:"pk"`
	Name string
}
type Grel struct {
	Id       int32   `orm:"pk"`
	User     *User   `orm:"rel(fk);column(uname);on_delete(do_nothing)"`
	Group    *Groups `orm:"rel(fk);column(gid);on_delete(do_nothing)"`
	Nickname string  `orm:"null"`
}

// 多字段唯一键
func (u *Grel) TableUnique() [][]string {
	return [][]string{
		[]string{"User", "Group"},
	}
}

type Message struct {
	Id        int32
	Uin       int32
	Uout      int32
	Content   []byte
	Timestamp int64
}

type Item struct {
	Id   int32
	Name string
}
type Friends struct {
	Uid  int32
	Iid  int32
	Fuid int32
}
type Login struct {
	Uid       int32
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
