import { BrowserRouter, Routes, Route, Navigate } from 'react-router-dom'
import { AuthProvider, useAuth } from './context/AuthContext'
import type { Role } from './types'
import Layout from './components/Layout'
import Login from './pages/Login'
import Dashboard from './pages/Dashboard'
import DepartmentsPage from './pages/hr/DepartmentsPage'
import PositionsPage from './pages/hr/PositionsPage'
import EmployeesPage from './pages/hr/EmployeesPage'
import CategoriesPage from './pages/inventory/CategoriesPage'
import ProductsPage from './pages/inventory/ProductsPage'
import StockPage from './pages/inventory/StockPage'
import CustomersPage from './pages/sales/CustomersPage'
import SalesOrdersPage from './pages/sales/SalesOrdersPage'
import SuppliersPage from './pages/purchase/SuppliersPage'
import PurchaseOrdersPage from './pages/purchase/PurchaseOrdersPage'
import UsersPage from './pages/admin/UsersPage'

function Guard({ children }: { children: React.ReactNode }) {
  const { token } = useAuth()
  return token ? <>{children}</> : <Navigate to="/login" replace />
}

function RoleGuard({ roles, children }: { roles: Role[]; children: React.ReactNode }) {
  const { user } = useAuth()
  if (!user) return <Navigate to="/login" replace />
  if (!roles.includes(user.role)) return <Navigate to="/" replace />
  return <>{children}</>
}

const HR_ROLES: Role[] = ['admin', 'hr']
const INV_ROLES: Role[] = ['admin', 'inventory']
const SALES_ROLES: Role[] = ['admin', 'sales']
const PUR_ROLES: Role[] = ['admin', 'purchase']
const ADMIN_ROLES: Role[] = ['admin']

export default function App() {
  return (
    <AuthProvider>
      <BrowserRouter>
        <Routes>
          <Route path="/login" element={<Login />} />
          <Route path="/" element={<Guard><Layout /></Guard>}>
            <Route index element={<Dashboard />} />

            <Route path="admin/users" element={<RoleGuard roles={ADMIN_ROLES}><UsersPage /></RoleGuard>} />

            <Route path="hr/departments" element={<RoleGuard roles={HR_ROLES}><DepartmentsPage /></RoleGuard>} />
            <Route path="hr/positions"   element={<RoleGuard roles={HR_ROLES}><PositionsPage /></RoleGuard>} />
            <Route path="hr/employees"   element={<RoleGuard roles={HR_ROLES}><EmployeesPage /></RoleGuard>} />

            <Route path="inventory/categories" element={<RoleGuard roles={INV_ROLES}><CategoriesPage /></RoleGuard>} />
            <Route path="inventory/products"   element={<RoleGuard roles={INV_ROLES}><ProductsPage /></RoleGuard>} />
            <Route path="inventory/stock"      element={<RoleGuard roles={INV_ROLES}><StockPage /></RoleGuard>} />

            <Route path="sales/customers" element={<RoleGuard roles={SALES_ROLES}><CustomersPage /></RoleGuard>} />
            <Route path="sales/orders"    element={<RoleGuard roles={SALES_ROLES}><SalesOrdersPage /></RoleGuard>} />

            <Route path="purchase/suppliers" element={<RoleGuard roles={PUR_ROLES}><SuppliersPage /></RoleGuard>} />
            <Route path="purchase/orders"    element={<RoleGuard roles={PUR_ROLES}><PurchaseOrdersPage /></RoleGuard>} />
          </Route>
          <Route path="*" element={<Navigate to="/" replace />} />
        </Routes>
      </BrowserRouter>
    </AuthProvider>
  )
}
