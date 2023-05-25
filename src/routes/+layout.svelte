<script>
	import '../app.css';
</script>

<div class="flex h-screen flex-col">
	<div class="flex-grow">
		<slot />
	</div>
	<footer class="bg-zinc-100 text-center">
		<p>&copy; 2023 Hoshii. All rights reserved.</p>
		<p class="text-gray-400">Designed by Arwyn Wyn & Alex McCune</p>
	</footer>
</div>
