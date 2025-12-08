package service_test

import (
	"context"
	"database/sql" // Нам нужен стандартный sql для Goose
	"path/filepath"
	"runtime"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib" // ВАЖНО: Регистрируем драйвер pgx для database/sql
	"github.com/pressly/goose/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
	"github.com/vashkevichjr/innowise-cart-api/internal/repository"
	"github.com/vashkevichjr/innowise-cart-api/internal/service"
)

func SetupTestDase(t *testing.T) (*pgxpool.Pool, func()) {
	ctx := context.Background()

	dbName := "dbName"
	dbUser := "testUser"
	dbPass := "testPass"

	postgresContainer, err := postgres.Run(ctx,
		"postgres:16-alpine",
		postgres.WithDatabase(dbName),
		postgres.WithUsername(dbUser),
		postgres.WithPassword(dbPass),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(5*time.Second),
		),
	)
	require.NoError(t, err)

	connStr, err := postgresContainer.ConnectionString(ctx)
	require.NoError(t, err)

	migrationDB, err := sql.Open("pgx", connStr)
	require.NoError(t, err)
	defer migrationDB.Close()

	_, filename, _, _ := runtime.Caller(0)
	migrationsPath := filepath.Join(filepath.Dir(filename), "../../migrations")

	err = goose.SetDialect("postgres")
	require.NoError(t, err)

	err = goose.Up(migrationDB, migrationsPath)
	require.NoError(t, err)

	pool, err := pgxpool.New(ctx, connStr)
	require.NoError(t, err)

	return pool, func() {
		pool.Close()
		postgresContainer.Terminate(ctx)
	}
}

func TestCart(t *testing.T) {
	pool, teardown := SetupTestDase(t)
	defer teardown()

	ctx := context.Background()

	repo := repository.NewCartRepo(pool)

	svc := service.NewCart(repo)

	cart, err := svc.CreateCart(ctx)
	require.NoError(t, err)

	item, err := svc.CreateItem(ctx, "MacBook Pro M3", 6000.0)
	require.NoError(t, err)

	_, err = svc.AddItemToCart(ctx, cart.Id, item.ID, 5)
	require.NoError(t, err)

	calculator, err := svc.CalculatePrice(ctx, cart.Id)
	require.NoError(t, err)

	assert.Equal(t, float32(30000.0), calculator.TotalPrice, "TotalPrice is not correct")
	assert.Equal(t, int32(10), calculator.DiscountPercent, "Discount is not correct")
	assert.Equal(t, float32(27000.0), calculator.FinalPrice, "FinalPrice is not correct")
}
