package entity

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type User struct {
	Username           string
	Password           string
	Email              string
	Phone              string
}

/*GetCurrentUserName 获取当前登陆用户*/
func GetCurrentUserName() (username string) {
	dir, err := os.Getwd()
	checkerr(err)
	b, err := ioutil.ReadFile(dir + "/entity/currentUserName.txt")
	checkerr(err)
	username = string(b)
	return username
}

/*SetCurrentUserName 获取当前正在操作的用户名字*/
func SetCurrentUserName(username string) {
	dir, err := os.Getwd()
	checkerr(err)
	b := []byte(username)
	err = ioutil.WriteFile(dir+"/entity/currentUserName.txt", b, 0777)
	checkerr(err)
}

func READUSERS() (user []User) {
	dir, err := os.Getwd()
	checkerr(err)
	b, err := ioutil.ReadFile(dir + "/entity/Users.txt")
	checkerr(err)
	//json转变为对象
	var users []User
	json.Unmarshal(b, &users)
	// log.Println("READUSER success")
	return users
}

func WRITEUSER(users []User) {
	dir, err := os.Getwd()
	checkerr(err)
	data, err := json.Marshal(users)
	checkerr(err)
	b := []byte(data)
	err = ioutil.WriteFile(dir+"/entity/Users.txt", b, 0777)
	checkerr(err)
	// log.Println("WRITEUSER success")
}


func checkerr(err error) {
	if err != nil {
		fmt.Print(err)
	}
}