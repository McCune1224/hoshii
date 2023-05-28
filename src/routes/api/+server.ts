import { json, redirect, type RequestHandler } from '@sveltejs/kit';

// Really just here to prevent someone from hitting the API thinking its a user account
export const GET = (async ({}) => {
	throw redirect(301, '/');
}) satisfies RequestHandler;
