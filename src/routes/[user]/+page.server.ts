import { SessionStore, type SessionData } from '$lib/sessions/redis';
import type { PageServerLoad } from './$types';

export const load = (async ({ params, cookies }) => {
    const seshID = cookies.get('sessionId');
    const userEndpoint = params.user;
    if (!seshID) {
        return {
            props: {
                currentUser: false
            }
        };
    }
    const currUser: SessionData = await SessionStore.hgetall(seshID);
    if (currUser.username === params.user) {
        return {
            props: {
                sameUser: true,
                currUser,
                userEndpoint
            }
        };
    }

    return {
        props: {
            sameUser: currUser.username === params.user,
            userEndpoint
        }
    };
}) satisfies PageServerLoad;
