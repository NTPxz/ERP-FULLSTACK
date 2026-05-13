import { createContext, useContext, useState, type ReactNode } from 'react'
import http, { apiGet } from '../api/client'
import type { User } from '../types'

interface AuthContextValue {
  token: string | null
  user: User | null
  login: (token: string) => Promise<void>
  logout: () => void
}

const AuthContext = createContext<AuthContextValue | null>(null)

export function AuthProvider({ children }: { children: ReactNode }) {
  const [token, setToken] = useState<string | null>(() => localStorage.getItem('token'))
  const [user, setUser] = useState<User | null>(() => {
    const stored = localStorage.getItem('user')
    return stored ? (JSON.parse(stored) as User) : null
  })

  async function login(tok: string) {
    localStorage.setItem('token', tok)
    setToken(tok)
    try {
      const u = await apiGet<User>('/me')
      localStorage.setItem('user', JSON.stringify(u))
      setUser(u)
    } catch {
      // proceed even if /me fails
    }
  }

  function logout() {
    localStorage.removeItem('token')
    localStorage.removeItem('user')
    setToken(null)
    setUser(null)
  }

  return (
    <AuthContext.Provider value={{ token, user, login, logout }}>
      {children}
    </AuthContext.Provider>
  )
}

export function useAuth(): AuthContextValue {
  const ctx = useContext(AuthContext)
  if (!ctx) throw new Error('useAuth must be used within AuthProvider')
  return ctx
}

export { http }
