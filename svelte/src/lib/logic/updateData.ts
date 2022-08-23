import { itemData } from '../store/itemData';
import { api_host } from '../store/api_host';

export async function updateData() {
	const url = `${api_host}/read`
	const response = await fetch(url);
	const data = await response.json();

	if (data !== null) {
		itemData.set(data);
	}
}