import { useState, useEffect } from 'react'
import { usePositions, useDepartments } from '../../hooks/useEmployees'
import Modal from '../../components/Modal'
import type { CreatePositionInput } from '../../types'

const blank: CreatePositionInput = { title: '', department_id: 0, min_salary: 0, max_salary: 0 }

export default function PositionsPage() {
  const { positions, loading, error: loadError, load } = usePositions()
  const { departments, load: loadDeps } = useDepartments()
  const [showModal, setShowModal] = useState(false)
  const [form, setForm] = useState<CreatePositionInput>(blank)
  const [saving, setSaving] = useState(false)
  const [error, setError] = useState('')
  const { create } = usePositions()

  useEffect(() => { void load(); void loadDeps() }, [load, loadDeps])

  async function handleSubmit(e: React.FormEvent) {
    e.preventDefault(); setSaving(true); setError('')
    try {
      await create(form)
      setShowModal(false)
    } catch (err: unknown) {
      setError(err instanceof Error ? err.message : 'Save failed')
    } finally { setSaving(false) }
  }

  return (
    <div className="page">
      <div className="page-header">
        <div>
          <div className="page-title">Positions</div>
          <div className="page-desc">Job titles and salary bands</div>
        </div>
        <button className="btn btn-primary" onClick={() => { setForm(blank); setError(''); setShowModal(true) }}>+ Add Position</button>
      </div>

      {(loadError || error) && !showModal && <div className="alert alert-error">{loadError || error}</div>}

      {loading ? <div className="loading">Loading…</div> : (
        <div className="table-wrap">
          <table className="data-table">
            <thead>
              <tr>
                <th>Title</th><th>Department</th>
                <th style={{ textAlign: 'right' }}>Min Salary</th>
                <th style={{ textAlign: 'right' }}>Max Salary</th>
              </tr>
            </thead>
            <tbody>
              {positions.length === 0
                ? <tr className="empty-row"><td colSpan={4}>No positions yet</td></tr>
                : positions.map(pos => (
                  <tr key={pos.id}>
                    <td style={{ fontWeight: 500 }}>{pos.title}</td>
                    <td style={{ color: '#888' }}>{pos.department?.name || '—'}</td>
                    <td style={{ textAlign: 'right' }}>{pos.min_salary?.toLocaleString()}</td>
                    <td style={{ textAlign: 'right' }}>{pos.max_salary?.toLocaleString()}</td>
                  </tr>
                ))}
            </tbody>
          </table>
        </div>
      )}

      {showModal && (
        <Modal title="New Position" onClose={() => setShowModal(false)}>
          <form onSubmit={handleSubmit}>
            <div className="modal-body">
              {error && <div className="alert alert-error">{error}</div>}
              <div className="form-group">
                <label className="form-label">Title</label>
                <input className="form-input" value={form.title}
                  onChange={e => setForm({ ...form, title: e.target.value })} required autoFocus />
              </div>
              <div className="form-group">
                <label className="form-label">Department</label>
                <select className="form-select" value={form.department_id}
                  onChange={e => setForm({ ...form, department_id: Number(e.target.value) })} required>
                  <option value={0}>Select department</option>
                  {departments.map(d => <option key={d.id} value={d.id}>{d.name}</option>)}
                </select>
              </div>
              <div className="form-row form-row-2">
                <div>
                  <label className="form-label">Min Salary</label>
                  <input className="form-input" type="number"
                    value={form.min_salary ?? 0}
                    onChange={e => setForm({ ...form, min_salary: Number(e.target.value) })} />
                </div>
                <div>
                  <label className="form-label">Max Salary</label>
                  <input className="form-input" type="number"
                    value={form.max_salary ?? 0}
                    onChange={e => setForm({ ...form, max_salary: Number(e.target.value) })} />
                </div>
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
