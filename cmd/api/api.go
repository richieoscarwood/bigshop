package api

import "database/sql"

type ApiServer struct {
	address  string
	database *sql.DB
}

func NewApiServer(address string, database *sql.DB) *ApiServer {
	return &ApiServer{
		address:  address,
		database: database,
	}
}

func (receiver *ApiServer) Run() error {
	return nil

}
