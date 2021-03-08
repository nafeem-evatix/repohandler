package repohandler

import (
	"errors"

	"github.com/google/uuid"
)

// InMemoryDB implements RepoHandler for a custom in memory database
type InMemoryDB struct {
	DB *map[string]interface{}
}

// Create inserts data by creating a generated uuid to track data
func (in *InMemoryDB) Create(data interface{}) error {
	uniqueID := uuid.New().String()
	database := *in.DB
	database[uniqueID] = data

	return nil
}

// Get retrieves all data from the map
func (in *InMemoryDB) Get() ([]interface{}, error) {
	database := *in.DB
	var values []interface{}

	for uniqueID := range database {
		singleValue := database[uniqueID]
		values = append(values, singleValue)
	}

	return values, nil
}

// GetSingle retrieves single data by using identifier
func (in *InMemoryDB) GetSingle(identifier string) (interface{}, error) {
	database := *in.DB
	singleValue := database[identifier]

	if singleValue == nil {
		return nil, errors.New("No Data Found")
	}

	return singleValue, nil
}

// Update edits data
func (in *InMemoryDB) Update(identifier string, data interface{}) error {
	database := *in.DB
	getValue := database[identifier]

	if getValue == nil {
		return errors.New("no data found")
	}

	database[identifier] = data

	return nil
}

// Delete removes data
func (in *InMemoryDB) Delete(identifier string) error {
	database := *in.DB
	getValue := database[identifier]

	if getValue == nil {
		return errors.New("no data found")
	}

	delete(database, identifier)

	return nil
}

// NewInMemoryDB returns pointer to an RepoHandler in this case
// returns InMemoryDB struct because it has implemented RepoHandler
func NewInMemoryDB() RepoHandler {
	inMemoryDBMap := make(map[string]interface{})

	return &InMemoryDB{DB: &inMemoryDBMap}
}
