<template>
    <div class="not-found">
        <img src="../assets/404.png" alt="404 Not Found" />
        <p>POV: Your Page was lost in Bielefeld and ceased to exist.</p>
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

.not-found > img {
    width: 35rem;
    max-width: 100%;
    
}

.not-found > p {
    font-size: 2rem;
    font-weight: bold;
    margin-top: 1rem;
}
</style>