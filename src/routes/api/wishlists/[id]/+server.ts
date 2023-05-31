import { json, type RequestHandler } from '@sveltejs/kit';

export const GET = (async ({ request, cookies, locals }) => {
	return json({
		data: {
			name: 'todo'
		}
	});
}) satisfies RequestHandler;
export const PUT = (async ({ request, cookies, locals }) => {
}) satisfies RequestHandler;
export const DELETE = (async ({ request, cookies, locals }) => {}) satisfies RequestHandler;
