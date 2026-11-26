<script setup lang="ts">
// คำสั่งซื้อของฉัน (ต้อง login)
//   เส้น 12: GET /api/orders
import { ORDER_STATUS_LABEL, type List, type Order } from '~/composables/types'

const { request } = useApi()
const orders = useApiData(() => request<List<Order>>('/api/orders'))
</script>

<template>
  <h1>คำสั่งซื้อของฉัน</h1>
  <ApiState
    :pending="orders.pending.value"
    :error="orders.error.value"
    :empty="!orders.data.value?.data.length"
    endpoint="GET /api/orders"
    empty-text="ยังไม่มีคำสั่งซื้อ"
  >
    <div class="card">
      <table class="table">
        <thead><tr><th>เลขที่</th><th>วันที่</th><th>สถานะ</th><th>ยอดรวม</th></tr></thead>
        <tbody>
          <tr v-for="o in orders.data.value?.data" :key="o.id">
            <td><NuxtLink :to="`/orders/${o.id}`">#{{ o.id }}</NuxtLink></td>
            <td>{{ formatDate(o.created_at) }}</td>
            <td><span class="badge">{{ ORDER_STATUS_LABEL[o.status] }}</span></td>
            <td>{{ formatBaht(o.total_satang) }}</td>
          </tr>
        </tbody>
      </table>
    </div>
  </ApiState>
</template>
