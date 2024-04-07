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
            <AvatarPreview :avatar="avatar" :size="15"/>
            <form @submit.prevent="newUserInfoSet">
                <div>
                    <label for="avatar">Avatar URL</label>
                    <input type="text" id="avatar" v-model="avatar" />
                </div>
                <div>
                    <label for="username">Username</label>
                    <input type="text" id="username" v-model="username" />
                </div>
                <div>
                    <label for="bio">About me</label>
                    <textarea id="bio" v-model="bio"></textarea>
                </div>
                <button type="submit">Save Changes</button>
            </form>
            <form @submit.prevent="editPassword">
                <div>
                    <label for="password">New Password</label>
                    <input type="password" id="password" v-model="password" />
                    <label for="confirmPassword">Confirm Password</label>
                    <input type="password" id="confirmPassword" v-model="confirmPassword" />
                </div>
                <button type="submit">Change Password</button>
            </form>
            <button @click="logout">Logout</button>
        </div>
    </div>
</template>

<script setup>
import { computed, ref } from 'vue';
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

const userStore = useUserStore()
const appStore = useAppStore()

const avatar = ref(userStore.user?.avatar)
const username = ref(userStore.user?.username)
const bio = ref(userStore.user?.bio)
const password = ref('')
const confirmPassword = ref('')

async function newUserInfoSet() {
    userStore.startEdit()

    userStore.user.avatar = avatar.value
    userStore.user.username = username.value
    userStore.user.bio = bio.value

    try {
        await userStore.updateUser()
    } catch (error) {
        alert(error)
        userStore.cancelEdit()
        avatar.value = userStore.user.avatar
        username.value = userStore.user.username
        bio.value = userStore.user.bio
    }
}

async function editPassword() {
    if (password.value !== '' && password.value === confirmPassword.value) {
        userStore.startEdit()
        userStore.user.password = password.value

        try {
            await userStore.updateUser()
            alert('Password changed successfully.')
        } catch (error) {
            alert(error)
            userStore.cancelEdit()
        } finally {
            password.value = ''
            confirmPassword.value = ''
        }
    } else {
        alert('Passwords do not match.')
    }
}

const isDarkMode = computed({
    get() {
        return appStore.darkMode
    },
    set(value) {
        appStore.darkMode = value
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
    width: 75%;
}

.user-settings > form > div {
    display: flex;
    flex-flow: column nowrap;
    margin-bottom: 1rem;
    border: 0.1rem solid var(--accent-color);
    border-radius: 0.25rem;
    padding: 1rem;
    width: 100%;
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

.user-settings > form > div > textarea {
    padding: 0.5rem;
    margin: 0.5rem;
    border: 0.1rem solid var(--accent-color);
    border-radius: 0.25rem;
    background-color: var(--background-color-secondary);
    color: var(--text-primary-color);
    resize: vertical;
    min-height: 8rem;
}

.user-settings > form > button {
    align-self: center;
}
</style>