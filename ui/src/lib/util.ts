import { MAX_INT, random } from '@aicacia/rand';

export function createInsecureID() {
	return (random() * MAX_INT) | 0;
}

export function getSearchTerms(search: string): string[] {
	return search
		.toLowerCase()
		.trim()
		.split(/\s+/)
		.filter((t) => !!t);
}

export function toURLSafe(value: string): string {
	return value
		.trim()
		.toLowerCase()
		.replace(/[\s]+/gi, '-')
		.replace(/[^\w\d\-_]+/gi, '');
}

export function parseJWTClaims(jwtToken: string): { exp: number } {
	return JSON.parse(atob(jwtToken.split('.')[1]).toString());
}

export function getHttpOrigin(host: string, ssl = true) {
	return `${ssl ? 'https' : 'http'}://${host}`;
}

export function getWSOrigin(host: string, ssl = true) {
	return `${ssl ? 'wss' : 'ws'}://${host}`;
}

export function waitMS(ms: number): Promise<void> {
	return new Promise((resolve) => setTimeout(resolve, ms));
}
