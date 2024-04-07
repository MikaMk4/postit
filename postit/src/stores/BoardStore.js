import { defineStore } from 'pinia'
import { ref } from 'vue'
import { useUserStore } from './UserStore'

const userStore = useUserStore()

export const useBoardStore = defineStore('boardStore', {
    state: () => ({
        boards: ref([])
    }),
    actions: {
        async fetchBoards() {
            const response = await fetch('http://localhost:3000/api/v1/boards')

            const data = await response.json()

            if (!response.ok) {
                throw new Error(data.error || 'Failed to fetch boards')
            }

            this.boards = data
        },
        async fetchBoard(id) {
            const response = await fetch(`http://localhost:3000/api/v1/boards/${id}`)

            const data = await response.json()

            if (!response.ok) {
                if (data.error.includes('no rows in result set')) {
                    throw "Board not found"
                }

                throw new Error(data.error || 'Failed to fetch board')
            }

            return data
        },
        async createBoard(board) {
            board = {
                ...board,
                creatorId: userStore.user.id
            }

            const response = await fetch('http://localhost:3000/api/v1/boards', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                    Authorization: `${userStore.user.token}`
                },
                body: JSON.stringify(board)
            })

            const data = await response.json()

            if (!response.ok) {
                throw new Error(data.error || 'Failed to create board')
            }

            this.boards.push(data)
        },
        async deleteBoard(id) {
            const response = await fetch(`http://localhost:3000/api/v1/boards/${id}`, {
                method: 'DELETE',
                headers: {
                    Authorization: `${userStore.user.token}`
                }
            })

            if (!response.ok) {
                throw new Error('Failed to delete board')
            }
        },
        async updateBoard(board) {
            const response = await fetch(`http://localhost:3000/api/v1/boards/${board.id}`, {
                method: 'PUT',
                headers: {
                    'Content-Type': 'application/json',
                    Authorization: `${userStore.user.token}`
                },
                body: JSON.stringify(board)
            })

            const data = await response.json()

            if (!response.ok) {
                throw new Error(data.error || 'Failed to update board')
            }
        }
    }
})