package dspattn

import (
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestBuildContainer(t *testing.T) {
	container := BuildContainer()

	err := container.Invoke(func(server *Server) {
		server.Run()
	})

	if err != nil {
		panic(err)
	}
}

func TestManual(t *testing.T) {
	config := NewConfig()

	db, err := ConnectDatabase(config)

	if err != nil {
		panic(err)
	}

	personRepository := NewPersonRepository(db)

	personService := NewPersonService(config, personRepository)

	server := NewServer(config, personService)

	server.Run()
}
