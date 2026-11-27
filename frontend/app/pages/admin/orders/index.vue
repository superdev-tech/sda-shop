<script setup lang="ts">
// หลังร้าน: คำสั่งซื้อทั้งหมด (ต้องเป็น admin)
//   เส้น 15: GET   /api/admin/orders
//   เส้น 16: PATCH /api/admin/orders/:id/status   { status }
import type { ApiError } from '~/composables/useApi'
import { ORDER_STATUS_LABEL, type List, type Order, type OrderStatus } from '~/composables/types'

const { request } = useApi()
const orders = useApiData(() => request<List<Order>>('/api/admin/orders'))
const statuses = Object.keys(ORDER_STATUS_LABEL) as OrderStatus[]

const busy = ref(false)
const actionError = ref<ApiError | null>(null)
const saved = ref<string | null>(null)

async function setStatus(id: number, status: OrderStatus) {
  busy.value = true
  actionError.value = null
  saved.value = null
  try {
    await request(`/api/admin/orders/${id}/status`, { method: 'PATCH', body: { status } })
    saved.value = `อัปเดตคำสั่งซื้อ #${id} เป็น "${ORDER_STATUS_LABEL[status]}" แล้ว`
    await orders.refresh()
  } catch (e) {
    actionError.value = e as ApiError
  } finally {
    busy.value = false
  }
}
</script>

<template>
  <div class="row">
    <h1>หลังร้าน · คำสั่งซื้อ</h1>
    <NuxtLink to="/admin/products/new" class="btn btn-ghost">+ เพิ่มสินค้า</NuxtLink>
  </div>
  <ApiState
    :pending="orders.pending.value"
    :error="orders.error.value"
    :empty="!orders.data.value?.data.length"
    endpoint="GET /api/admin/orders"
    empty-text="ยังไม่มีคำสั่งซื้อ"
  >
    <div class="card">
      <table class="table">
        <thead><tr><th>เลขที่</th><th>ลูกค้า</th><th>วันที่</th><th>ยอดรวม</th><th>สถานะ</th></tr></thead>
        <tbody>
          <tr v-for="o in orders.data.value?.data" :key="o.id">
            <td>#{{ o.id }}</td>
            <td>{{ o.customer?.name ?? '-' }}</td>
            <td>{{ formatDate(o.created_at) }}</td>
            <td>{{ formatBaht(o.total_satang) }}</td>
            <td>
              <select
                :value="o.status"
                :disabled="busy"
                @change="setStatus(o.id, ($event.target as HTMLSelectElement).value as OrderStatus)"
              >
                <option v-for="s in statuses" :key="s" :value="s">{{ ORDER_STATUS_LABEL[s] }}</option>
              </select>
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  </ApiState>
  <ActionResult :error="actionError" :success="saved" />
</template>
