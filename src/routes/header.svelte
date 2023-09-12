<script lang="ts">
	import type { Session } from '@supabase/supabase-js';

	export let session: Session;
</script>

<header>
	<h1>Welcome to the Wired</h1>
	<nav>
		<a href="/">Home</a>
		{#if session}
			<a href="/account">Profile</a>
		{:else}
			<a href="/login">Login</a>
			<a href="/register">Register</a>
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

	h1 {
		margin: 0;
	}

	a {
		background-color: black;
		text-decoration: none;
		font-family: sans-serif;
		color: white;
		padding: 0.5rem;
		border-top-right-radius: 15px;
		border-bottom-left-radius: 15px;
		border: 1px solid black;
		transition: all 0.1s ease-in-out;
	}

	a:hover {
		background-color: #fafffc;
		color: black;
		border: 1px solid black;
	}

	nav {
		margin: 1rem 0;
	}
</style>
