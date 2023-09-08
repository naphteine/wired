export const actions = {
	login: async ({ cookies, request }) => {
		const data = await request.formData();

		console.log('LOGIN REQUEST');
		console.log(data);
	},

	register: async ({ cookies, request }) => {
		const data = await request.formData();

		console.log('REGISTER REQUEST');
		console.log(data);
	}
};
