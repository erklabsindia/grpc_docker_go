package apis

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
)

type PostgresDb struct {
	DatabaseName string
	*pgx.Conn
}

// Connect connects to a running PostgreSQL instance.
func Connect(ctx context.Context, host string, user string, password string, database string, port int) (*PostgresDb, error) {
	connString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=require", host, port, user, password, database)
	conn, err := pgx.Connect(ctx, connString)
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect to PostgreSQL server")
	}
	err = conn.Ping(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to ping PostgreSQL server")
	}
	return &PostgresDb{
		DatabaseName: database,
		Conn:         conn,
	}, nil
}
