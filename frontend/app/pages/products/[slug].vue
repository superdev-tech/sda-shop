<script setup lang="ts">
// รายละเอียดสินค้า
//   เส้น 3: GET /api/products/:slug
//   เส้น 8: POST /api/cart/items  (ต้อง login)
import type { ApiError } from '~/composables/useApi'
import type { One, Product } from '~/composables/types'

const route = useRoute()
const { request } = useApi()
const { isLoggedIn } = useAuth()

const slug = computed(() => route.params.slug as string)
const product = useApiData(() => request<One<Product>>(`/api/products/${slug.value}`), [slug])

const quantity = ref(1)
const adding = ref(false)
const addError = ref<ApiError | null>(null)
const added = ref<string | null>(null)

async function addToCart() {
  if (!isLoggedIn.value) return navigateTo('/login?next=' + encodeURIComponent(route.fullPath))
  const p = product.data.value?.data
  if (!p) return
  adding.value = true
  addError.value = null
  added.value = null
  try {
    await request('/api/cart/items', { method: 'POST', body: { product_id: p.id, quantity: quantity.value } })
    added.value = `เพิ่ม ${p.name} ${quantity.value} ชิ้นลงตะกร้าแล้ว`
  } catch (e) {
    addError.value = e as ApiError
  } finally {
    adding.value = false
  }
}
</script>

<template>
  <NuxtLink to="/products" class="muted small">← กลับไปหน้าสินค้า</NuxtLink>
  <ApiState :pending="product.pending.value" :error="product.error.value" :endpoint="`GET /api/products/${slug}`">
    <div v-if="product.data.value" class="detail">
      <img :src="product.data.value.data.image_url || '/images/placeholder.svg'" :alt="product.data.value.data.name" />
      <div>
        <span class="muted">{{ product.data.value.data.category?.name }}</span>
        <h1>{{ product.data.value.data.name }}</h1>
        <p class="price">{{ formatBaht(product.data.value.data.price_satang) }}</p>
        <p>{{ product.data.value.data.description }}</p>
        <p class="muted small">
          {{ product.data.value.data.stock > 0 ? `เหลือ ${product.data.value.data.stock} ชิ้น` : 'สินค้าหมด' }}
        </p>
        <div class="row">
          <input v-model.number="quantity" type="number" min="1" :max="product.data.value.data.stock" class="qty" />
          <button class="btn" :disabled="adding || product.data.value.data.stock === 0" @click="addToCart">
            {{ adding ? 'กำลังเพิ่ม…' : 'ใส่ตะกร้า' }}
          </button>
        </div>
        <ActionResult :error="addError" :success="added" />
      </div>
    </div>
  </ApiState>
</template>
