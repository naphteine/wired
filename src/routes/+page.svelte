<script>
	import Header from '../components/header.svelte';
	import Post from './post.svelte';

	export let data;
	const theName = 'DOCTOR';
</script>

<svelte:head>
	<title>Home - Wired</title>
	<meta name="description" content="Connect to Wired" />
</svelte:head>

<h1>LATEST POSTS</h1>
<ul>
	{#each data.posts as post}
		<Post
			content={post.name}
			user={post.profiles.username}
			fullName={post.profiles.full_name}
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
