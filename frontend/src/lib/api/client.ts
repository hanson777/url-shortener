export type ShortenURLResponse = {
  shortUrl: string
  longUrl: string
}

export const shortenUrl = async (url: string) => {
  const response = await fetch('http://localhost:8080/api/shorten', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({ url }),
  })
  return response
}
