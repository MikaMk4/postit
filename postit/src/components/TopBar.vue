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
              <i class="fa fa-gear" :class="{ animated: appStore.animationsEnabled }" id="settings-button"></i>
            </router-link>
            <router-link :to="{ name: 'login' }" v-if="!isAuthed">
                <button>
                    Login
                </button>
            </router-link>
            <div v-else>
                <AvatarPreview :avatar="userStore.user.avatar" :pId="userStore.user.id"/>
            </div>
        </div>
    </div>
</template>

<script setup>
import { computed } from 'vue'
import AvatarPreview from '@/components/AvatarPreview.vue'
import { useUserStore } from '@/stores/UserStore.js'
import { useAppStore } from '@/stores/AppStore.js'

const userStore = useUserStore()
const appStore = useAppStore()

const isAuthed = computed(() => {
  return userStore.user !== null
})
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

#settings-button {
  font-size: 2.2rem;
  color: var(--text-primary-color);
  border: none;
}

#settings-button:hover {
  color: var(--accent-color-active);
}

#settings-button.animated:hover {
  animation: settings-button-animation .5s;
}

@keyframes settings-button-animation {
  0% {
    transform: rotate(0deg);
  }
  100% {
    transform: rotate(360deg);
  }
}
</style>