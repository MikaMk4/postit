<template>
  <div class="navbar">
    <div>
      <h1>PostIt</h1>
      <button>
        <router-link :to="{ name: 'home' }">Home</router-link>
      </button> |
      <button>
        <router-link :to="{ name: 'about' }">About</router-link>
      </button> |
      <button>
        <router-link :to="{ name: 'settings' }">Settings</router-link>
      </button>
    </div>
    <div>
      <button v-if="!isAuthed">
        <router-link :to="{ name: 'login' }">Login</router-link>
      </button>
      <AvatarPreview v-if="isAuthed" :avatar="userStore.user.avatar"/>
    </div>
  </div>
  <nav>
    
  </nav>
  <router-view/>
</template>

<script setup>
import { computed } from 'vue'
import { useUserStore } from '@/stores/UserStore.js'
import AvatarPreview from '@/components/AvatarPreview.vue'

const userStore = useUserStore()

const isAuthed = computed(() => {
  return userStore.user !== null
})

// function logout() {
//   userStore.logout()
// }
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

.navbar {
  display: flex;
  flex-flow: row nowrap;
  justify-content: space-between;
  align-items: center;
  margin: 10px;
  text-decoration: none;
  max-height: 50px;
}

.navbar > * {
  display: flex;
  flex-flow: row nowrap;
  align-items: center;
}

.navbar > * > * {
  margin: 0 10px;
}

button {
  background-color: transparent;
  border: 1px solid #ccc;
  border-radius: 5px;
  padding: 5px 10px;
  cursor: pointer;
}

button a {
  text-decoration: none;
  color: #2c3e50;
}

button:hover {
  background-color: #f0f0f0;
}
</style>
