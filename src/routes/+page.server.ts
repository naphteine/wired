import { supabase } from '$lib/supabaseClient';

export async function load() {
	const { data } = await supabase
		.from('posts')
		.select(`id, name, created, profiles (username)`)
		.order('created', { ascending: false });

	return {
		posts: data ?? []
	};
}
