// Strip prisma user object type of password, createdAt, updatedAt,

interface MeUserResponse {
	id: string;
	email: string;
	name: string;
	bio?: string;
}

interface WishlistResponse {
	id: string;
	name: string;
	userID: string;
	categoryID: string;
}

interface ItemsResponse {
	id: string;
	name: string;
	description?: string;
	link?: string;
	image?: string;
	price?: number;
	wishlistID: string;
	categoryID: string;
}

interface CategoryResponse {
	id: string;
	name: string;
}
