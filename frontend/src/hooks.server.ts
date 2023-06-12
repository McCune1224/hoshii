import { SessionStore } from '$lib/sessions/redis';
import { json, type Handle } from '@sveltejs/kit';

export const handle = (async ({ event, resolve }) => {
	// For right now everything under /api is protected, but maybe later down the line some will be public...
	const seshID = event.cookies.get('sessionId');

	if (seshID) {
		const currentSession = await SessionStore.hgetall(seshID as string);

		event.locals.activeUser = {
			userId: +currentSession.userId,
			email: currentSession.email,
			username: currentSession.username,
			token: seshID
		};
	}
	// Auth Middleware for API routes basically...
	if (event.url.pathname.includes('/api/')) {
		const authHeader = event.request.headers.get('Authorization');
		if (!authHeader) {
			return json(
				{
					message: 'Missing Authorization Header'
				},
				{
					status: 401
				}
			);
		}
		const token = authHeader.split(' ')[1];
		try {
			const currentSession = await SessionStore.hgetall(seshID as string);
			const sessionData = {
				userId: +currentSession.userId,
				email: currentSession.email,
				username: currentSession.username,
				token: seshID as string
			};
			if (!sessionData) {
				return json(
					{
						message: 'Invalid Token',
						token: token
					},
					{
						status: 401
					}
				);
			}

            console.log("API ROUTE HIT, SESSION DATA: ", sessionData)
			event.locals.activeUser = sessionData;
			event.locals.activeUser.userId = +event.locals.activeUser.userId;
			event.locals.activeUser.token = token;
            console.log("API ROUTE HIT, ACTIVE USER: ", event.locals.activeUser)
		} catch (error: any) {
			return json(
				{
					error: 'Invalid Token',
					message: error.message,
					token: token
				},
				{
					status: 401
				}
			);
		}
	}

	return await resolve(event);
}) satisfies Handle;
