package config

import (
	"fmt"
	"log"

	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	// Menggunakan Neon Database (Serverless Postgres)
	// Host biasanya format: ep-xyz-123.region.aws.neon.tech
	// Mohon masukkan Connection String lengkap dari Dashboard Neon Anda.
	// Contoh: postgres://alex:AbCdEf123@ep-cool-darkness-123456.us-east-2.aws.neon.tech/neondb?sslmode=require
	dsn := "postgresql://neondb_owner:npg_2dqsHUuxRL1r@ep-rapid-paper-ah8gmx9h-pooler.c-3.us-east-1.aws.neon.tech/neondb?sslmode=require"

	database, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{
		PrepareStmt: true, // Cache statements untuk performa lebih baik
	})
	if err != nil {
		log.Fatal("Gagal koneksi database Neon:", err)
	}

	// Set connection pool
	sqlDB, err := database.DB()
	if err != nil {
		log.Fatal("Gagal mendapatkan instance DB:", err)
	}

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)

	DB = database
	fmt.Println("Database Neon connected successfully!")
}
