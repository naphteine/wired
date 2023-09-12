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
