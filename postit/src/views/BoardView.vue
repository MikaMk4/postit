<template>
    <div class="board">
        <div class="board-header">
            <h1>{{ boardName }}</h1>
            <div class="delete-board" v-if="canEdit">
                <i class="fa fa-trash" :class="{ animated: appStore.animationsEnabled }" @click="deleteBoard"></i>
            </div>
        </div>
        <p>{{ boardDescription }}</p>
        <div v-if="canEdit">
            <router-link :to="{ name: 'edit-board', params: { id: boardId } }">
                <button>Edit this Board</button>
            </router-link>
        </div>
        <div>
            <router-link :to="{ name: 'create-post' }">
                <button>Create a Post</button>
            </router-link>
        </div>
        <div class="posts">
            <MiniPost v-for="post in posts" :key="post.id" :post="post" />
        </div>
    </div>
</template>

<script setup>
import { computed, onMounted, ref } from 'vue'
import MiniPost from '@/components/MiniPost.vue'
import { useMeta } from 'vue-meta'
import { useBoardStore } from '@/stores/BoardStore';
import { usePostStore } from '@/stores/PostStore';
import { useUserStore } from '@/stores/UserStore';
import { useAppStore } from '@/stores/AppStore';
import router from '@/router';

const appStore = useAppStore()

useMeta({
  title: 'Board',
  htmlAttrs: {
    lang: 'en',
    amp: true
  }
})

const boardStore = useBoardStore()
const postStore = usePostStore()

const boardId = router.currentRoute.value.params.id
const boardName = ref('')
const boardDescription = ref('')
const creatorId = ref(-1)

const posts = ref([])

onMounted(async () => {
    const id = router.currentRoute.value.params.id

    boardStore.fetchBoard(id).then((board) => {
        boardName.value = board.name
        boardDescription.value = board.description
        creatorId.value = board.creatorId
    }).catch((error) => {
        if (error === "Board not found") {
        router.push({ name: 'not-found' })
        }
    })

    await postStore.fetchPosts(id).then((ps) => {
        posts.value = posts.value.concat(ps)
    })
})

const userStore = useUserStore()
const canEdit = computed(() => {
  return userStore.user?.id === creatorId.value
})

function deleteBoard() {
    alert('Are you sure you want to delete this board?')
    boardStore.deleteBoard(boardId)
    router.push({ name: 'boards' })
}
</script>

<style>
.board-header {
    display: flex;
    flex-flow: row nowrap;
    justify-content: center;
    align-items: center;
}

.board-header > * {
    margin: 0 0.5rem;
}

.board {
    display: flex;
    flex-flow: column nowrap;
    justify-content: center;
}

.board > * {
    margin: 0.5rem;
}

.posts {
    display: flex;
    flex-flow: row wrap;
    justify-content: center;
}

.delete-board {
    display: flex;
    flex-flow: row nowrap;
    align-items: center;
    justify-content: center;
    font-size: 1.5rem;
}

.delete-board > i {
    color: grey;
    cursor: pointer;
}

.delete-board > i:hover {
    color: red;
}

.delete-board > i:hover.animated {
    animation: delete-animation 0.1s forwards linear;
}

@keyframes delete-animation {
    0% {
        transform: translate(0, 0);
    }
    25% {
        transform: translate(0.1rem, 0);
    }
    50% {
        transform: translate(0, 0);
    }
    75% {
        transform: translate(-0.1rem, 0);
    }
    100% {
        transform: translate(0, 0);
    }
}
</style>