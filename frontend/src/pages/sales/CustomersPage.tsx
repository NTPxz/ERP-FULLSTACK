import { useState, useEffect } from 'react'
import api from '../../api/client'
import Modal from '../../components/Modal'

const blank = { code: '', name: '', email: '', phone: '', address: '', tax_id: '', credit_limit: '' }

export default function CustomersPage() {
  const [rows, setRows] = useState([])
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState('')
  const [showModal, setShowModal] = useState(false)
  const [editing, setEditing] = useState(null)
  const [form, setForm] = useState(blank)
  const [saving, setSaving] = useState(false)

  useEffect(() => { load() }, [])

  async function load() {
    setLoading(true)
    try { setRows((await api.get('/customers')).data.data || []) }
    catch { setError('Failed to load') }
    finally { setLoading(false) }
  }

  function openCreate() { setEditing(null); setForm(blank); setError(''); setShowModal(true) }
  function openEdit(row) {
    setEditing(row)
    setForm({ code: row.code, name: row.name, email: row.email || '', phone: row.phone || '', address: row.address || '', tax_id: row.tax_id || '', credit_limit: String(row.credit_limit || 0) })
    setError(''); setShowModal(true)
  }

  async function handleSubmit(e) {
    e.preventDefault(); setSaving(true); setError('')
    const payload = { ...form, credit_limit: Number(form.credit_limit) }
    try {
      if (editing) await api.put(`/customers/${editing.id}`, payload)
      else await api.post('/customers', payload)
      setShowModal(false); load()
    } catch (err) {
      setError(err.response?.data?.error || 'Save failed')
    } finally { setSaving(false) }
  }

  async function handleDelete(id) {
    if (!confirm('Delete this customer?')) return
    try { await api.delete(`/customers/${id}`); load() }
    catch (err) { setError(err.response?.data?.error || 'Delete failed') }
  }

  return (
    <div className="page">
      <div className="page-header">
        <div>
          <div className="page-title">Customers</div>
          <div className="page-desc">Customer accounts and contacts</div>
        </div>
        <button className="btn btn-primary" onClick={openCreate}>+ Add Customer</button>
      </div>

      {error && !showModal && <div className="alert alert-error">{error}</div>}

      {loading ? <div className="loading">Loading…</div> : (
        <div className="table-wrap">
          <table className="data-table">
            <thead>
              <tr>
                <th>Code</th>
                <th>Name</th>
                <th>Email</th>
                <th>Phone</th>
                <th>Tax ID</th>
                <th style={{ textAlign: 'right' }}>Credit Limit</th>
                <th className="col-actions"></th>
              </tr>
            </thead>
            <tbody>
              {rows.length === 0 ? (
                <tr className="empty-row"><td colSpan={7}>No customers yet</td></tr>
              ) : rows.map(row => (
                <tr key={row.id}>
                  <td className="col-mono">{row.code}</td>
                  <td style={{ fontWeight: 500 }}>{row.name}</td>
                  <td style={{ color: '#888' }}>{row.email || '—'}</td>
                  <td style={{ color: '#888' }}>{row.phone || '—'}</td>
                  <td style={{ color: '#888' }}>{row.tax_id || '—'}</td>
                  <td style={{ textAlign: 'right' }}>{row.credit_limit?.toLocaleString()}</td>
                  <td className="col-actions">
                    <div className="row-actions">
                      <button className="btn btn-ghost btn-sm" onClick={() => openEdit(row)}>Edit</button>
                      <button className="btn btn-danger btn-sm" onClick={() => handleDelete(row.id)}>Delete</button>
                    </div>
                  </td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
      )}

      {showModal && (
        <Modal title={editing ? 'Edit Customer' : 'New Customer'} onClose={() => setShowModal(false)}>
          <form onSubmit={handleSubmit}>
            <div className="modal-body">
              {error && <div className="alert alert-error">{error}</div>}
              <div className="form-row form-row-2">
                <div>
                  <label className="form-label">Code</label>
                  <input className="form-input" value={form.code} onChange={e => setForm({ ...form, code: e.target.value })} required autoFocus />
                </div>
                <div>
                  <label className="form-label">Tax ID</label>
                  <input className="form-input" value={form.tax_id} onChange={e => setForm({ ...form, tax_id: e.target.value })} />
                </div>
              </div>
              <div className="form-group">
                <label className="form-label">Name</label>
                <input className="form-input" value={form.name} onChange={e => setForm({ ...form, name: e.target.value })} required />
              </div>
              <div className="form-row form-row-2">
                <div>
                  <label className="form-label">Email</label>
                  <input className="form-input" type="email" value={form.email} onChange={e => setForm({ ...form, email: e.target.value })} />
                </div>
                <div>
                  <label className="form-label">Phone</label>
                  <input className="form-input" value={form.phone} onChange={e => setForm({ ...form, phone: e.target.value })} />
                </div>
              </div>
              <div className="form-group">
                <label className="form-label">Address</label>
                <textarea className="form-textarea" value={form.address} onChange={e => setForm({ ...form, address: e.target.value })} />
              </div>
              <div className="form-group">
                <label className="form-label">Credit Limit</label>
                <input className="form-input" type="number" value={form.credit_limit} onChange={e => setForm({ ...form, credit_limit: e.target.value })} />
              </div>
            </div>
            <div className="modal-footer">
              <button type="button" className="btn btn-secondary" onClick={() => setShowModal(false)}>Cancel</button>
              <button type="submit" className="btn btn-primary" disabled={saving}>{saving ? 'Saving…' : editing ? 'Save' : 'Create'}</button>
            </div>
          </form>
        </Modal>
      )}
    </div>
  )
}
