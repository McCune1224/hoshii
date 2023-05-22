<script lang="ts">
	import ExampleCardListItem from './ExampleCardListItem.svelte';
</script>

<div class="bg-theme-primary">
	<header class="container mx-auto">
		<div class="sm:mx-auto sm:py-10">
			<nav
				class="mx-auto flex items-center justify-between bg-white py-4 sm:rounded-full sm:px-10 sm:text-xl"
			>
				<a href="/" class="text-3xl font-bold text-gray-800"> Hoshii [LOGO HERE]</a>
				<ul class="group flex sm:gap-48">
					<li
						class="rounded-lg px-2 py-2 text-center font-semibold hover:bg-gray-100 hover:text-gray-800 sm:px-4 sm:py-2"
					>
						<a href="#">FAQ</a>
					</li>
					<li
						class="rounded-lg px-2 py-2 text-center font-semibold hover:bg-gray-100 hover:text-gray-800 sm:px-4 sm:py-2"
					>
						<a href="" class="">
							<div class="flex gap-1">
								<span> Explore </span>
								<span class="hidden md:block"> Wishlists </span>
							</div>
						</a>
					</li>
				</ul>
				<div class="flex items-center px-5">
					<a
						href="/login"
						class="ml-4 rounded-full border-2 border-black px-2 py-2 text-center font-semibold hover:bg-theme-primary hover:opacity-80 sm:px-4 sm:py-2"
						>Log In</a
					>
					<a
						href="/signup"
						class="ml-4 rounded-full bg-theme-primary px-6 py-2 text-center font-semibold hover:bg-theme-primary hover:opacity-80 sm:px-4 sm:py-2"
						>Sign Up</a
					>
				</div>
			</nav>
		</div>
	</header>
</div>
<main class="bg-theme-primary">
	<section class="py-16">
		<div class="container mx-auto px-4 text-center">
			<h1 class="py-5 text-4xl font-extrabold sm:text-6xl">Sharing your Wishlist Made Easy.</h1>
			<p class="text-lg sm:text-xl">
				Hoshii is a platform for you to share your wishlist with your friends and family. You can
				also explore other people's wishlist and get inspired.
			</p>
		</div>
	</section>

	<section class="border-b-2 border-t-2 border-dashed border-black bg-zinc-100 py-16 text-center">
		<div class="container mx-auto">
			<h2 class="py-5 text-center text-3xl font-extrabold sm:text-5xl">Some Examples</h2>
			<ul class="grid grid-cols-3 gap-5">
				<ExampleCardListItem />
				<ExampleCardListItem />
				<ExampleCardListItem />
				<ExampleCardListItem />
				<ExampleCardListItem />
				<ExampleCardListItem />
			</ul>
		</div>
	</section>

	<section class="py-16">
		<div class="container mx-auto px-4">
			<div class="text-center">
				<h2 class="py-5 text-3xl font-extrabold sm:text-5xl">🔫 Sign Up Or else 🔫</h2>
				<a
					href="/signup"
					class="ml-4 rounded-full border-2 border-black px-2 py-2 text-center font-semibold hover:bg-theme-primary hover:opacity-80 sm:px-4 sm:py-2"
					>Make a List - It's free</a
				>
			</div>
		</div>
	</section>
</main>
