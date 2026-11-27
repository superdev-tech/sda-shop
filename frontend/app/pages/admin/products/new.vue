<script setup lang="ts">
// หลังร้าน: เพิ่มสินค้าใหม่ (ต้องเป็น admin)
//   เส้น 2:  GET  /api/categories
//   เส้น 17: POST /api/admin/products
import type { ApiError } from '~/composables/useApi'
import type { Category, List, One, Product } from '~/composables/types'

const { request } = useApi()
const categories = useApiData(() => request<List<Category>>('/api/categories'))

const form = reactive({
  name: '',
  slug: '',
  description: '',
  price_baht: 0,
  stock: 0,
  category_id: 0,
  image_url: '/images/placeholder.svg',
  is_featured: false,
})
const loading = ref(false)
const error = ref<ApiError | null>(null)
const success = ref<string | null>(null)

async function submit() {
  loading.value = true
  error.value = null
  success.value = null
  try {
    const { price_baht, ...rest } = form
    const res = await request<One<Product>>('/api/admin/products', {
      method: 'POST',
      // ส่งเงินเป็นสตางค์เสมอ
      body: { ...rest, price_satang: Math.round(price_baht * 100) },
    })
    success.value = `เพิ่มสินค้า ${res.data.name} แล้ว`
  } catch (e) {
    error.value = e as ApiError
  } finally {
    loading.value = false
  }
}
</script>

<template>
  <NuxtLink to="/admin/orders" class="muted small">← หลังร้าน</NuxtLink>
  <form class="card form" @submit.prevent="submit">
    <h1>เพิ่มสินค้า</h1>
    <label>ชื่อสินค้า <input v-model="form.name" required /></label>
    <label>slug (ใช้ใน URL) <input v-model="form.slug" pattern="[a-z0-9-]+" required /></label>
    <label>รายละเอียด <textarea v-model="form.description" rows="3" /></label>
    <div class="row">
      <label>ราคา (บาท) <input v-model.number="form.price_baht" type="number" min="0" step="0.01" required /></label>
      <label>จำนวนในคลัง <input v-model.number="form.stock" type="number" min="0" required /></label>
    </div>
    <label>
      หมวดหมู่
      <select v-model.number="form.category_id" required>
        <option :value="0" disabled>เลือกหมวดหมู่</option>
        <option v-for="c in categories.data.value?.data" :key="c.id" :value="c.id">{{ c.name }}</option>
      </select>
      <span v-if="categories.error.value" class="muted small">โหลดหมวดหมู่ไม่ได้ ({{ categories.error.value.status || 'ต่อไม่ได้' }})</span>
    </label>
    <label>รูปภาพ (URL) <input v-model="form.image_url" /></label>
    <label class="check"><input v-model="form.is_featured" type="checkbox" /> แสดงเป็นสินค้าแนะนำ</label>
    <button class="btn" :disabled="loading">{{ loading ? 'กำลังบันทึก…' : 'บันทึกสินค้า' }}</button>
    <ActionResult :error="error" :success="success" />
  </form>
</template>
