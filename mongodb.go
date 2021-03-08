package repohandler

import "github.com/globalsign/mgo"

// MongoDB ...
type MongoDB struct {
	Session      *mgo.Session
	DBName       string
	DBCollection string
}

// Create ...
// Todo : Implement Create
func (mdb *MongoDB) Create(data interface{}) error {
	return nil
}

// Get ...
// Todo : Implement Get
func (mdb *MongoDB) Get() ([]interface{}, error) {
	return nil, nil
}

// GetSingle ...
// Todo : Implement GetSingle
func (mdb *MongoDB) GetSingle(identifier string) (interface{}, error) {
	return nil, nil
}

// Update ...
// Todo : Implement Update
func (mdb *MongoDB) Update(identifier string, data interface{}) error {
	return nil
}

// Delete ...
// Todo : Implement Delete
func (mdb *MongoDB) Delete(identifier string) error {
	return nil
}

// NewMongoDB returns pointer to an RepoHandler in this case
// returns MongoDB struct because it has implemented RepoHandler
func NewMongoDB(dbSession *mgo.Session, DBName string, DBCollection string) RepoHandler {
	return &MongoDB{
		Session:      dbSession,
		DBName:       DBName,
		DBCollection: DBCollection,
	}
}
