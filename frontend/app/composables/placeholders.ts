// ข้อมูลตัวอย่าง (placeholder) ไว้แสดงแทนตอนที่ backend ยังไม่มีเส้นนั้น (404) หรือยังไม่ได้รัน
// รูปแบบตรงกับ types.ts และแท็บ Docs ใน Bruno ส่วนค่ามาจาก db/seed/seed.sql
// ทุกหน้าจะแปะป้าย "ข้อมูลตัวอย่าง" ไว้ จะได้ไม่สับสนกับข้อมูลจริง
import type { Cart, Category, Order, Paged, Product } from './types'
import type { AuthUser } from './useAuth'

export const PH_CATEGORIES: Category[] = [
  { id: 1, slug: 'coffee', name: 'กาแฟ' },
  { id: 2, slug: 'tea', name: 'ชา' },
  { id: 3, slug: 'snacks', name: 'ขนม' },
  { id: 4, slug: 'accessories', name: 'อุปกรณ์' },
]

const cat = (slug: string) => PH_CATEGORIES.find((c) => c.slug === slug)!

// [หมวด, slug, ชื่อ, คำอธิบาย, ราคา (สตางค์), สต็อก, รูป, แนะนำ]
const ROWS: [string, string, string, string, number, number, string, boolean][] = [
  ['coffee', 'doi-chang-arabica', 'กาแฟอาราบิก้าดอยช้าง 250 กรัม', 'คั่วกลาง หอมถั่วและช็อกโกแลต', 32000, 40, 'coffee', true],
  ['coffee', 'chiang-rai-espresso', 'เอสเพรสโซ่เบลนด์เชียงราย 500 กรัม', 'คั่วเข้ม เหมาะกับนมสด', 45000, 25, 'coffee', false],
  ['coffee', 'cold-brew-bottle', 'โคลด์บรูว์ขวด 350 มล.', 'สกัดเย็น 18 ชั่วโมง', 8900, 60, 'coffee', true],
  ['tea', 'thai-milk-tea-mix', 'ชาไทยปรุงสำเร็จ 400 กรัม', 'สูตรเข้มข้นแบบร้านดัง', 15900, 80, 'tea', true],
  ['tea', 'jasmine-green-tea', 'ชาเขียวมะลิ 50 ซอง', 'ชาเขียวกลิ่นมะลิอ่อน ๆ', 12900, 35, 'tea', false],
  ['tea', 'butterfly-pea-tea', 'ชาดอกอัญชัน 30 กรัม', 'ดอกอัญชันอบแห้ง ใส่มะนาวเปลี่ยนสี', 9900, 50, 'tea', false],
  ['snacks', 'durian-chips', 'ทุเรียนทอดกรอบ 100 กรัม', 'หมอนทองจันทบุรี ทอดกรอบ', 19900, 20, 'snack', true],
  ['snacks', 'mango-sticky-bar', 'ข้าวเหนียวมะม่วงอบแห้ง', 'ขนมทานเล่นรสมะม่วงน้ำดอกไม้', 5900, 100, 'snack', false],
  ['snacks', 'cashew-honey', 'เม็ดมะม่วงหิมพานต์อบน้ำผึ้ง 200 กรัม', 'เม็ดใหญ่ จากระนอง', 22900, 30, 'snack', false],
  ['accessories', 'drip-kettle', 'กาดริปคอห่าน 600 มล.', 'สแตนเลส คุมสายน้ำได้แม่น', 69000, 10, 'accessory', true],
  ['accessories', 'ceramic-mug', 'แก้วเซรามิกลายช้าง', 'ทำมือจากลำปาง ความจุ 300 มล.', 25000, 0, 'accessory', false],
  ['accessories', 'paper-filter', 'กระดาษกรองกาแฟ 100 แผ่น', 'ขนาด 02 ใช้กับดริปเปอร์ทรงกรวย', 7900, 120, 'accessory', false],
]

export const PH_PRODUCTS: Product[] = ROWS.map(([c, slug, name, description, price_satang, stock, img, is_featured], i) => ({
  id: i + 1,
  slug,
  name,
  description,
  price_satang,
  stock,
  image_url: `/images/${img}.svg`,
  is_featured,
  category: cat(c),
}))

// เส้น 3: กรองและแบ่งหน้าแบบเดียวกับที่ backend จะทำ
export function phProductPage(page: number, limit: number, category?: string, q?: string): Paged<Product> {
  const all = PH_PRODUCTS.filter(
    (p) => (!category || p.category.slug === category) && (!q || p.name.includes(q) || p.description.includes(q)),
  )
  return { data: all.slice((page - 1) * limit, page * limit), meta: { page, limit, total: all.length } }
}

export const PH_ME: AuthUser = { id: 2, email: 'somchai@example.com', name: 'สมชาย ใจดี', role: 'customer' }

const p = (slug: string) => PH_PRODUCTS.find((x) => x.slug === slug)!
const line = (slug: string, quantity: number) => {
  const x = p(slug)
  return { product_id: x.id, slug: x.slug, name: x.name, image_url: x.image_url, price_satang: x.price_satang, quantity, stock: x.stock }
}
const cartItems = [line('doi-chang-arabica', 2), line('drip-kettle', 1)]
export const PH_CART: Cart = {
  items: cartItems,
  total_satang: cartItems.reduce((s, i) => s + i.price_satang * i.quantity, 0),
}

function order(id: number, status: Order['status'], created_at: string, lines: [string, number][], customer?: Order['customer']): Order {
  const items = lines.map(([slug, quantity]) => ({ product_id: p(slug).id, product_name: p(slug).name, price_satang: p(slug).price_satang, quantity }))
  return { id, status, created_at, items, total_satang: items.reduce((s, i) => s + i.price_satang * i.quantity, 0), ...(customer ? { customer } : {}) }
}

// คำสั่งซื้อของฉัน (ใหม่สุดอยู่บน)
export const PH_ORDERS: Order[] = [
  order(2, 'pending', '2026-11-27T10:42:00+07:00', [['doi-chang-arabica', 2], ['cold-brew-bottle', 1]]),
  order(1, 'completed', '2026-11-20T15:05:00+07:00', [['thai-milk-tea-mix', 1]]),
]

// หลังร้าน: เห็นของทุกคน พร้อมบอกว่าใครสั่ง
const somchai = { id: 2, name: 'สมชาย ใจดี', email: 'somchai@example.com' }
const suda = { id: 3, name: 'สุดา รักดี', email: 'suda@example.com' }
export const PH_ADMIN_ORDERS: Order[] = [
  order(3, 'paid', '2026-11-27T11:10:00+07:00', [['durian-chips', 3]], suda),
  { ...PH_ORDERS[0]!, customer: somchai },
  { ...PH_ORDERS[1]!, customer: somchai },
]
