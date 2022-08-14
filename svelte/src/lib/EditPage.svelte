<script lang="ts">
	import { isEditmode} from './store/editmode';
	import SizedBox from './SizedBox.svelte';
	import { updateData } from './logic/updateData';

	let content: string = '';

	const disableEditmode = () => {
		isEditmode.set(false);
	}

	const addHandler = async () => {
		isEditmode.set(false);

		const url = "http://localhost:8002/add"
		const response = await fetch(url, {
			method: 'POST',
			body: `{"content": "${content}"}`,
		});

		if (response.status == 200) {
			await updateData();
		}
	}

</script>

<main on:click={disableEditmode}>
	<SizedBox height="50px"/>

	<div class="e1-1" on:click|stopPropagation={() => {}}>
		<textarea class="e2-1" bind:value={content} type="textbox" placeholder="Type here..."/>
		<div class="e2-2" on:click={addHandler}>Add</div>
	</div>

</main>

<style lang="scss">
	main {
		width: 100%;
		height: 100%;
		background-color: #000000aa;

		position: fixed;
		top: 0;
		left: 0;

		display: flex;
		flex-direction: column;
		justify-content: flex-start;
		align-items: center;
	}

	.e1-1 {
		width: 80vw;
		height: 250px;
		background-color: #2a292e;
		padding: 20px;
		border-radius: 20px;

		@media screen and (min-width: 1000px) {
			width: 800px;
		}

		display: flex;
		flex-direction: column;
		justify-content: center;
		align-items: flex-end;
	}

	.e2-1 {
		width: 90%;
		height: 180px;
		padding: 15px;
		margin: 10px;
		color: white;
		background-color: #2a292e;
		border: none;
		outline: none;
		border-radius: 20px;
		font-size: 20px;
		font-family: sans-serif;

		@extend .pressed;

		resize: none;
	}

	.e2-2 {
		border-radius: 20px;
		padding: 10px 30px;
		background-color: #2a292e;

		@extend .flat;

		&:active {
			@extend .pressed;
		}

		user-select: none;
	}

	// Soft UI Effect
	.flat {
		box-shadow:  8px 8px 8px #131215,
					-8px -8px 8px #414047;
	}

	.pressed {
		box-shadow: inset 8px 8px 8px #131215,
					inset -8px -8px 8px #414047;
	}
</style>