<template>
    <div class="login-view">
        <div class="login" v-if="wantsToLogin">
            <h1>Login</h1>
            <AuthInput @submit="login" submitText="Login"/>
            <p>Don't have an account? <br><a @click="wantsToLogin = false">Sign up</a></p>
        </div>
        <div class="login" v-else>
            <h1>Sign In</h1>
            <AuthInput submitText="Sign In" />
            <p>Already have an account? <br><a @click="wantsToLogin = true">Login</a></p>
        </div>
    </div>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import AuthInput from '@/components/AuthInput.vue'
import { useUserStore } from '@/stores/UserStore.js'

const router = useRouter();
const userStore = useUserStore();

const wantsToLogin = ref(true);

function login(data) {
    if (data.username !== '' && data.password !== '') {
        userStore.login(data.username);
        router.push({ name: 'home' });
    } else {
        alert('Cannot leave fields blank.');
    }
}
</script>

<style>
.login-view {
    display: flex;
    flex-flow: column nowrap;
    justify-content: center;
    align-items: center;
}

.login {
    border: 1px solid #ccc;
    border-radius: 5px;
    margin: 10px;
    padding: 10px;
    width: 250px;
}

.login a {
    cursor: pointer;
}
</style>