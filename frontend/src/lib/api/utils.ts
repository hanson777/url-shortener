export const normalizeUrl = (input: string): string => {
  input = input.trim()
  if (!input.startsWith('http://') && !input.startsWith('https://')) {
    input = 'https://' + input
  }
  return input
}

export const isValidUrl = (input: string): boolean => {
  try {
    const parsed = new URL(input)
    return /\.[a-z]{2,}$/i.test(parsed.hostname)
  } catch {
    return false
  }
}
