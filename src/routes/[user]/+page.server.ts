import type { PageServerLoad } from './$types';

export const load = (async ({ params, locals }) => {
	const userEndpoint = params.user;
	const user = await locals.activeUser;
	return {
		user: user,
		sameUser: user.username === userEndpoint
	};
}) satisfies PageServerLoad;
