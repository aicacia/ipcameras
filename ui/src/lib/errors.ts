import { ResponseError, type Errors } from '$lib/openapi/ipcameras';
import { get } from 'svelte/store';
import { createNotification } from './stores/notifications';
import LL from './i18n/i18n-svelte';

export async function handleError(error: unknown) {
	if (error instanceof ResponseError) {
		const errors = await error.response.json();
		if (errors) {
			notifyErrors(errors);
			return errors as Errors;
		}
	}
	console.error(error);
	createNotification(`${get(LL).errors.name.internal()}: ${get(LL).errors.message.application()}`);
	throw error;
}

export async function notifyErrors(errors: Errors) {
	const ll = get(LL);
	for (const [nameKey, messages] of Object.entries(errors.errors)) {
		for (const message of messages) {
			const name = (ll.errors.name as any)[nameKey]() as string;
			const body = (ll.errors.message as any)[message.error](...message.parameters) as string;
			createNotification(`${name}: ${body}`);
		}
	}
}
