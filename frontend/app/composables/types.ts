// รูปแบบข้อมูลที่หน้าเว็บคาดหวังจาก API
// เงินทุกช่องเป็นสตางค์ (number) หน้าเว็บจัดรูปแบบเป็นบาทเอง

export interface Category {
  id: number
  slug: string
  name: string
}

export interface Product {
  id: number
  slug: string
  name: string
  description: string
  price_satang: number
  stock: number
  image_url: string
  is_featured: boolean
  category: Category
}

export interface Paged<T> {
  data: T[]
  meta: { page: number; limit: number; total: number }
}

export interface One<T> {
  data: T
}

export interface List<T> {
  data: T[]
}

export interface CartItem {
  product_id: number
  slug: string
  name: string
  image_url: string
  price_satang: number
  quantity: number
  stock: number
}

export interface Cart {
  items: CartItem[]
  total_satang: number
}

export type OrderStatus = 'pending' | 'paid' | 'shipped' | 'completed' | 'cancelled'

export interface OrderItem {
  product_id: number
  product_name: string
  price_satang: number
  quantity: number
}

export interface Order {
  id: number
  status: OrderStatus
  total_satang: number
  created_at: string
  items: OrderItem[]
  customer?: { id: number; name: string; email: string }
}

export const ORDER_STATUS_LABEL: Record<OrderStatus, string> = {
  pending: 'รอชำระเงิน',
  paid: 'ชำระแล้ว',
  shipped: 'จัดส่งแล้ว',
  completed: 'สำเร็จ',
  cancelled: 'ยกเลิก',
}
