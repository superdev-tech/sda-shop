<script setup lang="ts">
// รายละเอียดคำสั่งซื้อ (ต้อง login)
//   เส้น 14: GET /api/orders/:id
import { ORDER_STATUS_LABEL, type One, type Order } from '~/composables/types'
import { PH_ORDERS } from '~/composables/placeholders'

const route = useRoute()
const { request } = useApi()
const id = computed(() => route.params.id as string)
const order = useApiData(
  () => request<One<Order>>(`/api/orders/${id.value}`),
  [id],
  () => ({ data: { ...(PH_ORDERS.find((o) => o.id === Number(id.value)) ?? PH_ORDERS[0]!), id: Number(id.value) || 1 } }),
)
</script>

<template>
  <NuxtLink to="/orders" class="muted small">← คำสั่งซื้อทั้งหมด</NuxtLink>
  <h1>คำสั่งซื้อ #{{ id }}</h1>
  <ApiState :pending="order.pending.value" :error="order.error.value" :endpoint="`GET /api/orders/${id}`" skeleton="rows" :placeholder="order.placeholder.value">
    <div v-if="order.data.value" class="card">
      <p>
        <span class="badge">{{ ORDER_STATUS_LABEL[order.data.value.data.status] }}</span>
        <span class="muted small"> · {{ formatDate(order.data.value.data.created_at) }}</span>
      </p>
      <table class="table">
        <thead><tr><th>สินค้า</th><th>ราคา</th><th>จำนวน</th><th>รวม</th></tr></thead>
        <tbody>
          <tr v-for="item in order.data.value.data.items" :key="item.product_id">
            <td>{{ item.product_name }}</td>
            <td>{{ formatBaht(item.price_satang) }}</td>
            <td>{{ item.quantity }}</td>
            <td>{{ formatBaht(item.price_satang * item.quantity) }}</td>
          </tr>
        </tbody>
      </table>
      <p class="row end"><strong>ยอดรวม {{ formatBaht(order.data.value.data.total_satang) }}</strong></p>
    </div>
  </ApiState>
</template>
