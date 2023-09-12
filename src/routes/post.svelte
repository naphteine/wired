<script lang="ts">
	import type { SupabaseClient } from '@supabase/supabase-js';
	import { createEventDispatcher } from 'svelte';

	export let content = '';
	export let user = '';
	export let date = '';
	export let fullName = '';

	export let size = 2;
	export let url: string;
	export let supabase: SupabaseClient;

	const objDate = new Date(date);
	const formattedDate = objDate.toLocaleDateString('tr-TR', {
		day: '2-digit',
		month: '2-digit',
		year: 'numeric',
		hour: '2-digit',
		minute: '2-digit'
	});

	let avatarUrl: string | null = null;
	let files: FileList;

	const dispatch = createEventDispatcher();

	const downloadImage = async (path: string) => {
		try {
			const { data, error } = await supabase.storage.from('avatars').download(path);

			if (error) {
				throw error;
			}

			const url = URL.createObjectURL(data);
			avatarUrl = url;
		} catch (error) {
			if (error instanceof Error) {
				console.log('Error downloading image: ', error.message);
			}
		}
	};

	$: if (url) downloadImage(url);
</script>

<article>
	{#if avatarUrl}
		<img
			src={avatarUrl}
			alt={avatarUrl ? 'Avatar' : 'No image'}
			class="avatar image"
			style="height: {size}em; width: {size}em; border-radius: 100px;"
		/>
	{:else}
		<div class="avatar no-image" style="height: 0em; width: 0em;" />
	{/if}
	<a href="/user/{user}"><b>{fullName}</b> @{user}</a>
	<em>{formattedDate}</em>
	<h2>{content}</h2>
</article>

<style>
	article {
		width: 80vw;
		max-width: 600px;
		margin: 1rem auto;
		padding: 1rem;
		background-color: #c5d1d6;
	}

	h2 {
		margin: 0;
		font-family: sans-serif;
	}

	a {
		text-decoration: none;
		color: blue;
	}
</style>
