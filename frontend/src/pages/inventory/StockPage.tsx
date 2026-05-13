import { useState, useEffect } from 'react'
import api from '../../api/client'

export default function StockPage() {
  const [products, setProducts] = useState([])
  const [movements, setMovements] = useState([])
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState('')
  const [success, setSuccess] = useState('')
  const [form, setForm] = useState({ product_id: '', quantity: '', note: '' })
  const [saving, setSaving] = useState(false)
  const [selectedProduct, setSelectedProduct] = useState('')

  useEffect(() => { loadProducts() }, [])

  async function loadProducts() {
    setLoading(true)
    try { setProducts((await api.get('/products')).data.data || []) }
    catch { setError('Failed to load products') }
    finally { setLoading(false) }
  }

  async function loadMovements(productId) {
    try { setMovements((await api.get(`/products/${productId}/movements`)).data.data || []) }
    catch { setMovements([]) }
  }

  async function handleProductSelect(id) {
    setSelectedProduct(id)
    if (id) await loadMovements(id)
    else setMovements([])
  }

  async function handleAdjust(e) {
    e.preventDefault(); setSaving(true); setError(''); setSuccess('')
    try {
      await api.post('/stock/adjust', {
        product_id: Number(form.product_id),
        quantity: Number(form.quantity),
        note: form.note,
      })
      setSuccess('Stock adjusted successfully')
      setForm({ product_id: '', quantity: '', note: '' })
      loadProducts()
      if (selectedProduct) loadMovements(selectedProduct)
    } catch (err) {
      setError(err.response?.data?.error || 'Adjust failed')
    } finally { setSaving(false) }
  }

  const typeStyle = { in: 'badge-green', out: 'badge-red', adjust: 'badge-blue' }

  return (
    <div className="page">
      <div className="page-header">
        <div>
          <div className="page-title">Stock Adjust</div>
          <div className="page-desc">Manual stock adjustments and movement history</div>
        </div>
      </div>

      <div style={{ display: 'grid', gridTemplateColumns: '340px 1fr', gap: 24, alignItems: 'start' }}>
        <div className="table-wrap" style={{ padding: 20 }}>
          <div className="section-title" style={{ marginBottom: 16 }}>Adjust Stock</div>
          {error && <div className="alert alert-error">{error}</div>}
          {success && <div className="alert" style={{ background: '#dcfce7', color: '#15803d', border: '1px solid #86efac', marginBottom: 12 }}>{success}</div>}
          <form onSubmit={handleAdjust}>
            <div className="form-group">
              <label className="form-label">Product</label>
              <select
                className="form-select"
                value={form.product_id}
                onChange={e => setForm({ ...form, product_id: e.target.value })}
                required
              >
                <option value="">Select product</option>
                {products.map(p => (
                  <option key={p.id} value={p.id}>{p.name} (stock: {p.stock_quantity})</option>
                ))}
              </select>
            </div>
            <div className="form-group">
              <label className="form-label">Quantity (+/−)</label>
              <input
                className="form-input"
                type="number"
                step="0.01"
                placeholder="e.g. 50 or -10"
                value={form.quantity}
                onChange={e => setForm({ ...form, quantity: e.target.value })}
                required
              />
            </div>
            <div className="form-group">
              <label className="form-label">Note</label>
              <input className="form-input" value={form.note} onChange={e => setForm({ ...form, note: e.target.value })} placeholder="Reason for adjustment" />
            </div>
            <button type="submit" className="btn btn-primary" style={{ width: '100%', justifyContent: 'center' }} disabled={saving || loading}>
              {saving ? 'Adjusting…' : 'Adjust Stock'}
            </button>
          </form>
        </div>

        <div>
          <div style={{ marginBottom: 12 }}>
            <div className="section-title">Movement History</div>
            <select
              className="form-select"
              style={{ marginTop: 8, maxWidth: 300 }}
              value={selectedProduct}
              onChange={e => handleProductSelect(e.target.value)}
            >
              <option value="">Select a product to view movements</option>
              {products.map(p => <option key={p.id} value={p.id}>{p.name}</option>)}
            </select>
          </div>

          <div className="table-wrap">
            <table className="data-table">
              <thead>
                <tr>
                  <th>Date</th>
                  <th>Type</th>
                  <th style={{ textAlign: 'right' }}>Qty</th>
                  <th>Reference</th>
                  <th>Note</th>
                </tr>
              </thead>
              <tbody>
                {!selectedProduct ? (
                  <tr className="empty-row"><td colSpan={5}>Select a product above</td></tr>
                ) : movements.length === 0 ? (
                  <tr className="empty-row"><td colSpan={5}>No movements found</td></tr>
                ) : movements.map(m => (
                  <tr key={m.id}>
                    <td className="col-mono">{new Date(m.created_at).toLocaleDateString()}</td>
                    <td><span className={`badge ${typeStyle[m.type] || 'badge-gray'}`}>{m.type}</span></td>
                    <td style={{ textAlign: 'right', fontWeight: 500 }}>{m.quantity > 0 ? `+${m.quantity}` : m.quantity}</td>
                    <td style={{ color: '#888', fontSize: 12 }}>{m.reference_type || '—'}</td>
                    <td style={{ color: '#888' }}>{m.note || '—'}</td>
                  </tr>
                ))}
              </tbody>
            </table>
          </div>
        </div>
      </div>
    </div>
  )
}
