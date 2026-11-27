<script setup lang="ts">
// แสดงผลหลังกดปุ่มที่ส่งข้อมูลไป API (สำเร็จ หรือ error)
import type { ApiError } from '~/composables/useApi'

defineProps<{ error: ApiError | null; success?: string | null }>()
</script>

<template>
  <p v-if="error" class="notice notice-error">
    <template v-if="error.status === 404">API ยังไม่พร้อม (404): </template>
    <template v-else-if="error.status">API ตอบ {{ error.status }}: </template>
    {{ error.message }}
  </p>
  <p v-else-if="success" class="notice notice-ok">{{ success }}</p>
</template>
