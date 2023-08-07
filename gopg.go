package gopg

import (
	"context"
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5/pgxpool"
	"path/filepath"
)

func DBPool(ctx context.Context, dbUrl string) (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(ctx, dbUrl)
	if err != nil {
		return nil, fmt.Errorf("database connection fail cause %v", err)
	}
	if err = pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("fail to ping database cause %v", err)
	}
	return pool, nil
}

func UpMigrations(dbUrl string, migrationFilesDir string) error {
	dir := filepath.ToSlash(migrationFilesDir)
	m, err := migrate.New(
		fmt.Sprintf("file://%s", dir),
		dbUrl)
	if err != nil {
		return fmt.Errorf("fail to prepare database migrations cause %v", err)
	}
	if err = m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("fail to up database migrations cause %+v", err)
	}
	return nil
}
