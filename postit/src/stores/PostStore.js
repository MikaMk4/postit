import { defineStore } from 'pinia'
import { ref } from 'vue'
import { useUserStore } from './UserStore'

const userStore = useUserStore()

export const usePostStore = defineStore('postStore', {
    state: () => ({
        posts: ref([])
    }),
    actions: {
        async fetchPosts(bid) {
            const response = await fetch(`http://localhost:3000/api/v1/board/${bid}/posts`)

            const data = await response.json()

            if (!response.ok) {
                throw new Error(data.error || 'Failed to fetch posts')
            }

            data.likeCount = parseInt(data.likeCount)

            return data
        },
        async fetchPostsByUser(uid) {
            const response = await fetch(`http://localhost:3000/api/v1/user/${uid}/posts`)

            const data = await response.json()

            if (!response.ok) {
                throw new Error(data.error || 'Failed to fetch posts')
            }

            data.likeCount = parseInt(data.likeCount)

            return data
        },
        async fetchPost(bid, pid) {
            const response = await fetch(`http://localhost:3000/api/v1/board/${bid}/posts/${pid}`)

            const data = await response.json()

            if (!response.ok) {
                if (data.error.includes('no rows in result set')) {
                    throw "Post not found"
                }

                throw new Error(data.error || 'Failed to fetch post')
            }

            data.likeCount = parseInt(data.likeCount)

            return data
        },
        async createPost(post) {
            const response = await fetch(`http://localhost:3000/api/v1/board/${post.boardId}/posts`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                    'Authorization': `${userStore.user.token}`
                },
                body: JSON.stringify(post)
            })

            const data = await response.json()

            if (!response.ok) {
                throw new Error(data.error || 'Failed to create post')
            }

            this.posts.push(data)
        },
        async updatePost(post) {
            post.id = parseInt(post.id)

            const response = await fetch(`http://localhost:3000/api/v1/board/${post.boardId}/posts/${post.id}`, {
                method: 'PUT',
                headers: {
                    'Content-Type': 'application/json',
                    'Authorization': `${userStore.user.token}`
                },
                body: JSON.stringify(post)
            })

            const data = await response.json()

            if (!response.ok) {
                throw new Error(data.error || 'Failed to update post')
            }
        },
        async deletePost(bid, pid) {
            const response = await fetch(`http://localhost:3000/api/v1/board/${bid}/posts/${pid}`, {
                method: 'DELETE',
                headers: {
                    'Authorization': `${userStore.user.token}`
                }
            })

            const data = await response.json()

            if (!response.ok) {
                throw new Error(data.error || 'Failed to delete post')
            }
        }
    }
})