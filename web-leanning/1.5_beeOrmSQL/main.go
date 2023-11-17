package main //使用beego 的orm框架操作数据库

import (
	//"database/sql"
	"fmt"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func init() {
	//注册驱动
	orm.RegisterDriver("mysql", orm.DRMySQL)
	//设置默认数据库
	orm.RegisterDataBase("default", "mysql", "root:admin@/test?charset=utf8", 30)
	//根据数据库的别名，设置数据库的最大空闲连接
	orm.SetMaxIdleConns("default", 30)
	//根据数据库的别名，设置数据库的最大数据库连接 (go >= 1.2)
	orm.SetMaxOpenConns("default", 30)
	//注册定义的model
	orm.RegisterModel(new(User))
	// 创建数据库table
	orm.RunSyncdb("default", false, true)
}

func main() {
	o := orm.NewOrm()

	user := User{Name: "slene"}

	//----------------插入表-----------
	id, err := o.Insert(&user)
	fmt.Printf("ID: %d,err: %v\n", id, err)
	//一次插入多个数据
	users := []User{
		{Name: "slene"},
		{Name: "astaxie"},
		{Name: "unknown"},
	}
	successNums, err := o.InsertMulti(10, users) //10表示可以10个数据一起插入，为 1 时，将会顺序插入 slice 中的数据
	fmt.Printf("successNums: %d,err: %v\n", successNums, err)

	//------------------更新表-------------
	user.Name = "asta"
	if o.Read(&user) == nil {
		user.Name = "MyName"
		if num, err := o.Update(&user); err == nil {
			fmt.Println(num)
		}
	}
	// 只更新 Name
	//o.Update(&user, "Name")
	// 指定多个字段
	// o.Update(&user, "Field1", "Field2", ...)

	//------------------读取(查询)------------
	u := User{Id: user.Id}
	err = o.Read(&u)

	if err == orm.ErrNoRows {
		fmt.Println("查询不到")
	} else if err == orm.ErrMissPK {
		fmt.Println("找不到主键")
	} else {
		fmt.Println(user.Id, user.Name)
	}
	//方法二
	qs := o.QueryTable(u) // 返回 QuerySeter
	qs.Filter("id", 1)    // WHERE id = 1
	//qs.Filter("profile__age", 18) // WHERE profile.age = 18
	//qs.Filter("profile__age__in", 18, 20).Exclude("profile__lt", 1000)// WHERE profile.age IN (18, 20) AND NOT profile_id < 1000

	//----------------删除表--------------------
	num, err := o.Delete(&u)
	fmt.Printf("NUM:%d,err:%v\n", num, err)

}

//复杂原生sql使用:
func Query(name string) (user []User) {
	var o orm.Ormer
	var rs orm.RawSeter
	o = orm.NewOrm()
	rs = o.Raw("SELECT * FROM user "+
		"WHERE name=? AND uid>10 "+
		"ORDER BY uid DESC "+
		"LIMIT 100", name)
	//var user []User
	num, err := rs.QueryRows(&user)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(num)
		//return user
	}
	return
}
