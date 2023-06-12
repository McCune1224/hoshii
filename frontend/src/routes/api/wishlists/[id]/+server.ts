import prisma from '$lib/prisma';
import { json, type RequestHandler } from '@sveltejs/kit';

export const GET = (async ({ request, params, locals }) => {
	const id = Number(params.id);
	if (isNaN(id)) {
		return json(
			{
				data: {
					error: 'Invalid id'
				}
			},
			{
				status: 400
			}
		);
	}
	const dbWishlist = await prisma.wishlist.findUnique({
		where: {
			id: id,
		}
	});

	return json({
		data: { ...dbWishlist }
	});
	return json({
		data: {
			name: 'todo'
		}
	});
}) satisfies RequestHandler;
export const PUT = (async ({ request, cookies, locals }) => {
	return json(
		{
			data: {
				name: 'todo PUT'
			}
		},
		{}
	);
}) satisfies RequestHandler;

export const DELETE = (async ({ request, cookies, locals }) => {
	return json(
		{
			data: {
				name: 'todo DELETE'
			}
		},
		{}
	);
}) satisfies RequestHandler;
