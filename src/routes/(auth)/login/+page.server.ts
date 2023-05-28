import { SessionStore, type SessionData } from '$lib/sessions/redis';
import { v4 as uuid } from 'uuid';
import type { PageServerLoad, Actions } from './$types';
import { ValidateFormData } from '$lib/utils/forms';
import bcrypt from 'bcrypt';
import { error, fail, redirect } from '@sveltejs/kit';
import prisma from '$lib/prisma';

/** 7 Days (measured in seconds) */
const SESSION_EXPIRY = 60 * 60 * 24 * 7;

export const load = (async ({ cookies, url }) => {
	// Check if session exists for user
	const redirectLocation = url.searchParams.get('redirect');
	const sessionCookieID = cookies.get('sessionId');
	if (!sessionCookieID) {
		// No session, proceed with normal client-side rendering
		return;
	}
	// Update and extend current session
	const expireUpdateOk = await SessionStore.expire(sessionCookieID, SESSION_EXPIRY);
	if (!expireUpdateOk) {
		// Session expired
		cookies.set('sessionId', '', {
			path: '/',
			maxAge: 0
		});
		return;
	}
	const seshUser: SessionData = await SessionStore.hgetall(sessionCookieID);
	throw redirect(301, `/${seshUser?.username}`);
}) satisfies PageServerLoad;

export const actions = {
	default: async ({ cookies, request, locals }) => {
		const data = await request.formData();
		const missingFields = ValidateFormData(data, ['email', 'password']);
		if (missingFields.length > 0) {
			return fail(400, { errors: missingFields });
		}
		const email = data.get('email') as string;
		const password = data.get('password') as string;
		const dbUser = await prisma.user.findUnique({
			where: {
				email
			}
		});
		if (!dbUser) {
			return fail(400, { errors: ['Cannot find Email'] });
		}

		const passwordMatch = await bcrypt.compare(password, dbUser.password);
		if (!passwordMatch) {
			return fail(400, { errors: ['Invalid Password'] });
		}

		if (!dbUser) {
			throw error(500, 'Failed to fetch user from email');
		}

		const sessionId = uuid();
		const sessionOK = await SessionStore.hset(sessionId, {
			userId: dbUser?.id,
			email: dbUser?.email,
			username: dbUser?.name
		});

		if (!sessionOK) {
			throw error(500, 'Failed to create session');
		}

		await SessionStore.expire(sessionId, SESSION_EXPIRY);
		cookies.set('sessionId', sessionId, {
			path: '/',
			// 7 days
			maxAge: SESSION_EXPIRY
		});
		locals.activeUser = {
			userId: dbUser.id,
			email: dbUser.email,
			username: dbUser.name
		};
	}
} satisfies Actions;
