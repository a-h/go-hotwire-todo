package todo

type Todo struct {
	ID       string
	Item     string
	Complete bool
}

// Simulate a database.
var todos map[string]*Todo

func init() {
	todos = map[string]*Todo{}
}

type DB struct {
}

func (db DB) List() ([]*Todo, error) {
	rv := make([]*Todo, len(todos))
	var i int
	for _, v := range todos {
		rv[i] = v
		i++
	}
	return rv, nil
}

func (db DB) Get(id string) (*Todo, error) {
	r, _ := todos[id]
	return r, nil
}

func (db DB) Upsert(id, item string, complete bool) error {
	todos[id] = &Todo{
		ID:       id,
		Item:     item,
		Complete: complete,
	}
	return nil
}
