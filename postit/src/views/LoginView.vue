<template>
    <div class="login-view">
        <div>
            <div class="login" v-if="wantsToLogin">
                <h1>Login</h1>
                <AuthInput @submit="login" submitText="Login"/>
                <p>Don't have an account? <br><a @click="toggleLogin(false)">Sign up</a></p>
            </div>
            <div class="login" v-else>
                <h1>Sign Up</h1>
                <AuthInput submitText="Sign Up" />
                <p>Already have an account? <br><a @click="toggleLogin(true)">Login</a></p>
            </div>
        </div>
        <div class="loader" v-if="loading">

        </div>
    </div>
</template>

<script setup>
import { ref, onBeforeMount } from 'vue'
import { useRouter } from 'vue-router'
import AuthInput from '@/components/AuthInput.vue'
import { useUserStore } from '@/stores/UserStore.js'
import { useAppStore } from '@/stores/AppStore.js'
import { useMeta } from 'vue-meta'

const router = useRouter();
const userStore = useUserStore();
const appStore = useAppStore();

const wantsToLogin = ref(true);
const loading = ref(false);

useMeta({
  title: 'Login',
  htmlAttrs: {
    lang: 'en',
    amp: true
  }
})

function toggleLogin(value) {
    wantsToLogin.value = value;
}

async function login(data) {
    if (data.username !== '' && data.password !== '') {
        loading.value = true;
        await userStore.login(data.username);
        loading.value = false;
        router.go(-1);
    } else {
        alert('Cannot leave fields blank.');
    }
}

onBeforeMount(() => {
    if (userStore.user !== null) {
        if (window.history.length <= appStore.historyCount)
            router.push({ name: 'home' });
        else
            router.go(-1);
    }
});
</script>

<style>
.login-view {
    display: flex;
    flex-flow: column nowrap;
    justify-content: center;
    align-items: center;
}

.login {
    border: 0.1rem solid var(--accent-color);
    border-radius: 0.4rem;
    padding: 0.6rem;
    width: 16rem;
    background-color: var(--background-color-secondary);
}

.login > * {
    margin: 0.5rem;
}

.login a {
    cursor: pointer;
    color: blue;
}

.loader {
  border: 1rem solid var(--background-color-secondary);
  border-top: 1rem solid var(--accent-color);
  border-radius: 50%;
  width: 7.5rem;
  height: 7.5rem;
  margin: 2rem auto;
  animation: spin 2s linear infinite;
}

@keyframes spin {
  0% { transform: rotate(0deg); }
  100% { transform: rotate(360deg); }
}
</style>