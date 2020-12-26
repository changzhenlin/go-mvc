package models

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
)


type Page struct {
	Id int
	Website string
	Email string
}

func init() {
	orm.RegisterDataBase("default","mysql","root:root@tcp(39.97.64.148)/test?charset=utf8")
	orm.RegisterModel(new(Page))
}

func GetPage() Page {
	//rtn := Page{Website:"changzhenlin.com", Email: "chang1051169769@126.com"}
	//return rtn
	o := orm.NewOrm()
	p := Page{Id:1}
	err := o.Read(&p)
	fmt.Print(err)
	return p
}

func UpdatePage() {
	p := Page{Id: 1, Email: "woshichang"}
	o := orm.NewOrm()
	o.Update(&p)
}
