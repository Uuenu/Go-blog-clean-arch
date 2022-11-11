package postgres

import "time"

type Option func(*Postgres)

func MaxPoolSize(poolSize int) Option {
	return func(p *Postgres) {
		p.maxPoolsize = poolSize
	}
}

func ConnAttempts(connAttempts int) Option {
	return func(p *Postgres) {
		p.connAttempts = connAttempts
	}
}

func ConnTimeout(connTimeout time.Duration) Option {
	return func(p *Postgres) {
		p.connTimeout = connTimeout
	}
}
