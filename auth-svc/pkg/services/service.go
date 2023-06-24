package services

import (
	"github.com/asadlive84/auth-svc/pkg/db"
	q "github.com/asadlive84/auth-svc/pkg/query"
)

type Server struct {
	H db.DbHandler
	Q q.Query
}
