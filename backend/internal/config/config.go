// Package config อ่านค่าตั้งค่าจากไฟล์ .env และ environment variable
package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

// Config คือค่าตั้งค่าทั้งหมดที่ backend ต้องใช้
type Config struct {
	DatabaseURL    string
	Port           string
	FrontendOrigin string
}

// ErrMissingDatabaseURL เกิดเมื่อยังไม่ได้ใส่ DATABASE_URL ใน .env
var ErrMissingDatabaseURL = errors.New("ไม่พบ DATABASE_URL: ไฟล์ .env ไม่ครบ หรือยังไม่ได้สร้าง ให้คัดลอก .env.example เป็น .env ที่โฟลเดอร์บนสุดของโปรเจกต์ แล้วใส่ connection string แบบ Session pooler จาก Supabase")

// Load อ่าน .env ที่โฟลเดอร์บนสุด (../.env) ก่อน แล้วค่อยอ่าน .env ในโฟลเดอร์ปัจจุบัน
// ไฟล์ไหนไม่มีก็ข้ามไป ค่าที่ตั้งไว้ใน environment อยู่แล้วจะไม่ถูกทับ
func Load() Config {
	_ = godotenv.Load("../.env")
	_ = godotenv.Load(".env")

	return Config{
		DatabaseURL:    os.Getenv("DATABASE_URL"),
		Port:           getenv("PORT", "8080"),
		FrontendOrigin: getenv("FRONTEND_ORIGIN", "http://localhost:3000"),
	}
}

// Validate ตรวจว่าค่าที่จำเป็นครบหรือยัง
func (c Config) Validate() error {
	if c.DatabaseURL == "" {
		return ErrMissingDatabaseURL
	}
	return nil
}

func getenv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
