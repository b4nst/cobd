package testable

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jackc/pgx/v4/stdlib"
	_ "github.com/nakagami/firebirdsql"
	_ "github.com/segmentio/go-athena"
)

type SQL struct {
	connstr string
	driver  string
	err     error
}

var DriverMap map[string]string = map[string]string{
	"mysql":    "mysql",
	"postgres": "pgx",
	"firebird": "firebirdsql",
	"athena":   "athena",
}

func SQLFrom(key, value string) *SQL {
	d := DriverMap[strings.ToLower(strings.TrimPrefix(key, "TEST_SQL_"))]
	if d == "" {
		return &SQL{err: fmt.Errorf("Unknown sql driver: %s", key)}
	}
	return &SQL{
		connstr: value,
		driver:  d,
	}
}

func (s *SQL) Error() error {
	return s.err
}

func (s *SQL) Name() string {
	return fmt.Sprintf("Sql %s", s.driver)
}

func (s *SQL) Test() error {
	if s.err != nil {
		return s.err
	}

	c, err := sql.Open(s.driver, s.connstr)
	if err != nil {
		s.err = fmt.Errorf("Unable to connect to database: %v\n", err)
		return s.err
	}
	defer c.Close()

	if err := c.Ping(); err != nil {
		s.err = err
		return s.err
	}

	return s.err
}
