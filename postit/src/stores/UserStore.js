import { defineStore } from 'pinia'

export const useUserStore = defineStore('userStore', {
    persist: true,
    state: () => ({
        loading: false,
        user: null
    }),
    actions: {
        async login(username) {
            this.loading = true

            // const response = await fetch('http://localhost:3000/api/v1/login', {
            //     method: 'POST',
            //     headers: {
            //         'Content-Type': 'application/json'
            //     },
            //     body: JSON.stringify({ username, password })
            // })

            // const data = await response.json()

            // if (!response.ok) {
            //     throw new Error(data.message || 'Failed to login')
            // }

            this.user = {
                name: username,
                avatar: 'https://cdn.discordapp.com/avatars/452415473687068672/1c6bad1e46a6612c40d20d1ac2f61c7e.webp?size=1024'
            }

            this.loading = false
        },
        async register(username) {
            this.loading = true

            // const response = await fetch('http://localhost:3000/api/v1/signup', {
            //     method: 'POST',
            //     headers: {
            //         'Content-Type': 'application/json'
            //     },
            //     body: JSON.stringify({ username, password })
            // })

            // const data = await response.json()

            // if (!response.ok) {
            //     throw new Error(data.message || 'Failed to register')
            // }

            this.user = {
                name: username,
                avatar: 'https://cdn.discordapp.com/avatars/452415473687068672/1c6bad1e46a6612c40d20d1ac2f61c7e.webp?size=1024'
            }

            this.loading = false
        },
        logout() {
            this.user = null
        }
    }
})