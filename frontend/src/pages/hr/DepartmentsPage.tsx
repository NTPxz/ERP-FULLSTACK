import { useState, useEffect } from 'react'
import { useDepartments } from '../../hooks/useEmployees'
import Modal from '../../components/Modal'
import type { Department, CreateDepartmentInput } from '../../types'

const blank: CreateDepartmentInput = { name: '', description: '' }

export default function DepartmentsPage() {
  const { departments, loading, error: loadError, load, create, update, remove } = useDepartments()
  const [showModal, setShowModal] = useState(false)
  const [editing, setEditing] = useState<Department | null>(null)
  const [form, setForm] = useState<CreateDepartmentInput>(blank)
  const [saving, setSaving] = useState(false)
  const [error, setError] = useState('')

  useEffect(() => { void load() }, [load])

  function openCreate() { setEditing(null); setForm(blank); setError(''); setShowModal(true) }
  function openEdit(dep: Department) {
    setEditing(dep); setForm({ name: dep.name, description: dep.description }); setError(''); setShowModal(true)
  }

  async function handleSubmit(e: React.FormEvent) {
    e.preventDefault(); setSaving(true); setError('')
    try {
      if (editing) await update(editing.id, form)
      else await create(form)
      setShowModal(false)
    } catch (err: unknown) {
      setError(err instanceof Error ? err.message : 'Save failed')
    } finally { setSaving(false) }
  }

  async function handleDelete(id: number) {
    if (!confirm('Delete this department?')) return
    try { await remove(id) }
    catch (err: unknown) { setError(err instanceof Error ? err.message : 'Delete failed') }
  }

  return (
    <div className="page">
      <div className="page-header">
        <div>
          <div className="page-title">Departments</div>
          <div className="page-desc">Manage organizational departments</div>
        </div>
        <button className="btn btn-primary" onClick={openCreate}>+ Add Department</button>
      </div>

      {(loadError || error) && !showModal && <div className="alert alert-error">{loadError || error}</div>}

      {loading ? <div className="loading">Loading…</div> : (
        <div className="table-wrap">
          <table className="data-table">
            <thead>
              <tr><th style={{ width: 48 }}>ID</th><th>Name</th><th>Description</th><th className="col-actions"></th></tr>
            </thead>
            <tbody>
              {departments.length === 0
                ? <tr className="empty-row"><td colSpan={4}>No departments yet</td></tr>
                : departments.map(dep => (
                  <tr key={dep.id}>
                    <td className="col-mono">{dep.id}</td>
                    <td style={{ fontWeight: 500 }}>{dep.name}</td>
                    <td style={{ color: '#888' }}>{dep.description || '—'}</td>
                    <td className="col-actions">
                      <div className="row-actions">
                        <button className="btn btn-ghost btn-sm" onClick={() => openEdit(dep)}>Edit</button>
                        <button className="btn btn-danger btn-sm" onClick={() => handleDelete(dep.id)}>Delete</button>
                      </div>
                    </td>
                  </tr>
                ))}
            </tbody>
          </table>
        </div>
      )}

      {showModal && (
        <Modal title={editing ? 'Edit Department' : 'New Department'} onClose={() => setShowModal(false)}>
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
              <button type="submit" className="btn btn-primary" disabled={saving}>
                {saving ? 'Saving…' : editing ? 'Save' : 'Create'}
              </button>
            </div>
          </form>
        </Modal>
      )}
    </div>
  )
}
