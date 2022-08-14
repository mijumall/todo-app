import { itemData } from '../store/itemData';

export async function updateData() {
	const url = "http://localhost:8002/read"
	const response = await fetch(url);
	const data = await response.json();

	if (data !== null) {
		itemData.set(data);
	}
}