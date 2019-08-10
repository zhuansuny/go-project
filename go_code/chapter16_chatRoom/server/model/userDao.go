package model

//Dao :data access object
//编写对User对象（实例）操作的各种方法，主要是增删改查
import (
	"encoding/json"
	"fmt"
	"go_code/chapter16_chatRoom/common/message"

	"github.com/garyburd/redigo/redis"
)

var (
	MyUserDao *UserDao
)

//结构体对User的各种操作
type UserDao struct {
	Pool *redis.Pool
}

//使用工厂模式，创建一个UserDao实例
func NewUserDao(pool *redis.Pool) (userDao *UserDao) {
	userDao = &UserDao{
		Pool: pool,
	}
	return
}

//1.根据用户的ID返回一个User实例
func (this *UserDao) getUserById(conn redis.Conn, id int) (user *User, err error) {
	//通过给定的Id去redis查询
	res, err := redis.String(conn.Do("hget", "user", id))
	if err != nil {
		if err == redis.ErrNil { //表示在users哈希中，没有找到对应的id
			err = ERROR_USER_NOTEXISTS
		}
		return
	}
	user = &User{}
	//这里我们需要把res 反序列化成User实例
	err = json.Unmarshal([]byte(res), user)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}
	fmt.Println(res)
	return
}

//完成登陆的校验

//如果用户的id和密码都正确，则返回一个user实例
//如果用户的id和密码错误，则返回一个error
func (this *UserDao) Login(userId int, userPwd string) (user *User, err error) {
	conn := this.Pool.Get()
	defer conn.Close()
	user, err = this.getUserById(conn, userId)
	if err != nil {
		return
	}
	if user.UserPwd != userPwd {
		err = ERROR_USER_PWD
		return
	}
	return
}

func (this *UserDao) Register(user *message.User) (err error) {
	conn := this.Pool.Get()
	defer conn.Close()
	_, err = this.getUserById(conn, user.UserId)
	if err == nil {
		err = ERROR_USER_EXISTS
		return
	}
	//这时，说明id在redis还没有，则可以完成注册
	data, err := json.Marshal(user)
	if err != nil {
		return
	}

	_, err = conn.Do("HSet", "user", user.UserId, string(data))
	if err != nil {
		fmt.Println("注册用户错误 err=", err)
		return
	}
	return
}
