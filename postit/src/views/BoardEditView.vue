<template>
    <div class="board-edit">
        <h1>Edit Board</h1>
            <form @submit.prevent="submitForm">
            <label for="title">Name</label>
            <input type="text" id="title" v-model="title" required />
            <label for="thumbnail">Thumbnail Url</label>
            <input type="text" id="thumbnail" v-model="thumbnail" />
            <label for="description">Description</label>
            <textarea id="description" v-model="description" required></textarea>
            <button type="submit" :disabled="!canEdit">Save</button>
        </form>
    </div>
</template>

<script setup>
import { ref, computed, onBeforeMount } from 'vue'
import { useRouter } from 'vue-router'
import { useMeta } from 'vue-meta'
import { useBoardStore } from '@/stores/BoardStore'

useMeta({
    title: 'Edit Board',
    htmlAttrs: {
        lang: 'en',
        amp: true
    }
})

const router = useRouter()

const title = ref('')
const thumbnail = ref('')
const description = ref('')
const canEdit = computed(() => {
    return title.value !== '' && description.value !== ''
})

function submitForm() {
    boardStore.updateBoard({
        id: router.currentRoute.value.params.id,
        name: title.value,
        thumbnail: thumbnail.value,
        description: description.value
    })
    router.go(-1)
}

const boardStore = useBoardStore()
onBeforeMount(() => {
    const id = router.currentRoute.value.params.id
    boardStore.fetchBoard(id).then((board) => {
        title.value = board.name,
        thumbnail.value = board.thumbnail,
        description.value = board.description
    })
})
</script>

<style>
.board-edit {
    display: flex;
    flex-flow: column nowrap;
    justify-content: center;
    max-width: 50rem;
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
    border: 0.1rem solid var(--accent-color);
    border-radius: 0.25rem;
    background-color: var(--background-color-secondary);
    color: var(--text-primary-color);
}

#description {
    height: 10rem;
    resize: vertical;
}
</style>