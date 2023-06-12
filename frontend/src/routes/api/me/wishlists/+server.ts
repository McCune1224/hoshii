import prisma from '$lib/prisma';
import { json, type RequestHandler } from '@sveltejs/kit';

export const GET = (async ({ request, params, locals, url }) => {
	let next: string | undefined = undefined;
	let previous: string | undefined = undefined;
	let offset: number = 0;
	let limit: number = 10;

	url.searchParams.forEach((value, key) => {
		if (key === 'offset') {
			offset = parseInt(value);
			return;
		} else if (key === 'limit') {
			limit = parseInt(value);
			return;
		}
	});

	if (limit > 20) {
		limit = 20;
	}

	const wishlists = await prisma.wishlist.findMany({
		where: {
			user: {
				id: locals.activeUser.userId
			}
		},
		skip: offset,
		// Need to get one more than the limit to see if there is a next page
		// If there is, we will return the next page url
		take: limit + 1,
		include: {
			items: true
		},
		orderBy: {
			createdAt: 'desc'
		}
	});

	if (!wishlists) {
		return json({ message: 'No wishlists found' }, { status: 404 });
	}

	// There is a next page, so pop the last element to match the limit length
	if (wishlists.length > limit) {
		next = `${request.url}?offset=${offset + limit}&limit=${limit}`;
		wishlists.pop();
	}
	// a previous page exists if we have an offset, so provide the url for that as well
	else if (offset > 0) {
		previous = `${request.url}?offset=${offset - limit}&limit=${limit}`;
	}

	return json({
		data: {
			offset: offset,
			limit: limit,
			// Return the limit number of wishlists
			wishlists: wishlists,
			next
		}
	});
}) satisfies RequestHandler;
