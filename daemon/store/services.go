package store

type Service struct {
}

type serviceStore struct {
	db *dbContext
}

func newServiceStore(db *dbContext) *serviceStore {
	return &serviceStore{db: db}
}
