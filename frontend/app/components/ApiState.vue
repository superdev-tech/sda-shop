<script setup lang="ts">
// แสดงสถานะการโหลดข้อมูลจาก API: กำลังโหลด / error / ว่าง
import type { ApiError } from '~/composables/useApi'

const props = defineProps<{
  pending: boolean
  error: ApiError | null
  empty?: boolean
  endpoint: string
  emptyText?: string
}>()
</script>

<template>
  <div v-if="props.pending" class="state">กำลังโหลด…</div>
  <div v-else-if="props.error" class="state state-error">
    <strong v-if="props.error.status === 404">ยังไม่มีข้อมูล — API ยังไม่พร้อม (404)</strong>
    <strong v-else-if="props.error.status === 0">ต่อ backend ไม่ได้</strong>
    <strong v-else>API ตอบกลับ {{ props.error.status }}</strong>
    <p>{{ props.error.message }}</p>
    <code>{{ props.endpoint }}</code>
  </div>
  <div v-else-if="props.empty" class="state">{{ props.emptyText ?? 'ยังไม่มีข้อมูล' }}</div>
  <slot v-else />
</template>
