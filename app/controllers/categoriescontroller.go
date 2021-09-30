package controllers

import "../models"

type CategoriesController struct {}

func (cc *CategoriesController) GetCategories(page, limit int) ([]models.Category, error) {
	return models.GetCategories(page, limit)
}