import { defineStore } from 'pinia'
import { ref } from 'vue'

export const useUserStore = defineStore('userStore', {
    persist: true,
    state: () => ({
        user: ref(null)
    }),
    actions: {
        async login(username, password) {
            const response = await fetch('http://localhost:3000/api/v1/login', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({ username, password })
            })

            const data = await response.json()

            if (!response.ok) {
                throw new Error(data.error || 'Failed to login')
            }

            this.user = {
                token: data.token,
            }
            await this.fetchUser()
        },
        async signUp(username, password) {
            const response = await fetch('http://localhost:3000/api/v1/signup', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({ username, password })
            })

            const data = await response.json()

            if (!response.ok) {
                if (data.error.includes('username')) {
                    throw "Username already exists"
                }

                throw new Error(data.error || 'Failed to signup')
            }

            this.user = {
                token: data.token,
            }
            await this.fetchUser()
            this.user.avatar = 'https://upload.wikimedia.org/wikipedia/commons/thumb/2/2c/Default_pfp.svg/510px-Default_pfp.svg.png'
            await this.updateUser()
            await this.fetchUser()
        },
        async fetchUser() {
            const response = await fetch('http://localhost:3000/api/v1/user', {
                headers: {
                    Authorization: `${this.user.token}`
                }
            })

            const data = await response.json()

            if (!response.ok) {
                if (data.error.includes('1062')) {
                    throw "Username already exists"
                }

                throw new Error(data.error || 'Failed to fetch user')
            }

            this.user = {
                ...this.user,
                ...data
            }
        },
        async fetchUserById(uid) {
            const response = await fetch(`http://localhost:3000/api/v1/users/${uid}`)

            const data = await response.json()

            if (!response.ok) {
                throw new Error(data.error || 'Failed to fetch user')
            }

            return data
        },
        async updateUser() {
            const response = await fetch('http://localhost:3000/api/v1/user', {
                method: 'PUT',
                headers: {
                    'Content-Type': 'application/json',
                    Authorization: `${this.user.token}`
                },
                body: JSON.stringify(this.user)
            })

            const res = await response.json()

            if (!response.ok) {
                if (res.error.includes('1062')) {
                    throw "Username already exists"
                }

                throw new Error(res.error || 'Failed to update user')
            }
        },
        async startEdit() {
            this.ogUser = { ...this.user }
        },
        async cancelEdit() {
            this.user = { ...this.ogUser }
        },

        logout() {
            this.user = null
        }
    }
})