import { SessionStore } from '$lib/sessions/redis';
import { redirect } from '@sveltejs/kit';
import type { PageServerLoad } from './$types';

export const load = (async ({ cookies, locals }) => {
	await SessionStore.del(cookies.get('sessionId') as string);
	cookies.delete('sessionId');
    //@ts-ignore - may god forgive me
    locals.activeUser = undefined;
	throw redirect(302, '/');
}) satisfies PageServerLoad;
