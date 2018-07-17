package cliente

import (
	"brajd/database"
	"encoding/json"
	"net/http"

	"github.com/gocql/gocql"
)

type Cliente struct {
	Username string
	password string
}

type Repository struct {
}

func (repository Repository) ObterTodos() []Cliente {
	var fetch = func(iter *gocql.Iter) []interface{} {
		var retorno []interface{}
		var username string
		var password string
		for iter.Scan(&username, &password) {
			var cliente Cliente
			cliente.Username = username
			cliente.password = password
			retorno = append(retorno, cliente)
		}
		return retorno
	}

	result := database.Query(
		"select username,password from cliente",
		fetch,
	)

	retorno := make([]Cliente, len(result))
	for index, cliente := range result {
		retorno[index] = cliente.(Cliente)
	}
	return retorno
}

type Resource struct {
}

func (resource Resource) GetTodosClientes(w http.ResponseWriter, r *http.Request) {
	var repository Repository
	json.NewEncoder(w).Encode(repository.ObterTodos())
}
