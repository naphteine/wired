import { supabase } from '$lib/supabaseClient';

export async function load() {
	const { data } = await supabase
		.from('posts')
		.select(`id, name, created, profiles (username, full_name, avatar_url)`)
		.order('created', { ascending: false });

	return {
		posts: data ?? []
	};
}
