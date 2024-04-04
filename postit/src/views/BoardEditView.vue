<template>
    <div class="board-edit">
        <h1>Edit Board</h1>
        <form @submit.prevent="submitForm">
        <label for="title">Title</label>
        <input type="text" id="title" v-model="title" required />
        <label for="description">Description</label>
        <textarea id="description" v-model="description" required></textarea>
        <button type="submit">Save</button>
        </form>
    </div>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useMeta } from 'vue-meta'

useMeta({
    title: 'Edit Board',
    htmlAttrs: {
        lang: 'en',
        amp: true
    }
})

const router = useRouter()

const title = ref('First Board')
const description = ref('This is the first board on the site.')
function submitForm() {
    console.log('Title:', title.value)
    console.log('Description:', description.value)
    router.push({ name: 'boards' })
}
</script>

<style>
.board-edit {
    display: flex;
    flex-flow: column nowrap;
    justify-content: center;
    max-width: 800px;
    margin: 0 auto;
}

.board-edit * {
    margin: 0.5rem;
}

.board-edit form {
    display: flex;
    flex-flow: column nowrap;
    justify-content: center;
    align-items: center;
}

.board-edit form * {
    margin: 0.25rem;
}

.board-edit button {
    padding: 0.5rem 1rem;
    background-color: #007bff;
    color: white;
    border: none;
    border-radius: 0.25rem;
    cursor: pointer;
}

.board-edit button:hover {
    background-color: #0056b3;
}

.board-edit input, .board-edit textarea {
    padding: 0.5rem;
    margin: 0.5rem;
}

.board-edit textarea {
    height: 10rem;
    resize: vertical;
}

.board-edit label {
    font-weight: bold;
}

.board-edit button:active {
    background-color: #0056b3;
}

.board-edit button:focus {
    outline: none;
}

.board-edit button:disabled {
    background-color: #ccc;
}

.board-edit button:disabled:hover {
    cursor: not-allowed;
}
</style>