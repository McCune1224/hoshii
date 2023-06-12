import prisma from '$lib/prisma';
import { fail, error, redirect } from '@sveltejs/kit';
import bcrypt from 'bcrypt';
import type { PageServerLoad, Actions } from './$types';
import { ValidateFormData } from '$lib/utils/forms';
import { Prisma } from '@prisma/client';

export const load = (async () => { }) satisfies PageServerLoad;

export const actions = {
    default: async ({ request }) => {
        const data = await request.formData();
        const email = data.get('email');
        const password = data.get('password');
        const username = data.get('username');

        const missingItems = ValidateFormData(data, ['email', 'password', 'username']);
        if (missingItems.length > 0) {
            return fail(400, { errors: missingItems });
        }

        if (username === 'api') {
            return fail(400, { errors: ['Username not available'] });
        }

        let passwordHash: string = '';
        try {
            passwordHash = await bcrypt.hash(password as string, 10);
        } catch (e) {
            throw error(500, `Failed to hash password\n${e}`);
        }
        try {
            const user = await prisma.user.create({
                data: {
                    email: email as string,
                    password: passwordHash,
                    name: username as string
                }
            });
            redirect(301, `/${username}`);
        } catch (e) {
            if (e instanceof Prisma.PrismaClientKnownRequestError) {
                if (e.code === 'P2002') {
                    return fail(400, { errors: ['Account with provided info already registered.'] });
                }
            }
            throw error(500, `Failed to create user\n${e}`);
        }
    }
} satisfies Actions;
