<template>
    <div class="auth-input">
        <label for="username">Username</label>
        <input type="text" id="username" v-model="username" />
        <label for="password">Password</label>
        <input type="password" id="password" v-model="password" />
        <button @click="submit">{{ props.submitText }}</button>
    </div>
</template>

<script setup>
import { ref, defineEmits, defineProps } from 'vue'

const username = ref('');
const password = ref('');

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
</script>

<style>
.auth-input {
    display: flex;
    flex-flow: column nowrap;
    justify-content: center;
    align-items: center;
}
</style>