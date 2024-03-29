<template>
  <div class="navbar">
    <div>
      <h1>PostIt</h1>
      <router-link :to="{ name: 'home' }">
        <button>
          Home
        </button>
      </router-link>
      <router-link :to="{ name: 'about' }">
        <button>
          About
        </button>
      </router-link>
      <router-link :to="{ name: 'settings' }">
        <button>
          Settings
        </button>
      </router-link>
    </div>
    <div>
      <router-link :to="{ name: 'login' }">
        <button v-if="!isAuthed">
          Login
        </button>
      </router-link>
      <div class="profile-dropdown">
        <AvatarPreview v-if="isAuthed" :avatar="userStore.user.avatar"/>
        <div v-if="isAuthed" class="profile-dropdown-content">
          <button @click="logout">
            Logout
          </button>
        </div>
      </div>
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

function logout() {
  userStore.logout()
}
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
  position: sticky;
  top: 0;
  padding: 10px 0;
  text-decoration: none;
  height: 60px;
  width: 100%;
  background-color: #ffffff;
  z-index: 1000;
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

.router-link-exact-active button {
  background-color: #6f6f6f;
  color: white;
}

.router-link-exact-active button:hover {
  background-color: #808080;
}

.profile-dropdown {
  position: relative;
  display: inline-block;
}

.profile-dropdown-content {
  display: none;
  position: absolute;
  background-color: #f9f9f9;
  box-shadow: 0px 8px 16px 0px rgba(0,0,0,0.2);
  z-index: 1;
  left: -50%;
}

.profile-dropdown-content button {
  color: black;
  padding: 12px 16px;
  text-decoration: none;
  display: block;
}

.profile-dropdown:hover .profile-dropdown-content {
  display: block;
}

.profile-dropdown AvatarPreview {
  background-color: #f9f9f9;
}
</style>
