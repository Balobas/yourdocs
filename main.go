package main

import (
	"./app/controllers"
	"./app/models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"./app"
)

func handler(w http.ResponseWriter, r *http.Request) {
	//doc := models.Document{
	//	Name:     "Doc",
	//	Number:   "1233123123",
	//	Category: models.Category{},
	//	Path:     "",
	//	Fields:   nil,
	//}
	//err := doc.Build(w)
	//if err != nil {
	//	fmt.Println(err)
	//}
	doc, err := models.CreateDocument("Document asasdasdasdasd", "1213asdasdad", "cat6c885e29-21c7-44d4-acd4-1fcbc420c597", map[string]string{"name": ""}, []byte("Document num 123123 {{.Name}} , {{.fields.name}}"))
	fmt.Println(doc, "\n ", err)
	//doc, err := models.GetDocument("doc921733ec-f979-46d6-8c71-27f904b18888")
	//if err != nil {
	//	fmt.Println("error get doc")
	//}
	//err = doc.Build(w)
	//if err != nil {
	//	fmt.Println("error build doc ", err)
	//}
}

func main() {

	cc := controllers.CategoriesController{}
	cat, _ := cc.GetCategories(0,0)
	fmt.Println(cat)

	c := controllers.DocumentController{}
	docs, _ := c.Search([]string{"6c885e29-21c7-44d4-acd4-1fcbc420c597", "2"}, []string{"Document asdasd", "name2"}, []string{"1213asdasdad", "num2"})
	//fmt.Println(docs, )
	b, _ := json.Marshal(docs)
	fmt.Println(string(b))
	doc, err := models.CreateDocument("Document huita", "1213asdasdad", "cat6c885e29-21c7-44d4-acd4-1fcbc420c597", map[string]string{"name": "", }, []byte("Document num 123123 {{.Name}} , {{.Fields.name}}"))
	b, _ = json.Marshal(doc)
	fmt.Println(string(b))

	router := mux.NewRouter()
	app.BindRoutes(router)
	err = http.ListenAndServe("localhost:8089", router)
	if err != nil {
		fmt.Println(err)
	}


}

func test() {
	category := models.Category{
		Name: "UR",
	}
	uid, err := models.PutCategory(category)
	fmt.Println(uid, " \n ", err)
}
