<template>
    <div class="profile">
        <h1>Profile</h1>
        <div class="profile-info">
            <AvatarPreview :avatar="userStore.user.avatar" :size="10" :expandable="true"/>
            <div class="profile-info-text">
                <div>
                    <h2>{{ userStore.user.username }}</h2>
                </div>
                <div class="profile-info-bio">
                    <h3>Bio:</h3>
                    <p>{{ userStore.user.bio }}</p>
                </div>
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
    width: 80%;
}

.profile-info {
    display: flex;
    flex-flow: row nowrap;
    align-items: top;
    border: 0.1rem solid var(--accent-color);
    border-radius: 0.4rem;
    width: 80%;
    margin: 1.5rem 0;
    padding: 1.5rem;
}

.profile-info-text {
    margin: 1.5rem 1.5rem;
    display: flex;
    flex-flow: column wrap;
    align-items: flex-start;
}
</style>