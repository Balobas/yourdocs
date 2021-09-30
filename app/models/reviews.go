package models

/*
Review - отзывы пользователей
ReviewAnswer - ответ техподдержки на отзывы
ReviewCategories - карта категорий отзывов:
	Пожелания
	Жалоба
	Вопрос
	Другое

Отзыв может быть закрыт, подразумевает, что ответ техподдержки удовлетворил
Пользователь может удалить свой отзыв DeleteReview
Пользователь может изменить содержимое своего отзыва
 */

import (
	"encoding/json"
	"github.com/pkg/errors"
	"../database"
	uuid "github.com/satori/go.uuid"
)

type Review struct {
	UID string `json:"uid"`
	UserUid string `json:"userUid"`
	Text string `json:"text"`
	CategoryNum int `json:"categoryNum"`
	CategoryName string `json:"categoryName"`
	Date int64 `json:"date"`
	IsCancel bool `json:"isCancel"`
	AnswerUid string `json:"answerUid"`
}

type ReviewAnswer struct {
	UID string `json:"uid"`
	Text string `json:"text"`
	UserName string `json:"userName"`
	UserLastName string `json:"userLastName"`
}

var ReviewCategories = map[int]string {
	1: "Пожелания",
	2: "Жалоба",
	3: "Вопрос",
}

func GetReviewCategoryByNumber(number int) string {
	if number > len(ReviewCategories) || number < 1 {
		return "Прочее"
	}
	return ReviewCategories[number]
}

func GetReviewCategories() map[int]string {
	return ReviewCategories
}

func AddReview(review Review) (string, error) {
	uid, err := uuid.NewV4()
	if err != nil {
		return "", err
	}
	review.UID = uid.String()
	review.CategoryName = GetReviewCategoryByNumber(review.CategoryNum)

	if _, err := uuid.FromString(review.UserUid); err != nil {
		return "", errors.New("invalid user uid")
	}
	review.IsCancel = false
	review.AnswerUid = ""
	err = database.DATABASE.Set(review.UID, review)
	if err != nil {
		return review.UID, err
	}
	return review.UID, nil
}

func UpdateReview(executorUid string, review Review) error {
	if executorUid != review.UserUid {
		return errors.New("Access denied")
	}
	if review.AnswerUid != "" {
		return errors.New("Access denied")
	}
	if _, err := uuid.FromString(review.UID); err != nil {
		return errors.New("invalid review uid")
	}
	err := database.DATABASE.Set(review.UID, review)
	if err != nil {
		return err
	}
	return nil
}

func GetReview(uid string) (Review, bool, error) {
	if _, err := uuid.FromString(uid); err != nil {
		return Review{}, false, err
	}
	reviewMap, err := database.DATABASE.Get(uid)
	if err != nil {
		if err.Error() == "status 404 - not found" {
			return Review{}, false, nil
		}
		return Review{}, false, err
	}
	b, err := json.Marshal(reviewMap)
	if err != nil {
		return Review{}, true, errors.New("Marshal map error")
	}
	var review Review
	err = json.Unmarshal(b, &review)
	if err != nil {
		return Review{}, true, errors.New("Unmarshal error")
	}
	return review, true, nil
}

func DeleteReview(userUid, reviewUid string) error {
	return nil
}

func AnswerOnReview(reviewUid string, answer ReviewAnswer) (string, error) {
	if _, err := uuid.FromString(reviewUid); err != nil {
		return "", errors.New("invalid review uid")
	}
	review, isFound, err := GetReview(reviewUid)
	if err != nil {
		return "", err
	}
	if !isFound {
		return "", errors.New("Review with uid : " + reviewUid + " not found")
	}
	if review.IsCancel {
		return "", errors.New("Review already is cancel")
	}
	uid, err := uuid.NewV4()
	if err != nil {
		return "", err
	}
	answer.UID = uid.String()
	review.AnswerUid = answer.UID
	err = database.DATABASE.Set(review.UID, review)
	if err != nil {
		return "", err
	}
	err = database.DATABASE.Set(answer.UID, answer)
	if err != nil {
		return "", err
	}
	return answer.UID, nil
}

func GetReviewAnswer(answerUid string) (ReviewAnswer, error) {
	if _, err := uuid.FromString(answerUid); err != nil {
		return ReviewAnswer{}, errors.New("invalid answer uid")
	}
	answerMap, err := database.DATABASE.Get(answerUid)
	if err != nil {
		return ReviewAnswer{}, err
	}
	b, err := json.Marshal(answerMap)
	if err != nil {
		return ReviewAnswer{}, errors.New("Marshal map error")
	}
	var answer ReviewAnswer
	err = json.Unmarshal(b, &answer)
	if err != nil {
		return ReviewAnswer{}, errors.New("Unmarshal error")
	}
	return answer, nil
}
