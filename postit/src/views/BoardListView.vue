<template>
    <div class="board-list">
        <h1>Boards</h1>
        <div class="board-creator" v-if="userStore.user != null">
            <RouterLink :to="{ name: 'create-board', params: { } }">
                <button>Create a Board</button>
            </RouterLink>
        </div>
        <div class="boards">
            <MiniBoard v-for="board in boards" :key="board.id" :board="board" />
        </div>
    </div>
</template>

<script setup>
import { computed, onMounted } from 'vue'
import MiniBoard from '@/components/MiniBoard.vue'
import { useMeta } from 'vue-meta'
import { useBoardStore } from '@/stores/BoardStore'
import { useUserStore } from '@/stores/UserStore'

const userStore = useUserStore()

useMeta({
    title: 'Boards',
    htmlAttrs: {
        lang: 'en',
        amp: true
    }
})

const boardStore = useBoardStore()
const boards = computed(() => boardStore.boards)

onMounted(() => {
    boardStore.fetchBoards()
})
</script>

<style>
.boards {
    display: flex;
    flex-flow: row wrap;
    justify-content: center;
}
</style>