<!-- src/routes/+page.svelte -->
<script lang="ts">
	import { Auth } from '@supabase/auth-ui-svelte';
	import { ThemeSupa } from '@supabase/auth-ui-shared';
	import Header from '../../components/header.svelte';

	export let data;
</script>

<svelte:head>
	<title>User Management</title>
</svelte:head>

<Header />
<div class="row flex-center flex">
	<div class="col-6 form-widget">
		<Auth
			supabaseClient={data.supabase}
			view="sign_in"
			redirectTo={`${data.url}/auth/callback`}
			showLinks={false}
			appearance={{ theme: ThemeSupa, style: { input: 'color: #000' } }}
		/>
	</div>
</div>

<a href="/reset">Forgot your password?</a>
