<template>
    <div class="board-create">
        <h1>Create a Board</h1>
        <form @submit.prevent="createBoard">
        <label for="name">Name</label>
        <input type="text" id="name" v-model="name" required />
        <button type="submit">Create</button>
        </form>
    </div>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useMeta } from 'vue-meta'

useMeta({
    title: 'Create Board',
    htmlAttrs: {
        lang: 'en',
        amp: true
    }
})

const router = useRouter()

const name = ref('')
function createBoard() {
    console.log('Name:', name.value)
    router.push({ name: 'boards' })
}
</script>

<style>
.board-create {
    display: flex;
    flex-flow: column nowrap;
    justify-content: center;
    max-width: 800px;
    margin: 0 auto;
}

.board-create * {
    margin: 0.5rem;
}

.board-create form {
    display: flex;
    flex-flow: column nowrap;
    justify-content: center;
    align-items: center;
}

.board-create form * {
    margin: 0.25rem;
}
</style>