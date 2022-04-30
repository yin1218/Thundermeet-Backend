//關於orm相關的變數以及function，這邊也是主要對資料庫進行資料CRUD的地方。
package service

import (
	// "backend/app/dao"
	// "backend/app/model"
	"errors"
	"fmt"
	"log"
	"thundermeet_backend/app/dao"
	"thundermeet_backend/app/middleware/crypto"
	"thundermeet_backend/app/model"
)

var UserFields = []string{"userId", "userName", "passwordHash", "passwordAnswer"}

func SelectOneUser(id int64) (*model.User, error) {
	userOne := &model.User{}
	err := dao.SqlSession.Select(UserFields).Where("id=?", id).First(&userOne).Error
	if err != nil {
		return nil, err
	} else {
		return userOne, nil
	}
}

func RegisterOneUser(userName string, password string, passwordAnswer string) error {
	fmt.Print("registering user")
	fmt.Print(userName, password, passwordAnswer)
	if !CheckOneUser(userName) {
		return fmt.Errorf("User exists.")
	}

	// hash password
	hash, err := crypto.Generate(password)
	if err != nil {
		log.Println(err)
	}

	user := model.User{
		UserName:       userName,
		PasswordHash:   hash,
		PasswordAnswer: passwordAnswer,
	}

	log.Print("user = ", user)

	insertErr := dao.SqlSession.Model(&model.User{}).Create(&user).Error
	return insertErr
}

func CheckOneUser(userName string) bool {
	result := false
	var user model.User
	fmt.Print(result)
	dbResult := dao.SqlSession.Where("user_name = ?", userName).Find(&user)
	fmt.Print(dbResult)
	if dbResult.Error != nil {
		fmt.Printf("Get User Info Failed:%v\n", dbResult.Error)
	} else {
		result = true
	}
	fmt.Print(result)
	return result
}

func GetOneUserUsernamePasswordHash(userName string) (int64, string, error) {
	var user model.User

	dbResult := dao.SqlSession.Where("user_name = ?", userName).Find(&user)

	if dbResult.Error != nil {
		fmt.Printf("Get User Info Failed:%v\n", dbResult.Error)
		return 0, "", errors.New("Can't find user")
	} else {
		userId, passwordHash := user.ID, user.PasswordHash
		return userId, passwordHash, nil
	}
}
