package controllers

import "../models"

type UsersController struct {}

func (uc *UsersController) CreateUser(user models.User) (string, error) {
	return models.CreateUser(user)
}

func (uc *UsersController) UpdateUser(user models.User) (string, error) {
	return models.UpdateUser(user)
}

func (uc *UsersController) GetUser(uid string) (models.User, bool, error) {
	return models.GetUser(uid)
}