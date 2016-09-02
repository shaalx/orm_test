package main

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	. "github.com/toukii/orm_test/user/model"
)

var ORM orm.Ormer

func init() {
	// orm.RegisterDataBase("default", "mysql", "root:1234@tcp(localhost:3306)/qq?charset=utf8")
	orm.RegisterDataBase("default", "mysql", "cdb_outerroot:root1234@tcp(55c354e17de4e.sh.cdb.myqcloud.com:7276)/qq?charset=utf8")
	orm.RegisterDriver("mysql", orm.DR_MySQL)
	orm.RegisterModel(new(User), new(Profile), new(Userz), new(Friends))
	ORM = orm.NewOrm()
}

func main() {
	add_user()
	add_profile()
	add_userz()
	add_friends()
}

func add_user() {
	profile := &Profile{Id: 1}
	user := &User{Name: "jackson", Profile: profile}
	n, err := ORM.Insert(user)
	fmt.Println(n, err)
}

func add_profile() {
	user := &User{Id: 1, Name: "jack"}
	profile := &Profile{User: user}

	n, err := ORM.Insert(profile)
	fmt.Println(n, err)
}

func add_userz() {
	// friend1 := &Friends{Id: 1}
	// friend2 := &Friends{Id: 2}
	// friends := []*Friends{friend1, friend2}
	userz := &Userz{Name: "tom" /*, Friends: friends*/}

	n, err := ORM.Insert(userz)
	fmt.Println(n, err)
}

func add_friends() {
	userz := &Userz{Id: 1, Name: "jack"}
	userzt := &Userz{Id: 2, Name: "jackson"}
	friends := &Friends{Userz: userz, Userzt: userzt}

	n, err := ORM.Insert(friends)
	fmt.Println(n, err)
}
