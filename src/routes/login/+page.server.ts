import { fail, redirect } from '@sveltejs/kit';
import type { PageServerLoad } from '../$types.js';

export const load: PageServerLoad = async ({ url, locals: { getSession } }) => {
	const session = await getSession();

	// if the user is already logged in return them to the account page
	if (session) {
		throw redirect(303, '/account');
	}

	return { url: url.origin };
};

export const actions = {
	default: async ({ request, locals: { supabase } }) => {

		const formData = await request.formData();
		const email = formData.get('email') as string;
		const password = formData.get('password') as string;

		const { error } = await supabase.auth.signInWithPassword({ email, password });

		if (error) {
			return fail(500, { message: `Error: ${error.message}`, success: false, email });
		}

		return {
			message: 'YEY LOGIN',
			success: true
		};
	}
};
