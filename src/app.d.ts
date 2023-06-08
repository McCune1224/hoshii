// See https://kit.svelte.dev/docs/types#app

import type { HoshiiAPI } from '$lib/api/client';
import type { SessionData } from '$lib/sessions/redis';

// for information about these interfaces
declare global {
    namespace App {
        interface Locals {
            activeUser: {
                userId: number;
                email: string;
                username: string;
                token: string;
            };
        }
        // interface Error {}
        // interface Locals {}
        // interface PageData {}
        // interface Platform {}
    }
}

export { };
