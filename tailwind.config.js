/** @type {import('tailwindcss').Config} */
export default {
	content: ['./src/**/*.{html,js,svelte,ts}'],
	theme: {
		extend: {
			colors: {
				/* 'theme-primary': '#FAA6A6', */
				'theme-primary': '#FAF1A6',
				/* 'theme-primary': '#A6F1FA', */
				'theme-secondary': '#FAF1A6',
				'theme-tertiary': '#A6F1FA'
			}
		}
	},
	plugins: []
};
