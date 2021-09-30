package database

import (
	"fmt"
	"github.com/leesper/couchdb-golang"
	"github.com/pkg/errors"
	"strings"
)

type CouchDB struct {
	Server couchdb.Server
	DB     couchdb.Database
}

var DATABASE CouchDB

func init() {
	DATABASE = CouchDB{}
	err := DATABASE.Init("balobas", "balobas", "http://localhost:5984")
	if err != nil {
		panic("db init error " + err.Error())
	}
}

func (c *CouchDB) Init(login, password, url string) error {
	server, err := couchdb.NewServer(url)
	if err != nil {
		panic(err)
	}
	c.Server = *server
	token, err := server.Login(login, password)
	if err != nil {
		return errors.New("login db error " + err.Error())
	}
	err = server.VerifyToken(token)
	if err != nil {
		return errors.New("verify token error " + err.Error())
	}
	DB, err := server.Get("ydocs")
	if DB == nil {
		return errors.New("nil pointer database ")
	}
	c.DB = *DB
	return err
}

func (c *CouchDB) Set(key string, obj interface{}) error {
	objMap, err := couchdb.ToJSONCompatibleMap(obj)
	if err != nil {
		return errors.Wrap(err, "marshalling error ")
	}
	savedObj, err := c.DB.Get(key, nil)
	if err != nil {
		if !strings.Contains(err.Error(), "not found") {
			fmt.Println("tut")
			return err
		}
	}
	if savedObj != nil {
		objMap["_rev"] = savedObj["_rev"]
	}
	return c.DB.Set(key, objMap)
}

func (c *CouchDB) Get(key string) (map[string]interface{}, error) {
	return c.DB.Get(key, nil)
}

func (c *CouchDB) Unmarshal(obj interface{}, objMap map[string]interface{}) (interface{}, error) {
	err := couchdb.FromJSONCompatibleMap(obj, objMap)
	return obj, err
}

func (c *CouchDB) Delete(key string) error {
	return c.DB.Delete(key)
}

func (c *CouchDB) QueryAllFieldsWithSelector(selector string) ([]map[string]interface{}, error) {
	return c.DB.Query(nil, selector, nil, nil, nil, nil)
}

func (c *CouchDB) QueryFieldsWithSelector(fields []string, selector string) ([]map[string]interface{}, error) {
	return c.DB.Query(fields, selector, nil, nil, nil, nil)
}

func (c *CouchDB) JSONQuery(query string) ([]map[string]interface{}, error) {
	return c.DB.QueryJSON(query)
}
