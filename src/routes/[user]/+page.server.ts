import type { PageServerLoad } from './$types';

export const load = (async ({ cookies, params, locals }) => {
	const userEndpoint = params.user;
	const user = locals.activeUser;
	return {
		user: user,
		sameUser: user.username === userEndpoint
	};
}) satisfies PageServerLoad;
