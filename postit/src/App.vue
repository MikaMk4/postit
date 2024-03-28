<template>
  <nav>
    <router-link :to="{ name: 'home' }">Home</router-link> |
    <router-link :to="{ name: 'about' }">About</router-link> |
    <router-link :to="{ name: 'settings' }">Settings</router-link> |
    <router-link :to="{ name: 'login' }" v-if="!isAuthed">Login</router-link>
    <span v-else>Log off</span>
  </nav>
  <router-view/>
</template>

<script setup>
import { computed } from 'vue'
import { useUserStore } from '@/stores/UserStore.js'

const userStore = useUserStore()

const isAuthed = computed(() => {
  return userStore.user !== null
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
}

nav {
  padding: 30px;

  a {
    font-weight: bold;
    color: #2c3e50;

    &.router-link-exact-active {
      color: #42b983;
    }
  }
}
</style>
