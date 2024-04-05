<template>
    <div class="settings">
        <h1>Settings</h1>
        <div class="app-settings">
            <h2>App Settings</h2>
            <div>
                <h3>Dark Mode</h3>
                <ToggleSwitch v-model="isDarkMode" :animationsEnabled="animationsEnabled" />
            </div>
            <div>
                <h3>Enable Animations</h3>
                <ToggleSwitch v-model="animationsEnabled" :animationsEnabled="animationsEnabled" />
            </div>
        </div>
        <div class="user-settings" v-if="isAuthed">
            <h2>User Settings</h2>
            <AvatarPreview :avatar="userStore.user.avatar" :size="200" :isEditable="true" @changeAvatar="changeAvatar"/>
            <form>
                <div>
                    <label for="username">New Username</label>
                    <input type="text" id="username" v-model="userStore.user.name" />
                </div>
                <div>
                    <label for="password">New Password</label>
                    <input type="password" id="password" />
                    <label for="confirm-password">Confirm New Password</label>
                    <input type="password" id="confirm-password" />
                </div>
                <button type="submit">Save Changes</button>
            </form>
            <button @click="logout">Logout</button>
        </div>
    </div>
</template>

<script setup>
import { computed } from 'vue';
import { useMeta } from 'vue-meta';
import AvatarPreview from '@/components/AvatarPreview.vue';
import ToggleSwitch from '@/components/ToggleSwitch.vue';
import { useUserStore } from '@/stores/UserStore.js';
import { useAppStore } from '@/stores/AppStore.js';

useMeta({
    title: 'Settings',
    htmlAttrs: {
        lang: 'en',
        amp: true
    }
});

const userStore = useUserStore();
const appStore = useAppStore();

const isDarkMode = computed({
    get() {
        return appStore.darkMode;
    },
    set(value) {
        appStore.darkMode = value;
    }
});

const animationsEnabled = computed({
    get() {
        return appStore.animationsEnabled;
    },
    set(value) {
        appStore.animationsEnabled = value;
    }
});

const isAuthed = computed(() => {
    return userStore.user !== null;
});

function changeAvatar() {
    console.log('Change avatar');
}

function logout() {
    userStore.logout();
}
</script>

<style>
.settings {
    display: flex;
    flex-flow: column nowrap;
    justify-content: center;
    align-items: center;
}

.settings > * {
    margin-bottom: 4rem;
}

.settings > * > * {
    margin: 0 0 1.25rem 0;
}

.settings > div {
    display: flex;
    flex-flow: column nowrap;
    align-items: center;
    padding: 1.25rem;
    border: 0.1rem solid var(--accent-color);
    border-radius: 0.5rem;
    width: 100%;
}

.app-settings > div {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: 1rem;
    padding-top: 1rem;
    width: 95%;
}

.app-settings > div > h3 {
    max-width: 50%;
    text-align: left;
}

.app-settings > *:not(:first-child) + div:not(:first-child) {
    border-top: 0.1rem solid var(--accent-color);
}

.user-settings > form {
    display: flex;
    flex-flow: column nowrap;
    align-items: flex-start;
}

.user-settings > form > div {
    display: flex;
    flex-flow: column nowrap;
    margin-bottom: 1rem;
    border: 0.1rem solid var(--accent-color);
    border-radius: 0.25rem;
    padding: 1rem;
}

.user-settings > form > div > label {
    font-weight: bold;
    color: var(--text-primary-color);
}

.user-settings > form > div > input {
    padding: 0.5rem;
    margin: 0.5rem;
    border: 0.1rem solid var(--accent-color);
    border-radius: 0.25rem;
    background-color: var(--background-color-secondary);
    color: var(--text-primary-color);
}

.user-settings > form > button {
    align-self: center;
}
</style>