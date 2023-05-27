import { SessionStore, type SessionData } from '$lib/sessions/redis';
import { error, redirect } from '@sveltejs/kit';
import type { PageServerLoad } from './$types';

export const load = (async () => { }) satisfies PageServerLoad;
