<template>
    <div class="board-list">
        <h1>Boards</h1>
        <div class="boards">
            <MiniBoard v-for="board in boards" :key="board.id" :board="board" />
        </div>
    </div>
</template>

<script setup>
import { ref } from 'vue'
import MiniBoard from '@/components/MiniBoard.vue'
import { useMeta } from 'vue-meta'

useMeta({
    title: 'Boards',
    htmlAttrs: {
        lang: 'en',
        amp: true
    }
})

const boards = ref(
    [
        {
            id: 1,
            title: 'First Board',
            description: 'This is the first board on the site.',
            thumbnail: 'https://via.placeholder.com/150'
        },
        {
            id: 2,
            title: 'Second Board',
            description: 'This is the second board on the site.',
            thumbnail: 'https://via.placeholder.com/150'
        },
        {
            id: 3,
            title: 'Third Board',
            description: 'This is the third board on the site.',
            thumbnail: 'https://via.placeholder.com/150'
        }
    ]
)
</script>

<style>
.boards {
    display: flex;
    flex-flow: row wrap;
    justify-content: center;
}
</style>