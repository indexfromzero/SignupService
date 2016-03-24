package service

import (
	"github.com/gocql/gocql"
	"github.com/microbusinesses/SignupService/domain"
)

var cluster *gocql.ClusterConfig
var isClusterInitialized bool
var isKeyspaceCreated bool

func initializeCluster() {
	cluster = gocql.NewCluster("172.17.0.2")
	cluster.Keyspace = "microbusinesses"
	cluster.Consistency = gocql.Quorum
	cluster.ProtoVersion = 4

	isClusterInitialized = true
}

func createKeyspace() {
	if !isClusterInitialized {
		panic("Cluster must be initialized before create the keyspace")
	}

	session, _ := cluster.CreateSession()

	session.Query(`create keyspace microbusinesses with replication = { 'class' : 'SimpleStrategy', 'replication_factor' : 1 };
                 create table microbusinesses.signup(id int, email text, username text, password text, PRIMARY KEY(id));`).Exec()

	isKeyspaceCreated = true
}

func signupUser(user domain.User) error {
	if !isKeyspaceCreated {
		panic("Keyspace must be created before attempting to signup a user")
	}

	session, _ := cluster.CreateSession()
	if err := session.Query(`INSERT INTO signup (email, username, password, id) VALUES (?, ?, ?, ?)`, user.Email, user.Username, user.Password, gocql.TimeUUID()).Exec(); err != nil {
		return err
	}
	return nil
}
