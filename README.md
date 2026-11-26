# sda-shop

โค้ดประกอบคอร์ส **ออกแบบและสร้าง API ด้วย Go Fiber (Screen to API)** ของ Superdev Academy

repo นี้คือร้านค้าออนไลน์ **sda-shop** ที่หน้าเว็บ (frontend) พร้อมใช้แล้ว แต่ backend ยังไม่มี API สักเส้น
ตลอดคอร์สเราจะแตกหน้าจอออกมาเป็น API 17 เส้น ออกแบบใน Bruno แล้วเขียนด้วย Go จนร้านใช้งานได้จริง

| โฟลเดอร์ / ไฟล์ | คืออะไร |
|---|---|
| `frontend/` | หน้าเว็บร้านค้า เขียนด้วย Nuxt พร้อมใช้แล้ว ไม่ต้องแก้ |
| `backend/` | ส่วนที่คุณจะเขียนตลอดทั้งคอร์ส (Go + Fiber) |
| `db/migrations/` | ไฟล์ SQL ที่ใช้สร้างตารางในฐานข้อมูล |
| `db/seed/` | ข้อมูลตัวอย่าง: หมวดหมู่ สินค้า และ admin 1 คน |
| `bruno/` | collection สำหรับออกแบบและทดสอบ API ตอนนี้มีตัวอย่าง 2 request ส่วนเส้นจริงเติมในโมดูล 3 |
| `Makefile` | รวมคำสั่งที่ใช้บ่อย พิมพ์ `make help` เพื่อดูทั้งหมด |

## เริ่มต้นใช้งาน (ตามโมดูล 2)

### 1. ติดตั้งเครื่องมือ

**macOS** (ใช้ Homebrew จาก [brew.sh](https://brew.sh) ซึ่งจะลง Command Line Tools ที่มี Git และ Make มาให้ด้วย)

```bash
brew install go
brew install --cask visual-studio-code
```

Node ติดตั้งผ่าน [nvm](https://github.com/nvm-sh/nvm) ไม่ใช่ `brew install node` เพื่อให้สลับเวอร์ชันตาม repo ได้

```bash
curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.41.0/install.sh | bash
# ปิด terminal แล้วเปิดใหม่ ก่อนพิมพ์คำสั่งต่อไป
nvm install --lts
nvm alias default 'lts/*'
```

**Windows** (เปิด PowerShell แล้วใช้ winget จากนั้นปิดแล้วเปิด terminal ใหม่)

```powershell
winget install --id Git.Git
winget install --id GoLang.Go
winget install --id CoreyButler.NVMforWindows
winget install --id Microsoft.VisualStudioCode
winget install --id ezwinports.make
```

ปิดแล้วเปิด PowerShell ใหม่ จากนั้นติดตั้ง Node LTS ผ่าน [nvm-windows](https://github.com/coreybutler/nvm-windows) (ถ้าเคยลง Node แบบอื่นไว้ ให้ถอนออกก่อน)

```powershell
nvm install lts
nvm use lts
```

เช็กว่าพร้อม: `go version` · `node -v` · `npm -v` · `git --version` · `make --version`

> repo นี้ lock Node ไว้ที่ **26** (LTS) ในไฟล์ [`.nvmrc`](.nvmrc) พอ clone แล้วให้พิมพ์ `nvm use` ในโฟลเดอร์ repo เพื่อสลับให้ตรง (Windows พิมพ์ `nvm use 26` ตามเลขในไฟล์)

ใน VS Code ให้ติดตั้ง extension **Go**, **Bruno** และแนะนำ **[Git Graph](https://marketplace.visualstudio.com/items?itemName=mhutchie.git-graph)** ไว้ดูประวัติ commit

### 2. Fork แล้ว clone

ยังไม่มีบัญชี GitHub ให้สมัครที่ [github.com](https://github.com) ก่อน

เปิด [github.com/superdev-tech/sda-shop](https://github.com/superdev-tech/sda-shop) กด **Fork** แล้ว clone จาก fork ของคุณเอง (ไม่ใช่ของต้นทาง)

```bash
git clone https://github.com/<บัญชีของคุณ>/sda-shop.git
cd sda-shop
nvm use        # อ่านเวอร์ชันจาก .nvmrc (ถ้ายังไม่มีเวอร์ชันนี้ พิมพ์ nvm install)
code .
```

### 3. ติดตั้ง dependency

```bash
make setup
```

### 4. ตั้งค่า Supabase และไฟล์ `.env`

1. สร้างโปรเจกต์บน [supabase.com](https://supabase.com) ชื่อ `sda-shop-dev` และ **จดรหัสผ่านฐานข้อมูลไว้**
   Region ค่าเริ่มต้นเป็น Auto ให้เลือกเองให้ใกล้ไทย เช่น Singapore
2. กดปุ่ม **Connect** ด้านบนของโปรเจกต์ หน้าต่างจะเปิดมาที่ Direct connection ให้เปลี่ยน **Connection Method** เป็น **Session pooler** ส่วน **Type** ใช้ `URI` แล้วคัดลอก connection string
   (host ลงท้าย `pooler.supabase.com`, port `5432`, username เป็น `postgres.<รหัสโปรเจกต์>` — แบบ Direct connection ใช้ได้แค่ IPv6 เน็ตบ้านหลายที่จะต่อไม่ติด)
3. คัดลอก `.env.example` เป็น `.env` แล้ววางค่าลงใน `DATABASE_URL` อย่าลืมเปลี่ยน `[YOUR-DATABASE-PASSWORD]` เป็นรหัสผ่านจริง

```bash
cp .env.example .env      # Windows: copy .env.example .env
```

> `.env` คือกุญแจฐานข้อมูลของคุณ อยู่ใน `.gitignore` แล้ว **ห้าม commit เด็ดขาด**

### 5. สร้างตารางและข้อมูลตัวอย่าง

```bash
make db-migrate
```

เปิด Table Editor ใน Supabase จะเห็น 7 ตาราง `products` `categories` `users` มีข้อมูลตัวอย่าง (admin 1 คน) ส่วน `cart_items` `orders` `order_items` ยังว่าง ซึ่งเป็นเรื่องปกติ

### 6. รันทั้งระบบ

```bash
make dev
```

- หน้าเว็บ: http://localhost:3000
- backend: http://localhost:8080/healthz ต้องตอบ `{"status":"ok"}` (เส้นนี้เช็กการต่อฐานข้อมูลด้วย ถ้าต่อไม่ได้จะตอบ 503)
- Bruno: เปิดแถบ Bruno ใน VS Code เลือก environment `local` แล้วกด Send ในโฟลเดอร์ `00 · ตัวอย่าง`

**หน้าเว็บจะขึ้น แต่ยังไม่มีสินค้า** — ตั้งใจให้เป็นแบบนี้
เปิด DevTools > Network จะเห็น request ไปที่ `/api/...` ตอบกลับ `404`
เพราะ frontend เรียก API แล้ว แต่ backend ยังไม่มีเส้นนั้น เราจะเริ่มออกแบบกันในโมดูล 3

## error ที่พบบ่อย

| อาการ | วิธีแก้ |
|---|---|
| `ไม่พบ DATABASE_URL` | ยังไม่ได้สร้าง `.env` หรือใส่ค่าไม่ครบ |
| `network is unreachable` / timeout | ใช้ connection string แบบ Session pooler |
| `password authentication failed` | เปลี่ยน `[YOUR-DATABASE-PASSWORD]` เป็นรหัสผ่านจริง |
| `address already in use` | มีโปรแกรมอื่นใช้ port 3000 หรือ 8080 อยู่ ปิดโปรแกรมนั้น หรือเปลี่ยน `PORT` ใน `.env` |
| `nuxt: command not found` | ยังไม่ได้รัน `make setup` |

## คำสั่งทั้งหมด

| คำสั่ง | ทำอะไร |
|---|---|
| `make help` | แสดงรายการคำสั่ง |
| `make setup` | ติดตั้ง dependency ของ frontend และ backend |
| `make dev` | รัน frontend และ backend พร้อมกัน |
| `make dev-fe` | รันเฉพาะ frontend ที่ port 3000 |
| `make dev-be` | รันเฉพาะ backend ที่ port 8080 |
| `make db-migrate` | สร้างตาราง (รันเฉพาะไฟล์ที่ยังไม่เคยรัน) แล้วใส่ข้อมูลตัวอย่าง |
| `make db-seed` | ใส่ข้อมูลตัวอย่างซ้ำ รันกี่ครั้งก็ได้ |
| `make test` | รัน test ของ backend |
| `make fmt` | จัดรูปแบบโค้ด Go |
