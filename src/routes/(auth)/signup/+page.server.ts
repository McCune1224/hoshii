import prisma from '$lib/prisma';
import { json, fail, error } from '@sveltejs/kit';
import bcrypt from 'bcrypt';
import type { PageServerLoad, Actions } from './$types';
import { ValidateFormData } from '$lib/utils/forms';

export const load = (async ({ cookies }) => {
    /* const user = await SessionRedis.get(`session:${cookies.sessionToken}`); */
    /* return user */
}) satisfies PageServerLoad;

export const actions = {
    default: async ({ cookies, request }) => {
        const data = await request.formData();
        const email = data.get('email');
        const password = data.get('password');
        const name = data.get('username');

        const missingItems = ValidateFormData(data, ['email', 'password', 'username']);
        if (missingItems.length > 0) {
            return fail(400, { errors: missingItems });
        }

        const existingUser = await prisma.user.findUnique({
            where: {
                // TS is cringe here because it won't be nullable at this point in the code
                email: email as string
            }
        });
        if (existingUser) {
            return fail(400, { errors: ['Email already registered'] });
        }

        let passwordHash: string = '';
        try {
            passwordHash = await bcrypt.hash(password as string, 10);
        } catch (e) {
            throw error(500, `Failed to hash password\n${e}`);
        }
        const user = await prisma.user.create({
            data: {
                email: email as string,
                password: passwordHash,
                name: name as string
            }
        });
    }
} satisfies Actions;
