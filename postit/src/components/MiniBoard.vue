<template>
    <div class="mini-board" @click="miniBoardClicked">
        <div class="mini-board__title">
        <h2>{{ board.title }}</h2>
        </div>
        <div class="mini-board__description">
        <p>{{ board.description }}</p>
        </div>
        <div class="mini-board__thumbnail">
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

const miniBoardClicked = () => {
    router.push({ name: 'board', params: { id: props.board.id } })
}
</script>

<style>
.mini-board {
    border: 1px solid var(--accent-color);
    border-radius: 5px;
    margin: 10px;
    padding: 10px;
    width: 500px;
}

.mini-board:hover {
    background-color: var(--accent-color);
    cursor: pointer;
}

.mini-board > * {
    margin: 0.5rem;
}
</style>