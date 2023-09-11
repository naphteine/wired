<script lang="ts">
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
	h1 {
		margin: 0;
		text-align: center;
	}

	ul {
		margin: 0;
		padding: 0;
	}
</style>
