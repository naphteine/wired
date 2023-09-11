<script>
	import Header from '../components/header.svelte';

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
		<li>{post.name} - {post.profiles.username} - {post.created}</li>
	{/each}
</ul>

<style>
	h1 {
		margin: 0;
		text-align: center;
	}
</style>
