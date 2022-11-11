package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	_defaultMaxPoolSize  = 1
	_defaultConnAttempts = 3
	_defaultConnTimeout  = 10
)

type Postgres struct {
	maxPoolsize  int
	connAttempts int
	connTimeout  time.Duration

	Builder squirrel.StatementBuilderType
	Pool    *pgxpool.Pool
}

func New(uri string, opts ...Option) (*Postgres, error) {
	pg := &Postgres{
		maxPoolsize:  _defaultMaxPoolSize,
		connAttempts: _defaultConnAttempts,
		connTimeout:  _defaultConnTimeout,
	}

	for _, opt := range opts {
		opt(pg)
	}

	pg.Builder = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

	poolConfig, err := pgxpool.ParseConfig(uri)
	if err != nil {
		return nil, fmt.Errorf("postgres - New - pgxpool.ParseConfig. error: %v", err)
	}

	poolConfig.MaxConns = int32(pg.maxPoolsize)

	for pg.connAttempts > 0 {
		pg.Pool, err = pgxpool.ConnectConfig(context.TODO(), poolConfig)
		if err == nil {
			break
		}

		//TODO logger
		fmt.Printf("Try to connect to Postgresql, try: %d", pg.connAttempts)

		time.Sleep(pg.connTimeout)

		pg.connAttempts--
	}

	if err != nil {
		return nil, fmt.Errorf("postgres - New - pgxpool.ConnectConfig. error: %v", err)
	}

	return pg, nil
}

//Close
func (p *Postgres) Close() {
	if p.Pool != nil {
		p.Pool.Close()
	}
}
