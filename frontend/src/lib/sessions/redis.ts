import { Redis } from 'ioredis';
import { SESSION_DB } from '$env/static/private';

// Global session store to be used throughout the app
export const SessionStore = new Redis(SESSION_DB);
