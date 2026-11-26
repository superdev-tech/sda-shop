// คำสั่ง api คือ backend ของร้าน รันด้วย make dev-be
package main

import (
	"context"
	"errors"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/superdev-tech/sda-shop/backend/internal/config"
	"github.com/superdev-tech/sda-shop/backend/internal/database"
	"github.com/superdev-tech/sda-shop/backend/internal/httpserver"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, "backend เริ่มไม่ได้:", err)
		os.Exit(1)
	}
}

func run() error {
	cfg := config.Load()
	if err := cfg.Validate(); err != nil {
		return err
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	pool, err := database.Connect(ctx, cfg.DatabaseURL)
	if err != nil {
		return err
	}
	defer pool.Close()

	app := httpserver.New(cfg.FrontendOrigin, pool)

	go func() {
		<-ctx.Done()
		_ = app.Shutdown()
	}()

	// เปิด port ก่อน แล้วค่อยบอกว่าพร้อม จะได้ไม่ขึ้นว่าพร้อมทั้งที่ port ถูกใช้อยู่
	ln, err := net.Listen("tcp", ":"+cfg.Port)
	if err != nil {
		if errors.Is(err, syscall.EADDRINUSE) {
			return fmt.Errorf("port %s ถูกใช้อยู่แล้ว อาจเปิด make dev ค้างไว้อีกหน้าต่าง ให้ปิดหน้าต่างนั้นด้วย Ctrl+C หรือเปลี่ยน PORT ใน .env", cfg.Port)
		}
		return err
	}

	fmt.Printf("backend พร้อมแล้วที่ http://localhost:%s (ลองเปิด /healthz)\n", cfg.Port)
	return app.Listener(ln)
}
