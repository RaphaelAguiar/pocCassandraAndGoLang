package database

import (
	"fmt"
	"log"

	"github.com/gocql/gocql"
)

var cluster *gocql.ClusterConfig

func Init() {
	cluster = gocql.NewCluster("172.17.0.1", "172.17.0.2")
	cluster.Keyspace = "brajd"
}

func Query(cql string, query func(*gocql.Iter) []interface{}) []interface{} {
	session, _ := cluster.CreateSession()
	defer session.Close()

	fmt.Println("Executando query: " + cql)
	iter := session.Query(cql).Iter()
	retorno := query(iter)
	if err := iter.Close(); err != nil {
		log.Fatal("Error: " + err.Error())
	}
	return retorno
}
