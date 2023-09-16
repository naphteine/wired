<!-- src/routes/+layout.svelte -->
<script lang="ts">
	import './styles.css';
	import { invalidate } from '$app/navigation';
	import { onMount } from 'svelte';
	import { navigating } from '$app/stores';
	import { loading } from '$lib/loading';
	import Loading from '$lib/Loading.svelte';
	import Header from './header.svelte';

	import SvelteLogo from 'virtual:icons/logos/svelte-icon';
	import VercelLogo from 'virtual:icons/logos/vercel-icon';

	export let data;

	let { supabase, session } = data;
	$: ({ supabase, session } = data);

	$: loading.setNavigate(!!$navigating);

	onMount(() => {
		const { data } = supabase.auth.onAuthStateChange((event, _session) => {
			if (_session?.expires_at !== session?.expires_at) {
				invalidate('supabase:auth');
			}
		});

		return () => data.subscription.unsubscribe();
	});
</script>

<Header {session} />

<main>
	<slot />
</main>

<footer class="flex my-20 items-center justify-center">
	<SvelteLogo />
	<VercelLogo />
	Made with *love* <a class="hover:underline" href="https://gokaygultekin.dev">Gökay Gültekin</a>
</footer>

<Loading />
