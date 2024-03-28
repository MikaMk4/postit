import { defineStore } from 'pinia'

export const useUserStore = defineStore('userStore', {
    state: () => ({
        user: null
    }),
    actions: {
        login(username) {
            this.user = {
                name: username,
                avatar: 'https://cdn.discordapp.com/avatars/452415473687068672/1c6bad1e46a6612c40d20d1ac2f61c7e.webp?size=1024'
            }
        }
    }
})