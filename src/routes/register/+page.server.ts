import { fail, redirect } from '@sveltejs/kit'
import type { PageServerLoad } from '../$types.js';

export const load: PageServerLoad = async ({ url, locals: { getSession } }) => {
	const session = await getSession();

	// if the user is already logged in return them to the account page
	if (session) {
		redirect(303, '/account');
	}

	return { url: url.origin };
};

export const actions = {
	default: async ({ request, url, locals: { supabase } }) => {
	  const formData = await request.formData()
	  const email = formData.get('email') as string
	  const password = formData.get('password') as string
  
	  const { error } = await supabase.auth.signUp({
		email,
		password,
		options: {
		  emailRedirectTo: `${url.origin}/auth/callback`,
		},
	  })
  
	  if (error) {
		return fail(500, { message: `Error: ${error.message}`, success: false, email })
	  }
  
	  return {
		message: 'Please check your email for a magic link to log into the website.',
		success: true,
	  }
	},
  }