import { useState } from 'react'
import { useNavigate } from 'react-router-dom'
import { apiPost } from '../api/client'
import { useAuth } from '../context/AuthContext'

interface LoginResponse { token: string }

export default function Login() {
  const { login } = useAuth()
  const navigate = useNavigate()
  const [form, setForm] = useState({ email: '', password: '' })
  const [error, setError] = useState('')
  const [loading, setLoading] = useState(false)

  async function handleSubmit(e: React.FormEvent) {
    e.preventDefault()
    setError('')
    setLoading(true)
    try {
      const res = await apiPost<LoginResponse>('/auth/login', form)
      await login(res.token)
      navigate('/')
    } catch {
      setError('Login failed. Check your email and password.')
    } finally {
      setLoading(false)
    }
  }

  return (
    <div className="login-page">
      <div className="login-left">
        <div className="login-left-glow" />
        <div className="login-left-glow2" />

        <div className="login-left-content">
          <div className="login-logo">
            <div className="login-logo-icon">
              <svg width={18} height={18} viewBox="0 0 24 24" fill="none" stroke="#fff" strokeWidth="2.2" strokeLinecap="round" strokeLinejoin="round">
                <rect x="2" y="3" width="20" height="14" rx="2" />
                <line x1="8" y1="21" x2="16" y2="21" />
                <line x1="12" y1="17" x2="12" y2="21" />
              </svg>
            </div>
            <span className="login-logo-text">ERP System</span>
          </div>

          <div className="login-headline">
            Manage your<br />business smarter.
          </div>
          <div className="login-subheadline">
            All-in-one platform for HR, inventory, sales and purchasing — built for teams that move fast.
          </div>
        </div>

        <div className="login-left-bottom">
          <ul className="login-feature-list">
            {['Role-based access control', 'Real-time inventory tracking', 'Sales & purchase orders', 'HR & employee management'].map(f => (
              <li key={f} className="login-feature-item">
                <span className="login-feature-dot" />
                {f}
              </li>
            ))}
          </ul>
        </div>
      </div>

      <div className="login-right">
        <div className="login-card">
          <div className="login-title">Welcome back</div>
          <div className="login-sub">Sign in to your account to continue</div>

          {error && <div className="alert alert-error">{error}</div>}

          <form onSubmit={handleSubmit}>
            <div className="form-group">
              <label className="form-label">Email address</label>
              <input
                className="form-input"
                type="email"
                value={form.email}
                onChange={e => setForm({ ...form, email: e.target.value })}
                placeholder="you@company.com"
                required
                autoFocus
              />
            </div>
            <div className="form-group" style={{ marginBottom: 24 }}>
              <label className="form-label">Password</label>
              <input
                className="form-input"
                type="password"
                value={form.password}
                onChange={e => setForm({ ...form, password: e.target.value })}
                placeholder="••••••••"
                required
              />
            </div>
            <button
              className="btn btn-primary"
              style={{ width: '100%', justifyContent: 'center', padding: '10px 16px' }}
              disabled={loading}
            >
              {loading ? 'Signing in…' : 'Sign in'}
            </button>
          </form>
        </div>
      </div>
    </div>
  )
}
