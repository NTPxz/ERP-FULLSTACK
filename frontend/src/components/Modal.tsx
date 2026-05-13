import { useEffect, type ReactNode } from 'react'

interface ModalProps {
  title: string
  onClose: () => void
  large?: boolean
  children: ReactNode
}

export default function Modal({ title, onClose, large, children }: ModalProps) {
  useEffect(() => {
    function onKey(e: KeyboardEvent) { if (e.key === 'Escape') onClose() }
    document.addEventListener('keydown', onKey)
    return () => document.removeEventListener('keydown', onKey)
  }, [onClose])

  return (
    <div className="modal-backdrop" onClick={e => e.target === e.currentTarget && onClose()}>
      <div className={`modal-box${large ? ' modal-lg' : ''}`}>
        <div className="modal-header">
          <span className="modal-title">{title}</span>
          <button className="modal-close" onClick={onClose}>×</button>
        </div>
        {children}
      </div>
    </div>
  )
}
