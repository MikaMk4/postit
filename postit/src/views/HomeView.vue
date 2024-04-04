<template>
  <div class="home">
    <h1>Home</h1>
    <p>Welcome to the home page.</p>
  </div>
</template>

<script setup>
import { useMeta } from 'vue-meta'

useMeta({
  title: 'Home',
  htmlAttrs: {
    lang: 'en',
    amp: true
  }
})
</script>

<style>
.home {
  display: flex;
  flex-flow: column nowrap;
  justify-content: center;
}

.home * {
  margin: 0.5rem;
}
</style>