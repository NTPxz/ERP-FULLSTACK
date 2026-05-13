import { useState, useEffect } from 'react'
import api from '../../api/client'
import Modal from '../../components/Modal'

const ROLES = ['admin', 'hr', 'inventory', 'sales', 'purchase']

const roleBadge = {
  admin:     'badge-purple',
  hr:        'badge-blue',
  inventory: 'badge-green',
  sales:     'badge-yellow',
  purchase:  'badge-red',
}

const emptyCreate = { name: '', email: '', password: '', role: 'hr' }
const emptyEdit   = { name: '', email: '', role: '' }

export default function UsersPage() {
  const [users, setUsers]   = useState([])
  const [mode, setMode]     = useState(null) // 'create' | 'edit'
  const [selected, setSel]  = useState(null)
  const [form, setForm]     = useState(emptyCreate)
  const [error, setError]   = useState('')
  const [loading, setLoading] = useState(true)

  function load() {
    setLoading(true)
    api.get('/users')
      .then(r => setUsers(r.data.data || []))
      .finally(() => setLoading(false))
  }

  useEffect(() => { load() }, [])

  function openCreate() {
    setForm(emptyCreate)
    setError('')
    setMode('create')
  }

  function openEdit(u) {
    setSel(u)
    setForm({ name: u.name, email: u.email, role: u.role })
    setError('')
    setMode('edit')
  }

  function closeModal() { setMode(null); setSel(null) }

  async function handleCreate(e) {
    e.preventDefault()
    setError('')
    try {
      await api.post('/users', form)
      closeModal()
      load()
    } catch (err) {
      setError(err.response?.data?.error || 'Failed to create user')
    }
  }

  async function handleEdit(e) {
    e.preventDefault()
    setError('')
    try {
      await api.patch(`/users/${selected.id}/role`, { name: form.name, email: form.email, role: form.role })
      closeModal()
      load()
    } catch (err) {
      setError(err.response?.data?.error || 'Failed to update user')
    }
  }

  async function handleDelete(u) {
    if (!window.confirm(`Delete user "${u.name}"?`)) return
    try {
      await api.delete(`/users/${u.id}`)
      load()
    } catch (err) {
      alert(err.response?.data?.error || 'Failed to delete user')
    }
  }

  return (
    <div className="page">
      <div className="page-header">
        <div>
          <div className="page-title">Users & Roles</div>
          <div className="page-desc">Manage system accounts and access permissions</div>
        </div>
        <button className="btn btn-primary" onClick={openCreate}>+ New User</button>
      </div>

      {loading ? <div className="loading">Loading…</div> : (
        <div className="table-wrap">
          <table className="data-table">
            <thead>
              <tr>
                <th>Name</th>
                <th>Email</th>
                <th>Role</th>
                <th>Created</th>
                <th style={{ width: 120 }}></th>
              </tr>
            </thead>
            <tbody>
              {users.length === 0 ? (
                <tr className="empty-row"><td colSpan={5}>No users found</td></tr>
              ) : users.map(u => (
                <tr key={u.id}>
                  <td>{u.name}</td>
                  <td className="text-muted">{u.email}</td>
                  <td>
                    <span className={`badge ${roleBadge[u.role] || 'badge-gray'}`}>{u.role}</span>
                  </td>
                  <td className="text-muted">{u.created_at ? new Date(u.created_at).toLocaleDateString() : '—'}</td>
                  <td>
                    <div style={{ display: 'flex', gap: 6 }}>
                      <button className="btn btn-sm" onClick={() => openEdit(u)}>Edit</button>
                      <button className="btn btn-sm btn-danger" onClick={() => handleDelete(u)}>Del</button>
                    </div>
                  </td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
      )}

      {mode === 'create' && (
        <Modal title="New User" onClose={closeModal}>
          <form onSubmit={handleCreate}>
            {error && <div className="alert alert-error" style={{ marginBottom: 12 }}>{error}</div>}
            <div className="form-group">
              <label className="form-label">Name</label>
              <input className="form-input" value={form.name} onChange={e => setForm({ ...form, name: e.target.value })} required />
            </div>
            <div className="form-group">
              <label className="form-label">Email</label>
              <input className="form-input" type="email" value={form.email} onChange={e => setForm({ ...form, email: e.target.value })} required />
            </div>
            <div className="form-group">
              <label className="form-label">Password</label>
              <input className="form-input" type="password" value={form.password} onChange={e => setForm({ ...form, password: e.target.value })} required minLength={6} />
            </div>
            <div className="form-group">
              <label className="form-label">Role</label>
              <select className="form-input" value={form.role} onChange={e => setForm({ ...form, role: e.target.value })}>
                {ROLES.map(r => <option key={r} value={r}>{r}</option>)}
              </select>
            </div>
            <div style={{ display: 'flex', gap: 8, justifyContent: 'flex-end', marginTop: 16 }}>
              <button type="button" className="btn" onClick={closeModal}>Cancel</button>
              <button type="submit" className="btn btn-primary">Create</button>
            </div>
          </form>
        </Modal>
      )}

      {mode === 'edit' && selected && (
        <Modal title={`Edit — ${selected.name}`} onClose={closeModal}>
          <form onSubmit={handleEdit}>
            {error && <div className="alert alert-error" style={{ marginBottom: 12 }}>{error}</div>}
            <div className="form-group">
              <label className="form-label">Name</label>
              <input className="form-input" value={form.name} onChange={e => setForm({ ...form, name: e.target.value })} />
            </div>
            <div className="form-group">
              <label className="form-label">Email</label>
              <input className="form-input" type="email" value={form.email} onChange={e => setForm({ ...form, email: e.target.value })} />
            </div>
            <div className="form-group">
              <label className="form-label">Role</label>
              <select className="form-input" value={form.role} onChange={e => setForm({ ...form, role: e.target.value })}>
                {ROLES.map(r => <option key={r} value={r}>{r}</option>)}
              </select>
            </div>
            <div style={{ display: 'flex', gap: 8, justifyContent: 'flex-end', marginTop: 16 }}>
              <button type="button" className="btn" onClick={closeModal}>Cancel</button>
              <button type="submit" className="btn btn-primary">Save</button>
            </div>
          </form>
        </Modal>
      )}
    </div>
  )
}
