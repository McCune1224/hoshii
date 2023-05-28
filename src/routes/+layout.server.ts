import type { LayoutServerLoad } from './$types';

export const load = (async ({ locals }) => {
	return {
		activeUser: locals.activeUser
	};
}) satisfies LayoutServerLoad;
