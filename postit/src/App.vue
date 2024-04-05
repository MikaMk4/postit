<template>
    <metainfo>
    <template v-slot:title="{ content }">
    {{ content }}
    </template>
    </metainfo>

    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/4.7.0/css/font-awesome.min.css">

    <TopBar />
    <router-view class="router-view"/>
</template>

<script setup>
import { onMounted, watch } from 'vue'
import { useAppStore } from '@/stores/AppStore.js'
import TopBar from '@/components/TopBar.vue'
import { useMeta } from 'vue-meta'

const appStore = useAppStore()

useMeta({
  title: '',
  htmlAttrs: {
    lang: 'en',
    amp: true
  }
})

onMounted(() => {
  appStore.setHistoryCount(window.history.length)

  if (appStore.darkMode) {
    document.documentElement.classList.add('dark-theme')
  }
})

watch(() => appStore.darkMode, (value) => {
  if (value) {
    document.documentElement.classList.add('dark-theme')
  } else {
    document.documentElement.classList.remove('dark-theme')
  }
})
</script>

<style>
:root {
  --background-color-primary: #ebebeb;
  --background-color-secondary: #fafafa;
  --accent-color: #cacaca;
  --accent-color-hover: #b0b0b0;
  --accent-color-active: #a0a0a0;
  --accent-color-inactive: #d0d0d0;
  --text-primary-color: #222;
  --element-size: 4rem;
}

/* Define styles for the root window with dark - mode preference */
:root.dark-theme {
  --background-color-primary: #1e1e1e;
  --background-color-secondary: #2d2d30;
  --accent-color: #3f3f3f;
  --accent-color-hover: #4f4f4f;
  --accent-color-active: #5f5f5f;
  --accent-color-inactive: #2f2f2f;
  --text-primary-color: #ddd;
}

h1, h2, h3, h4, h5, h6 {
  color: var(--text-primary-color);
}

p {
  color: var(--text-primary-color);
}

a {
  text-decoration: none;
  color: var(--text-primary-color);
}

button {
  background-color: transparent;
  color: var(--text-primary-color);
  border: 1px solid var(--accent-color);
  border-radius: 5px;
  padding: 0.5rem 1rem;
  cursor: pointer;
}

button:hover {
  background-color: var(--accent-color);
}

button:active {
  background-color: var(--accent-color-active);
}

button:focus {
  outline: none;
}

button:disabled {
  color: #666;
  cursor: not-allowed;
}

button:disabled:hover {
  background-color: var(--accent-color-inactive);
}

#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
}

* {
  box-sizing: border-box;
  margin: 0;
  padding: 0;
}

body {
  background-color: var(--background-color-primary);
  color: var(--text-primary-color);
  font-size: 1rem;
  line-height: 1.5;
  margin: 0;
  padding: 0;
}

.router-view {
  margin: 3rem auto;
  width: 80%;
}
</style>
