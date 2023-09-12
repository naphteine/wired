<script lang="ts">
	import { enhance } from '$app/forms';
	import Post from './post.svelte';

	export let data;
	export let form;

	let { session, supabase } = data;
	$: ({ session, supabase } = data);
</script>

<svelte:head>
	<title>Home - Wired</title>
	<meta name="description" content="Connect to Wired" />
</svelte:head>

<h1>LATEST POSTS</h1>

{#if session}
	<h2>CONNECTED TO WIRED</h2>
	<h3>{session.user.email}</h3>

	<form method="post" action="?/createPost" use:enhance>
		<input placeholder="new post" name="post" type="text" />
		<button>Submit</button>
	</form>
{/if}

<ul>
	{#each data.posts as post}
		<Post
			{supabase}
			content={post.name}
			user={post.profiles.username}
			fullName={post.profiles.full_name}
			bind:url={post.profiles.avatar_url}
			date={post.created}
		/>
	{/each}
</ul>

<style>
	h1,
	h2,
	h3 {
		margin: 0;
		text-align: center;
	}

	ul {
		margin: 0;
		padding: 0;
	}

	form {
		background-color: rgb(177, 177, 177);
		max-width: 400px;
		margin: 0 auto;
		display: flex;
		flex-direction: column;
		align-items: center;
		padding: 1rem;
	}

	form input {
		width: 90%;
		height: 3rem;
	}
</style>
