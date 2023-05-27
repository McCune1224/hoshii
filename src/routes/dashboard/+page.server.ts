import { SessionStore, type SessionData } from '$lib/sessions/redis';
import { redirect } from '@sveltejs/kit';
import type { PageServerLoad } from './$types';

export const load = (async ({ cookies }) => {
	const sessionID = cookies.get('sessionID');
	console.log('SESSION ID', sessionID);
	console.log('Session hash', await SessionStore.hgetall(sessionID as string));
	// Login route should handle a redirect back to the dashboard after a successful login
	if (!sessionID) {
		throw redirect(301, '/login?redirect=dashboard');
	}
	const sessionMap = await SessionStore.hgetall(sessionID);
	//Convert record to a SessionData object
	const currentUser = {} as SessionData;

	// Basically a for-each loop over the object keys and values
	// to assign them to the currentUser object
	// Really just here to make TS happy and give us typing on the Redis hash
	for (const [key, value] of Object.entries(sessionMap)) {
		currentUser[key as keyof SessionData] = value;
	}

	return {
		currentUser
	};
}) satisfies PageServerLoad;
