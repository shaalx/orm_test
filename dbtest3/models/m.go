package models

// import (
// 	"fmt"
// )

// var ORM orm.Ormer

// func init() {
// 	orm.RegisterDataBase("default", "mysql", "cdb_outerroot:root1234@tcp(55c354e17de4e.sh.cdb.myqcloud.com:7276)/dbt?charset=utf8")
// 	orm.RegisterModel( /*new(User) new(Profile), ,*/ new(Post), new(Tag))
// 	ORM = orm.NewOrm()
// }
// func main() {
// 	// cascadeUserProfile()
// 	// cascadeUserPost()
// 	insertPostTag()
// }

// func cascadeUserProfile() {
// 	createUser()
// 	// deleteUser(2)
// 	// deleteProfile()
// }
// func deleteProfile() {
// 	n, err := ORM.QueryTable((*Profile)(nil)).Filter("Id", 2).Delete()
// 	fmt.Println(n, err)
// }
// func deleteUser(id int) {
// 	n, err := ORM.QueryTable((*User)(nil)).Filter("Id", id).Delete()
// 	fmt.Println(n, err)
// }
// func createUser() {
// 	user := &User{Id: 2, Name: "jack"}
// 	profile := &Profile{Id: 2, Name: "jack profile"}
// 	user.Profile = profile
// 	profile.User = user
// 	ORM.Insert(user)
// 	ORM.Insert(profile)
// }

// func cascadeUserPost() {
// 	createPost()
// 	// deleteUser(1)
// 	queryUserPost(1)
// }

// func createPost() {
// 	user := &User{Id: 1, Name: "tom"}
// 	post1 := &Post{Id: 1, User: user}
// 	post2 := &Post{Id: 2, User: user}
// 	ORM.Insert(user)
// 	ORM.Insert(post1)
// 	ORM.Insert(post2)
// }
// func queryUserPost(id int) {
// 	var posts []*Post
// 	ORM.Raw("select * from post where uid=?", id).QueryRows(&posts)
// 	for _, it := range posts {
// 		fmt.Println(it.Id)
// 	}
// 	var maps []orm.Params
// 	ORM.Raw("select * from post where uid=?", id).Values(&maps, "id", "uid")
// 	fmt.Println(maps)
// 	for _, it := range maps {
// 		fmt.Println(it["id"], it["uid"])
// 	}
// 	ORM.Raw("select * from post where uid=?", id).Values(&maps, "id")
// 	fmt.Println(maps)
// 	for _, it := range maps {
// 		fmt.Println(it["id"])
// 	}

// 	var lists []orm.ParamsList
// 	ORM.Raw("select * from post where uid=?", id).ValuesList(&lists)
// 	fmt.Println(lists)
// 	ORM.Raw("select * from post where uid=?", id).ValuesList(&lists, "id")
// 	fmt.Println(lists)
// 	ORM.QueryTable((*User)(nil)).Filter("Id", id).Update(
// 		orm.Params{
// 			"name": "yocob",
// 		})
// }

// func insertPostTag() {
// 	// user := &User{Id: 1, Name: "tom"}
// 	post1 := &Post{ /*Id: 5 , User: user*/ }
// 	post2 := &Post{ /*Id: 6 , User: user*/ }
// 	posts := []*Post{post1, post2}
// 	tag1 := &Tag{ /*Id: 7 ,*/ Posts: posts}
// 	tag2 := &Tag{ /*Id: 8 ,*/ Posts: posts}
// 	tags := []*Tag{tag1, tag2}
// 	post1.Tags = tags
// 	post2.Tags = tags
// 	n, err := ORM.Insert(post1)
// 	fmt.Println(n, err)
// 	n, err = ORM.Insert(post2)
// 	fmt.Println(n, err)
// 	n, err = ORM.Insert(tag1)
// 	fmt.Println(n, err)
// 	n, err = ORM.Insert(tag2)
// 	fmt.Println(n, err)
// }

// type User struct {
// 	Id      int      `orm:"pk"`
// 	Name    string   `orm:"column(name)"`
// 	Profile *Profile `orm:"null;rel(one);reverse(one);on_delete(set_null);column(pid)"`
// 	Posts   []*Post  `orm:"reverse(many)" json:"-"`
// }

// type Profile struct {
// 	Id   int    `orm:"pk"`
// 	Name string `orm:"column(name)"`
// 	User *User  `orm:"reverse(one);column(uid);on_delete(cascade)"`
// }

type Post struct {
	Id int `orm:"pk"`
	// User *User  `orm:"rel(fk);column(uid);reverse(one);on_delete(cascade)"`
	Tags []*Tag `orm:"rel(m2m);rel_through(github.com/toukii/orm_test/dbtest3/models.PostTags);column(psid)"`
	// Tags []*Tag `orm:"rel(m2m);column(tsid)"`
}

func NewPost() *Post {
	obj := new(Post)
	return obj
}

type Tag struct {
	Id    int   `orm:"pk"`
	Posts *Post `orm:"reverse(one);column(psid)"`
}

func NewTag() *Tag {
	obj := new(Tag)
	return obj
}

type TagPosts struct {
	Id   int
	Post *Post `orm:"rel(fk);column(post_id)"`
	Tag  *Tag  `orm:"rel(fk);column(tag_id)"`
}

func (m *TagPosts) TableName() string {
	return "prefix_tag_posts"
}

type PostTags struct {
	Id   int
	Post *Post `orm:"rel(fk);column(post_id)"`
	Tag  *Tag  `orm:"rel(fk);column(tag_id)"`
}

func (m *PostTags) TableName() string {
	return "prefix_post_tags"
}
