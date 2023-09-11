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

<h2>{theName}</h2>
<h3>{data.session?.user.email}</h3>

<ul>
	{#each data.countries as country}
		<li>{country.name}</li>
	{/each}
</ul>
