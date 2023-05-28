import { SessionStore, type SessionData } from '$lib/sessions/redis';
import type { PageServerLoad } from './$types';

export const load = (async ({ params, locals }) => {
    const userEndpoint = params.user;
    const user: SessionData = await locals.activeUser;
    return {
        user: user,
        sameUser: user.username === userEndpoint
    };
}) satisfies PageServerLoad;
