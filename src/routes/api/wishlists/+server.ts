import prisma from '$lib/prisma';
import { json, type HandleServerError, type RequestHandler } from '@sveltejs/kit';


export const POST = (async ({ request, locals }) => {
    const body: WishlistPostBody = await request.json();
    if (!body) {
        return json({ message: 'Missing body' }, { status: 400 });
    }
    if (body.name === undefined) {
        return json({ message: 'Missing name' }, { status: 400 });
    }

    try {
        //check if wishlist with name already exists
        const existingWishlist = await prisma.wishlist.findFirst({
            where: {
                name: body.name
            }
        });
        if (existingWishlist) {
            return json({ message: 'Wishlist with that name already exists' }, { status: 400 });
        }

        const newWishlist = await prisma.wishlist.create({
            data: {
                name: body.name,
                user: {
                    connect: {
                        id: locals.activeUser.userId
                    }
                },
                items: {
                    create: {
                        name: 'My First Item!',
                        description: 'This is my first item!'
                    }
                }
            }
        });
        return json({
            message: 'Wishlist created',
            wishlist_id: newWishlist.id,
            wishlist_name: newWishlist.name
        });
    } catch (e) {
        const error = e as Error;
        return json({ message: `Error creating wishlist\n${error.message}` }, { status: 500 });
    }
}) satisfies RequestHandler;
