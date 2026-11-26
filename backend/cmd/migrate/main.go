// คำสั่ง migrate สร้างตารางจาก db/migrations แล้วใส่ข้อมูลตัวอย่างจาก db/seed
//
//	make db-migrate  -> go run ./cmd/migrate
//	make db-seed     -> go run ./cmd/migrate -seed-only
package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/superdev-tech/sda-shop/backend/internal/config"
	"github.com/superdev-tech/sda-shop/backend/internal/database"
)

func main() {
	dbDir := flag.String("dir", filepath.Join("..", "db"), "โฟลเดอร์ที่มี migrations/ และ seed/")
	seedOnly := flag.Bool("seed-only", false, "ใส่เฉพาะข้อมูลตัวอย่าง ไม่รัน migration")
	flag.Parse()

	if err := run(*dbDir, *seedOnly); err != nil {
		fmt.Fprintln(os.Stderr, "migrate ไม่สำเร็จ:", err)
		os.Exit(1)
	}
}

func run(dbDir string, seedOnly bool) error {
	cfg := config.Load()
	if err := cfg.Validate(); err != nil {
		return err
	}

	ctx := context.Background()
	pool, err := database.Connect(ctx, cfg.DatabaseURL)
	if err != nil {
		return err
	}
	defer pool.Close()

	if !seedOnly {
		if err := migrate(ctx, pool, filepath.Join(dbDir, "migrations")); err != nil {
			return err
		}
	}
	return seed(ctx, pool, filepath.Join(dbDir, "seed", "seed.sql"))
}

func migrate(ctx context.Context, pool *pgxpool.Pool, dir string) error {
	if _, err := pool.Exec(ctx, `CREATE TABLE IF NOT EXISTS schema_migrations (
		version    TEXT PRIMARY KEY,
		applied_at TIMESTAMPTZ NOT NULL DEFAULT now()
	)`); err != nil {
		return fmt.Errorf("สร้างตาราง schema_migrations ไม่ได้: %w", err)
	}

	files, err := filepath.Glob(filepath.Join(dir, "*.sql"))
	if err != nil {
		return err
	}
	if len(files) == 0 {
		return fmt.Errorf("ไม่พบไฟล์ .sql ใน %s ให้รันคำสั่งจากโฟลเดอร์บนสุดของโปรเจกต์ด้วย make db-migrate", dir)
	}
	sort.Strings(files)

	applied := 0
	for _, f := range files {
		version := strings.TrimSuffix(filepath.Base(f), ".sql")

		var exists bool
		if err := pool.QueryRow(ctx, `SELECT EXISTS (SELECT 1 FROM schema_migrations WHERE version = $1)`, version).Scan(&exists); err != nil {
			return err
		}
		if exists {
			fmt.Println("ข้าม (รันไปแล้ว):", version)
			continue
		}

		sqlBytes, err := os.ReadFile(f)
		if err != nil {
			return err
		}

		err = pgx.BeginFunc(ctx, pool, func(tx pgx.Tx) error {
			if _, err := tx.Exec(ctx, string(sqlBytes)); err != nil {
				return err
			}
			_, err := tx.Exec(ctx, `INSERT INTO schema_migrations (version) VALUES ($1)`, version)
			return err
		})
		if err != nil {
			return fmt.Errorf("รัน %s ไม่ผ่าน: %w", version, err)
		}
		fmt.Println("สร้างแล้ว:", version)
		applied++
	}
	fmt.Printf("migration เสร็จ: รันใหม่ %d ไฟล์\n", applied)
	return nil
}

func seed(ctx context.Context, pool *pgxpool.Pool, file string) error {
	sqlBytes, err := os.ReadFile(file)
	if err != nil {
		return fmt.Errorf("อ่านไฟล์ seed ไม่ได้: %w", err)
	}
	if _, err := pool.Exec(ctx, string(sqlBytes)); err != nil {
		return fmt.Errorf("ใส่ข้อมูลตัวอย่างไม่ผ่าน: %w", err)
	}

	var products, categories, admins int
	_ = pool.QueryRow(ctx, `SELECT
		(SELECT count(*) FROM products),
		(SELECT count(*) FROM categories),
		(SELECT count(*) FROM users WHERE role = 'admin')`).Scan(&products, &categories, &admins)
	fmt.Printf("ข้อมูลตัวอย่างพร้อม: สินค้า %d ชิ้น, หมวดหมู่ %d หมวด, admin %d คน\n", products, categories, admins)
	return nil
}
