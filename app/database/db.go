package database

//
//var (
//	server couchdb.Server
//	DB *couchdb.Database
//)
//
//func Init() {
//	server, err := couchdb.NewServer("http://localhost:5984")
//	if err != nil {
//		panic(err)
//	}
//	token, err := server.Login("balobas", "balobas")
//	if err != nil {
//		panic("login db error " + err.Error())
//	}
//	err = server.VerifyToken(token)
//	if err != nil {
//		panic("verify token error " + err.Error())
//	}
//	DB, err = server.Get("aio")
//	//DB.UpdateDoc()
//	if err != nil {
//		panic(err)
//	}
//}
//
//func Set(key string, obj interface{}) error {
//	objMap, err := couchdb.ToJSONCompatibleMap(obj)
//	if err != nil {
//		return errors.Wrap(err, "marshalling error ")
//	}
//	return DB.Set(key, objMap)
//}
//
//
