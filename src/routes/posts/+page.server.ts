// src/routes/posts/+page.server.ts
import { error, fail, redirect } from '@sveltejs/kit';
import type { PageServerLoad } from '../$types.js';

export const load: PageServerLoad = async ({ url, locals: { getSession } }) => {
	const session = await getSession();

	if (!session) {
		throw redirect(303, '/login');
	}

	return { url: url.origin };
};

export const actions = {
	createPost: async ({ request, locals: { supabase, getSession } }) => {
		const session = await getSession();

		console.log('111');

		if (!session) {
			console.log('193');
			throw error(401, { message: 'Unauthorized' });
		}

		const formData = await request.formData();
		const content = formData.get('post');
		const created = new Date().toUTCString();

		const { error: createPostError, data: newPost } = await supabase
			.from('posts')
			.insert({ name: content, user: session.user.id, created: created });

		if (createPostError) {
			console.log('9570');
			return fail(500, {
				supabaseErrorMessage: createPostError.message
			});
		}
		return {
			newPost
		};
	}
};
