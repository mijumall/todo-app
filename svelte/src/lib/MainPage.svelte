<script lang="ts">
	import Item from './Item.svelte';
	import SizedBox from './SizedBox.svelte';
	import AddButton from './AddButton.svelte'; 
	import EditPage from './EditPage.svelte';
	import { isEditmode } from './store/editmode';
	import { itemData } from './store/itemData';
	import { updateData } from './logic/updateData';


	updateData();

</script>

<main>
	{#if $isEditmode}
		<EditPage/>
	{/if}
	<p class="e1-1">To Do List</p>
	<SizedBox height="10px"/>

	{#each $itemData as item}
		<Item 
			id={item.id}
			content={item.content}
			createdAt={item.created_at}
		/>
	{/each}

	<SizedBox height="10px"/>
	<AddButton/>
</main>

<style lang="scss">
	main {
		display: flex;
		flex-direction: column;
		justify-content: flex-start;
		align-items: center;

		width: 100%;
		height: 100%;

		overscroll-behavior-y: unset;

		user-select: none;
	}

	.e1-1 {
		font-size: 40px;
	}
</style>