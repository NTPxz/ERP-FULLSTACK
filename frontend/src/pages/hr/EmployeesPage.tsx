import { useState, useEffect } from 'react'
import { useEmployees, useDepartments, usePositions } from '../../hooks/useEmployees'
import Modal from '../../components/Modal'
import type { Employee, CreateEmployeeInput, EmployeeStatus } from '../../types'

const blank: CreateEmployeeInput = {
  employee_code: '', name: '', email: '', phone: '',
  department_id: 0, position_id: 0, salary: 0, status: 'active',
}

export default function EmployeesPage() {
  const { employees, loading, error: loadError, load, create, update, remove } = useEmployees()
  const { departments, load: loadDeps } = useDepartments()
  const { positions, load: loadPos } = usePositions()
  const [showModal, setShowModal] = useState(false)
  const [editing, setEditing] = useState<Employee | null>(null)
  const [form, setForm] = useState<CreateEmployeeInput>(blank)
  const [saving, setSaving] = useState(false)
  const [error, setError] = useState('')

  useEffect(() => { void load(); void loadDeps(); void loadPos() }, [load, loadDeps, loadPos])

  function openCreate() { setEditing(null); setForm(blank); setError(''); setShowModal(true) }
  function openEdit(emp: Employee) {
    setEditing(emp)
    setForm({
      employee_code: emp.employee_code, name: emp.name, email: emp.email,
      phone: emp.phone || '', department_id: emp.department_id,
      position_id: emp.position_id, salary: emp.salary, status: emp.status,
    })
    setError(''); setShowModal(true)
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
    if (!confirm('Delete this employee?')) return
    try { await remove(id) }
    catch (err: unknown) { setError(err instanceof Error ? err.message : 'Delete failed') }
  }

  return (
    <div className="page">
      <div className="page-header">
        <div>
          <div className="page-title">Employees</div>
          <div className="page-desc">All staff records</div>
        </div>
        <button className="btn btn-primary" onClick={openCreate}>+ Add Employee</button>
      </div>

      {(loadError || error) && !showModal && <div className="alert alert-error">{loadError || error}</div>}

      {loading ? <div className="loading">Loading…</div> : (
        <div className="table-wrap">
          <table className="data-table">
            <thead>
              <tr>
                <th>Code</th><th>Name</th><th>Email</th>
                <th>Department</th><th>Position</th>
                <th style={{ textAlign: 'right' }}>Salary</th>
                <th>Status</th><th className="col-actions"></th>
              </tr>
            </thead>
            <tbody>
              {employees.length === 0
                ? <tr className="empty-row"><td colSpan={8}>No employees yet</td></tr>
                : employees.map(emp => (
                  <tr key={emp.id}>
                    <td className="col-mono">{emp.employee_code}</td>
                    <td style={{ fontWeight: 500 }}>{emp.name}</td>
                    <td style={{ color: '#888' }}>{emp.email}</td>
                    <td>{emp.department?.name || '—'}</td>
                    <td>{emp.position?.title || '—'}</td>
                    <td style={{ textAlign: 'right' }}>{emp.salary?.toLocaleString()}</td>
                    <td>
                      <span className={`badge ${emp.status === 'active' ? 'badge-green' : 'badge-gray'}`}>{emp.status}</span>
                    </td>
                    <td className="col-actions">
                      <div className="row-actions">
                        <button className="btn btn-ghost btn-sm" onClick={() => openEdit(emp)}>Edit</button>
                        <button className="btn btn-danger btn-sm" onClick={() => handleDelete(emp.id)}>Delete</button>
                      </div>
                    </td>
                  </tr>
                ))}
            </tbody>
          </table>
        </div>
      )}

      {showModal && (
        <Modal title={editing ? 'Edit Employee' : 'New Employee'} onClose={() => setShowModal(false)}>
          <form onSubmit={handleSubmit}>
            <div className="modal-body">
              {error && <div className="alert alert-error">{error}</div>}
              <div className="form-row form-row-2">
                <div>
                  <label className="form-label">Employee Code</label>
                  <input className="form-input" value={form.employee_code}
                    onChange={e => setForm({ ...form, employee_code: e.target.value })} required autoFocus />
                </div>
                <div>
                  <label className="form-label">Status</label>
                  <select className="form-select" value={form.status}
                    onChange={e => setForm({ ...form, status: e.target.value as EmployeeStatus })}>
                    <option value="active">Active</option>
                    <option value="inactive">Inactive</option>
                  </select>
                </div>
              </div>
              <div className="form-group">
                <label className="form-label">Full Name</label>
                <input className="form-input" value={form.name}
                  onChange={e => setForm({ ...form, name: e.target.value })} required />
              </div>
              <div className="form-row form-row-2">
                <div>
                  <label className="form-label">Email</label>
                  <input className="form-input" type="email" value={form.email}
                    onChange={e => setForm({ ...form, email: e.target.value })} required />
                </div>
                <div>
                  <label className="form-label">Phone</label>
                  <input className="form-input" value={form.phone ?? ''}
                    onChange={e => setForm({ ...form, phone: e.target.value })} />
                </div>
              </div>
              <div className="form-row form-row-2">
                <div>
                  <label className="form-label">Department</label>
                  <select className="form-select" value={form.department_id}
                    onChange={e => setForm({ ...form, department_id: Number(e.target.value) })} required>
                    <option value={0}>Select</option>
                    {departments.map(d => <option key={d.id} value={d.id}>{d.name}</option>)}
                  </select>
                </div>
                <div>
                  <label className="form-label">Position</label>
                  <select className="form-select" value={form.position_id}
                    onChange={e => setForm({ ...form, position_id: Number(e.target.value) })} required>
                    <option value={0}>Select</option>
                    {positions.map(p => <option key={p.id} value={p.id}>{p.title}</option>)}
                  </select>
                </div>
              </div>
              <div className="form-group">
                <label className="form-label">Salary</label>
                <input className="form-input" type="number" value={form.salary ?? 0}
                  onChange={e => setForm({ ...form, salary: Number(e.target.value) })} />
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
