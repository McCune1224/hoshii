import { Redis } from 'ioredis';
import { SESSION_DB } from '$env/static/private';

export const SessionStore = new Redis(SESSION_DB);

/** Datatype for what is expected from the Session store hash 
    Normally retrieved by SessionStore.hgetall(sessionID)
*/
export interface SessionData {
    UserId: string;
    email: string;
    username: string;
}
