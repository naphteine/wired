<script lang="ts">
	import { loading } from '$lib/loading';
</script>

{#if $loading}
	<div />
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
