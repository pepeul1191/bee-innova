import '../stylesheets/styles.css';
import '../stylesheets/admin.css';
import App from '../components/layouts/Admin.svelte';

const app = new App({
	target: document.body,
	props: {}
});

export default app;