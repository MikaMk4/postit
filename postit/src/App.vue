<template>
    <metainfo>
    <template v-slot:title="{ content }">
    {{ content }}
    </template>
    </metainfo>
    <TopBar />
    <router-view />
</template>

<script setup>
import { onMounted } from 'vue'
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
})
</script>

<style lang="scss">
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}

* {
  box-sizing: border-box;
  margin: 0;
  padding: 0;
}

.router-view {
  display: flex;
  flex-flow: row nowrap;
  justify-content: center;
  align-items: center;
  height: 100vh;
  margin: 0 0;
}
</style>
