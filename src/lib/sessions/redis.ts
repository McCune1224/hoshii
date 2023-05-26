import { Redis } from 'ioredis';
import { SESSION_DB } from '$env/static/private';





export const SessionStore = new Redis(SESSION_DB);
