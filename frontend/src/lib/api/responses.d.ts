interface UserResponse {
	id: number;
	email: string;
	name: string;
	bio?: string;
}

interface MeWishlistsResponse {
	offset: number;
	limit: number;
	next?: string;
	wishlists: SimpleWishlist[];
}

interface SimpleWishlist {
	id: number;
	name: string;
	userID: number;
	categoryID: number;
}

interface ItemsResponse {
	id: number;
	name: string;
	description?: string;
	link?: string;
	image?: string;
	price?: number;
	wishlistID: number;
	categoryID: number;
}

interface SimpleItemsResponse {
	id: number;
}

interface SimpleWishlist {
	id: number;
	name: string;
}
