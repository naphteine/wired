<script>
	import Avatar from './account/Avatar.svelte';

	export let content = '';
	export let user = '';
	export let date = '';
	export let fullName = '';
</script>

<article>
	<a href="/user/{user}"><b>{fullName}</b> @{user}</a>
	<em>{date}</em>
	<h2>{content}</h2>
</article>

<style>
	article {
		width: 80vw;
		max-width: 896px;
		margin: 0 auto;
		padding: 1rem;
	}
	h2 {
		margin: 0;
		font-family: sans-serif;
	}
</style>
