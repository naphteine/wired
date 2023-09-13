import { error, fail } from '@sveltejs/kit';

export const actions = {
	createPost: async ({ request, locals: { supabase, getSession } }) => {
		const session = await getSession();

		if (!session) {
			throw error(401, { message: 'Unauthorized' });
		}

		const formData = await request.formData();
		const content = formData.get('post');
		const created = new Date().toUTCString();

		const { error: createPostError, data: newPost } = await supabase
			.from('posts')
			.insert({ name: content, user: session.user.id, created: created });

		if (createPostError) {
			return fail(500, {
				supabaseErrorMessage: createPostError.message
			});
		}
		return {
			newPost
		};
	},

	upvote: async ({ request, locals: { supabase, getSession }}) => {
		const formData = await request.formData();
		const postId = formData.get("postId");
		const session = await getSession();

		if (!session) {
			throw error(401, { message: 'Unauthorized' });
		}

		if (!postId) {
			throw error(404, { message: 'Post not found' });
		}

		const { error: createLikeError, data: newLike } = await supabase
			.from('likes')
			.insert({ user: session.user.id, post: postId, liked: true });

		if (createLikeError) {
			return fail(500, {
				supabaseErrorMessage: createLikeError.message
			});
		}
		return {
			newLike
		};
	},

	downvote: async ({ request, locals: { supabase, getSession }}) => {
		const formData = await request.formData();
		const postId = formData.get("postId");
		const session = await getSession();

		if (!session) {
			throw error(401, { message: 'Unauthorized' });
		}

		if (!postId) {
			throw error(404, { message: 'Post not found' });
		}

		const { error: createLikeError, data: newLike } = await supabase
			.from('likes')
			.insert({ user: session.user.id, post: postId, liked: false });

		if (createLikeError) {
			return fail(500, {
				supabaseErrorMessage: createLikeError.message
			});
		}
		return {
			newLike
		};
	}
};

export const load = async ({ locals: { supabase } }) => {
  
	const { data: postsData } = await supabase
		.from('posts')
		.select(`id, name, created, profiles (username, full_name, avatar_url), likes (user, liked)`)
		.order('created', { ascending: false });

	return {
		posts: postsData ?? []
	};
  }
