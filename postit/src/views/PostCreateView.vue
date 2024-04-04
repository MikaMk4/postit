<template>
    <div class="post-create">
        <h1>Create Post</h1>
        <form @submit.prevent="submitForm">
        <label for="title" ref="titleInput">Title</label>
        <input type="text" id="title" v-model="title" />
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
const isCreatable = computed(() => title.value.length > 0 && content.value.length > 0)
const boardName = computed(() => 'First Board')

function submitForm() {
    console.log('Title:', title.value)
    console.log('Content:', content.value)
    router.go(-1)
}

onMounted(() => {
    titleInput.value.focus()
})
</script>

<style>
.post-create {
    display: flex;
    flex-flow: column nowrap;
    justify-content: center;
    max-width: 800px;
    margin: 0 auto;
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
}

.post-create input, .post-create textarea {
    padding: 0.5rem;
    margin: 0.5rem;
}

#content {
    height: 10rem;
    resize: vertical;
}

.post-create button {
    padding: 0.5rem;
    margin: 0.5rem;
    background-color: #2c3e50;
    color: white;
    border: none;
    border-radius: 0.25rem;
}

.post-create button:hover {
    cursor: pointer;
    background-color: #34495e;
}

.post-create button:active {
    background-color: #2c3e50;
}

.post-create button:focus {
    outline: none;
}

.post-create button:disabled {
    background-color: #ccc;
    color: #666;
    cursor: not-allowed;
}

.post-create button:disabled:hover {
    background-color: #ccc;
}

.post-create button:disabled:active {
    background-color: #ccc;
}

.post-create button:disabled:focus {
    outline: none;
}
</style>