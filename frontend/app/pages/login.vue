<script setup lang="ts">
// เข้าสู่ระบบ
//   เส้น 6: POST /api/auth/login  -> { data: { token, user } }
import type { ApiError } from '~/composables/useApi'
import type { AuthUser } from '~/composables/useAuth'
import type { One } from '~/composables/types'

const route = useRoute()
const { request } = useApi()
const { setSession } = useAuth()
const form = reactive({ email: '', password: '' })
const loading = ref(false)
const error = ref<ApiError | null>(null)

async function submit() {
  loading.value = true
  error.value = null
  try {
    const res = await request<One<{ token: string; user: AuthUser }>>('/api/auth/login', {
      method: 'POST',
      body: { ...form },
    })
    setSession(res.data.token, res.data.user)
    navigateTo((route.query.next as string) || '/')
  } catch (e) {
    error.value = e as ApiError
  } finally {
    loading.value = false
  }
}
</script>

<template>
  <form class="card form" @submit.prevent="submit">
    <h1>เข้าสู่ระบบ</h1>
    <label>อีเมล <input v-model="form.email" type="email" required /></label>
    <label>รหัสผ่าน <input v-model="form.password" type="password" required /></label>
    <button class="btn" :disabled="loading">{{ loading ? 'กำลังเข้าสู่ระบบ…' : 'เข้าสู่ระบบ' }}</button>
    <ActionResult :error="error" />
    <p class="muted small">ยังไม่มีบัญชี? <NuxtLink to="/register">สมัครสมาชิก</NuxtLink></p>
  </form>
</template>
