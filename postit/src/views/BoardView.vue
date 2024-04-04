<template>
    <div class="board">
        <h1>{{ board.title }}</h1>
        <div class="postCreator">
            <RouterLink :to="{ name: 'create-post' }">
                <button>Create a Post</button>
            </RouterLink>
        </div>
        <div class="posts">
            <MiniPost v-for="post in board.posts" :key="post.id" :post="post" />
        </div>
    </div>
</template>

<script setup>
import { ref } from 'vue'
import MiniPost from '@/components/MiniPost.vue'
import { useMeta } from 'vue-meta'

useMeta({
  title: 'Board',
  htmlAttrs: {
    lang: 'en',
    amp: true
  }
})

const board = ref(
    {
        id: 1,
        title: 'First Board',
        posts: [
            {
                id: 1,
                title: 'First Post',
                content: 'This is the first post on this board.'
            },
            {
                id: 2,
                title: 'Second Post',
                content: 'This is the second post on this board.'
            },
            {
                id: 3,
                title: 'Third Post',
                content: 'This is the third post on this board.'
            }
        ]
    }
)
</script>

<style>
.posts {
    display: flex;
    flex-flow: row wrap;
    justify-content: center;
}
</style>