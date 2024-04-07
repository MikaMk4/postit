<template>
    <div class="mini-board" @click="miniBoardClicked">
        <div class="mini-board__title">
        <h2>{{ board.name }}</h2>
        </div>
        <div class="mini-board__description">
        <p>{{ board.description }}</p>
        </div>
        <div class="mini-board__thumbnail" v-if="hasThumbnail">
        <img :src="board.thumbnail" alt="Thumbnail" />
        </div>
    </div>
</template>

<script setup>
import { useRouter } from 'vue-router';

const router = useRouter()

const props = defineProps({
    board: {
        type: Object,
        required: true
    }
})

const hasThumbnail = props.board.thumbnail != null

const miniBoardClicked = () => {
    router.push({ name: 'board', params: { id: props.board.id } })
}
</script>

<style>
.mini-board {
    border: 0.1rem solid var(--accent-color);
    border-radius: 0.4rem;
    margin: 1rem;
    padding: 1rem;
    width: 35rem;
}

.mini-board:hover {
    background-color: var(--accent-color);
    cursor: pointer;
}

.mini-board > * {
    margin: 0.5rem;
}
</style>