import { NavLink, useNavigate } from 'react-router-dom'
import { useAuth } from '../context/AuthContext'

function Icon({ children }) {
  return (
    <svg
      className="sidebar-link-icon"
      width={16} height={16}
      viewBox="0 0 24 24"
      fill="none"
      stroke="currentColor"
      strokeWidth="1.75"
      strokeLinecap="round"
      strokeLinejoin="round"
    >
      {children}
    </svg>
  )
}

const NavIcon = {
  home: <Icon><path d="M3 9l9-7 9 7v11a2 2 0 01-2 2H5a2 2 0 01-2-2z" /><polyline points="9 22 9 12 15 12 15 22" /></Icon>,
  shield: <Icon><path d="M12 22s8-4 8-10V5l-8-3-8 3v7c0 6 8 10 8 10z" /></Icon>,
  users: <Icon><path d="M17 21v-2a4 4 0 00-4-4H5a4 4 0 00-4 4v2" /><circle cx="9" cy="7" r="4" /><path d="M23 21v-2a4 4 0 00-3-3.87M16 3.13a4 4 0 010 7.75" /></Icon>,
  building: <Icon><rect x="3" y="9" width="18" height="12" rx="2" /><path d="M9 22V12h6v10M3 9l9-6 9 6" /></Icon>,
  briefcase: <Icon><rect x="2" y="7" width="20" height="14" rx="2" ry="2" /><path d="M16 21V5a2 2 0 00-2-2h-4a2 2 0 00-2 2v16" /></Icon>,
  person: <Icon><path d="M20 21v-2a4 4 0 00-4-4H8a4 4 0 00-4 4v2" /><circle cx="12" cy="7" r="4" /></Icon>,
  box: <Icon><path d="M21 16V8a2 2 0 00-1-1.73l-7-4a2 2 0 00-2 0l-7 4A2 2 0 003 8v8a2 2 0 001 1.73l7 4a2 2 0 002 0l7-4A2 2 0 0021 16z" /><polyline points="3.27 6.96 12 12.01 20.73 6.96" /><line x1="12" y1="22.08" x2="12" y2="12" /></Icon>,
  tag: <Icon><path d="M20.59 13.41l-7.17 7.17a2 2 0 01-2.83 0L2 12V2h10l8.59 8.59a2 2 0 010 2.82z" /><line x1="7" y1="7" x2="7.01" y2="7" /></Icon>,
  sliders: <Icon><line x1="4" y1="21" x2="4" y2="14" /><line x1="4" y1="10" x2="4" y2="3" /><line x1="12" y1="21" x2="12" y2="12" /><line x1="12" y1="8" x2="12" y2="3" /><line x1="20" y1="21" x2="20" y2="16" /><line x1="20" y1="12" x2="20" y2="3" /><line x1="1" y1="14" x2="7" y2="14" /><line x1="9" y1="8" x2="15" y2="8" /><line x1="17" y1="16" x2="23" y2="16" /></Icon>,
  adduser: <Icon><path d="M16 21v-2a4 4 0 00-4-4H5a4 4 0 00-4 4v2" /><circle cx="8.5" cy="7" r="4" /><line x1="20" y1="8" x2="20" y2="14" /><line x1="23" y1="11" x2="17" y2="11" /></Icon>,
  doc: <Icon><path d="M14 2H6a2 2 0 00-2 2v16a2 2 0 002 2h12a2 2 0 002-2V8z" /><polyline points="14 2 14 8 20 8" /><line x1="16" y1="13" x2="8" y2="13" /><line x1="16" y1="17" x2="8" y2="17" /></Icon>,
  truck: <Icon><rect x="1" y="3" width="15" height="13" /><polygon points="16 8 20 8 23 11 23 16 16 16 16 8" /><circle cx="5.5" cy="18.5" r="2.5" /><circle cx="18.5" cy="18.5" r="2.5" /></Icon>,
  cart: <Icon><circle cx="9" cy="21" r="1" /><circle cx="20" cy="21" r="1" /><path d="M1 1h4l2.68 13.39a2 2 0 002 1.61h9.72a2 2 0 002-1.61L23 6H6" /></Icon>,
  logout: <Icon><path d="M9 21H5a2 2 0 01-2-2V5a2 2 0 012-2h4" /><polyline points="16 17 21 12 16 7" /><line x1="21" y1="12" x2="9" y2="12" /></Icon>,
}

const allSections = [
  {
    key: 'admin',
    label: 'Admin',
    roles: ['admin'],
    links: [
      { to: '/admin/users', label: 'Users & Roles', icon: NavIcon.shield },
    ],
  },
  {
    key: 'hr',
    label: 'Human Resources',
    roles: ['admin', 'hr'],
    links: [
      { to: '/hr/departments', label: 'Departments', icon: NavIcon.building },
      { to: '/hr/positions',   label: 'Positions',   icon: NavIcon.briefcase },
      { to: '/hr/employees',   label: 'Employees',   icon: NavIcon.person },
    ],
  },
  {
    key: 'inventory',
    label: 'Inventory',
    roles: ['admin', 'inventory'],
    links: [
      { to: '/inventory/products',   label: 'Products',       icon: NavIcon.box },
      { to: '/inventory/categories', label: 'Categories',     icon: NavIcon.tag },
      { to: '/inventory/stock',      label: 'Stock Adjust',   icon: NavIcon.sliders },
    ],
  },
  {
    key: 'sales',
    label: 'Sales',
    roles: ['admin', 'sales'],
    links: [
      { to: '/sales/customers', label: 'Customers', icon: NavIcon.adduser },
      { to: '/sales/orders',    label: 'Orders',    icon: NavIcon.doc },
    ],
  },
  {
    key: 'purchase',
    label: 'Purchase',
    roles: ['admin', 'purchase'],
    links: [
      { to: '/purchase/suppliers', label: 'Suppliers', icon: NavIcon.truck },
      { to: '/purchase/orders',    label: 'Orders',    icon: NavIcon.cart },
    ],
  },
]

const avatarColors = {
  admin:     '#6366f1',
  hr:        '#0369a1',
  inventory: '#059669',
  sales:     '#d97706',
  purchase:  '#be185d',
}

function getInitials(name, email) {
  const src = name || email || '?'
  return src.split(' ').map(w => w[0]).join('').toUpperCase().slice(0, 2)
}

export default function Sidebar() {
  const { user, logout } = useAuth()
  const navigate = useNavigate()
  const role = user?.role || ''

  const visibleSections = allSections.filter(s => s.roles.includes(role))

  function handleLogout() {
    logout()
    navigate('/login')
  }

  const initials = getInitials(user?.name, user?.email)
  const avatarColor = avatarColors[role] || '#475569'

  return (
    <nav className="sidebar">
      <div className="sidebar-logo">
        <div className="sidebar-logo-icon">
          <svg width={16} height={16} viewBox="0 0 24 24" fill="none" stroke="#fff" strokeWidth="2.2" strokeLinecap="round" strokeLinejoin="round">
            <rect x="2" y="3" width="20" height="14" rx="2" />
            <line x1="8" y1="21" x2="16" y2="21" />
            <line x1="12" y1="17" x2="12" y2="21" />
          </svg>
        </div>
        <span className="sidebar-logo-text">ERP System</span>
      </div>

      <div className="sidebar-nav">
        <NavLink to="/" end className={({ isActive }) => 'sidebar-link-top' + (isActive ? ' active' : '')}>
          {NavIcon.home}
          Dashboard
        </NavLink>

        {visibleSections.map(sec => (
          <div key={sec.key}>
            <div className="sidebar-section-label">{sec.label}</div>
            {sec.links.map(link => (
              <NavLink
                key={link.to}
                to={link.to}
                className={({ isActive }) => 'sidebar-link' + (isActive ? ' active' : '')}
              >
                {link.icon}
                {link.label}
              </NavLink>
            ))}
          </div>
        ))}
      </div>

      <div className="sidebar-footer">
        <div className="sidebar-footer-user">
          <div className="sidebar-avatar" style={{ background: avatarColor }}>
            {initials}
          </div>
          <div className="sidebar-footer-info">
            <div className="sidebar-footer-name">{user?.name || user?.email || 'User'}</div>
            <div className="sidebar-footer-email">{role}</div>
          </div>
        </div>
        <button
          className="btn btn-ghost btn-sm"
          style={{ color: '#475569', paddingLeft: 0, gap: 6 }}
          onClick={handleLogout}
        >
          {NavIcon.logout}
          Sign out
        </button>
      </div>
    </nav>
  )
}
