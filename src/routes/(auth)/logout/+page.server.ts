import { SessionStore } from '$lib/sessions/redis';
import { redirect } from '@sveltejs/kit';
import type { PageServerLoad } from './$types';

export const load = (async ({ cookies }) => {
    await SessionStore.del(cookies.get('sessionId') as string);
    cookies.delete('sessionId');
    console.log('Session deleted')
    throw redirect(301, '/');
}) satisfies PageServerLoad;
