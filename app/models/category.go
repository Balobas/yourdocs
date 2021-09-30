package models

import (
	"../database"
	"encoding/json"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
)

type Category struct {
	UID         string
	Name        string
	Description string
}

func GetCategory(uid string) (Category, error) {
	categoryMap, err := database.DATABASE.Get(uid)
	if err != nil {
		return Category{}, err
	}
	var category Category
	b, err := json.Marshal(categoryMap)
	if err != nil {
		return Category{}, err
	}
	err = json.Unmarshal(b, &category)
	if err != nil {
		return Category{}, err
	}
	return category, nil
}

func PutCategory(category Category) (string, error) {
	if category.Name == "" {
		return "", errors.New("Empty name")
	}
	uid, err := uuid.NewV4()
	if err != nil {
		return "", err
	}
	category.UID = uid.String()
	err = database.DATABASE.Set("cat"+category.UID, category)
	if err != nil {
		return "", err
	}
	return category.UID, nil
}

func GetCategories(page, limit int) ([]Category, error) {
	objectsMap, err := database.DATABASE.QueryAllFieldsWithSelector("")
	if err != nil {
		return []Category{}, errors.Wrap(err, "query error")
	}
	b, err := json.Marshal(objectsMap)
	if err != nil {
		return nil, errors.Wrap(err, " marshal error ")
	}
	var categories = make([]Category, 0)
	err = json.Unmarshal(b, &categories)
	if err != nil {
		return []Category{}, errors.Wrap(err, "unmarshal error")
	}
	return categories, nil
}
