<!-- // src/routes/login/+page.svelte -->
<script lang="ts">
	import { enhance } from '$app/forms'
  import Header from '../../components/header.svelte';
	export let form
</script>

<Header />
<form method="post" use:enhance>
  <input name="email" value={form?.email ?? ''} />
  <input type="password" name="password" />
  <button>Sign up</button>
</form>