import { SessionStore } from '$lib/sessions/redis';
import { redirect } from '@sveltejs/kit';
import type { PageServerLoad } from './$types';

export const load = (async ({ cookies }) => {
	// Check session
	const sessionCookieID = cookies.get('sessionID');
	if (!sessionCookieID) {
		throw redirect(301, '/');
	}
	await SessionStore.del(sessionCookieID);
	sessionCookieID &&
		cookies.set('sessionID', '', {
			path: '/',
			maxAge: 0
		});

	throw redirect(301, '/');
}) satisfies PageServerLoad;
