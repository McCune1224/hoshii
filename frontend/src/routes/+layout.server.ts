import type { LayoutServerLoad } from './$types';

export const load = (async ({ locals }) => {
	if (locals.activeUser) {
		return {
			activeUser: locals.activeUser
		};
	}
	return {
		activeUser: undefined
	};
}) satisfies LayoutServerLoad;
