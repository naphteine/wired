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

	import IconSun from 'virtual:icons/basil/sun-solid';
	import IconMoon from 'virtual:icons/basil/moon-solid';
	import { browser } from '$app/environment';

	export let data;

	let currentThemeDark = false;

	// Theme
	if (browser) {
		currentThemeDark = window.matchMedia('(prefers-color-scheme: dark)').matches;
		updateTheme(currentThemeDark);
	}

	let { supabase, session } = data;
	$: ({ supabase, session } = data);

	$: loading.setNavigate(!!$navigating);

	onMount(() => {
		// Session
		const { data } = supabase.auth.onAuthStateChange((event, _session) => {
			if (_session?.expires_at !== session?.expires_at) {
				invalidate('supabase:auth');
			}
		});

		return () => data.subscription.unsubscribe();
	});

	function updateTheme(isDark: boolean) {
		document.documentElement.classList.remove(isDark ? 'light' : 'dark');
		document.documentElement.classList.add(isDark ? 'dark' : 'light');
		currentThemeDark = isDark;
	}

	function changeTheme() {
		updateTheme(!currentThemeDark);
	}
</script>

<a href="/" class="flex items-end mt-3" on:click|preventDefault={changeTheme}
	><IconSun class="hidden dark:block w-8 h-8" />
	<IconMoon class="block dark:hidden w-8 h-8" /></a
>

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
