# Makefile ของ sda-shop
# รองรับ GNU Make 3.81 (macOS) และ make บน Windows (cmd.exe)
# จึงใช้เฉพาะรูปแบบ cd โฟลเดอร์ && คำสั่ง และ @echo ที่ไม่มีเครื่องหมายคำพูดหรือวงเล็บ
# (cmd.exe จะพิมพ์เครื่องหมายคำพูดออกมาด้วย ส่วน sh จะ error ถ้ามีวงเล็บ)

.PHONY: help setup dev dev-fe dev-be db-migrate db-seed test fmt

help:
	@echo คำสั่งที่ใช้ได้ - พิมพ์ make ตามด้วยชื่อคำสั่ง
	@echo help       - แสดงรายการคำสั่งทั้งหมด
	@echo setup      - ติดตั้ง dependency ของ frontend และ backend
	@echo dev        - รัน frontend และ backend พร้อมกัน
	@echo dev-fe     - รันเฉพาะ frontend Nuxt ที่ http://localhost:3000
	@echo dev-be     - รันเฉพาะ backend Go ที่ http://localhost:8080
	@echo db-migrate - สร้างตารางในฐานข้อมูล แล้วใส่ข้อมูลตัวอย่าง
	@echo db-seed    - ใส่ข้อมูลตัวอย่างซ้ำ รันกี่ครั้งก็ได้
	@echo test       - รัน test ของ backend
	@echo fmt        - จัดรูปแบบโค้ด Go

setup:
	cd frontend && npm install
	cd backend && go mod download

dev:
	$(MAKE) -j 2 dev-be dev-fe

dev-fe:
	cd frontend && npm run dev

dev-be:
	cd backend && go run ./cmd/api

db-migrate:
	cd backend && go run ./cmd/migrate

db-seed:
	cd backend && go run ./cmd/migrate -seed-only

test:
	cd backend && go test ./...

fmt:
	cd backend && go fmt ./...
