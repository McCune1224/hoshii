import type { PageServerLoad } from './$types';

export const load = (async ({ cookies, params, locals }) => {
    const userEndpoint = params.user;
    const user = locals.activeUser;
    if (!user) {
        return {
            user: undefined,
            sameUser: false,
            userEndpoint,
        };
    }
    return {
        user: user,
        sameUser: user.username === userEndpoint,
        userEndpoint
    };
}) satisfies PageServerLoad;
