<script lang="ts">
	import Icon from './Icon.svelte';
	import {isEditmode} from './store/editmode';

	const activateEditmode = () => {
		isEditmode.set(true);
	}
</script>

<main on:click={activateEditmode}>
	<Icon size={60} type="add"/>
</main>

<style lang="scss">
	main {
		border-radius: 100px;

		@extend .flat;

		&:active {
			@extend .pressed;
		}
	}
	// Soft UI Effect
	.flat {
		box-shadow:  8px 8px 12px #131215,
					-8px -8px 12px #414047;
	}

	.pressed {
		box-shadow: inset 8px 8px 12px #131215,
					inset -8px -8px 12px #414047;
	}
</style>