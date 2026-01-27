package config

import (
	"fmt"
	"log"

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

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal koneksi database Neon:", err)
	}

	DB = database
	fmt.Println("Database Neon connected successfully!")
}
