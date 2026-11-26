<script setup lang="ts">
import type { Product } from '~/composables/types'

defineProps<{ product: Product }>()
</script>

<template>
  <NuxtLink :to="`/products/${product.slug}`" class="card product-card">
    <img :src="product.image_url || '/images/placeholder.svg'" :alt="product.name" />
    <div class="product-body">
      <span class="muted small">{{ product.category?.name }}</span>
      <h3>{{ product.name }}</h3>
      <div class="row">
        <strong>{{ formatBaht(product.price_satang) }}</strong>
        <span v-if="product.stock === 0" class="badge badge-off">สินค้าหมด</span>
      </div>
    </div>
  </NuxtLink>
</template>
