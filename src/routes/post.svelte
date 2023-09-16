<script lang="ts">
	import type { Session, SupabaseClient } from '@supabase/supabase-js';
	import { createEventDispatcher } from 'svelte';

	import IconParkTwotoneUpC from 'virtual:icons/icon-park-twotone/up-c';
	import IconParkTwotoneDownC from 'virtual:icons/icon-park-twotone/down-c';
	import { error, fail } from '@sveltejs/kit';

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
	let likeCount = likes.filter((like: { liked: any }) => like.liked).length;
	let dislikeCount = likes.filter((like: { liked: any }) => !like.liked).length;

	let isUserLiked = false;
	let isUserDisliked = false;

	if (loggedInUserId) {
		isUserLiked = likes.some(
			(like: { user: string; liked: any }) => like.user === loggedInUserId && like.liked
		);
		isUserDisliked = likes.some(
			(like: { user: string; liked: any }) => like.user === loggedInUserId && !like.liked
		);
	}

	let isUserActed = isUserLiked || isUserDisliked;

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

	const removeLikeDislike = async (user_id: string, post_id: number) => {
		const { error: createRemoveError, data: removeLike } = await supabase
			.from('likes')
			.delete()
			.eq('user', user_id)
			.eq('post', post_id);

		if (createRemoveError) {
			return fail(500, {
				supabaseErrorMessage: createRemoveError.message
			});
		}

		if (isUserActed) {
			if (isUserLiked) {
				likeCount -= 1;
			} else {
				dislikeCount -= 1;
			}
		}

		isUserActed = false;
		isUserLiked = false;
		isUserDisliked = false;

		return {
			createRemoveError
		};
	};

	const likePost = async (user_id: string, post_id: number, isLike: boolean) => {
		const { error: createLikeError, data: newLike } = await supabase
			.from('likes')
			.insert({ user: session.user.id, post: postId, liked: isLike });

		if (createLikeError) {
			return fail(500, {
				supabaseErrorMessage: createLikeError.message
			});
		}

		isUserActed = true;

		if (isLike) {
			likeCount += 1;
		} else {
			dislikeCount += 1;
		}

		isUserLiked = isLike;
		isUserDisliked = !isLike;

		return {
			newLike
		};
	};

	const upvote = async () => {
		if (!session) {
			alert('You are not logged in!');
			return;
		}

		if (!postId) {
			alert('Post not found!');
			return;
		}

		if (isUserActed && isUserLiked) {
			await removeLikeDislike(session.user.id, +postId);
		} else {
			await removeLikeDislike(session.user.id, +postId);
			await likePost(session.user.id, +postId, true);
		}
	};

	const dislike = async () => {
		if (!session) {
			alert('You are not logged in!');
			return;
		}

		if (!postId) {
			alert('Post not found!');
			return;
		}

		if (isUserActed && isUserDisliked) {
			await removeLikeDislike(session.user.id, +postId);
		} else {
			await removeLikeDislike(session.user.id, +postId);
			await likePost(session.user.id, +postId, false);
		}
	};
</script>

<article class="bg-gray-100 bg-opacity-70 rounded p-5 mx-auto my-3">
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
			<button on:click={upvote} class="flex items-center">
				<IconParkTwotoneUpC class={isUserActed && isUserLiked ? 'mx-1 text-green-600' : 'mx-1'} />
				{#if likeCount > dislikeCount}
					<b>{likeCount}</b>
				{:else}
					{likeCount}
				{/if}
			</button>

			<button on:click={dislike} class="flex items-center">
				<IconParkTwotoneDownC class={isUserActed && !isUserLiked ? 'mx-1 text-red-600' : 'mx-1'} />
				{#if likeCount < dislikeCount}
					<b>{-dislikeCount}</b>
				{:else}
					{dislikeCount}
				{/if}
			</button>
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
