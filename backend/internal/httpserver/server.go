// Package httpserver ประกอบ Fiber app: middleware, route และ handler 404 กลาง
package httpserver

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// Pinger คือสิ่งที่ ping ฐานข้อมูลได้ (pgxpool.Pool ทำได้)
type Pinger interface {
	Ping(ctx context.Context) error
}

// ErrorBody คือรูปแบบ error กลางที่ทุกเส้น API ตอบกลับ
type ErrorBody struct {
	Error ErrorDetail `json:"error"`
}

// ErrorDetail คือรายละเอียดของ error
type ErrorDetail struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// New สร้าง Fiber app
//
// ตอนนี้มีแค่ GET /healthz ส่วนเส้น /api/* ทั้งหมดยังไม่มี จึงตอบ 404
// เราจะออกแบบเส้นเหล่านี้ใน Bruno (โมดูล 3) แล้วค่อยเขียนจริงในโมดูล 4
func New(frontendOrigin string, db Pinger) *fiber.App {
	app := fiber.New(fiber.Config{
		AppName:               "sda-shop",
		DisableStartupMessage: true,
	})

	app.Use(logger.New())

	// ตั้งค่า CORS แบบแคบไว้ก่อน: อนุญาตเฉพาะ GET/HEAD และ header Accept
	// request ที่ส่ง JSON (Content-Type: application/json), มี Authorization
	// หรือใช้ PATCH/DELETE จะติด CORS preflight ในเบราว์เซอร์
	// ซึ่งเราจะมาแก้กันในโมดูล 4D (ตอนต่อหน้า login จริง)
	app.Use(cors.New(cors.Config{
		AllowOrigins: frontendOrigin,
		AllowMethods: "GET,HEAD",
		AllowHeaders: "Accept",
	}))

	app.Get("/healthz", func(c *fiber.Ctx) error {
		ctx, cancel := context.WithTimeout(c.UserContext(), 3*time.Second)
		defer cancel()
		if db == nil || db.Ping(ctx) != nil {
			return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{"status": "error", "database": "down"})
		}
		return c.JSON(fiber.Map{"status": "ok"})
	})

	// เส้นที่ไม่มี (รวมทุกเส้นใต้ /api) ตอบ 404 ในรูปแบบ error กลาง
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(ErrorBody{Error: ErrorDetail{
			Code:    "not_found",
			Message: "ยังไม่มีเส้นนี้: " + c.Method() + " " + c.Path(),
		}})
	})

	return app
}
