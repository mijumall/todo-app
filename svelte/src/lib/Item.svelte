<script lang="ts">
	import DropdownMenu from './DropdownMenu.svelte';

	export let id: string = "asdfasdf1234"
	export let content: string = 'Sample todo content';
	export let createdAt: string = '2022-01-01';

	let isShowDropdownMenu = false;

	let dropdownMenuLeft = 0;
	let dropdownMenuTop = 0;

	function showDropdownMenu(e: MouseEvent): void {
		dropdownMenuLeft = e.pageX;
		dropdownMenuTop = e.pageY;
		isShowDropdownMenu = true;
	}
</script>

<main>
	<div class="e1-1" on:click|stopPropagation={showDropdownMenu}>
		<span class="e2-1">{content}</span>
		<span class="e2-2">{createdAt}</span>
	</div>
	{#if isShowDropdownMenu}
		<DropdownMenu 
			left={dropdownMenuLeft} 
			top={dropdownMenuTop} 
			bind:isActive={isShowDropdownMenu}
			bind:content={content}
			bind:id={id}
		/>
	{/if}
</main>

<style lang="scss">
	.e1-1 {
		display: flex;
		flex-direction: column;

		width: 70vw;

		padding: 18px;
		margin: 20px;

		@extend .flat;

		&:active {
			@extend .pressed;
		}

		@media screen and (min-width: 1000px) {
			width: 700px;
		}

		.e2-1 {
			font-size: 25px;
		}

		.e2-2 {
			font-size: 18px;
		}

		user-select: none;
	}

	// Soft UI Effect
	.flat {
		border-radius: 20px;
		box-shadow:  8px 8px 8px #131215,
					-8px -8px 8px #414047;
	}

	.pressed {
		border-radius: 20px;
		box-shadow: inset 8px 8px 8px #131215,
					inset -8px -8px 8px #414047;
	}
</style>
