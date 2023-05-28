import { dev } from '$app/environment';
import type { ErrorResponse } from './errors';
interface APIResponse<T> {
	data: T;
}
export class HoshiiAPI {
	private authToken: string = '';

	private devServer: boolean = dev;
	private BASE_URL = 'http://localhost:5173/api';

	/** 
    Wrapper for API routes to easily call endpoints and get typed responses 
    */
	constructor(token?: string) {
		if (!this.devServer) {
			this.BASE_URL = '';
		}
		if (token) {
			this.authToken = token;
		}
	}

	private async makeRequest<T>(method: string, url: string, payload?: any): Promise<T> {
		const requestOptions: RequestInit = {
			method,
			headers: {
				'Content-Type': 'application/json',
				Authorization: 'Bearer ' + this.authToken
			}
		};

		if (payload) {
			requestOptions.body = JSON.stringify(payload);
		}

		try {
			const response = await fetch(this.BASE_URL + url, requestOptions);

			if (!response.ok) {
				const responseJson = await response.json();
				const errorResponse: ErrorResponse = {
					error: {
						status: response.status,
						message: responseJson.message
					}
				};
				console.log('API REQ: ', this.BASE_URL + url, requestOptions, response);
				throw errorResponse;
			}

			const responseData: APIResponse<T> = await response.json();
			console.log('API REQ: ', this.BASE_URL + url, responseData.data);
			return responseData.data;
		} catch (error: any) {
			throw { message: 'An error occurred while making the request: ' + error.message };
		}
	}

	// ME Endpoints

	async GetMe(): Promise<MeUserResponse> {
		return this.makeRequest<MeUserResponse>('GET', '/me/users');
	}

	// Wishlist Endpoints

	async GetWishlist(id: string): Promise<WishlistResponse> {
		return this.makeRequest<WishlistResponse>('GET', `/wishlists/${id}`);
	}

	async GetWishlistItems(id: string): Promise<ItemsResponse[]> {
		return this.makeRequest<ItemsResponse[]>('GET', `/wishlists/${id}/items`);
	}

	// User Endpoints
}
