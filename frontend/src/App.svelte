<script lang="ts">
	import { shortenUrl } from './lib/api/client'
	import { normalizeUrl, isValidUrl } from './lib/api/utils'
	import type { ShortenURLResponse } from './lib/api/client';

	let url = ''
	let shortUrl = ''
	let error = ''

	const getShortUrl = async (e: SubmitEvent) => {
		e.preventDefault()
		
		error = '';
    const normalized = normalizeUrl(url);

    if (!isValidUrl(normalized)) {
      error = 'Enter a valid URL';
      return;
    }
		
		try {
			const response = await shortenUrl(normalized)

			if (response.status === 429) {
				error = 'You have reached the rate limit. Try again later'
				return
			}

			if (!response.ok) {
				error = 'Something went wrong'
				return
			}

			const data: ShortenURLResponse = await response.json()
			shortUrl = data.shortUrl
		} catch {
			error = 'Network error'
		}
	}
</script>


<h1>URL Shortener</h1>

<form>
	<input type="text" bind:value={url} placeholder="Enter URL" />
<button type="submit" on:submit={getShortUrl}>Shorten</button>
</form>
{#if shortUrl}
<div>Short URL: {shortUrl}</div>
{/if}
{#if error}
  <p class="error">{error}</p>
{/if}