package controllers

import "../models"

type ReviewsController struct{}

func (rc *ReviewsController) AddReview(review models.Review) (string, error) {
	return models.AddReview(review)
}

func (rc *ReviewsController) UpdateReview(executorUid string, review models.Review) error {
	return models.UpdateReview(executorUid, review)
}

func (rc *ReviewsController) DeleteReview(userUid, reviewUid string) error {
	return models.DeleteReview(userUid, reviewUid)
}

func (rc *ReviewsController) GetReviewCategories() map[int]string {
	return models.GetReviewCategories()
}

func (rc *ReviewsController) GetReview(uid string) (models.Review, bool, error) {
	return models.GetReview(uid)
}

func (rc *ReviewsController) AnswerOnReview(reviewUid string, answer models.ReviewAnswer) (string, error) {
	return models.AnswerOnReview(reviewUid, answer)
}

func (rc *ReviewsController) GetReviewAnswer(answerUid string) (models.ReviewAnswer, error) {
	return models.GetReviewAnswer(answerUid)
}

