import { json, type RequestHandler } from '@sveltejs/kit';

export const GET = (async ({ request, locals }) => {
	return json({ message: `GET HIT FROM ${locals.activeUser}` });
}) satisfies RequestHandler;

export const POST = (async ({ request, locals }) => {
	return json({ message: `POST HIT FROM ${locals.activeUser}` });
}) satisfies RequestHandler;
