<script setup lang="ts">
// ตะกร้าสินค้า (ต้อง login)
//   เส้น 8:  GET    /api/cart
//   เส้น 10: PATCH  /api/cart/items/:productId   { quantity }
//   เส้น 11: DELETE /api/cart/items/:productId
//   เส้น 12: POST   /api/orders                  สร้างคำสั่งซื้อจากตะกร้า
import type { ApiError } from '~/composables/useApi'
import type { Cart, One, Order } from '~/composables/types'
import { PH_CART } from '~/composables/placeholders'

const { request } = useApi()
const { isLoggedIn } = useAuth()
const cart = useApiData(() => request<One<Cart>>('/api/cart'), [], () => ({ data: PH_CART }))

const busy = ref(false)
const actionError = ref<ApiError | null>(null)

async function act(fn: () => Promise<unknown>) {
  busy.value = true
  actionError.value = null
  try {
    await fn()
    await cart.refresh()
  } catch (e) {
    actionError.value = e as ApiError
  } finally {
    busy.value = false
  }
}

const setQty = (productId: number, quantity: number) =>
  act(() => request(`/api/cart/items/${productId}`, { method: 'PATCH', body: { quantity } }))

const remove = (productId: number) => act(() => request(`/api/cart/items/${productId}`, { method: 'DELETE' }))

async function checkout() {
  busy.value = true
  actionError.value = null
  try {
    const res = await request<One<Order>>('/api/orders', { method: 'POST' })
    navigateTo(`/orders/${res.data.id}`)
  } catch (e) {
    actionError.value = e as ApiError
  } finally {
    busy.value = false
  }
}
</script>

<template>
  <h1>ตะกร้าสินค้า</h1>
  <p v-if="!isLoggedIn" class="notice">ต้อง <NuxtLink to="/login?next=/cart">เข้าสู่ระบบ</NuxtLink> ก่อนจึงจะดูตะกร้าได้</p>
  <ApiState
    :pending="cart.pending.value"
    :error="cart.error.value"
    :empty="!cart.data.value?.data.items.length"
    endpoint="GET /api/cart"
    :placeholder="cart.placeholder.value"
    skeleton="rows"
    empty-text="ตะกร้ายังว่าง"
  >
    <div v-if="cart.data.value" class="card">
      <table class="table">
        <thead>
          <tr><th>สินค้า</th><th>ราคา</th><th>จำนวน</th><th>รวม</th><th /></tr>
        </thead>
        <tbody>
          <tr v-for="item in cart.data.value.data.items" :key="item.product_id">
            <td><NuxtLink :to="`/products/${item.slug}`">{{ item.name }}</NuxtLink></td>
            <td>{{ formatBaht(item.price_satang) }}</td>
            <td>
              <input
                type="number"
                min="1"
                :max="item.stock"
                :value="item.quantity"
                class="qty"
                :disabled="busy"
                @change="setQty(item.product_id, Number(($event.target as HTMLInputElement).value))"
              />
            </td>
            <td>{{ formatBaht(item.price_satang * item.quantity) }}</td>
            <td><button class="link danger" :disabled="busy" @click="remove(item.product_id)">ลบ</button></td>
          </tr>
        </tbody>
      </table>
      <div class="row end">
        <strong>ยอดรวม {{ formatBaht(cart.data.value.data.total_satang) }}</strong>
        <button class="btn" :disabled="busy" @click="checkout">สั่งซื้อ</button>
      </div>
    </div>
  </ApiState>
  <ActionResult :error="actionError" />
</template>
