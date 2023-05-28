import { HoshiiAPI } from '$lib/api/client';
import { SessionStore, type SessionData } from '$lib/sessions/redis';
import { json, type Handle } from '@sveltejs/kit';

export const handle = (async ({ event, resolve }) => {
	// For right now everything under /api is protected, but maybe later down the line some will be public...
	const seshID = event.cookies.get('sessionId');

	if (seshID && event.locals.activeUser === undefined) {
		event.locals.activeUser = await SessionStore.hgetall(seshID as string);
	}
	// Auth Middleware for API routes basically...
	if (event.url.pathname.includes('/api')) {
		console.log('PROTECTED API ROUTE: ', event.url.pathname);
		const authHeader = event.request.headers.get('Authorization');
		console.log('AUTH HEADER: ', authHeader);
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
			const sessionData: SessionData = await SessionStore.hgetall(token);
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

			event.locals.activeUser = sessionData;
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
