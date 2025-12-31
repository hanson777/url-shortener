<script lang="ts">
	import { shortenUrl } from './lib/api/client'
	import { normalizeUrl, isValidUrl } from './lib/api/utils'

	let url = ''
	let code = ''
	let error = ''

	const getShortUrl = async () => {
		error = '';
    const normalized = normalizeUrl(url);

    if (!isValidUrl(normalized)) {
      error = 'Enter a valid URL';
      return;
    }
		code = await shortenUrl(normalized)
	}
</script>


<form on:submit|preventDefault={getShortUrl}>
	<input type="text" bind:value={url} placeholder="Enter URL" />
<button type="submit">Shorten</button>
</form>
{#if code}
<div>Short URL: {code}</div>
{/if}
{#if error}
  <p class="error">{error}</p>
{/if}