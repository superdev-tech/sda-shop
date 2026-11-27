<script setup lang="ts">
// โปรไฟล์ของฉัน
//   เส้น 7: GET /api/me  (ต้อง login)
import type { AuthUser } from '~/composables/useAuth'
import type { One } from '~/composables/types'

const { request } = useApi()
const me = useApiData(() => request<One<AuthUser>>('/api/me'))
</script>

<template>
  <h1>โปรไฟล์ของฉัน</h1>
  <ApiState :pending="me.pending.value" :error="me.error.value" endpoint="GET /api/me">
    <div v-if="me.data.value" class="card">
      <p><span class="muted">ชื่อ</span> {{ me.data.value.data.name }}</p>
      <p><span class="muted">อีเมล</span> {{ me.data.value.data.email }}</p>
      <p><span class="muted">บทบาท</span> {{ me.data.value.data.role === 'admin' ? 'ผู้ดูแลร้าน' : 'ลูกค้า' }}</p>
    </div>
  </ApiState>
</template>
