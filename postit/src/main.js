import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import { createPinia } from 'pinia'
import piniaPluginPersistedState from 'pinia-plugin-persistedstate'
import { createMetaManager } from 'vue-meta'

const pinia = createPinia()
pinia.use(piniaPluginPersistedState)

const metaManager = createMetaManager()

const app = createApp(App)
app.use(router)
app.use(pinia)
app.use(metaManager)
app.mount('#app')
