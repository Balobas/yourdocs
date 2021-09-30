package models

import (
	"../database"
	"../templates"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"io"
	"os"
	"text/template"
)

//TODO: add export formats

type Document struct {
	UID         string `json:"uid"`
	Name        string `json:"name"`
	Number      string `json:"number"`
	Category    `json:"Category"`
	Path        string            `json:"path"`
	Fields      map[string]string `json:"fields"`
	Doc         []byte            `json:"doc"`
	Description string            `json:"description"`
}

//tested
//Before using this method, you must enter Fields map by values without save to DB
func (doc *Document) Build(w io.Writer) error {
	file, path, err := templates.CreateTempTemplate(doc.Doc)
	defer func() {
		err := file.Close()
		if err != nil {
			fmt.Println("error close file: ", err.Error())
		}
		err = os.Remove(path)
		if err != nil {
			fmt.Println("error remove file: ", err.Error())
		}
	}()
	if err != nil {
		return err
	}
	fmt.Println(path)
	t, err := template.ParseFiles(path)
	if err != nil {
		return err
	}
	err = t.Execute(w, doc)
	fmt.Println(doc)
	if err != nil {
		return err
	}
	return nil
}

//tested
//Fields values must be empty
func CreateDocument(name, number, categoryUID string, fields map[string]string, document []byte) (Document, error) {
	var doc Document
	uid, err := uuid.NewV4()
	if err != nil {
		return Document{}, err
	}
	doc.UID = "doc" + uid.String()
	doc.Name = name
	doc.Number = number
	doc.Fields = fields
	doc.Doc = document
	category, err := GetCategory(categoryUID)
	doc.Category = category
	doc.Doc = document
	if err != nil {
		return Document{}, err
	}
	err = database.DATABASE.Set(doc.UID, doc)
	if err != nil {
		return Document{}, err
	}
	return doc, nil
}

//tested
func GetDocument(uid string) (Document, error) {
	docMap, err := database.DATABASE.Get(uid)
	if err != nil {
		return Document{}, err
	}
	var doc Document
	b, err := json.Marshal(docMap)
	if err != nil {
		return Document{}, err
	}
	err = json.Unmarshal(b, &doc)
	if err != nil {
		return Document{}, err
	}
	return doc, nil
}

//tested
func Search(categories, names, docNumbers []string) ([]Document, error) {
	//TODO: refactor this shit
	selector := ``
	part1 := ``
	if len(categories) != 0 {
		for _, category := range categories {
			part1 += `Category.UID=="` + category + `" ||`
		}
		part1 = part1[:len(part1)-2]
		part1 += "&&"
	}
	part2 := ``
	if len(names) != 0 {
		for _, name := range names {
			part2 += `Name=="` + name + `" ||`
		}
		part2 = part2[:len(part2)-2]
		part2 += "&&"
	}
	part3 := ``
	if len(docNumbers) != 0 {
		for _, num := range docNumbers {
			part3 += `Number=="` + num + `" ||`
		}
		part3 = part3[:len(part3)-2]
		part3 += ""
	}
	//
	selector += part1 + part2 + part3
	var objectsMap []map[string]interface{}
	var err error
	if selector != "" {
		objectsMap, err = database.DATABASE.QueryFieldsWithSelector([]string{"Name", "UID", "Number", "Category", "Description"}, selector)
	} else {
		objectsMap, err = database.DATABASE.JSONQuery("{}")
	}
	if err != nil {
		return nil, errors.Wrap(err, "query error")
	}
	b, err := json.Marshal(objectsMap)
	if err != nil {
		return nil, errors.Wrap(err, " marshal error ")
	}
	var docs = make([]Document, 0)
	err = json.Unmarshal(b, &docs)
	if err != nil {
		return nil, errors.Wrap(err, "unmarshal error")
	}
	return docs, nil
}
