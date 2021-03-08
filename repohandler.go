package repohandler

//  Repohandler is a general interface
//  for working with database

// RepoHandler ...
type RepoHandler interface {
	Create(data interface{}) error
	Get() ([]interface{}, error)
	GetSingle(identifier string) (interface{}, error)
	Update(identifier string, data interface{}) error
	Delete(identifier string) error
}
