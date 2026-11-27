# Bruno collection ของ sda-shop

เปิดด้วย Bruno extension ใน VS Code แล้วเลือก environment `local`

ตอนนี้มีโฟลเดอร์ `00 · ตัวอย่าง` ไว้ลองกด Send ก่อน (ต้องรัน `make dev` อยู่)

- `Health check` เรียก `GET /healthz` ต้องได้ 200
- `Categories (ยังไม่มี)` เรียก `GET /api/categories` ตอนนี้ได้ 404

ทั้งสองไฟล์มี `assert` ไว้ตรวจผลอัตโนมัติ และใช้เป็นแบบเวลาสร้าง request ใหม่ได้

ส่วน API จริงครบ 17 เส้น เราจะเติมกันในโมดูล 3
