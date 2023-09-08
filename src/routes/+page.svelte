<script>
	import Header from '../components/header.svelte';

	export let data;

	const theName = data.myName.toUpperCase();
</script>

<svelte:head>
	<title>Home - Wired</title>
	<meta name="description" content="Svelte demo app" />
</svelte:head>

<Header />

<h2>{ theName }</h2>

<ul>
	<li>Hello</li>
</ul>
