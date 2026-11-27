// ตัวช่วยเรียก API ของ backend
// - ใส่ baseURL จาก runtimeConfig.public.apiBase
// - แนบ Authorization: Bearer <token> ถ้า login แล้ว
// - แปลง error ทุกแบบให้เป็น ApiError ที่หน้าเว็บแสดงผลได้

export class ApiError extends Error {
  constructor(
    public status: number,
    public code: string,
    message: string,
  ) {
    super(message)
  }
}

type Method = 'GET' | 'POST' | 'PATCH' | 'DELETE'

interface RequestOptions {
  method?: Method
  body?: Record<string, unknown>
  query?: Record<string, string | number | undefined>
}

function toApiError(e: unknown): ApiError {
  const err = e as {
    response?: { status?: number }
    statusCode?: number
    data?: { error?: { code?: string; message?: string } }
  }
  const status = err?.response?.status ?? err?.statusCode ?? 0
  if (!status) {
    return new ApiError(
      0,
      'network_error',
      'ต่อ backend ไม่ได้ (backend ยังไม่รัน หรือติด CORS) ลองเปิด Network tab ดู',
    )
  }
  const code = err.data?.error?.code ?? 'http_' + status
  const message =
    err.data?.error?.message ??
    (status === 404 ? 'ยังไม่มีเส้น API นี้ใน backend' : 'API ตอบกลับด้วย error')
  return new ApiError(status, code, message)
}

export function useApi() {
  const config = useRuntimeConfig()
  const { token } = useAuth()

  async function request<T>(path: string, options: RequestOptions = {}): Promise<T> {
    try {
      return await $fetch<T>(path, {
        baseURL: config.public.apiBase as string,
        method: options.method ?? 'GET',
        body: options.body,
        query: options.query,
        headers: token.value ? { Authorization: `Bearer ${token.value}` } : {},
      })
    } catch (e) {
      throw toApiError(e)
    }
  }

  return { request }
}

// โหลดข้อมูลตอนเปิดหน้า และโหลดใหม่เมื่อค่าใน watch เปลี่ยน
export function useApiData<T>(loader: () => Promise<T>, watchSources: unknown[] = []) {
  const data = ref<T | null>(null) as Ref<T | null>
  const error = ref<ApiError | null>(null)
  const pending = ref(true)

  async function refresh() {
    pending.value = true
    error.value = null
    try {
      data.value = await loader()
    } catch (e) {
      data.value = null
      error.value = e instanceof ApiError ? e : new ApiError(0, 'unknown', String(e))
    } finally {
      pending.value = false
    }
  }

  onMounted(refresh)
  if (watchSources.length) watch(watchSources as never, refresh)

  return { data, error, pending, refresh }
}
