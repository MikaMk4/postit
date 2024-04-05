<template>
    <div class="not-found">
        <img src="/assets/404.svg" alt="404 Not Found" />
        <p>Sorry, the page you are looking for does not exist.</p>
    </div>
</template>

<script setup>
import { useMeta } from 'vue-meta'

useMeta({
  title: 'Not Found',
  htmlAttrs: {
    lang: 'en',
    amp: true
  }
})
</script>

<style>
.not-found {
    display: flex;
    flex-flow: column nowrap;
    justify-content: center;
    align-items: center;
}
</style>