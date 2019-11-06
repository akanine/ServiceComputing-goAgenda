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

/*GetCurrentUserName ��ȡ��ǰ��½�û�*/
func GetCurrentUserName() (username string) {
	dir, err := os.Getwd()
	checkerr(err)
	b, err := ioutil.ReadFile(dir + "/entity/currentUserName.txt")
	checkerr(err)
	username = string(b)
	return username
}

/*SetCurrentUserName ��ȡ��ǰ���ڲ������û�����*/
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
	//jsonת��Ϊ����
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