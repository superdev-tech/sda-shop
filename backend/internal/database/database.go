// Package database เปิดการเชื่อมต่อไปยัง PostgreSQL บน Supabase
package database

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Connect เปิด connection pool แล้ว ping หนึ่งครั้งเพื่อเช็กว่าต่อติดจริง
func Connect(ctx context.Context, databaseURL string) (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(ctx, databaseURL)
	if err != nil {
		return nil, fmt.Errorf("อ่าน DATABASE_URL ไม่ได้: รูปแบบ connection string ไม่ถูกต้อง ให้คัดลอกใหม่จากหน้า Connect ของ Supabase (แบบ Session pooler): %w", err)
	}

	pingCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	if err := pool.Ping(pingCtx); err != nil {
		pool.Close()
		return nil, fmt.Errorf("%s\nรายละเอียด: %w", Hint(err), err)
	}
	return pool, nil
}

// Hint แปล error ตอนต่อฐานข้อมูลเป็นคำแนะนำภาษาไทยว่าควรแก้ตรงไหน
func Hint(err error) string {
	msg := strings.ToLower(err.Error())
	switch {
	case strings.Contains(msg, "password authentication failed"),
		strings.Contains(msg, "sasl"),
		strings.Contains(msg, "tenant or user not found"):
		return "ต่อฐานข้อมูลไม่ได้: รหัสผ่านหรือชื่อผู้ใช้ไม่ถูกต้อง ให้เช็กว่าเปลี่ยน [YOUR-DATABASE-PASSWORD] ใน DATABASE_URL เป็นรหัสผ่านจริงแล้ว"
	case strings.Contains(msg, "network is unreachable"),
		strings.Contains(msg, "no route to host"),
		strings.Contains(msg, "no such host"),
		strings.Contains(msg, "timeout"),
		strings.Contains(msg, "deadline exceeded"),
		strings.Contains(msg, "connection refused"):
		return "ต่อฐานข้อมูลไม่ได้: เครือข่ายไปไม่ถึงเซิร์ฟเวอร์ ให้เช็กว่าใช้ connection string แบบ Session pooler (host ลงท้าย pooler.supabase.com, port 5432) ไม่ใช่แบบ Direct connection ที่ใช้ได้แค่ IPv6"
	default:
		return "ต่อฐานข้อมูลไม่ได้: ให้เช็ก 3 ข้อ 1) .env ครบไหม 2) ใช้ connection string แบบ Session pooler หรือยัง 3) เปลี่ยน [YOUR-DATABASE-PASSWORD] เป็นรหัสผ่านจริงหรือยัง"
	}
}
