import type { BaseTranslation } from '../i18n-types';

const en = {
	errors: {
		name: {
			internal: 'Application Error'
		},
		message: {
			application: 'if it presists please contact support',
			tooManyRequests: 'Too many requests',
			notFound: 'Not found',
			invalid: 'Invalid',
			required: 'Required',
			noChange: 'No change',
			mismatch: 'Passwords do not match',
			alreayUsed: 'Already used',
			disabled: 'Disabled'
		}
	},
	auth: {
		signIn: 'Sign in',
		signOut: 'Sign out',
		usernameLabel: 'Username',
		usernamePlaceholder: 'Enter your Username',
		passwordLabel: 'Password',
		passwordPlaceholder: 'Enter your Password'
	},
	connect: {
		connect: 'Connect',
		hostLabel: 'Host',
		hostPlaceholder: 'Host and Port (localhost:8080)',
		sslLabel: 'SSL',
		idLabel: 'Id',
		idPlaceholder: 'Enter your Id',
		passwordLabel: 'Password',
		passwordPlaceholder: 'Enter your Password',
		httpAccess: 'HTTP Access',
		p2pAccess: 'P2P Access'
	},
	camera: {
		hardwareIdLabel: 'Hardware ID',
		hardwareIdPlaceholder: 'Device Hardware ID',
		nameLabel: 'Name',
		namePlaceholder: 'Name',
		recordLabel: 'Record',
		recordWindowLabel: 'Record Window',
		recordWindowPlaceholder: '1h, 30d, 1y, etc...',
		save: 'Save'
	},
	header: {
		title: 'IP Cameras'
	},
	cameras: {
		title: 'Cameras'
	},
	health: {
		title: 'Health',
		header: 'Health Check',
		body: 'Healthy'
	}
} satisfies BaseTranslation;

export default en;
