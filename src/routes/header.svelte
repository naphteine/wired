<script lang="ts">
	import type { Session } from '@supabase/supabase-js';

	export let session: Session;
</script>

<header>
	<img class="dark:hidden" src="/logo.svg" alt="killme" />
	<img class="hidden dark:block" src="/lightlogo.svg" alt="killme" />
	<nav>
		<a
			class="bg-slate-600 p-2 text-slate-100 rounded hover:text-blue-100 hover:bg-blue-600 transition"
			href="/"
		>
			Home
		</a>
		{#if session}
			<a
				class="bg-slate-600 p-2 text-slate-100 rounded hover:text-blue-100 hover:bg-blue-600 transition"
				href="/account">Profile</a
			>
		{:else}
			<a
				class="bg-slate-600 p-2 text-slate-100 rounded hover:text-blue-100 hover:bg-blue-600 transition"
				href="/login">Login</a
			>
			<a
				class="bg-slate-600 p-2 text-slate-100 rounded hover:text-blue-100 hover:bg-blue-600 transition"
				href="/register">Register</a
			>
		{/if}
	</nav>
</header>

<style>
	header {
		padding: 2rem 0;
		display: flex;
		flex-direction: column;
		align-items: center;
	}

	nav {
		margin: 1rem 0;
	}

	img {
		pointer-events: none;
		user-select: none;
	}
</style>
