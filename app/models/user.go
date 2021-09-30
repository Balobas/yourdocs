package models

import (
	"encoding/json"
	"github.com/pkg/errors"
	"../database"
	uuid "github.com/satori/go.uuid"
)

//TODO: добавить подтверждение аккаунта по почте

type User struct {
	UID        string `json:"uid"`
	Name       string `json:"name"`
	LastName   string `json:"lastName"`
	Patronymic string `json:"patronymic"`
	Email      string `json:"email"`
	Address    string `json:"address"`
	IsConfirmed bool  `json:"isConfirmed"`
	//uid ы оформленных документов
	Docs []string `json:"docs"`
}

//TODO: добавить валидацию email
func CreateUser(user User) (string, error) {
	uid, err := uuid.NewV4()
	if err != nil {
		return "", err
	}
	for _, isFound, _ := GetUser(uid.String()); isFound == true; {
		uid, err = uuid.NewV4()
		if err != nil {
			return "", err
		}
	}
	user.UID = uid.String()
	if user.Name == "" {
		return "", errors.New("Name is empty")
	}
	if user.LastName == "" {
		return "", errors.New("Last name is empty")
	}
	user.IsConfirmed = false
	err = database.DATABASE.Set(user.UID, user)
	if err != nil {
		return "", errors.WithStack(err)
	}
	return user.UID, err
}

/*
При обновлении данных, новые данные сохраняются с uid temp_
Это сделано для того, чтобы при обновлении данных, если пользователь не подтверждает действие по почте,
была возможность восстановить старые данные
 */
func UpdateUser(user User) (string, error) {
	if _, err := uuid.FromString(user.UID); err != nil {
		return "", errors.New("invalid uid")
	}
	oldUserData, isFound, err := GetUser(user.UID)
	if err != nil {
		return user.UID, err
	}
	if !isFound {
		return "", errors.New("User does not exists")
	}
	oldUserData.Name = user.Name
	oldUserData.LastName = user.LastName
	oldUserData.Patronymic = user.Patronymic
	oldUserData.Address = user.Address
	if user.Email != oldUserData.Email {
		//валидация email
	}
	//нужно подтвердить изменения по почте
	oldUserData.IsConfirmed = false
	tempUid := "temp_" + user.UID
	err = database.DATABASE.Set(tempUid, oldUserData)
	return user.UID, nil
}

func ConfirmUser(uid, token string) {

}

func GetUser(uid string) (User, bool, error) {
	if _, err := uuid.FromString(uid); err != nil {
		return User{}, false, err
	}
	userMap, err := database.DATABASE.Get(uid)
	if err != nil {
		if err.Error() == "status 404 - not found" {
			return User{}, false, nil
		}
		return User{}, false, err
	}
	b, err := json.Marshal(userMap)
	if err != nil {
		return User{}, true, errors.New("Marshal map error")
	}
	var user User
	err = json.Unmarshal(b, &user)
	if err != nil {
		return User{}, true, errors.New("Unmarshal error")
	}
	return user, true, nil
}