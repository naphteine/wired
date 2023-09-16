<!-- // src/routes/login/+page.svelte -->
<script lang="ts">
	import { enhance } from '$app/forms';
	export let form;
</script>

<h1>LOGIN</h1>
<form method="post" use:enhance>
	<input placeholder="email" name="email" value={form?.email ?? ''} />
	<input placeholder="password" type="password" name="password" />
	<button>LOGIN</button>
</form>

<audio hidden autoplay>
	<source src="/whoistheuser.mp3" type="audio/mpeg">
</audio>
{#if form?.message}
	<h2>{form?.message}</h2>
{/if}

<style>
	h1 {
		text-align: center;
		margin: 0;
	}

	h2 {
		text-align: center;
	}

	form {
		margin: 0 auto;
		width: 150px;
	}
</style>
