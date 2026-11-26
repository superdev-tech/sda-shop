// แปลงสตางค์เป็นข้อความเงินบาท เช่น 12900 -> ฿129.00
export function formatBaht(satang: number): string {
  return new Intl.NumberFormat('th-TH', { style: 'currency', currency: 'THB' }).format(satang / 100)
}

export function formatDate(iso: string): string {
  return new Date(iso).toLocaleString('th-TH', { dateStyle: 'medium', timeStyle: 'short' })
}
