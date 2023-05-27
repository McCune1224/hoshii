import { SessionStore } from '$lib/sessions/redis';
import { redirect } from '@sveltejs/kit';
import type { PageServerLoad } from './$types';

export const load = (async ({ cookies }) => {
    await SessionStore.del(cookies.get('sessionID') as string);
    cookies.delete('sessionID');
    console.log('Session deleted')
    throw redirect(301, '/');
}) satisfies PageServerLoad;
