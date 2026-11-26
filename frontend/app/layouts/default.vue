<script setup lang="ts">
const { isLoggedIn, isAdmin, user, logout } = useAuth()
const config = useRuntimeConfig()

function onLogout() {
  logout()
  navigateTo('/')
}
</script>

<template>
  <div class="shell">
    <header class="topbar">
      <NuxtLink to="/" class="brand">sda-shop</NuxtLink>
      <nav>
        <NuxtLink to="/products">สินค้า</NuxtLink>
        <NuxtLink to="/cart">ตะกร้า</NuxtLink>
        <template v-if="isLoggedIn">
          <NuxtLink to="/orders">คำสั่งซื้อ</NuxtLink>
          <NuxtLink to="/profile">{{ user?.name ?? 'โปรไฟล์' }}</NuxtLink>
          <NuxtLink v-if="isAdmin" to="/admin/orders">หลังร้าน</NuxtLink>
          <button class="link" @click="onLogout">ออกจากระบบ</button>
        </template>
        <template v-else>
          <NuxtLink to="/login">เข้าสู่ระบบ</NuxtLink>
          <NuxtLink to="/register" class="btn btn-small">สมัครสมาชิก</NuxtLink>
        </template>
      </nav>
    </header>
    <main class="container">
      <slot />
    </main>
    <footer class="footer muted small">
      sda-shop · ร้านตัวอย่างของคอร์ส Screen to API · API: <code>{{ config.public.apiBase }}</code>
    </footer>
  </div>
</template>
