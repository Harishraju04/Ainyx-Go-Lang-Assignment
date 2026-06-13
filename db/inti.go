package db

import (
	"context"
	"time"

	"github.com/Harishraju04/Ainyx-Go-Lang-Assignment/internal/logger"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

func InitDB(dbURL string) *pgxpool.Pool {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pool, err := pgxpool.New(ctx, dbURL)
	if err != nil {
		logger.Logger.Fatal("Failed to create database pool", zap.Error(err))
	}

	if err := pool.Ping(ctx); err != nil {
		logger.Logger.Fatal("Failed to ping database", zap.Error(err))
	}

	logger.Logger.Info("Database connection successful")
	return pool
}
