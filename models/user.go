package models

import (
	"time"
)

type User struct {
	ID int
	Name string
	Age int
	Sex int
	CreateTime int64
	UpdateTime int64
}

func (user *User) First(id int) *User {
	orm.Where(&User{ID: id}).First(user)
	return user
}

func (_ *User) List() []User {
	var user []User
	orm.Select("id,name,age,sex,create_time,update_time").Order("id desc").Find(&user)
	return user
}

func (_ *User) Insert(name string, age int, sex int) int {
	createTime := time.Now().Unix()
	updateTime := time.Now().Unix()
	user := &User{Name:name, Age:age, Sex:sex, CreateTime:createTime, UpdateTime:updateTime}
	orm.Create(user)
	return user.ID
}

func (user *User) Edit(id int, name string, age int, sex int) int64 {
	ret := user.First(id)
	if ret.ID == 0 {
		return 0
	}
	updateTime := time.Now().Unix()
	rowsAffected := orm.Model(ret).Updates(map[string]interface{}{"name":name, "age":age, "sex": sex, "updateTime": updateTime}).RowsAffected
	return rowsAffected
}

func (user *User) Del(id int) int64 {
	ret := user.First(id)
	if ret.ID == 0 {
		return 0
	}
	rowsAffected := orm.Delete(ret).RowsAffected
	return rowsAffected
}


