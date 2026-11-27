<script setup lang="ts">
// แสดงผลหลังกดปุ่มที่ส่งข้อมูลไป API (สำเร็จ หรือ error)
import type { ApiError } from '~/composables/useApi'

defineProps<{ error: ApiError | null; success?: string | null }>()
</script>

<template>
  <!-- 404 = backend ยังไม่มีเส้นนี้ เป็นเรื่องปกติก่อนโมดูล 5 จึงไม่ใช้สีแดง -->
  <p v-if="error && error.status === 404" class="notice notice-wait">รอ API · ยังไม่มีเส้นนี้ (404): {{ error.message }}</p>
  <p v-else-if="error" class="notice notice-error">
    <template v-if="error.status">API ตอบ {{ error.status }}: </template>
    {{ error.message }}
  </p>
  <p v-else-if="success" class="notice notice-ok">{{ success }}</p>
</template>
