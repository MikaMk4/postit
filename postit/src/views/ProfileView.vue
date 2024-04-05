<template>
    <div class="profile">
        <h1>Profile</h1>
        <div class="profile-info">
            <AvatarPreview :avatar="userStore.user.avatar" :size="100" :expandable="true"/>
            <div class="profile-info-text">
                <p>{{ userStore.user.username }}</p>
                <p>{{ userStore.user.email }}</p>
            </div>
        </div>
        <div class="profile-posts">
            <h2>Posts</h2>
            <div class="profile-posts-list">
                <MiniPost v-for="post in userStore.user.posts" :post="post" :key="post.id"/>
            </div>
        </div>
    </div>
</template>

<script setup>
import { useUserStore } from '@/stores/UserStore.js'
import MiniPost from '@/components/MiniPost.vue'
import AvatarPreview from '@/components/AvatarPreview.vue'

const userStore = useUserStore()
</script>

<style>
.profile {
    display: flex;
    flex-flow: column nowrap;
    align-items: center;
}
</style>