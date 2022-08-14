<script lang="ts">
	import Icon from './Icon.svelte';
	import SizedBox from './SizedBox.svelte';
	import { itemData } from './store/itemData';

	export let top: number = 0;
	export let left: number = 0;
	export let content: string = '';
	export let id: string = "";

	// Handle DropdownMenu Position and others
	if (left > window.innerWidth / 2) {
		left -= 235;
	}
	export let isActive = true;
	window.onclick = _ => isActive = false; 

	// Copy content
	async function copy() {
		await navigator.clipboard.writeText(content)
	}

	// Delete content 
	async function deleteItem() {
		const url = `http://localhost:8002/delete?id=${id}`;
		const response = await fetch(url, {method: 'DELETE'});
		if (response.status >= 200 && response.status < 300) {
			// Remove local corresponding item data
			let copyItemData = $itemData;
			for (let i in copyItemData) {
				if (copyItemData[i].id === id) {
					copyItemData.splice(Number(i), 1);
					break;
				}
			}

			itemData.set(copyItemData);
		}
	}
</script>

<main>
	<div class="e1-1" style="top: {top}px; left: {left}px">
		<div class="e2-1" on:click={copy}>
			<Icon size={32} type="copy"/>
			<SizedBox width="20px"/>
			<p>Copy</p>
		</div>

		<div style="border: 1px solid #ffffff88; width: 96%; flex: 0.1;"></div>

		<div class="e2-1" on:click={deleteItem}>
			<Icon size={32} type="deleteIcon"/>
			<SizedBox width="20px"/>
			<p>Delete</p>
		</div>
	</div>
</main>

<style lang="scss">
	main {
		position: fixed;
		top: 0;
		left: 0;
		width: 100%;
		height: 100%;
		overflow: scroll;
	}

	.e1-1{
		position: absolute;

		width: 180px;
		height: 80px;

		background-color: #2a292e;

		border-radius: 20px;

		@extend .flat;

		display: flex;
		flex-direction: column;
		justify-content: center;
		align-items: center;

		padding: 15px;

		.e2-1 {
			width: 100%;

			display: flex;
			flex-direction: row;
			justify-content: flex-start;
			align-items: center;

			border-radius: 20px;

			&:hover {
				background-color: #333333;
			}
		}
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