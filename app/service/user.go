//關於orm相關的變數以及function，這邊也是主要對資料庫進行資料CRUD的地方。
package service

import (
	"errors"
	"fmt"
	"log"
	"thundermeet_backend/app/dao"
	"thundermeet_backend/app/middleware/crypto"
	"thundermeet_backend/app/model"
	"time"
)

var UserFields = []string{"user_Id", "user_Name", "password_Hash", "password_Answer"}

func SelectOneUser(id string) (*model.User, error) {
	userOne := &model.User{}
	err := dao.SqlSession.Select(UserFields).Where("User_Id=?", id).First(&userOne).Error
	if err != nil {
		return nil, err
	} else {
		return userOne, nil
	}
}

func RegisterOneUser(userId string, userName string, password string, passwordAnswer string) error {
	fmt.Print("registering user")
	fmt.Print(userName, password, passwordAnswer)
	if !CheckOneUser(userId) {
		return fmt.Errorf("User exists.")
	}

	// hash password
	hash, err := crypto.Generate(password)
	if err != nil {
		log.Println(err)
	}

	user := model.User{
		UserId:         userId,
		UserName:       userName,
		PasswordHash:   hash,
		PasswordAnswer: passwordAnswer,
	}

	log.Print("user = ", user)

	insertErr := dao.SqlSession.Model(&model.User{}).Create(&user).Error
	return insertErr
}

func CheckOneUser(userId string) bool {
	result := false
	var user model.User
	fmt.Print(result)
	dbResult := dao.SqlSession.Where("user_id = ?", userId).Find(&user)
	fmt.Print(dbResult)
	if dbResult.Error != nil {
		fmt.Printf("Get User Info Failed:%v\n", dbResult.Error)
	} else {
		result = true
	}
	fmt.Print(result)
	return result
}

func GetOneUserUsernamePasswordHash(userId string) (string, string, error) {
	var user model.User

	dbResult := dao.SqlSession.Where("user_id = ?", userId).Find(&user)

	if dbResult.Error != nil {
		fmt.Printf("Get User Info Failed:%v\n", dbResult.Error)
		return "", "", errors.New("Can't find user")
	} else {
		userId, passwordHash := user.UserId, user.PasswordHash
		return userId, passwordHash, nil
	}
}

func UpdateOneUser(userId string, userName string, password string, passwordAnswer string) error {
	fmt.Print("Updating user")
	fmt.Print("userID = ", userId, " ")
	fmt.Print(userName, password, passwordAnswer)
	if !CheckOneUser(userId) {
		return fmt.Errorf("User Not exists.")
	}

	//hash password
	hash, err := crypto.Generate(password)
	if err != nil {
		log.Println(err)
		return err
	}
	user := model.User{
		UserName:       userName,
		PasswordHash:   hash,
		PasswordAnswer: passwordAnswer,
	}

	log.Print("user = ", user)
	updateErr := dao.SqlSession.Model(&model.User{}).Where("user_id = ?", userId).Updates(map[string]interface{}{"UserName": userName, "PasswordHash": hash, "PasswordAnswer": passwordAnswer}).Error
	return updateErr
}

func ResetUserPassword(userId string, password string, passwordAnswer string) error {
	//Check whether user exist
	if !CheckOneUser(userId) {
		return fmt.Errorf("User Not exists.")
	}
	//Check whether the answer is correct

	type pwdField struct {
		PasswordAnswer string `gorm:"size:100;not null" json:"passwordAnswer"`
	}
	userOne := &model.User{}
	err := dao.SqlSession.Select(UserFields).Where("User_Id=?", userId).First(&userOne).Error
	if err != nil {
		return err
	}

	fmt.Print("Correct Pwd = ", userOne.PasswordAnswer)
	fmt.Print("Correct Pwd = ", passwordAnswer)

	if userOne.PasswordAnswer != passwordAnswer {
		return fmt.Errorf("Incorrect answer")
	}

	//hash password
	hash, err := crypto.Generate(password)
	if err != nil {
		log.Println(err)
		return err
	}

	//update password
	updateErr := dao.SqlSession.Model(&model.User{}).Where("user_id = ?", userId).Update("password_hash", hash).Error
	return updateErr

}

func CreateEvent(eventName string, isPriorityEnabled bool, startTime time.Time, endTime time.Time, dateOrDays bool, startDay string, endDay string, startDate time.Time, endDate time.Time, adminId string) error {

	event := model.Event{
		EventName:         eventName,
		IsPriorityEnabled: isPriorityEnabled,
		IsConfirmed:       false,
		StartTime:         startTime,
		EndTime:           endTime,
		DateOrDays:        dateOrDays,
		StartDay:          startDay,
		EndDay:            endDay,
		StartDate:         startDate,
		EndDate:           endDate,
		AdminId:           adminId,
	}

	insertErr := dao.SqlSession.Model(&model.Event{}).Create(&event).Error
	return insertErr
}
