<template>
    <div class="post-create">
        <h1>Create Post</h1>
        <form @submit.prevent="submitForm">
            <label for="title" ref="titleInput">Title</label>
            <input type="text" id="title" v-model="title" />
            
            <label for="thumbnail">Thumbnail</label>
            <input type="text" id="thumbnail" v-model="thumbnail" />

            <label for="content">Content</label>
            <textarea id="content" v-model="content"></textarea>
        <button type="submit" :disabled="!isCreatable">Publish to {{ boardName }}</button>
        </form>
    </div>
</template>

<script setup>
import { ref, onMounted, computed } from 'vue'
import { useRouter } from 'vue-router'
import { useMeta } from 'vue-meta'
import { usePostStore } from '@/stores/PostStore'
import { useBoardStore } from '@/stores/BoardStore';

useMeta({
    title: 'Create Post',
    htmlAttrs: {
        lang: 'en',
        amp: true
    }
})

const router = useRouter()
const titleInput = ref(null)

const title = ref('')
const content = ref('')
const thumbnail = ref('')
const isCreatable = computed(() => title.value.length > 0 && content.value.length > 0)

const boardStore = useBoardStore()
const boardName = ref('')

const postStore = usePostStore()
function submitForm() {
    postStore.createPost({
        title: title.value,
        content: content.value,
        thumbnail: thumbnail.value,
        boardId: router.currentRoute.value.params.id
    })
    router.go(-1)
}

onMounted(() => {
    titleInput.value.focus()

    const id = router.currentRoute.value.params.id
    boardStore.fetchBoard(id).then((board) => {
        boardName.value = board.name
    }).catch((error) => {
        if (error === "Board not found") {
            router.push({ name: 'not-found' })
        }
    })
})
</script>

<style>
.post-create {
    display: flex;
    flex-flow: column nowrap;
    justify-content: center;
    max-width: 50rem;
}

.post-create * {
    margin: 0.5rem;
}

.post-create form {
    display: flex;
    flex-flow: column nowrap;
    justify-content: center;
}

.post-create label {
    font-weight: bold;
    color: var(--text-primary-color);
}

.post-create input, .post-create textarea {
    padding: 0.5rem;
    margin: 0.5rem;
    border: 0.1rem solid var(--accent-color);
    border-radius: 0.25rem;
    background-color: var(--background-color-secondary);
    color: var(--text-primary-color);
}

#content {
    height: 10rem;
    resize: vertical;
}
</style>