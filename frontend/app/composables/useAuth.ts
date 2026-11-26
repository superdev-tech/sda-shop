// เก็บ token ที่ได้จากการ login ไว้ในเบราว์เซอร์ (localStorage)
const TOKEN_KEY = 'sda-shop-token'
const USER_KEY = 'sda-shop-user'

export interface AuthUser {
  id: number
  email: string
  name: string
  role: 'customer' | 'admin'
}

export function useAuth() {
  const token = useState<string | null>('auth-token', () =>
    import.meta.client ? localStorage.getItem(TOKEN_KEY) : null,
  )
  const user = useState<AuthUser | null>('auth-user', () => {
    if (!import.meta.client) return null
    const raw = localStorage.getItem(USER_KEY)
    try {
      return raw ? (JSON.parse(raw) as AuthUser) : null
    } catch {
      return null
    }
  })

  function setSession(newToken: string, newUser: AuthUser | null) {
    token.value = newToken
    user.value = newUser
    localStorage.setItem(TOKEN_KEY, newToken)
    if (newUser) localStorage.setItem(USER_KEY, JSON.stringify(newUser))
  }

  function logout() {
    token.value = null
    user.value = null
    localStorage.removeItem(TOKEN_KEY)
    localStorage.removeItem(USER_KEY)
  }

  const isLoggedIn = computed(() => !!token.value)
  const isAdmin = computed(() => user.value?.role === 'admin')

  return { token, user, isLoggedIn, isAdmin, setSession, logout }
}
