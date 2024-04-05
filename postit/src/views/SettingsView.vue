<template>
    <div class="settings">
        <h1>Settings</h1>
        <div class="user-settings" v-if="isAuthed">
            <h2>User Settings</h2>
            <AvatarPreview :avatar="userStore.user.avatar" :size="200" :isEditable="true"/>
            <p>{{ userStore.user.name }}</p>
            <h3>Change Username</h3>
            <h3>Change Password</h3>
        </div>
        <div class="preferences">
            <div class="preferences-item">
                <h3>Dark Mode</h3>
                <ToggleSwitch v-model="isDarkMode" />
            </div>
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
        console.log(appStore.darkMode);
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
    width: 80%;
    margin: 0 auto;
}

.settings > div {
    margin: 0 0 500px 0;
}

.settings > h1 {
    margin: 0 0 50px 0;
}

.settings > * > * {
    margin: 0 0 20px 0;
}

.preferences {
    display: flex;
    flex-flow: column;
    width: 100%;
}

.preferences-item {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: 10px;
    width: 100%;
}
</style>