import { supabase } from '$lib/supabaseClient';


export const load = async ({ locals: { supabase } }) => {
  
	const { data: postsData } = await supabase
		.from('posts')
		.select(`id, name, created, profiles (username, full_name, avatar_url)`)
		.order('created', { ascending: false });

	return {
		posts: postsData ?? []
	};
  }
