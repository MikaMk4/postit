<template>
    <div class="board-edit">
        <h1>Edit Board</h1>
        <form @submit.prevent="submitForm">
        <label for="title">Title</label>
        <input type="text" id="title" v-model="title" required />
        <label for="description">Description</label>
        <textarea id="description" v-model="description" required></textarea>
        <button type="submit" :disabled="!canEdit">Save</button>
        </form>
    </div>
</template>

<script setup>
import { ref, computed } from 'vue'
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
const canEdit = computed(() => {
    return title.value !== '' && description.value !== ''
})

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
}

.board-edit label {
    font-weight: bold;
    color: var(--text-primary-color);
}

.board-edit input, .board-edit textarea {
    padding: 0.5rem;
    margin: 0.5rem;
    border: 1px solid var(--accent-color);
    border-radius: 0.25rem;
    background-color: var(--background-color-secondary);
    color: var(--text-primary-color);
}

#description {
    height: 10rem;
    resize: vertical;
}
</style>