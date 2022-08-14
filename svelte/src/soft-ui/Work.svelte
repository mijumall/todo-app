<script lang="ts">

</script>

<main>
	<div class="search">Search</div>
	<div class="box">This is a box</div>
</main>

<style lang="scss">
	main {
		display: flex;
		flex-direction: column;
		justify-content: center;
		align-items: center;

		width: 100%;
		height: 100%;
	}

	.search {
		margin-bottom: 50px;
		padding: 5px 30px;

		display: flex;
		justify-content: flex-start;
		align-items: center;

		width: 300px;
		height: 50px;

		user-select: none;

		border-radius: 18px;

		@extend .pressed;
	}

	.box {
		width: 200px;
		height: 200px;
		border-radius: 18px;

		display: flex;
		flex-direction: column;
		justify-content: center;
		align-items: center;

		user-select: none;

		@extend .flat;

		&:active {
			@extend .pressed;
		}
	}

	// Soft UI Effect
	.flat {
		box-shadow:  8px 8px 16px #131215,
					-8px -8px 16px #414047;
	}

	.pressed {
		box-shadow: inset 8px 8px 16px #131215,
					inset -8px -8px 16px #414047;
	}

</style>