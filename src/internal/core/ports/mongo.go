package ports

type MongoRepo interface {
	FindAll(collection string, filter interface{}, result interface{}) error
}
