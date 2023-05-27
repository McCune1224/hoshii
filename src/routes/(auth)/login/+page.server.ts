import { SessionStore } from '$lib/sessions/redis';
import { v4 as uuid } from 'uuid';
import type { PageServerLoad, Actions } from './$types';
import { ValidateFormData } from '$lib/utils/forms';
import bcrypt from 'bcrypt';
import { error, fail, redirect } from '@sveltejs/kit';
import prisma from '$lib/prisma';

/** 7 Days (measured in seconds) */
const SESSION_EXPIRY = 60 * 60 * 24 * 7;

interface SessionData {
    userID: string;
    email: string;
    username: string;
}

export const load = (async ({ cookies, url }) => {
    // Check if session exists for user
    const redirectLocation = url.searchParams.get('redirect');
    const sessionCookieID = cookies.get('sessionID');
    if (!sessionCookieID) {
        // No session, proceed with normal client-side rendering
        return;
    }
    if (sessionCookieID) {
        // Update and extend current session
        const expireUpdateOk = await SessionStore.expire(sessionCookieID, SESSION_EXPIRY);
        if (!expireUpdateOk) {
            // Session expired
            cookies.set('sessionID', '', {
                path: '/',
                maxAge: 0
            });
        }
        console.log(await SessionStore.hgetall(sessionCookieID));
        console.log(await SessionStore.ttl(sessionCookieID));
        // Redirect since all the session stuff is now refreshed
        if (redirectLocation) {
            throw redirect(301, `/${redirectLocation}`);
        }
        throw redirect(301, '/dashboard');
    }
}) satisfies PageServerLoad;

export const actions = {
    default: async ({ cookies, request }) => {
        const data = await request.formData();
        const missingFields = ValidateFormData(data, ['email', 'password']);
        if (missingFields.length > 0) {
            return fail(400, { errors: missingFields });
        }
        const email = data.get('email') as string;
        const password = data.get('password') as string;
        const hashedPassword = await bcrypt.hash(password, 10);
        const user = await prisma.user.findUnique({
            where: {
                email
            }
        });
        if (!user) {
            return fail(400, { errors: ['Email not Registered'] });
        }

        const passwordMatch = await bcrypt.compare(password, user.password);
        if (!passwordMatch) {
            return fail(400, { errors: ['Invalid Password'] });
        }

        const sessionKey = uuid();
        //Pull down user data from main DB
        const dbUser = await prisma.user.findUnique({
            where: {
                email
            }
        });
        if (!dbUser) {
            throw error(500, 'Failed to fetch user from email');
        }

        const sessionID = uuid();
        const sessionOK = await SessionStore.hset(sessionID, {
            userId: dbUser?.id,
            email: dbUser?.email,
            username: dbUser?.name
        });
        if (!sessionOK) {
            throw error(500, 'Failed to create session');
        }
        await SessionStore.expire(sessionID, SESSION_EXPIRY);
        cookies.set('sessionID', sessionID, {
            path: '/',
            // 7 days
            maxAge: SESSION_EXPIRY
        });
    }
} satisfies Actions;
