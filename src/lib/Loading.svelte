<script lang="ts">
	import { loading } from '$lib/loading';

	$: if ($loading.status === 'NAVIGATING') {
		setTimeout(() => {
			if ($loading.status === 'NAVIGATING') {
				$loading.status = 'LOADING';
			}
		}, 400);
	}
</script>

{#if $loading.status === 'LOADING'}
	<div />
	{#if $loading.message}
		<p>{$loading.message}</p>
	{/if}
{/if}

<style>
	div {
		--size: 100px;
		position: fixed;
		top: 0;
		left: 0;
		inset: calc(50% - calc(var(--size) / 2));
		background: purple;
		border-radius: var(--size);
		height: var(--size);
		width: var(--size);
		animation: moveLoader 2s infinite alternate;
	}

	@keyframes moveLoader {
		from {
			transform: translate3d(100px, 0, 0);
		}

		to {
			transform: translate3d(-100px, 0, 0);
		}
	}
</style>
