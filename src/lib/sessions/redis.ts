import { Redis } from 'ioredis';
import { SESSION_DB } from '$env/static/private';

export const SessionStore = new Redis(SESSION_DB);

/** Datatype for what is expected from the Session store hash 
    Normally retrieved by SessionStore.hgetall(sessionID)
*/

export type SessionData = {
    // Redis stores verything as strings so for userID needs to be able to convert between the two without ORM's (Prisma) being cranky
	userId: string | number;
	email: string;
	username: string;
};
