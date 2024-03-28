<template>
  <div class="home">
    <h1>Home</h1>
    <p>Welcome to the home page.</p>
    <div class="posts">
      <MiniPost v-for="post in posts" :key="post.id" :post="post" />
    </div>
  </div>
</template>

<script setup>

import { ref } from 'vue'
import MiniPost from '@/components/MiniPost.vue'

const posts = ref(
  [
    {
      id: 1,
      title: 'First Post',
      content: 'This is the first post on the site.'
    },
    {
      id: 2,
      title: 'Second Post',
      content: 'This is the second post on the site.'
    },
    {
      id: 3,
      title: 'Third Post',
      content: 'This is the third post on the site.'
    }
  ]
)
</script>

<style>

.home {
  display: flex;
  flex-flow: column nowrap;
  justify-content: center;
}

.posts {
  display: flex;
  flex-flow: row wrap;
  justify-content: center;
}
</style>