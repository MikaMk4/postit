<template>
    <div class="auth-input">
        <form @submit.prevent="submit">
            <label for="username" ref="usernameInput">Username</label>
            <input type="text" id="username" v-model="username" />
            <label for="password">Password</label>
            <input type="password" id="password" v-model="password" />
            <button @click="submit">{{ props.submitText }}</button>
        </form>
    </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'

const username = ref('');
const password = ref('');

const usernameInput = ref(null);

const props = defineProps({
    submitText: {
        type: String,
        default: 'Submit'
    }
});

const emit = defineEmits(['submit'])

function submit() {
    emit('submit', {
        username: username.value,
        password: password.value
    });
}

onMounted(() => {
    usernameInput.value.focus();
});
</script>

<style>
.auth-input {
    display: flex;
    flex-flow: column nowrap;
    justify-content: center;
    align-items: center;
    max-width: 800px;
    margin: 0 auto;
}

.auth-input * {
    margin: 0.25rem;
}
</style>