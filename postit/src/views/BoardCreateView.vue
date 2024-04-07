<template>
    <div class="board-create">
        <h1>Create a Board</h1>
        <form @submit.prevent="createBoard">
            <label for="name">Name</label>
            <input type="text" id="name" v-model="name" required />
            <label for="b-description">Description</label>
            <textarea id="b-description" v-model="description" required></textarea>
            <button type="submit" :disabled="!canEdit">Save</button>
        </form>
    </div>
</template>

<script setup>
import { computed, ref } from 'vue'
import { useRouter } from 'vue-router'
import { useMeta } from 'vue-meta'
import { useBoardStore } from '@/stores/BoardStore';

const boardStore = useBoardStore()

useMeta({
    title: 'Create Board',
    htmlAttrs: {
        lang: 'en',
        amp: true
    }
})

const router = useRouter()

const description = ref('')
const name = ref('')
async function createBoard() {
    await boardStore.createBoard({ name: name.value, description: description.value })
    router.go(-1)
}

const canEdit = computed(() => name.value.length > 0 && description.value.length > 0)
</script>

<style>
.board-create {
    display: flex;
    flex-flow: column nowrap;
    justify-content: center;
    max-width: 50rem;
}

.board-create * {
    margin: 0.5rem;
}

.board-create form {
    display: flex;
    flex-flow: column nowrap;
    justify-content: center;
}

.board-create label {
    font-weight: bold;
    color: var(--text-primary-color);
}

.board-create input, .board-create textarea {
    padding: 0.5rem;
    margin: 0.5rem;
    border: 0.1rem solid var(--accent-color);
    border-radius: 0.25rem;
    background-color: var(--background-color-secondary);
    color: var(--text-primary-color);
}

#b-description {
    height: 10rem;
    resize: vertical;
}
</style>