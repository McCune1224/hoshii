import { json } from '@sveltejs/kit';
import type { RequestHandler } from './$types';
import { SessionStore, type SessionData } from '$lib/sessions/redis';
import prisma from '$lib/prisma';

export const GET = (async ({ request, cookies, locals }) => {
	const sessionUser: SessionData = SessionStore.hgetall(cookies.get('sessionId') as string)
	const dbUser = await prisma.user.findUnique({
		where: {
			email: sessionUser.email
		}
	});
	if (!dbUser) {
		return json({
			error: 'User not found',
			status: 404
		});
	}

	return json({
		data: {
			id: dbUser.id,
			name: dbUser.name,
			email: dbUser.email,
			bio: dbUser.bio || undefined
		} as MeUserResponse
	});
}) satisfies RequestHandler;
