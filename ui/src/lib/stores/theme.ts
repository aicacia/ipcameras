import { browser } from '$app/environment';
import { localstorageWritable } from 'svelte-localstorage-writable';
import { derived } from 'svelte/store';

const themeWritable = localstorageWritable<'dark' | 'light' | null>('theme', null);

export const theme = derived(themeWritable, (theme) => theme);

if (browser) {
	theme.subscribe((currentTheme) => {
		const classList = document.documentElement.classList;
		if (currentTheme === 'light') {
			classList.remove('dark');
			classList.add('light');
		} else if (
			currentTheme === 'dark' ||
			window.matchMedia('(prefers-color-scheme: dark)').matches
		) {
			classList.remove('light');
			classList.add('dark');
		}
	});

	const currentTheme = localStorage.getItem('theme');
	if (!currentTheme) {
		if (window.matchMedia('(prefers-color-scheme: dark)').matches) {
			themeWritable.set('dark');
		} else {
			themeWritable.set('light');
		}
	}
}
