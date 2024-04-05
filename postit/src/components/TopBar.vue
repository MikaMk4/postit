<template>
    <div class="navbar">
        <div class="left-nav">
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
        <div class="right-nav">
            <router-link :to="{ name: 'settings' }">
                <button>
                    Settings
                </button>
            </router-link>
            <router-link :to="{ name: 'login' }" v-if="!isAuthed">
                <button>
                    Login
                </button>
            </router-link>
            <div class="profile-dropdown" v-else>
                <AvatarPreview :avatar="userStore.user.avatar"/>
                <div class="profile-dropdown-content">
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
  padding: 0.5rem 0;
  text-decoration: none;
  height: 4rem;
  width: 100%;
  background-color: var(--background-color-secondary);
  z-index: 1000;
}

.navbar > * {
  display: flex;
  flex-flow: row nowrap;
  align-items: center;
}

.left-nav > * {
  margin-left: 1.25rem;
}

.right-nav > * {
  margin-right: 1.25rem;
}

#logout-button {
  color: red;
  font-weight: bold;
}

.profile-dropdown {
  position: relative;
  display: inline-block;
}

.profile-dropdown-content {
  display: none;
  position: absolute;
  background-color: var(--background-color-secondary);
  box-shadow: 0 0.5rem 1rem 0 rgba(0,0,0,0.2);
  z-index: 1;
  left: -50%;
}

.profile-dropdown-content button {
  padding: 0.75rem 1rem;
  border-radius: 0%;
  overflow: hidden;
  border: none;
  display: block;
  width: 100%;
}

.profile-dropdown:hover .profile-dropdown-content {
  display: block;
}
</style>