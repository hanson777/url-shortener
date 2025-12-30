<script lang="ts">
	let url = ''
	let code = ''
	let error = ''

	type ShortenURLResponse = {
		ShortURL: string
		LongURL: string
	}

	const normalizeUrl = (input: string): string => {
    input = input.trim();
    if (!/^https?:\/\//i.test(input)) {
      input = 'https://' + input;
    }
    return input;
  }

	const isValidUrl = (input: string): boolean => {
		try {
			const parsed = new URL(input);
			return /\.[a-z]{2,}$/i.test(parsed.hostname);
		} catch {
			return false;
		}
	}

	const getShortUrl = async () => {
		error = '';
    const normalized = normalizeUrl(url);

    if (!isValidUrl(normalized)) {
      error = 'Enter a valid URL';
      return;
    }

		try {
			const response = await fetch('http://localhost:8080/api/shorten', {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify({ url })
		})

		const data: ShortenURLResponse = await response.json()
		code = data.ShortURL
		} catch (err) {
			console.error(err)
		}
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