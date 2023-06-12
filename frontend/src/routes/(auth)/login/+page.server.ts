import { SessionStore } from '$lib/sessions/redis';
import { v4 as uuid } from 'uuid';
import type { PageServerLoad, Actions } from './$types';
import { ValidateFormData } from '$lib/utils/forms';
import bcrypt from 'bcrypt';
import { error, fail, redirect } from '@sveltejs/kit';
import prisma from '$lib/prisma';

/** 7 Days (measured in seconds) */
const SESSION_EXPIRY = 60 * 60 * 24 * 7;

export const load = (async ({ cookies, locals, url }) => {
	if (locals.activeUser) {
		throw redirect(301, `/${locals.activeUser.username}`);
	}
	const sessionCookieID = cookies.get('sessionId');
	if (!sessionCookieID) {
		return;
	}

	const expireUpdateOk = await SessionStore.expire(sessionCookieID, SESSION_EXPIRY);
	if (!expireUpdateOk) {
		cookies.delete('sessionId');
		return;
	}
}) satisfies PageServerLoad;

export const actions = {
	default: async ({ cookies, request, locals }) => {
		//get form data
		const data = await request.formData();
		const missingFields = ValidateFormData(data, ['email', 'password']);
		if (missingFields.length > 0) {
			return fail(400, { errors: missingFields });
		}
		const email = data.get('email') as string;
		const password = data.get('password') as string;

		// try to find user in db
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

		// other unknown errors
		if (!dbUser) {
			throw error(500, 'Failed to fetch user from email');
		}

		// create session
		const sessionId = uuid();
		const sessionOK = await SessionStore.hset(sessionId, {
			userId: dbUser?.id,
			email: dbUser?.email,
			username: dbUser?.name
		});

		if (!sessionOK) {
			throw error(500, 'Failed to create session');
		}

		// setting session includes store session in redis, and setting cookie, and for the app to quickly use set in locals
		await SessionStore.expire(sessionId, SESSION_EXPIRY);
		cookies.set('sessionId', sessionId, {
			path: '/',
			maxAge: SESSION_EXPIRY
		});
		locals.activeUser = {
			userId: dbUser.id,
			email: dbUser.email,
			username: dbUser.name
		};
	}
} satisfies Actions;
