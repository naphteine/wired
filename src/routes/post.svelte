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

<article class="bg-slate-300 rounded-xl p-5 mx-auto my-3">
	<header class="flex items-center">
		{#if avatarUrl}
			<img src={avatarUrl} alt={avatarUrl ? 'Avatar' : 'No image'} class="h-20 w-20 rounded-full" />
		{:else}
			<div class="bg-slate-400 w-20 h-20 rounded-full" />
		{/if}
		<div class="flex flex-col mx-3">
			<a href="/user/{user}" class="text-slate-800 hover:underline"><b>{fullName}</b> @{user}</a>
			<em>{formattedDate}</em>
		</div>
	</header>
	<p class="my-4">{content}</p>
</article>

<style>
	article {
		width: 80vw;
		max-width: 600px;
	}
</style>
