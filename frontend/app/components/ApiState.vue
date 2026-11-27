<script setup lang="ts">
// แสดงสถานะการโหลดข้อมูลจาก API: กำลังโหลด / error / ว่าง
// ถ้าส่ง skeleton มา: ตอนโหลด และตอนที่ API ยังไม่มี (404) หรือต่อ backend ไม่ได้
// จะแสดงโครงของหน้าจริงแบบไม่มีข้อมูล พร้อมป้ายบอกว่ารอเส้นไหน แทนกล่องสีแดง
import type { ApiError } from '~/composables/useApi'

type Skeleton = 'chips' | 'cards' | 'detail' | 'fields' | 'rows'

const props = defineProps<{
  pending: boolean
  error: ApiError | null
  empty?: boolean
  endpoint: string
  emptyText?: string
  skeleton?: Skeleton
  count?: number
  placeholder?: boolean
}>()

// 404 = backend ยังไม่มีเส้นนี้ · 0 = ต่อ backend ไม่ได้ · error อื่นยังแสดงกล่องแดงเหมือนเดิม
const waiting = computed(() => !!props.error && (props.error.status === 404 || props.error.status === 0))
const showSkeleton = computed(() => !!props.skeleton && (props.pending || waiting.value))
const n = computed(() => props.count ?? { chips: 4, cards: 4, detail: 1, fields: 3, rows: 3 }[props.skeleton ?? 'rows'])
</script>

<template>
  <div v-if="props.placeholder && !props.pending" class="ph">
    <p class="skel-tag ph-tag">
      <b>ข้อมูลตัวอย่าง</b>
      <template v-if="props.error?.status === 404">ยังไม่มีเส้นนี้ (404)</template>
      <template v-else>ต่อ backend ไม่ได้</template>
      <code>{{ props.endpoint }}</code>
    </p>
    <div class="ph-body"><slot /></div>
  </div>

  <div v-else-if="showSkeleton" class="skel" :class="{ 'skel-wait': waiting }" :aria-busy="props.pending">
    <p v-if="waiting" class="skel-tag">
      <template v-if="props.error?.status === 404">รอ API · ยังไม่มีเส้นนี้ (404)</template>
      <template v-else>ต่อ backend ไม่ได้</template>
      <code>{{ props.endpoint }}</code>
    </p>

    <div v-if="props.skeleton === 'chips'" class="chips">
      <span v-for="i in n" :key="i" class="chip skel-chip" />
    </div>

    <div v-else-if="props.skeleton === 'cards'" class="grid">
      <div v-for="i in n" :key="i" class="card product-card skel-card">
        <div class="skel-img" />
        <div class="skel-body"><i class="skel-line s" /><i class="skel-line" /><i class="skel-line m" /></div>
      </div>
    </div>

    <div v-else-if="props.skeleton === 'detail'" class="detail">
      <div class="skel-img" />
      <div class="skel-body"><i class="skel-line s" /><i class="skel-line l" /><i class="skel-line m" /><i class="skel-line" /><i class="skel-line" /></div>
    </div>

    <div v-else-if="props.skeleton === 'fields'" class="card skel-body">
      <p v-for="i in n" :key="i" class="skel-field"><i class="skel-line s" /><i class="skel-line" /></p>
    </div>

    <div v-else class="card skel-body">
      <p v-for="i in n" :key="i" class="skel-row"><i class="skel-line" /><i class="skel-line s" /><i class="skel-line s" /></p>
    </div>
  </div>

  <div v-else-if="props.pending" class="state">กำลังโหลด…</div>
  <div v-else-if="props.error" class="state state-error">
    <strong v-if="props.error.status === 404">ยังไม่มีข้อมูล — API ยังไม่พร้อม (404)</strong>
    <strong v-else-if="props.error.status === 0">ต่อ backend ไม่ได้</strong>
    <strong v-else>API ตอบกลับ {{ props.error.status }}</strong>
    <p>{{ props.error.message }}</p>
    <code>{{ props.endpoint }}</code>
  </div>
  <div v-else-if="props.empty" class="state">{{ props.emptyText ?? 'ยังไม่มีข้อมูล' }}</div>
  <slot v-else />
</template>
