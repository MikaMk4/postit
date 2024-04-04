<template>
    <div class="navbar">
        <div>
            <router-link :to="{ name: 'home' }">
                <h1>PostIt</h1>
            </router-link>
            <router-link :to="{ name: 'about' }">
                <button>
                    About
                </button>
            </router-link>
            <router-link :to="{ name: 'boards' }">
                <button>
                    Boards
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
                    <router-link :to="{ name: 'settings' }">
                        <button>
                            Settings
                        </button>
                    </router-link>
                    <button @click="logout" id="logout-button">
                        Logout
                    </button>
                </div>
            </div>
        </div>
    </div>
</template>

<script setup>
import { computed } from 'vue'
import AvatarPreview from '@/components/AvatarPreview.vue'
import { useUserStore } from '@/stores/UserStore.js'

const userStore = useUserStore()

const isAuthed = computed(() => {
  return userStore.user !== null
})

function logout() {
  userStore.logout()
}
</script>

<style>
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

.navbar h1 {
  margin: 20px;
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

#logout-button {
  background-color: red;
  color: white;
}

a {
  text-decoration: none;
  color: #2c3e50;
}

button:hover {
  background-color: #f0f0f0;
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
  border-radius: 0%;
  border: none;
  display: block;
  width: 100%;
}

.profile-dropdown:hover .profile-dropdown-content {
  display: block;
}

.profile-dropdown AvatarPreview {
  background-color: #f9f9f9;
}
</style>