<script>
	import Header from '../../components/header.svelte';
</script>

<svelte:head>
	<title>Login - Wired</title>
</svelte:head>

<Header />

<h2>Login</h2>
<form>
	<input type="text" placeholder="e-mail" />
	<input type="password" name="password" id="password" placeholder="password" />
	<button type="submit">LOGIN</button>
</form>

<h2>Register</h2>
<form>
	<input type="text" placeholder="e-mail" />
	<input type="password" name="password" id="password" placeholder="password" />
	<input type="password" name="reenter" id="reenter" placeholder="re-enter password " />
	<button type="submit">REGISTER</button>
</form>
