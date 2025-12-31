type ShortenURLResponse = {
  ShortURL: string
  LongURL: string
}

export const shortenUrl = async (url: string): Promise<string> => {
  const response = await fetch('http://localhost:8080/api/shorten', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify({ url })
  })

  const data: ShortenURLResponse = await response.json()

  return data.ShortURL
}