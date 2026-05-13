import { useState, useEffect } from 'react'
import api from '../../api/client'
import Modal from '../../components/Modal'

const blank = { sku: '', name: '', description: '', category_id: '', unit: '', cost_price: '', sale_price: '', reorder_level: '' }

export default function ProductsPage() {
  const [rows, setRows] = useState([])
  const [categories, setCategories] = useState([])
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState('')
  const [showModal, setShowModal] = useState(false)
  const [editing, setEditing] = useState(null)
  const [form, setForm] = useState(blank)
  const [saving, setSaving] = useState(false)

  useEffect(() => { load() }, [])

  async function load() {
    setLoading(true)
    try {
      const [prod, cats] = await Promise.all([api.get('/products'), api.get('/categories')])
      setRows(prod.data.data || [])
      setCategories(cats.data.data || [])
    } catch { setError('Failed to load') }
    finally { setLoading(false) }
  }

  function openCreate() {
    setEditing(null); setForm(blank); setError(''); setShowModal(true)
  }
  function openEdit(row) {
    setEditing(row)
    setForm({
      sku: row.sku, name: row.name, description: row.description || '',
      category_id: String(row.category_id), unit: row.unit || '',
      cost_price: String(row.cost_price), sale_price: String(row.sale_price),
      reorder_level: String(row.reorder_level),
    })
    setError(''); setShowModal(true)
  }

  async function handleSubmit(e) {
    e.preventDefault(); setSaving(true); setError('')
    const payload = {
      ...form,
      category_id: Number(form.category_id),
      cost_price: Number(form.cost_price),
      sale_price: Number(form.sale_price),
      reorder_level: Number(form.reorder_level),
    }
    try {
      if (editing) await api.put(`/products/${editing.id}`, payload)
      else await api.post('/products', payload)
      setShowModal(false); load()
    } catch (err) {
      setError(err.response?.data?.error || 'Save failed')
    } finally { setSaving(false) }
  }

  async function handleDelete(id) {
    if (!confirm('Delete this product?')) return
    try { await api.delete(`/products/${id}`); load() }
    catch (err) { setError(err.response?.data?.error || 'Delete failed') }
  }

  return (
    <div className="page">
      <div className="page-header">
        <div>
          <div className="page-title">Products</div>
          <div className="page-desc">Product catalog and stock levels</div>
        </div>
        <button className="btn btn-primary" onClick={openCreate}>+ Add Product</button>
      </div>

      {error && !showModal && <div className="alert alert-error">{error}</div>}

      {loading ? <div className="loading">Loading…</div> : (
        <div className="table-wrap">
          <table className="data-table">
            <thead>
              <tr>
                <th>SKU</th>
                <th>Name</th>
                <th>Category</th>
                <th>Unit</th>
                <th style={{ textAlign: 'right' }}>Cost</th>
                <th style={{ textAlign: 'right' }}>Sale Price</th>
                <th style={{ textAlign: 'right' }}>Stock</th>
                <th style={{ textAlign: 'right' }}>Reorder</th>
                <th className="col-actions"></th>
              </tr>
            </thead>
            <tbody>
              {rows.length === 0 ? (
                <tr className="empty-row"><td colSpan={9}>No products yet</td></tr>
              ) : rows.map(row => (
                <tr key={row.id}>
                  <td className="col-mono">{row.sku}</td>
                  <td style={{ fontWeight: 500 }}>{row.name}</td>
                  <td style={{ color: '#888' }}>{row.category?.name || '—'}</td>
                  <td style={{ color: '#888' }}>{row.unit || '—'}</td>
                  <td style={{ textAlign: 'right' }}>{row.cost_price?.toLocaleString(undefined, { minimumFractionDigits: 2 })}</td>
                  <td style={{ textAlign: 'right' }}>{row.sale_price?.toLocaleString(undefined, { minimumFractionDigits: 2 })}</td>
                  <td style={{ textAlign: 'right', fontWeight: row.stock_quantity <= row.reorder_level ? 600 : 400, color: row.stock_quantity <= row.reorder_level ? '#dc2626' : '#111' }}>
                    {row.stock_quantity}
                  </td>
                  <td style={{ textAlign: 'right', color: '#aaa' }}>{row.reorder_level}</td>
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
        <Modal title={editing ? 'Edit Product' : 'New Product'} onClose={() => setShowModal(false)}>
          <form onSubmit={handleSubmit}>
            <div className="modal-body">
              {error && <div className="alert alert-error">{error}</div>}
              <div className="form-row form-row-2">
                <div>
                  <label className="form-label">SKU</label>
                  <input className="form-input" value={form.sku} onChange={e => setForm({ ...form, sku: e.target.value })} required autoFocus />
                </div>
                <div>
                  <label className="form-label">Unit</label>
                  <input className="form-input" value={form.unit} onChange={e => setForm({ ...form, unit: e.target.value })} placeholder="pcs, kg, box…" />
                </div>
              </div>
              <div className="form-group">
                <label className="form-label">Name</label>
                <input className="form-input" value={form.name} onChange={e => setForm({ ...form, name: e.target.value })} required />
              </div>
              <div className="form-group">
                <label className="form-label">Category</label>
                <select className="form-select" value={form.category_id} onChange={e => setForm({ ...form, category_id: e.target.value })} required>
                  <option value="">Select category</option>
                  {categories.map(c => <option key={c.id} value={c.id}>{c.name}</option>)}
                </select>
              </div>
              <div className="form-group">
                <label className="form-label">Description</label>
                <textarea className="form-textarea" value={form.description} onChange={e => setForm({ ...form, description: e.target.value })} />
              </div>
              <div className="form-row form-row-3">
                <div>
                  <label className="form-label">Cost Price</label>
                  <input className="form-input" type="number" step="0.01" value={form.cost_price} onChange={e => setForm({ ...form, cost_price: e.target.value })} />
                </div>
                <div>
                  <label className="form-label">Sale Price</label>
                  <input className="form-input" type="number" step="0.01" value={form.sale_price} onChange={e => setForm({ ...form, sale_price: e.target.value })} required />
                </div>
                <div>
                  <label className="form-label">Reorder Level</label>
                  <input className="form-input" type="number" value={form.reorder_level} onChange={e => setForm({ ...form, reorder_level: e.target.value })} />
                </div>
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
