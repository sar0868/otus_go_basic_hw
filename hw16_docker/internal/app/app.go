package app

import (
	"context"
	"fmt"
	"log"
	"net"
	"strconv"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sar0868/otus_go_basic_hw/hw16_docker/internal/config"
	"github.com/sar0868/otus_go_basic_hw/hw16_docker/internal/repository"
)

var (
	Ctx  context.Context
	Repo repository.Querier
	DB   *pgxpool.Pool
)

func Init() *config.Cfg {
	Ctx = context.Background()
	conf, err := config.Init()
	if err != nil {
		log.Panicln(err.Error())
	}
	db, errDB := NewDB(Ctx, conf.DB)
	if errDB != nil {
		log.Panicln(errDB.Error())
	}
	DB = db
	log.Println("Connected to database")
	Repo = repository.New(DB)
	return conf
}

func NewDB(ctx context.Context, dbCfg config.DB) (*pgxpool.Pool, error) {
	connConfig, err := pgx.ParseConfig(
		fmt.Sprintf(
			"postgres://%s:%s@%s/%s?TimeZone=Europe/Moscow",
			dbCfg.User,
			dbCfg.Password,
			net.JoinHostPort(dbCfg.Host, strconv.Itoa(dbCfg.Port)),
			dbCfg.Database,
		),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create DSN for DB connection: %w", err)
	}
	dbc, errPgx := pgxpool.New(ctx, connConfig.ConnString())
	if errPgx != nil {
		return nil, fmt.Errorf("failed to connect to DB: %w", errPgx)
	}
	if errPing := dbc.Ping(ctx); errPing != nil {
		return nil, fmt.Errorf("failed to ping DB: %w", errPing)
	}
	return dbc, nil
}
