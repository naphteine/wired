<script lang="ts">
	import type { Session, SupabaseClient } from '@supabase/supabase-js';
	import { createEventDispatcher } from 'svelte';

	import IconParkTwotoneUpC from 'virtual:icons/icon-park-twotone/up-c';
	import IconParkTwotoneDownC from 'virtual:icons/icon-park-twotone/down-c';

	export let postId = '';
	export let content = '';
	export let user = '';
	export let date = '';
	export let fullName = '';
	export let likes;
	export let size = 2;
	export let url: string;
	export let supabase: SupabaseClient;
	export let session: Session;

	const loggedInUserId = session ? session.user.id : null;
	const likeCount = likes.filter((like) => like.liked).length;
	const dislikeCount = likes.filter((like) => !like.liked).length;

	let isUserLiked = false;
	let isUserDisliked = false;

	if (loggedInUserId) {
		isUserLiked = likes.some((like) => like.user === loggedInUserId && like.liked);
		isUserDisliked = likes.some((like) => like.user === loggedInUserId && !like.liked);
	}

	const isUserActed = isUserLiked || isUserDisliked;

	const objDate = new Date(date);
	const formattedDate = objDate.toLocaleDateString('tr-TR', {
		day: '2-digit',
		month: '2-digit',
		year: 'numeric',
		hour: '2-digit',
		minute: '2-digit'
	});

	let avatarUrl: string | null = null;
	let files: FileList;

	const dispatch = createEventDispatcher();

	const downloadImage = async (path: string) => {
		try {
			const { data, error } = await supabase.storage.from('avatars').download(path);

			if (error) {
				throw error;
			}

			const url = URL.createObjectURL(data);
			avatarUrl = url;
		} catch (error) {
			if (error instanceof Error) {
				console.log('Error downloading image: ', error.message);
			}
		}
	};

	$: if (url) downloadImage(url);
</script>

<article class="bg-slate-300 rounded-xl p-5 mx-auto my-3">
	<header class="flex items-center">
		{#if avatarUrl}
			<img src={avatarUrl} alt={avatarUrl ? 'Avatar' : 'No image'} class="h-20 w-20 rounded-full" />
		{:else}
			<div class="bg-slate-400 w-20 h-20 rounded-full" />
		{/if}
		<div class="flex flex-col mx-3">
			<a href="/user/{user}" class="text-slate-800 hover:underline"><b>{fullName}</b> @{user}</a>
			<em>{formattedDate}</em>
		</div>
	</header>
	<p class="my-4">{content}</p>
	<footer class="flex items-center">
		{#if loggedInUserId}
			<form method="post" action="?/upvote">
				<input type="hidden" name="postId" value={postId} />
				<button class="flex">
					<IconParkTwotoneUpC class={isUserActed && isUserLiked ? 'mx-1 text-green-600' : 'mx-1'} />
					{#if likeCount > dislikeCount}
						<b>{likeCount}</b>
					{:else}
						{likeCount}
					{/if}
				</button>
			</form>
			<form method="post" action="?/downvote">
				<input type="hidden" name="postId" value={postId} />
				<button class="flex">
					<IconParkTwotoneDownC
						class={isUserActed && !isUserLiked ? 'mx-1 text-red-600' : 'mx-1'}
					/>
					{#if likeCount < dislikeCount}
						<b>{-dislikeCount}</b>
					{:else}
						{dislikeCount}
					{/if}
				</button>
			</form>
		{:else}
			<IconParkTwotoneUpC class="mx-1 text-gray-600" />
			{#if likeCount > dislikeCount}
				<b>{likeCount}</b>
			{:else}
				{likeCount}
			{/if}

			<IconParkTwotoneDownC class="mx-1 text-gray-600" />
			{#if likeCount < dislikeCount}
				<b>{-dislikeCount}</b>
			{:else}
				{dislikeCount}
			{/if}
		{/if}
	</footer>
</article>

<style>
	article {
		width: 80vw;
		max-width: 600px;
	}
</style>
