<template>
    <div class="settings">
        <h1>Settings</h1>
        <div class="user-settings" v-if="isAuthed">
            <h2>User Settings</h2>
            <AvatarPreview :avatar="userStore.user.avatar" :height="200" :width="200"/>
            <p>{{ userStore.user.name }}</p>
            <h3>Change Username</h3>
            <h3>Change Password</h3>
        </div>
        <div class="app-settings">
            <h2>App Settings</h2>
            <h3>Change Theme</h3>
            <h3>Change Language</h3>
        </div>
    </div>
</template>

<script setup>
import { computed } from 'vue';
import AvatarPreview from '@/components/AvatarPreview.vue';
import { useUserStore } from '@/stores/UserStore.js';

const userStore = useUserStore();

const isAuthed = computed(() => {
    return userStore.user !== null;
});
</script>

<style>
.settings {
    display: flex;
    flex-flow: column nowrap;
    justify-content: center;
    align-items: center;
}
</style>