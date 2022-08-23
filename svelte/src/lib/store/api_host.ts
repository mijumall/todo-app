export let api_host: string = 'https://todo-api.anchorboard.net/v1';

if (location.hostname === 'localhost') {
	api_host = 'http://localhost:8002/v1';
}
