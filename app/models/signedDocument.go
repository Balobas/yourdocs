package models

import (
	"../database"
	"encoding/json"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
)

type SignedDocument struct {
	UID         string            `json:"uid"`
	UserUID     string            `json:"userUid"`
	DocumentUID string            `json:"documentUid"`
	Fields      map[string]string `json:"fields"`
}

func PutSignedDocument(signedDoc SignedDocument) (string, error) {
	if signedDoc.UID == "" {
		uid, err := uuid.NewV4()
		if err != nil {
			return "", err
		}
		signedDoc.UID = uid.String()
	} else if _, err := uuid.FromString(signedDoc.UID); err != nil {
		return "", errors.New("invalid uid")
	}
	err := database.DATABASE.Set(signedDoc.UID, signedDoc)
	if err != nil {
		return "", errors.WithStack(err)
	}
	return signedDoc.UID, err
}

func GetSignedDocument(uid string) (SignedDocument, error) {
	if _, err := uuid.FromString(uid); err != nil {
		return SignedDocument{}, err
	}
	docMap, err := database.DATABASE.Get(uid)
	if err != nil {
		return SignedDocument{}, err
	}
	b, err := json.Marshal(docMap)
	if err != nil {
		return SignedDocument{}, errors.New("Marshal map error")
	}
	var signedDoc SignedDocument
	err = json.Unmarshal(b, &signedDoc)
	if err != nil {
		return SignedDocument{}, errors.New("Unmarshal error")
	}
	return signedDoc, nil
}

func GetSignedDocuments(uids []string) ([]SignedDocument, error) {
	selector := ``
	for _, uid := range uids {
		selector += `uid=="` + uid + `" && `
	}
	selector = selector[:len(selector)-3]
	objectsMap, err := database.DATABASE.QueryAllFieldsWithSelector(selector)
	if err != nil {
		return []SignedDocument{}, errors.WithStack(err)
	}
	b, err := json.Marshal(objectsMap)
	if err != nil {
		return nil, errors.Wrap(err, " marshal error ")
	}
	var docs = make([]SignedDocument, 0)
	err = json.Unmarshal(b, &docs)
	if err != nil {
		return nil, errors.Wrap(err, "unmarshal error")
	}
	return docs, nil
}
