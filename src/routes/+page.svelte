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

	<form method="post" action="?/createPost" class="bg-slate-300" use:enhance>
		<input
			class="rounded bg-slate-600 p-3 text-slate-100"
			placeholder="new post"
			name="post"
			type="text"
			autocomplete="off"
		/>
		<button
			class="bg-slate-600 p-2 rounded my-3 text-slate-100 hover:bg-green-600 hover:text-green-100 transition"
			>Submit</button
		>
	</form>
{/if}

<ul>
	{#each data.posts as post}
		<Post
			{supabase}
			{session}
			postId={post.id}
			content={post.name}
			user={post.profiles.username}
			fullName={post.profiles.full_name}
			bind:url={post.profiles.avatar_url}
			date={post.created}
			likes={post.likes}
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
		border-radius: 20px;
		max-width: 400px;
		margin: 1rem auto;
		display: flex;
		flex-direction: column;
		align-items: center;
		padding: 1rem;
	}

	form input {
		width: 90%;
		height: 3rem;
		border: none;
	}
</style>
