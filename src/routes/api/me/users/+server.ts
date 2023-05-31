import { json } from '@sveltejs/kit';
import type { RequestHandler } from './$types';
import { SessionStore, type SessionData } from '$lib/sessions/redis';
import prisma from '$lib/prisma';

export const GET = (async ({ request,  locals }) => {

	const dbUser = await prisma.user.findUnique({
		where: {
			id: locals.activeUser.userId as number
		}
	});

	if (!dbUser) {
		return json({ error: { status: 404, message: 'User not found' } });
	}

	return json({
		data: {
			id: dbUser.id,
			name: dbUser.name,
			email: dbUser.email,
			bio: dbUser.bio || undefined
		} as UserResponse
	});
}) satisfies RequestHandler;
