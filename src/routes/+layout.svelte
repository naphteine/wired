<!-- src/routes/+layout.svelte -->
<script lang="ts">
	import './styles.css';
	import { invalidate } from '$app/navigation';
	import { onMount } from 'svelte';
	import { navigating } from '$app/stores';
	import { loading } from '$lib/loading';
	import Loading from '$lib/Loading.svelte';
	import Header from './header.svelte';

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
<slot />

<Loading />
