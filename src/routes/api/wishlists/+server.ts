import prisma from '$lib/prisma';
import { json, type HandleServerError, type RequestHandler } from '@sveltejs/kit';

export const POST = (async ({ request, locals }) => {
    const body: WishlistPostBody = await request.json();
    if (!body || body.name === undefined) {
        return json({ message: 'Missing required params' }, { status: 400 });
    }
    const dbUser = await prisma.user.findUnique({
        where: {
            id: locals.activeUser.userId as number
        }
    });
    try {
        console.log(locals.activeUser.userId);
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
                        id: dbUser?.id as number
                    }
                },
                items: {
                    create: {
                        name: 'First Item',
                        description: 'My first item is cool!'
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
