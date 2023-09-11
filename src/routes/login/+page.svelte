<!-- // src/routes/login/+page.svelte -->
<script lang="ts">
	import { enhance } from '$app/forms';
	import Header from '../../components/header.svelte';
	export let form;
</script>

<h1>LOGIN</h1>
<form method="post" use:enhance>
	<input placeholder="email" name="email" value={form?.email ?? ''} />
	<input placeholder="password" type="password" name="password" />
	<button>LOGIN</button>
</form>

<style>
	h1 {
		text-align: center;
		margin: 0;
	}

	form {
		margin: 0 auto;
		width: 150px;
	}
</style>