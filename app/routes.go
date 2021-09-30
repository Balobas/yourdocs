package app

import (
	"github.com/gorilla/mux"
	"./api"
)

func BindRoutes(router *mux.Router) {

	//Documents
	router.HandleFunc("/api/buildDoc", api.BuildDocHandler)
	router.HandleFunc("/api/getDocFields", api.GetDocumentFieldsHandler)
	router.HandleFunc("/api/searchDocs", api.SearchDocumentsHandler)
	router.HandleFunc("/api/getDoc", api.GetDocumentHandler)
	//

	//Signed Documents
	router.HandleFunc("/api/getSignedDoc", api.GetSignedDocumentHandler)
	router.HandleFunc("/api/putSignedDoc", api.PutSignedDocumentHandler)
	//

	//Categories
	router.HandleFunc("/api/getDocumentCategories", api.GetCategoriesHandler)
	//

	//Reviews
	router.HandleFunc("/api/addReview", api.AddReviewHandler)
	router.HandleFunc("/api/updateReview", api.UpdateReviewHandler)
	router.HandleFunc("/api/deleteReview", api.DeleteReviewHandler)
	router.HandleFunc("/api/getReview", api.GetReviewHandler)
	router.HandleFunc("/api/answerOnReview", api.AnswerOnReviewHandler)
	router.HandleFunc("/api/getReviewCategories", api.GetReviewCategoriesHandler)
	//

	//Users

	//


}