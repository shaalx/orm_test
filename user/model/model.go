package model

type User struct {
	Id      int      `orm:"pk;auto"`
	Name    string   `orm:"unique"`
	Profile *Profile `orm:"rel(one);reverse(one)"`
}

type Profile struct {
	Id   int   `orm:"pk"`
	User *User `orm:"reverse(one);rel(one)"`
}

// one 2 many
type Userz struct {
	Id      int        `orm:"auto"`
	Name    string     `orm:"unique"`
	Friends []*Friends `orm:"reverse(many)"`
}

type Friends struct {
	Id     int
	Userz  *Userz `orm:"rel(fk)"`
	Userzt *Userz `orm:"rel(fk)"`
}
