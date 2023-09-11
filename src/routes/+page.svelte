<script>
	import Header from '../components/header.svelte';

	export let data;
	const theName = 'DOCTOR';
</script>

<svelte:head>
	<title>Home - Wired</title>
	<meta name="description" content="Svelte demo app" />
</svelte:head>

<Header />

<h1>LATEST POSTS</h1>
<ul>
	{#each data.posts as post}
		<li>{post.name} - {post.profiles.username} - {post.created}</li>
	{/each}
</ul>
