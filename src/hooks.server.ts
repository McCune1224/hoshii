import { SessionStore, type SessionData } from '$lib/sessions/redis';
import { json, type Handle } from '@sveltejs/kit';

export const handle = (async ({ event, resolve }) => {
	// For right now everything under /api is protected, but maybe later down the line some will be public...
	const protectedApiRoutes = 'api/';
	const seshID = event.cookies.get('sessionID');

	if (!seshID && protectedApiRoutes.includes(event.url.pathname)) {
		return json({
			error: 'You must be logged in to access this route',
			status: 401
		});
	}
	// https://kit.svelte.dev/docs/hooks
	const sessionUser: SessionData = await SessionStore.hgetall(seshID as string);
	event.locals.user = sessionUser;

	return await resolve(event);
}) satisfies Handle;
