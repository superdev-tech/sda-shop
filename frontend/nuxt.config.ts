// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  compatibilityDate: '2025-07-15',
  devtools: { enabled: false },

  // ปิด SSR ให้หน้าเว็บเรียก API จากเบราว์เซอร์โดยตรง
  // ผู้เรียนจึงเห็นทุก request (และ 404) ใน Network tab ของ DevTools
  ssr: false,

  css: ['~/assets/main.css'],

  app: {
    head: {
      htmlAttrs: { lang: 'th' },
      title: 'sda-shop',
      meta: [{ name: 'viewport', content: 'width=device-width, initial-scale=1' }],
    },
  },

  runtimeConfig: {
    public: {
      // เปลี่ยนได้ด้วย env NUXT_PUBLIC_API_BASE (ใช้ตอน deploy บน Vercel ในโมดูล 5)
      apiBase: 'http://localhost:8080',
    },
  },
})
