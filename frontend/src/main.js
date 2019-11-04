import 'babel-polyfill';
import Vue from 'vue';
import VueFlashMessage from 'vue-flash-message';
import router from './router'

// Setup Vuetify
import Vuetify from 'vuetify';
Vue.use(Vuetify);
import 'vuetify/dist/vuetify.min.css';
import 'material-design-icons-iconfont';

import App from './App.vue';

Vue.use(VueFlashMessage, {
	messageOptions: {
		timeout: 6000,
		important: false
	}
});

Vue.config.productionTip = false;
Vue.config.devtools = true;

router.beforeEach((to, from, next) => {
	document.title = "Go Desktop"
	next()
})

import Wails from '@wailsapp/runtime';

Wails.Init(() => {
	new Vue({
		router,
		render: h => h(App)
	}).$mount('#app');
});
