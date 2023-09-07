<script>
	import Header from '../components/header.svelte';

	export let data;
</script>

<svelte:head>
	<title>Home - Wired</title>
	<meta name="description" content="Svelte demo app" />
</svelte:head>

<Header />

<ul>
	{#each data.countries as country}
		<li>{country.name}</li>
	{/each}
</ul>
