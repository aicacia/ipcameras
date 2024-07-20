import { derived, get } from 'svelte/store';
import { localstorageWritable } from 'svelte-localstorage-writable';
import { isOnline } from './online';
import EventEmitter from 'eventemitter3';
import { goto } from '$app/navigation';
import { base } from '$app/paths';
import type { Token, User } from '$lib/openapi/ipcameras';
import {
	currentUserApi,
	getAuthToken,
	ipcamerasConfiguration,
	setAuthToken,
	tokenApi
} from '$lib/openapi';

const tokenWritable = localstorageWritable<Token | null>('token', null);
const userWritable = localstorageWritable<User | null>('user', null);

export const currentUser = derived(userWritable, (user) => user);
export const signedIn = derived(userWritable, (user) => !!user);

export const userEmitter = new EventEmitter<{
	user(user: User): void;
	signOut(): void;
}>();

export function waitForUser() {
	const user = get(userWritable);
	if (getAuthToken() && user) {
		return Promise.resolve(user);
	} else {
		return new Promise<User>((resolve) => userEmitter.once('user', resolve));
	}
}

export function updateCurrentUser(user: User) {
	if (get(currentUser)?.username === user.username) {
		userWritable.update((currentUser) => (currentUser ? { ...currentUser, ...user } : null));
	}
}

export function isSignedIn() {
	return get(signedIn);
}

export function getCurrentUser() {
	return get(currentUser);
}

export async function signIn(username: string, password: string) {
	const token = await tokenApi.token({
		username,
		password
	});
	return signInWithToken(token);
}

export async function signInWithToken(token: Token) {
	setAuthToken(token);
	const user = await currentUserApi.currentUser();
	userWritable.set(user);
	tokenWritable.set(token);
	userEmitter.emit('user', user);
	return user;
}

export function signOut() {
	userWritable.set(null);
	tokenWritable.set(null);
	setAuthToken(undefined);
	userEmitter.emit('signOut');
}

let initialCall = true;
export async function tryGetCurrentUser() {
	try {
		let user = get(userWritable);
		if (initialCall) {
			if (isOnline()) {
				const token = get(tokenWritable);
				if (token) {
					setAuthToken(token);
					user = await currentUserApi.currentUser();
					userWritable.set(user);
					userEmitter.emit('user', user);
				} else {
					signOut();
					user = null;
				}
			} else if (user) {
				userEmitter.emit('user', user);
			}
			initialCall = false;
		}
		return user;
	} catch (error) {
		console.error(error);
		signOut();
		return null;
	}
}

ipcamerasConfiguration.middleware?.push({
	async post(context) {
		switch (context.response.status) {
			case 401: {
				signOut();
				await goto(`${base}/signin`);
				break;
			}
			case 503: {
				await goto(`${base}/maintenance`);
				break;
			}
		}
	}
});
