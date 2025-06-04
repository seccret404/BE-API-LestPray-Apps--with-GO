package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitPostgre(){
	dsn := os.Getenv("DATABASE_URL")

	if dsn == "" {
		log.Fatal("Tidak menemukan variabel di env")
	}

	db,err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil{
		log.Fatal("Gagal connect ke supabase", err)
	}

	//set jumlah koneksi
	sqlDB, err := db.DB()
	if err != nil{
		log.Fatal("gagal set koneksi ke database", err)
	}
	sqlDB.SetMaxIdleConns(10)//10 koneksi
	sqlDB.SetMaxOpenConns(100)//100 koneksi pertama
	sqlDB.SetConnMaxLifetime(time.Hour)//set time koneksi

	DB = db

	fmt.Println("Berhasil terhubung ke database")
}