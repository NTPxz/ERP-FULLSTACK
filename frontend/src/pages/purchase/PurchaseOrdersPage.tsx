import { useState, useEffect } from 'react'
import api from '../../api/client'
import Modal from '../../components/Modal'

const blankItem = { product_id: '', quantity: 1, unit_price: 0 }

const statusBadge = { draft: 'badge-gray', sent: 'badge-blue', received: 'badge-green', cancelled: 'badge-red' }

export default function PurchaseOrdersPage() {
  const [orders, setOrders] = useState([])
  const [suppliers, setSuppliers] = useState([])
  const [products, setProducts] = useState([])
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState('')
  const [mode, setMode] = useState(null)
  const [selected, setSelected] = useState(null)
  const [form, setForm] = useState({ supplier_id: '', note: '', items: [{ ...blankItem }] })
  const [saving, setSaving] = useState(false)

  useEffect(() => { load() }, [])

  async function load() {
    setLoading(true)
    try {
      const [po, sup, prod] = await Promise.all([
        api.get('/purchase-orders'), api.get('/suppliers'), api.get('/products'),
      ])
      setOrders(po.data.data || [])
      setSuppliers(sup.data.data || [])
      setProducts(prod.data.data || [])
    } catch { setError('Failed to load') }
    finally { setLoading(false) }
  }

  function openCreate() {
    setForm({ supplier_id: '', note: '', items: [{ ...blankItem }] })
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

  const total = form.items.reduce((s, it) => s + Number(it.unit_price) * Number(it.quantity), 0)

  async function handleCreate(e) {
    e.preventDefault(); setSaving(true); setError('')
    try {
      await api.post('/purchase-orders', {
        supplier_id: Number(form.supplier_id),
        note: form.note,
        items: form.items.map(it => ({
          product_id: Number(it.product_id),
          quantity: Number(it.quantity),
          unit_price: Number(it.unit_price),
        })),
      })
      closeModal(); load()
    } catch (err) {
      setError(err.response?.data?.error || 'Create failed')
    } finally { setSaving(false) }
  }

  async function changeStatus(id, status) {
    setError('')
    try { await api.patch(`/purchase-orders/${id}/status`, { status }); load() }
    catch (err) { setError(err.response?.data?.error || 'Status update failed') }
  }

  async function receiveOrder(id) {
    if (!confirm('Receive this order? Stock will be updated automatically.')) return
    setError('')
    try { await api.post(`/purchase-orders/${id}/receive`); load() }
    catch (err) { setError(err.response?.data?.error || 'Receive failed') }
  }

  async function handleDelete(id) {
    if (!confirm('Delete this order?')) return
    try { await api.delete(`/purchase-orders/${id}`); load() }
    catch (err) { setError(err.response?.data?.error || 'Delete failed') }
  }

  return (
    <div className="page">
      <div className="page-header">
        <div>
          <div className="page-title">Purchase Orders</div>
          <div className="page-desc">Procurement and goods receipt</div>
        </div>
        <button className="btn btn-primary" onClick={openCreate}>+ New PO</button>
      </div>

      {error && !mode && <div className="alert alert-error">{error}</div>}

      {loading ? <div className="loading">Loading…</div> : (
        <div className="table-wrap">
          <table className="data-table">
            <thead>
              <tr>
                <th>PO No</th>
                <th>Supplier</th>
                <th>Status</th>
                <th>Order Date</th>
                <th>Expected</th>
                <th style={{ textAlign: 'right' }}>Total</th>
                <th className="col-actions"></th>
              </tr>
            </thead>
            <tbody>
              {orders.length === 0 ? (
                <tr className="empty-row"><td colSpan={7}>No purchase orders yet</td></tr>
              ) : orders.map(o => (
                <tr key={o.id}>
                  <td className="col-mono">{o.po_no}</td>
                  <td>{o.supplier?.name || '—'}</td>
                  <td><span className={`badge ${statusBadge[o.status] || 'badge-gray'}`}>{o.status}</span></td>
                  <td style={{ color: '#888' }}>{o.order_date ? new Date(o.order_date).toLocaleDateString() : '—'}</td>
                  <td style={{ color: '#888' }}>{o.expected_date ? new Date(o.expected_date).toLocaleDateString() : '—'}</td>
                  <td style={{ textAlign: 'right' }}>{o.total_amount?.toLocaleString(undefined, { minimumFractionDigits: 2 })}</td>
                  <td className="col-actions">
                    <div className="row-actions">
                      <button className="btn btn-ghost btn-sm" onClick={() => openView(o)}>View</button>
                      {o.status === 'draft' && <button className="btn btn-ghost btn-sm" style={{ color: '#1d4ed8' }} onClick={() => changeStatus(o.id, 'sent')}>Send</button>}
                      {(o.status === 'draft' || o.status === 'sent') && (
                        <button className="btn btn-ghost btn-sm" style={{ color: '#15803d' }} onClick={() => receiveOrder(o.id)}>Receive</button>
                      )}
                      {(o.status === 'draft' || o.status === 'sent') && (
                        <button className="btn btn-ghost btn-sm" style={{ color: '#888' }} onClick={() => changeStatus(o.id, 'cancelled')}>Cancel</button>
                      )}
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
        <Modal title="New Purchase Order" onClose={closeModal} large>
          <form onSubmit={handleCreate}>
            <div className="modal-body">
              {error && <div className="alert alert-error">{error}</div>}
              <div className="form-row form-row-2">
                <div>
                  <label className="form-label">Supplier</label>
                  <select className="form-select" value={form.supplier_id} onChange={e => setForm({ ...form, supplier_id: e.target.value })} required>
                    <option value="">Select supplier</option>
                    {suppliers.map(s => <option key={s.id} value={s.id}>{s.name}</option>)}
                  </select>
                </div>
                <div>
                  <label className="form-label">Note</label>
                  <input className="form-input" value={form.note} onChange={e => setForm({ ...form, note: e.target.value })} />
                </div>
              </div>

              <div className="items-section">
                <div className="items-section-title">Items</div>
                <div className="item-row item-row-po" style={{ marginBottom: 4 }}>
                  <div className="item-col-label">Product</div>
                  <div className="item-col-label">Qty</div>
                  <div className="item-col-label">Unit Price</div>
                  <div></div>
                </div>
                {form.items.map((it, i) => (
                  <div key={i} className="item-row item-row-po">
                    <select className="form-select" value={it.product_id} onChange={e => updateItem(i, 'product_id', e.target.value)} required>
                      <option value="">Select product</option>
                      {products.map(p => <option key={p.id} value={p.id}>{p.name}</option>)}
                    </select>
                    <input className="form-input" type="number" min="1" step="0.01" value={it.quantity} onChange={e => updateItem(i, 'quantity', e.target.value)} required />
                    <input className="form-input" type="number" min="0" step="0.01" value={it.unit_price} onChange={e => updateItem(i, 'unit_price', e.target.value)} required />
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
              <button type="submit" className="btn btn-primary" disabled={saving}>{saving ? 'Creating…' : 'Create PO'}</button>
            </div>
          </form>
        </Modal>
      )}

      {mode === 'view' && selected && (
        <Modal title={`PO ${selected.po_no}`} onClose={closeModal} large>
          <div className="modal-body">
            <div className="detail-meta">
              <div>
                <div className="detail-field-label">Supplier</div>
                <div className="detail-field-value">{selected.supplier?.name || '—'}</div>
              </div>
              <div>
                <div className="detail-field-label">Status</div>
                <div className="detail-field-value"><span className={`badge ${statusBadge[selected.status] || 'badge-gray'}`}>{selected.status}</span></div>
              </div>
              <div>
                <div className="detail-field-label">Order Date</div>
                <div className="detail-field-value">{selected.order_date ? new Date(selected.order_date).toLocaleDateString() : '—'}</div>
              </div>
              <div>
                <div className="detail-field-label">Expected Date</div>
                <div className="detail-field-value">{selected.expected_date ? new Date(selected.expected_date).toLocaleDateString() : '—'}</div>
              </div>
              <div>
                <div className="detail-field-label">Total</div>
                <div className="detail-field-value" style={{ fontWeight: 600 }}>{selected.total_amount?.toLocaleString(undefined, { minimumFractionDigits: 2 })}</div>
              </div>
              {selected.note && <div>
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
                    <th style={{ textAlign: 'right' }}>Ordered</th>
                    <th style={{ textAlign: 'right' }}>Received</th>
                    <th style={{ textAlign: 'right' }}>Unit Price</th>
                    <th style={{ textAlign: 'right' }}>Subtotal</th>
                  </tr>
                </thead>
                <tbody>
                  {(selected.items || []).map(it => (
                    <tr key={it.id}>
                      <td>{it.product?.name || `Product #${it.product_id}`}</td>
                      <td style={{ textAlign: 'right' }}>{it.quantity}</td>
                      <td style={{ textAlign: 'right' }}>{it.received_qty}</td>
                      <td style={{ textAlign: 'right' }}>{it.unit_price?.toLocaleString(undefined, { minimumFractionDigits: 2 })}</td>
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
