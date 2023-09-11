<script lang="ts">
	import { enhance } from '$app/forms';
	import Header from '../../components/header.svelte';

	export let data;
</script>

<svelte:head>
	<title>User Management</title>
</svelte:head>

<Header />
<form method="post" action="?/createPost" use:enhance>
	<input placeholder="new post" name="post" type="text" />
	<button>Submit</button>
</form>
