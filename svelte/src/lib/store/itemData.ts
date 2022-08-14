import { writable } from 'svelte/store';

interface Item {
	id: string,
	content: string,
	created_at: string,
}

export const itemData = writable<Item[]>([]);
