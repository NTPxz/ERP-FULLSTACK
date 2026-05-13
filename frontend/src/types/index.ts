export type Role = 'admin' | 'hr' | 'inventory' | 'sales' | 'purchase'

export interface User {
  id: number
  name: string
  email: string
  role: Role
  created_at: string
  updated_at: string
}

export interface Department {
  id: number
  name: string
  description: string
}

export interface Position {
  id: number
  title: string
  department_id: number
  department: Department
  min_salary: number
  max_salary: number
}

export type EmployeeStatus = 'active' | 'inactive'

export interface Employee {
  id: number
  employee_code: string
  name: string
  email: string
  phone: string
  department_id: number
  department: Department
  position_id: number
  position: Position
  salary: number
  hire_date: string
  status: EmployeeStatus
}

export interface Category {
  id: number
  name: string
  description: string
}

export interface Product {
  id: number
  sku: string
  name: string
  description: string
  category_id: number
  category: Category
  unit: string
  cost_price: number
  sale_price: number
  stock_quantity: number
  reorder_level: number
}

export type StockMovementType = 'in' | 'out' | 'adjust'

export interface StockMovement {
  id: number
  product_id: number
  product: Product
  type: StockMovementType
  quantity: number
  reference_type: string
  reference_id: number
  note: string
  created_at: string
}

export interface Customer {
  id: number
  code: string
  name: string
  email: string
  phone: string
  address: string
  tax_id: string
  credit_limit: number
}

export type SalesOrderStatus = 'draft' | 'confirmed' | 'completed' | 'cancelled'

export interface SalesOrderItem {
  id: number
  sales_order_id: number
  product_id: number
  product: Product
  quantity: number
  unit_price: number
  discount: number
  subtotal: number
}

export interface SalesOrder {
  id: number
  order_no: string
  customer_id: number
  customer: Customer
  status: SalesOrderStatus
  order_date: string
  total_amount: number
  note: string
  items: SalesOrderItem[]
}

export interface Supplier {
  id: number
  code: string
  name: string
  email: string
  phone: string
  address: string
  tax_id: string
  payment_terms: string
}

export type PurchaseOrderStatus = 'draft' | 'sent' | 'received' | 'cancelled'

export interface PurchaseOrderItem {
  id: number
  purchase_order_id: number
  product_id: number
  product: Product
  quantity: number
  unit_price: number
  subtotal: number
}

export interface PurchaseOrder {
  id: number
  order_no: string
  supplier_id: number
  supplier: Supplier
  status: PurchaseOrderStatus
  order_date: string
  total_amount: number
  note: string
  items: PurchaseOrderItem[]
}

// Request types
export interface CreateDepartmentInput { name: string; description?: string }
export interface CreatePositionInput { title: string; department_id: number; min_salary?: number; max_salary?: number }
export interface CreateEmployeeInput {
  employee_code: string; name: string; email: string; phone?: string
  department_id: number; position_id: number; salary?: number
  hire_date?: string; status?: EmployeeStatus
}
export interface UpdateEmployeeInput {
  name?: string; email?: string; phone?: string; department_id?: number
  position_id?: number; salary?: number; status?: EmployeeStatus
}
export interface CreateCategoryInput { name: string; description?: string }
export interface CreateProductInput {
  sku: string; name: string; description?: string; category_id: number
  unit?: string; cost_price?: number; sale_price: number; reorder_level?: number
}
export interface UpdateProductInput {
  name?: string; description?: string; category_id?: number
  unit?: string; cost_price?: number; sale_price?: number; reorder_level?: number
}
export interface StockAdjustInput { product_id: number; quantity: number; note?: string }
export interface CreateCustomerInput {
  code: string; name: string; email?: string; phone?: string
  address?: string; tax_id?: string; credit_limit?: number
}
export interface UpdateCustomerInput {
  name?: string; email?: string; phone?: string
  address?: string; tax_id?: string; credit_limit?: number
}
export interface SalesOrderItemInput { product_id: number; quantity: number; unit_price: number; discount?: number }
export interface CreateSalesOrderInput {
  customer_id: number; order_date?: string; note?: string
  items: SalesOrderItemInput[]
}
export interface CreateSupplierInput {
  code: string; name: string; email?: string; phone?: string; address?: string; tax_id?: string
}
export interface UpdateSupplierInput {
  name?: string; email?: string; phone?: string; address?: string; tax_id?: string
}
export interface PurchaseOrderItemInput { product_id: number; quantity: number; unit_price: number }
export interface CreatePurchaseOrderInput {
  supplier_id: number; order_date?: string; note?: string
  items: PurchaseOrderItemInput[]
}
export interface AdminCreateUserInput { name: string; email: string; password: string; role: Role }
export interface AdminUpdateUserInput { name?: string; email?: string; role?: Role }
