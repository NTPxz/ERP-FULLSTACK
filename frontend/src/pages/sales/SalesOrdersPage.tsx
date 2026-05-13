import { useState, useEffect } from 'react'
import api from '../../api/client'
import Modal from '../../components/Modal'

const blankItem = { product_id: '', quantity: 1, unit_price: 0, discount: 0 }

const statusBadge = { draft: 'badge-gray', confirmed: 'badge-blue', completed: 'badge-green', cancelled: 'badge-red' }

export default function SalesOrdersPage() {
  const [orders, setOrders] = useState([])
  const [customers, setCustomers] = useState([])
  const [products, setProducts] = useState([])
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState('')
  const [mode, setMode] = useState(null) // 'create' | 'view'
  const [selected, setSelected] = useState(null)
  const [form, setForm] = useState({ customer_id: '', note: '', items: [{ ...blankItem }] })
  const [saving, setSaving] = useState(false)

  useEffect(() => { load() }, [])

  async function load() {
    setLoading(true)
    try {
      const [so, cust, prod] = await Promise.all([
        api.get('/sales-orders'), api.get('/customers'), api.get('/products'),
      ])
      setOrders(so.data.data || [])
      setCustomers(cust.data.data || [])
      setProducts(prod.data.data || [])
    } catch { setError('Failed to load') }
    finally { setLoading(false) }
  }

  function openCreate() {
    setForm({ customer_id: '', note: '', items: [{ ...blankItem }] })
    setError(''); setMode('create')
  }

  function openView(order) { setSelected(order); setMode('view') }

  function closeModal() { setMode(null); setSelected(null); setError('') }

  function addItem() { setForm({ ...form, items: [...form.items, { ...blankItem }] }) }
  function removeItem(i) { setForm({ ...form, items: form.items.filter((_, idx) => idx !== i) }) }
  function updateItem(i, field, val) {
    const items = [...form.items]
    items[i] = { ...items[i], [field]: val }
    setForm({ ...form, items })
  }

  const total = form.items.reduce((s, it) => s + (Number(it.unit_price) * Number(it.quantity)) - Number(it.discount || 0), 0)

  async function handleCreate(e) {
    e.preventDefault(); setSaving(true); setError('')
    try {
      await api.post('/sales-orders', {
        customer_id: Number(form.customer_id),
        note: form.note,
        items: form.items.map(it => ({
          product_id: Number(it.product_id),
          quantity: Number(it.quantity),
          unit_price: Number(it.unit_price),
          discount: Number(it.discount || 0),
        })),
      })
      closeModal(); load()
    } catch (err) {
      setError(err.response?.data?.error || 'Create failed')
    } finally { setSaving(false) }
  }

  async function changeStatus(id, status) {
    setError('')
    try { await api.patch(`/sales-orders/${id}/status`, { status }); load() }
    catch (err) { setError(err.response?.data?.error || 'Status update failed') }
  }

  async function handleDelete(id) {
    if (!confirm('Delete this order?')) return
    try { await api.delete(`/sales-orders/${id}`); load() }
    catch (err) { setError(err.response?.data?.error || 'Delete failed') }
  }

  return (
    <div className="page">
      <div className="page-header">
        <div>
          <div className="page-title">Sales Orders</div>
          <div className="page-desc">Manage customer orders and stock allocation</div>
        </div>
        <button className="btn btn-primary" onClick={openCreate}>+ New Order</button>
      </div>

      {error && !mode && <div className="alert alert-error">{error}</div>}

      {loading ? <div className="loading">Loading…</div> : (
        <div className="table-wrap">
          <table className="data-table">
            <thead>
              <tr>
                <th>Order No</th>
                <th>Customer</th>
                <th>Status</th>
                <th>Date</th>
                <th style={{ textAlign: 'right' }}>Total</th>
                <th className="col-actions"></th>
              </tr>
            </thead>
            <tbody>
              {orders.length === 0 ? (
                <tr className="empty-row"><td colSpan={6}>No sales orders yet</td></tr>
              ) : orders.map(o => (
                <tr key={o.id}>
                  <td className="col-mono">{o.order_no}</td>
                  <td>{o.customer?.name || '—'}</td>
                  <td><span className={`badge ${statusBadge[o.status] || 'badge-gray'}`}>{o.status}</span></td>
                  <td style={{ color: '#888' }}>{o.order_date ? new Date(o.order_date).toLocaleDateString() : '—'}</td>
                  <td style={{ textAlign: 'right' }}>{o.total_amount?.toLocaleString(undefined, { minimumFractionDigits: 2 })}</td>
                  <td className="col-actions">
                    <div className="row-actions">
                      <button className="btn btn-ghost btn-sm" onClick={() => openView(o)}>View</button>
                      {o.status === 'draft' && <button className="btn btn-ghost btn-sm" style={{ color: '#1d4ed8' }} onClick={() => changeStatus(o.id, 'confirmed')}>Confirm</button>}
                      {o.status === 'confirmed' && <button className="btn btn-ghost btn-sm" style={{ color: '#15803d' }} onClick={() => changeStatus(o.id, 'completed')}>Complete</button>}
                      {(o.status === 'draft' || o.status === 'confirmed') && <button className="btn btn-ghost btn-sm" style={{ color: '#888' }} onClick={() => changeStatus(o.id, 'cancelled')}>Cancel</button>}
                      {o.status === 'draft' && <button className="btn btn-danger btn-sm" onClick={() => handleDelete(o.id)}>Delete</button>}
                    </div>
                  </td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
      )}

      {mode === 'create' && (
        <Modal title="New Sales Order" onClose={closeModal} large>
          <form onSubmit={handleCreate}>
            <div className="modal-body">
              {error && <div className="alert alert-error">{error}</div>}
              <div className="form-row form-row-2">
                <div>
                  <label className="form-label">Customer</label>
                  <select className="form-select" value={form.customer_id} onChange={e => setForm({ ...form, customer_id: e.target.value })} required>
                    <option value="">Select customer</option>
                    {customers.map(c => <option key={c.id} value={c.id}>{c.name}</option>)}
                  </select>
                </div>
                <div>
                  <label className="form-label">Note</label>
                  <input className="form-input" value={form.note} onChange={e => setForm({ ...form, note: e.target.value })} />
                </div>
              </div>

              <div className="items-section">
                <div className="items-section-title">Items</div>
                <div className="item-row item-row-so" style={{ marginBottom: 4 }}>
                  <div className="item-col-label">Product</div>
                  <div className="item-col-label">Qty</div>
                  <div className="item-col-label">Unit Price</div>
                  <div className="item-col-label">Discount</div>
                  <div></div>
                </div>
                {form.items.map((it, i) => (
                  <div key={i} className="item-row item-row-so">
                    <select className="form-select" value={it.product_id} onChange={e => updateItem(i, 'product_id', e.target.value)} required>
                      <option value="">Select product</option>
                      {products.map(p => <option key={p.id} value={p.id}>{p.name} (stock: {p.stock_quantity})</option>)}
                    </select>
                    <input className="form-input" type="number" min="1" step="0.01" value={it.quantity} onChange={e => updateItem(i, 'quantity', e.target.value)} required />
                    <input className="form-input" type="number" min="0" step="0.01" value={it.unit_price} onChange={e => updateItem(i, 'unit_price', e.target.value)} required />
                    <input className="form-input" type="number" min="0" step="0.01" value={it.discount} onChange={e => updateItem(i, 'discount', e.target.value)} />
                    <button type="button" className="item-remove" onClick={() => removeItem(i)}>×</button>
                  </div>
                ))}
                <div className="items-add">
                  <button type="button" className="btn btn-ghost btn-sm" onClick={addItem}>+ Add Item</button>
                </div>
                <div className="items-total">Total: {total.toLocaleString(undefined, { minimumFractionDigits: 2 })}</div>
              </div>
            </div>
            <div className="modal-footer">
              <button type="button" className="btn btn-secondary" onClick={closeModal}>Cancel</button>
              <button type="submit" className="btn btn-primary" disabled={saving}>{saving ? 'Creating…' : 'Create Order'}</button>
            </div>
          </form>
        </Modal>
      )}

      {mode === 'view' && selected && (
        <Modal title={`Order ${selected.order_no}`} onClose={closeModal} large>
          <div className="modal-body">
            <div className="detail-meta">
              <div>
                <div className="detail-field-label">Customer</div>
                <div className="detail-field-value">{selected.customer?.name || '—'}</div>
              </div>
              <div>
                <div className="detail-field-label">Status</div>
                <div className="detail-field-value"><span className={`badge ${statusBadge[selected.status] || 'badge-gray'}`}>{selected.status}</span></div>
              </div>
              <div>
                <div className="detail-field-label">Date</div>
                <div className="detail-field-value">{selected.order_date ? new Date(selected.order_date).toLocaleDateString() : '—'}</div>
              </div>
              <div>
                <div className="detail-field-label">Total</div>
                <div className="detail-field-value" style={{ fontWeight: 600 }}>{selected.total_amount?.toLocaleString(undefined, { minimumFractionDigits: 2 })}</div>
              </div>
              {selected.note && <div style={{ gridColumn: '1 / -1' }}>
                <div className="detail-field-label">Note</div>
                <div className="detail-field-value">{selected.note}</div>
              </div>}
            </div>
            <div className="items-section-title" style={{ marginBottom: 8 }}>Items</div>
            <div className="table-wrap">
              <table className="data-table">
                <thead>
                  <tr>
                    <th>Product</th>
                    <th style={{ textAlign: 'right' }}>Qty</th>
                    <th style={{ textAlign: 'right' }}>Unit Price</th>
                    <th style={{ textAlign: 'right' }}>Discount</th>
                    <th style={{ textAlign: 'right' }}>Subtotal</th>
                  </tr>
                </thead>
                <tbody>
                  {(selected.items || []).map(it => (
                    <tr key={it.id}>
                      <td>{it.product?.name || `Product #${it.product_id}`}</td>
                      <td style={{ textAlign: 'right' }}>{it.quantity}</td>
                      <td style={{ textAlign: 'right' }}>{it.unit_price?.toLocaleString(undefined, { minimumFractionDigits: 2 })}</td>
                      <td style={{ textAlign: 'right' }}>{it.discount?.toLocaleString(undefined, { minimumFractionDigits: 2 })}</td>
                      <td style={{ textAlign: 'right', fontWeight: 500 }}>{it.subtotal?.toLocaleString(undefined, { minimumFractionDigits: 2 })}</td>
                    </tr>
                  ))}
                </tbody>
              </table>
            </div>
          </div>
          <div className="modal-footer">
            <button className="btn btn-secondary" onClick={closeModal}>Close</button>
          </div>
        </Modal>
      )}
    </div>
  )
}
