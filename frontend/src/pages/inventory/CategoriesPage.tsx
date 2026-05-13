import { useState, useEffect } from 'react'
import { useCategories } from '../../hooks/useInventory'
import Modal from '../../components/Modal'
import type { CreateCategoryInput } from '../../types'

const blank: CreateCategoryInput = { name: '', description: '' }

export default function CategoriesPage() {
  const { categories, loading, error: loadError, load, create } = useCategories()
  const [showModal, setShowModal] = useState(false)
  const [form, setForm] = useState<CreateCategoryInput>(blank)
  const [saving, setSaving] = useState(false)
  const [error, setError] = useState('')

  useEffect(() => { void load() }, [load])

  async function handleSubmit(e: React.FormEvent) {
    e.preventDefault(); setSaving(true); setError('')
    try { await create(form); setShowModal(false) }
    catch (err: unknown) { setError(err instanceof Error ? err.message : 'Save failed') }
    finally { setSaving(false) }
  }

  return (
    <div className="page">
      <div className="page-header">
        <div>
          <div className="page-title">Categories</div>
          <div className="page-desc">Product categories</div>
        </div>
        <button className="btn btn-primary" onClick={() => { setForm(blank); setError(''); setShowModal(true) }}>+ Add Category</button>
      </div>

      {(loadError || error) && !showModal && <div className="alert alert-error">{loadError || error}</div>}

      {loading ? <div className="loading">Loading…</div> : (
        <div className="table-wrap">
          <table className="data-table">
            <thead>
              <tr><th style={{ width: 48 }}>ID</th><th>Name</th><th>Description</th></tr>
            </thead>
            <tbody>
              {categories.length === 0
                ? <tr className="empty-row"><td colSpan={3}>No categories yet</td></tr>
                : categories.map(cat => (
                  <tr key={cat.id}>
                    <td className="col-mono">{cat.id}</td>
                    <td style={{ fontWeight: 500 }}>{cat.name}</td>
                    <td style={{ color: '#888' }}>{cat.description || '—'}</td>
                  </tr>
                ))}
            </tbody>
          </table>
        </div>
      )}

      {showModal && (
        <Modal title="New Category" onClose={() => setShowModal(false)}>
          <form onSubmit={handleSubmit}>
            <div className="modal-body">
              {error && <div className="alert alert-error">{error}</div>}
              <div className="form-group">
                <label className="form-label">Name</label>
                <input className="form-input" value={form.name}
                  onChange={e => setForm({ ...form, name: e.target.value })} required autoFocus />
              </div>
              <div className="form-group">
                <label className="form-label">Description</label>
                <textarea className="form-textarea" value={form.description ?? ''}
                  onChange={e => setForm({ ...form, description: e.target.value })} />
              </div>
            </div>
            <div className="modal-footer">
              <button type="button" className="btn btn-secondary" onClick={() => setShowModal(false)}>Cancel</button>
              <button type="submit" className="btn btn-primary" disabled={saving}>{saving ? 'Saving…' : 'Create'}</button>
            </div>
          </form>
        </Modal>
      )}
    </div>
  )
}
