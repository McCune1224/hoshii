import { json, type RequestHandler } from '@sveltejs/kit';

export const GET = (async ({ request, cookies, locals, params }) => {
	return json({
		data: {
			name: 'todo item'
		}
	});
}) satisfies RequestHandler;
