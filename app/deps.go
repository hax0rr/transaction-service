package app

import (
	"github.com/hax0rr/transaction-service/config"
	"github.com/hax0rr/transaction-service/db"
	"github.com/hax0rr/transaction-service/httpserver"
	"github.com/hax0rr/transaction-service/internal/repository"
	"github.com/hax0rr/transaction-service/internal/service/account"
	"github.com/hax0rr/transaction-service/internal/service/transaction"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog"
	"os"
)

type Dependencies struct {
	Logger             *zerolog.Logger
	AccountService     account.IService
	TransactionService transaction.IService
}

func NewDependencies(conf *config.Config) (*Dependencies, error) {
	db, err := db.NewDB(&conf.Database)
	if err != nil {
		return nil, err
	}

	dbRepo := repository.New(db, conf.Database.TimeoutInMs)
	accountService := account.New(dbRepo)
	txnService := transaction.New(dbRepo)

	logger := setupLogger(conf.Server)

	return &Dependencies{
		Logger:             logger,
		AccountService:     accountService,
		TransactionService: txnService,
	}, nil
}

func setupLogger(cfg httpserver.Config) *zerolog.Logger {
	logLevel, err := zerolog.ParseLevel(cfg.LogLevel)
	if err != nil || len(cfg.LogLevel) == 0 {
		logLevel = zerolog.InfoLevel
	}

	logger := zerolog.New(os.Stdout).With().Timestamp().Logger().Level(logLevel)

	return &logger
}
