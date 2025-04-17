export function fromISO(isoStr) {
  if (!isoStr) return ''
  const date = new Date(isoStr)
  const year = date.getFullYear().toString().padStart(4, '0')
  const month = (date.getMonth() + 1).toString().padStart(2, '0')
  const day = date.getDate().toString().padStart(2, '0')
  const hours = date.getHours().toString().padStart(2, '0')
  const minutes = date.getMinutes().toString().padStart(2, '0')
  return `${year}-${month}-${day}T${hours}:${minutes}`
}

export function toISO(localDateTimeStr) {
  if (!localDateTimeStr) return null
  const date = new Date(localDateTimeStr)
  return date.toISOString()
}
