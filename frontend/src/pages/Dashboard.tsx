import { useState, useEffect, type ReactNode } from 'react'
import { apiGet } from '../api/client'
import { useAuth } from '../context/AuthContext'
import type { Employee, Product, Customer, SalesOrder, PurchaseOrder } from '../types'

async function tryGet<T>(url: string): Promise<T[]> {
  try { return await apiGet<T[]>(url) }
  catch { return [] }
}

interface Stats { employees: number; products: number; customers: number; openPO: number }

function StatIcon({ children, bg, color }: { children: ReactNode; bg: string; color: string }) {
  return (
    <div className="stat-icon" style={{ background: bg }}>
      <svg width={18} height={18} viewBox="0 0 24 24" fill="none" stroke={color} strokeWidth="2" strokeLinecap="round" strokeLinejoin="round">
        {children}
      </svg>
    </div>
  )
}

const soStatusBadge: Record<string, string> = {
  draft: 'badge-gray', confirmed: 'badge-blue', completed: 'badge-green', cancelled: 'badge-red',
}
const poStatusBadge: Record<string, string> = {
  draft: 'badge-gray', sent: 'badge-blue', received: 'badge-green', cancelled: 'badge-red',
}

export default function Dashboard() {
  const { user } = useAuth()
  const role = user?.role ?? ''

  const [stats, setStats] = useState<Stats>({ employees: 0, products: 0, customers: 0, openPO: 0 })
  const [recentSO, setRecentSO] = useState<SalesOrder[]>([])
  const [recentPO, setRecentPO] = useState<PurchaseOrder[]>([])
  const [loading, setLoading] = useState(true)

  useEffect(() => {
    async function fetchData() {
      const [emps, prods, custs, pos, sos] = await Promise.all([
        tryGet<Employee>('/employees'),
        tryGet<Product>('/products'),
        tryGet<Customer>('/customers'),
        tryGet<PurchaseOrder>('/purchase-orders'),
        tryGet<SalesOrder>('/sales-orders'),
      ])

      const openPO = pos.filter(o => o.status === 'draft' || o.status === 'sent').length
      setStats({ employees: emps.length, products: prods.length, customers: custs.length, openPO })
      setRecentSO(sos.slice(0, 6))
      setRecentPO(pos.filter(o => o.status === 'draft' || o.status === 'sent').slice(0, 6))
      setLoading(false)
    }
    void fetchData()
  }, [])

  if (loading) return <div className="page"><div className="loading">Loading dashboard…</div></div>

  const isHR    = ['admin', 'hr'].includes(role)
  const isInv   = ['admin', 'inventory'].includes(role)
  const isSales = ['admin', 'sales'].includes(role)
  const isPur   = ['admin', 'purchase'].includes(role)

  const greeting = (() => {
    const h = new Date().getHours()
    if (h < 12) return 'Good morning'
    if (h < 17) return 'Good afternoon'
    return 'Good evening'
  })()

  return (
    <div className="page">
      <div className="page-header">
        <div>
          <div className="page-title">Dashboard</div>
          <div className="page-desc">{greeting}, {user?.name ?? 'User'} — here's what's happening today.</div>
        </div>
      </div>

      <div className="stats-grid">
        {isHR && (
          <div className="stat-card">
            <StatIcon bg="#eef2ff" color="#6366f1">
              <path d="M17 21v-2a4 4 0 00-4-4H5a4 4 0 00-4 4v2" />
              <circle cx="9" cy="7" r="4" />
              <path d="M23 21v-2a4 4 0 00-3-3.87M16 3.13a4 4 0 010 7.75" />
            </StatIcon>
            <div className="stat-label">Employees</div>
            <div className="stat-value">{stats.employees}</div>
            <div className="stat-hint">Total headcount</div>
          </div>
        )}
        {isInv && (
          <div className="stat-card">
            <StatIcon bg="#ecfdf5" color="#10b981">
              <path d="M21 16V8a2 2 0 00-1-1.73l-7-4a2 2 0 00-2 0l-7 4A2 2 0 003 8v8a2 2 0 001 1.73l7 4a2 2 0 002 0l7-4A2 2 0 0021 16z" />
              <polyline points="3.27 6.96 12 12.01 20.73 6.96" />
              <line x1="12" y1="22.08" x2="12" y2="12" />
            </StatIcon>
            <div className="stat-label">Products</div>
            <div className="stat-value">{stats.products}</div>
            <div className="stat-hint">In catalog</div>
          </div>
        )}
        {isSales && (
          <div className="stat-card">
            <StatIcon bg="#fffbeb" color="#f59e0b">
              <path d="M17 21v-2a4 4 0 00-4-4H5a4 4 0 00-4 4v2" />
              <circle cx="8.5" cy="7" r="4" />
              <line x1="20" y1="8" x2="20" y2="14" />
              <line x1="23" y1="11" x2="17" y2="11" />
            </StatIcon>
            <div className="stat-label">Customers</div>
            <div className="stat-value">{stats.customers}</div>
            <div className="stat-hint">Total accounts</div>
          </div>
        )}
        {isPur && (
          <div className="stat-card">
            <StatIcon bg="#eff6ff" color="#3b82f6">
              <circle cx="9" cy="21" r="1" />
              <circle cx="20" cy="21" r="1" />
              <path d="M1 1h4l2.68 13.39a2 2 0 002 1.61h9.72a2 2 0 002-1.61L23 6H6" />
            </StatIcon>
            <div className="stat-label">Open POs</div>
            <div className="stat-value">{stats.openPO}</div>
            <div className="stat-hint">Draft or sent</div>
          </div>
        )}
      </div>

      {isSales && (
        <>
          <div className="section-title" style={{ marginTop: 8 }}>Recent Sales Orders</div>
          <div className="table-wrap" style={{ marginBottom: 32 }}>
            <table className="data-table">
              <thead>
                <tr><th>Order No</th><th>Customer</th><th>Status</th><th>Date</th><th style={{ textAlign: 'right' }}>Total</th></tr>
              </thead>
              <tbody>
                {recentSO.length === 0
                  ? <tr className="empty-row"><td colSpan={5}>No sales orders yet</td></tr>
                  : recentSO.map(o => (
                    <tr key={o.id}>
                      <td className="col-mono">{o.order_no}</td>
                      <td style={{ fontWeight: 500 }}>{o.customer?.name ?? '—'}</td>
                      <td><span className={`badge ${soStatusBadge[o.status] ?? 'badge-gray'}`}>{o.status}</span></td>
                      <td className="text-muted">{o.order_date ? new Date(o.order_date).toLocaleDateString() : '—'}</td>
                      <td style={{ textAlign: 'right', fontWeight: 600 }}>
                        {o.total_amount?.toLocaleString(undefined, { minimumFractionDigits: 2 })}
                      </td>
                    </tr>
                  ))}
              </tbody>
            </table>
          </div>
        </>
      )}

      {isPur && recentPO.length > 0 && (
        <>
          <div className="section-title">Open Purchase Orders</div>
          <div className="table-wrap">
            <table className="data-table">
              <thead>
                <tr><th>PO No</th><th>Supplier</th><th>Status</th><th>Date</th><th style={{ textAlign: 'right' }}>Total</th></tr>
              </thead>
              <tbody>
                {recentPO.map(o => (
                  <tr key={o.id}>
                    <td className="col-mono">{o.order_no}</td>
                    <td style={{ fontWeight: 500 }}>{o.supplier?.name ?? '—'}</td>
                    <td><span className={`badge ${poStatusBadge[o.status] ?? 'badge-gray'}`}>{o.status}</span></td>
                    <td className="text-muted">{o.order_date ? new Date(o.order_date).toLocaleDateString() : '—'}</td>
                    <td style={{ textAlign: 'right', fontWeight: 600 }}>
                      {o.total_amount?.toLocaleString(undefined, { minimumFractionDigits: 2 })}
                    </td>
                  </tr>
                ))}
              </tbody>
            </table>
          </div>
        </>
      )}
    </div>
  )
}
