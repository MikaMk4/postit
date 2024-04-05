<template>
    <div class="settings">
        <h1>Settings</h1>
        <div class="app-settings">
            <h2>App Settings</h2>
            <div>
                <h3>Dark Mode</h3>
                <ToggleSwitch v-model="isDarkMode" />
            </div>
            <div>
                <h3>Setting 2</h3>
                <ToggleSwitch />
            </div>
            <div>
                <h3>This is a super long text for setting 3</h3>
                <ToggleSwitch />
            </div>
        </div>
        <div class="user-settings" v-if="isAuthed">
            <h2>User Settings</h2>
            <AvatarPreview :avatar="userStore.user.avatar" :size="200" :isEditable="true"/>
            <p>{{ userStore.user.name }}</p>
            <h3>Change Username</h3>
            <h3>Change Password</h3>
        </div>
    </div>
</template>

<script setup>
import { computed } from 'vue';
import AvatarPreview from '@/components/AvatarPreview.vue';
import ToggleSwitch from '@/components/ToggleSwitch.vue';
import { useUserStore } from '@/stores/UserStore.js';
import { useAppStore } from '@/stores/AppStore.js';

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
</style>