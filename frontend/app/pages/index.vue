<script setup lang="ts">
// หน้าแรก
//   เส้น 1: GET /api/products/featured
//   เส้น 2: GET /api/categories
import type { Category, List, Product } from '~/composables/types'
import { PH_CATEGORIES, PH_PRODUCTS } from '~/composables/placeholders'

const { request } = useApi()

const featured = useApiData(
  () => request<List<Product>>('/api/products/featured'),
  [],
  () => ({ data: PH_PRODUCTS.filter((p) => p.is_featured) }),
)
const categories = useApiData(() => request<List<Category>>('/api/categories'), [], () => ({ data: PH_CATEGORIES }))
</script>

<template>
  <section class="hero">
    <h1>ของดีจากทั่วไทย ส่งถึงบ้าน</h1>
    <p class="muted">กาแฟ ชา ขนม และอุปกรณ์ คัดมาให้แล้ว</p>
    <NuxtLink to="/products" class="btn">ดูสินค้าทั้งหมด</NuxtLink>
  </section>

  <h2>หมวดหมู่</h2>
  <ApiState
    :pending="categories.pending.value"
    :error="categories.error.value"
    :empty="!categories.data.value?.data.length"
    endpoint="GET /api/categories"
    :placeholder="categories.placeholder.value"
    skeleton="chips"
  >
    <div class="chips">
      <NuxtLink
        v-for="c in categories.data.value?.data"
        :key="c.id"
        :to="{ path: '/products', query: { category: c.slug } }"
        class="chip"
      >
        {{ c.name }}
      </NuxtLink>
    </div>
  </ApiState>

  <h2>สินค้าแนะนำ</h2>
  <ApiState
    :pending="featured.pending.value"
    :error="featured.error.value"
    :empty="!featured.data.value?.data.length"
    endpoint="GET /api/products/featured"
    :placeholder="featured.placeholder.value"
    skeleton="cards"
    :count="4"
    empty-text="ยังไม่มีสินค้าแนะนำ"
  >
    <div class="grid">
      <ProductCard v-for="p in featured.data.value?.data" :key="p.id" :product="p" />
    </div>
  </ApiState>
</template>
