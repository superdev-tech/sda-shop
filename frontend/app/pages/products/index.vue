<script setup lang="ts">
// รายการสินค้า
//   เส้น 2: GET /api/categories (ตัวกรองหมวดหมู่)
//   เส้น 3: GET /api/products?page=&limit=&category=&q=
import type { Category, List, Paged, Product } from '~/composables/types'

const route = useRoute()
const router = useRouter()
const { request } = useApi()

const limit = 8
const page = computed(() => Number(route.query.page ?? 1) || 1)
const category = computed(() => (route.query.category as string) || undefined)
const q = computed(() => (route.query.q as string) || undefined)
const search = ref(q.value ?? '')

const categories = useApiData(() => request<List<Category>>('/api/categories'))
const products = useApiData(
  () =>
    request<Paged<Product>>('/api/products', {
      query: { page: page.value, limit, category: category.value, q: q.value },
    }),
  [page, category, q],
)

const totalPages = computed(() => {
  const meta = products.data.value?.meta
  return meta ? Math.max(1, Math.ceil(meta.total / meta.limit)) : 1
})

function go(query: Record<string, string | number | undefined>) {
  router.push({ path: '/products', query: { ...route.query, ...query } })
}
</script>

<template>
  <h1>สินค้าทั้งหมด</h1>

  <form class="toolbar" @submit.prevent="go({ q: search || undefined, page: 1 })">
    <input v-model="search" type="search" placeholder="ค้นหาสินค้า…" />
    <button class="btn" type="submit">ค้นหา</button>
  </form>

  <div v-if="categories.data.value?.data.length" class="chips">
    <button class="chip" :class="{ active: !category }" @click="go({ category: undefined, page: 1 })">
      ทั้งหมด
    </button>
    <button
      v-for="c in categories.data.value.data"
      :key="c.id"
      class="chip"
      :class="{ active: category === c.slug }"
      @click="go({ category: c.slug, page: 1 })"
    >
      {{ c.name }}
    </button>
  </div>

  <ApiState
    :pending="products.pending.value"
    :error="products.error.value"
    :empty="!products.data.value?.data.length"
    endpoint="GET /api/products?page=&limit=&category=&q="
    empty-text="ไม่พบสินค้าที่ค้นหา"
  >
    <div class="grid">
      <ProductCard v-for="p in products.data.value?.data" :key="p.id" :product="p" />
    </div>
    <div class="pager">
      <button class="btn btn-ghost" :disabled="page <= 1" @click="go({ page: page - 1 })">ก่อนหน้า</button>
      <span class="muted">หน้า {{ page }} / {{ totalPages }}</span>
      <button class="btn btn-ghost" :disabled="page >= totalPages" @click="go({ page: page + 1 })">ถัดไป</button>
    </div>
  </ApiState>
</template>
