<script setup lang="ts">
// สมัครสมาชิก
//   เส้น 5: POST /api/auth/register
import type { ApiError } from '~/composables/useApi'

const { request } = useApi()
const form = reactive({ name: '', email: '', password: '' })
const loading = ref(false)
const error = ref<ApiError | null>(null)
const success = ref<string | null>(null)

async function submit() {
  loading.value = true
  error.value = null
  success.value = null
  try {
    await request('/api/auth/register', { method: 'POST', body: { ...form } })
    success.value = 'สมัครสมาชิกสำเร็จ เข้าสู่ระบบได้เลย'
    setTimeout(() => navigateTo('/login'), 800)
  } catch (e) {
    error.value = e as ApiError
  } finally {
    loading.value = false
  }
}
</script>

<template>
  <form class="card form" @submit.prevent="submit">
    <h1>สมัครสมาชิก</h1>
    <label>ชื่อ <input v-model="form.name" required /></label>
    <label>อีเมล <input v-model="form.email" type="email" required /></label>
    <label>รหัสผ่าน <input v-model="form.password" type="password" minlength="8" required /></label>
    <button class="btn" :disabled="loading">{{ loading ? 'กำลังสมัคร…' : 'สมัครสมาชิก' }}</button>
    <ActionResult :error="error" :success="success" />
    <p class="muted small">มีบัญชีแล้ว? <NuxtLink to="/login">เข้าสู่ระบบ</NuxtLink></p>
  </form>
</template>
