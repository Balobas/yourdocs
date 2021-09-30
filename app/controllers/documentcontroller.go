package controllers

import (
	"../models"
	"net/http"
)

type DocumentController struct{}

func (d *DocumentController) Search(categories, names, docNumbers []string) ([]models.Document, error) {
	docs, err := models.Search(categories, names, docNumbers)
	if err != nil {
		return nil, err
	}
	return docs, nil
}

func (d *DocumentController) GetDocumentFields(uid string) ([]string, error) {
	doc, err := models.GetDocument(uid)
	if err != nil {
		return []string{}, err
	}
	var fields = make([]string, 0)
	for fieldName, _ := range doc.Fields {
		fields = append(fields, fieldName)
	}
	return fields, nil
}

func (d *DocumentController) BuildDocument(w http.ResponseWriter, uid string, fields map[string]string) error {
	doc, err := models.GetDocument(uid)
	if err != nil {
		return err
	}
	for key, _ := range doc.Fields {
		if val, ok := fields[key]; ok {
			//TODO: Add val validation
			doc.Fields[key] = val
		}
	}
	return doc.Build(w)
}

func (d *DocumentController) GetDocument(uid string) (models.Document, error) {
	return models.GetDocument(uid)
}

func (d *DocumentController) PutSignedDocument(doc models.SignedDocument) (string, error) {
	return models.PutSignedDocument(doc)
}

func (d *DocumentController) GetSignedDocument(uid string) (models.SignedDocument, error) {
	return models.GetSignedDocument(uid)
}