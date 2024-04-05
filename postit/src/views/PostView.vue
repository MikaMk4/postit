<template>
    <div class="post-view">
        <div class="post">
            <h1>{{ post.title }}</h1>
            <div class="options-container">
                <div class="like-container">
                    <i class="fa fa-heart" :class="{ liked: isLiked, animated: appStore.animationsEnabled }" @click="toggleLike"></i>
                    <p>{{ post.likes }}</p>
                </div>
                <div class="delete-post-container" v-if="canDelete">
                    <i class="fa fa-trash" :class="{ animated: appStore.animationsEnabled }" @click="deletePost"></i>
                </div>
            </div>
            <p>{{ post.content }}</p>
        </div>
        <div class="comment-container">
            <h2>Comments</h2>
            <div class="comments">
                <div v-for="comment in comments" :key="comment.id">
                    <span>
                        <AvatarPreview :avatar="comment.author.avatar" :size="40"/>
                        <h3>{{ comment.author.name }}</h3>
                    </span>
                    <p>{{ comment.content }}</p>
                </div>
            </div>
        </div>
    </div>
</template>

<script setup>
import { computed, ref } from 'vue'
import { useMeta } from 'vue-meta'
import AvatarPreview from '@/components/AvatarPreview.vue';
import { useUserStore } from '@/stores/UserStore.js';
import { useAppStore } from '@/stores/AppStore';

const userStore = useUserStore()
const appStore = useAppStore()

const post = ref({
    authorId: '1',
    title: 'Post Title',
    content: 'Post Content',
    likes: 12
})

const comments = ref([
    { id: 1, content: 'Comment 1', author: { name: 'Author 1', avatar: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTYloKopOZ_oudmWTNK-xVmdVPxdKsgKniHbr8Vr0hk1g&s' } },
    { id: 2, content: 'Comment 2', author: { name: 'Author 2', avatar: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTYloKopOZ_oudmWTNK-xVmdVPxdKsgKniHbr8Vr0hk1g&s' } },
    { id: 3, content: 'Comment 3', author: { name: 'Author 3', avatar: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTYloKopOZ_oudmWTNK-xVmdVPxdKsgKniHbr8Vr0hk1g&s' } }
])

useMeta({
    title: 'Post: ' + post.value.title + ' - PostIt',
    htmlAttrs: {
        lang: 'en',
        amp: true
    }
})

const isLiked = ref(false)

function toggleLike() {
    isLiked.value = !isLiked.value
    post.value.likes += isLiked.value ? 1 : -1
}

const canDelete = computed(() => {
    if (userStore.user === null) return false

    return userStore.user.id === post.value.authorId
})

function deletePost() {
    console.log('Post deleted')
}
</script>

<style>
.post-view {
    display: flex;
    flex-flow: column nowrap;
    align-items: center;
    width: 50rem;
}

.post-view > * {
    border: 0.1rem solid var(--accent-color);
    border-radius: 0.4rem;
    padding: 1rem;
    width: 100%;
}

.post > p {
    text-align: left;
}

.post {
    min-height: 25rem;
}

.options-container {
    display: flex;
    flex-flow: row nowrap;
    justify-content: center;
    align-items: center;
    margin-bottom: 2rem;
}

.options-container > * {
    margin: 0 1rem;
}

.like-container {
    display: flex;
    flex-flow: row nowrap;
    align-items: center;
    justify-content: center;
    font-size: 1.5rem;
}

.like-container > * {
    margin: 0 0.5rem;
}

.like-container > i {
    color: grey;
}

.like-container > i.liked {
    color: red;
}

.like-container > i:hover:not(.liked) {
    color: red;
    cursor: pointer;
}

.like-container > i:hover:not(.liked).animated {
    animation: like-animation 0.5s forwards;
}

.like-container > i.liked:hover {
    color: grey;
    cursor: pointer;
}

.like-container > i.liked:hover.animated {
    animation: unlike-animation 0.5s forwards;
}

@keyframes like-animation {
    0% {
        transform: scale(1);
    }
    50% {
        transform: scale(1.2);
    }
    100% {
        transform: scale(1);
    }
}

@keyframes unlike-animation {
    0% {
        transform: translate(0, 0) rotate(0);
    }
    100% {
        transform: translate(5px, 5px) rotate(5deg);
    }
}

.like-container > p {
    font-weight: bold;
}

.delete-post-container {
    display: flex;
    flex-flow: row nowrap;
    align-items: center;
    justify-content: center;
    font-size: 1.5rem;
}

.delete-post-container > i {
    color: grey;
    cursor: pointer;
}

.delete-post-container > i:hover {
    color: red;
}

.delete-post-container > i:hover.animated {
    animation: delete-animation 0.1s forwards linear;
}

@keyframes delete-animation {
    0% {
        transform: translate(0, 0);
    }
    25% {
        transform: translate(0.1rem, 0);
    }
    50% {
        transform: translate(0, 0);
    }
    75% {
        transform: translate(-0.1rem, 0);
    }
    100% {
        transform: translate(0, 0);
    }
}

.comment-container {
    margin-top: 1rem;
}

.comments {
    display: flex;
    flex-flow: column nowrap;
    align-items: stretch;
}

.comments > * {
    margin: 0.5rem;
    text-align: left;
    background-color: var(--background-color-secondary);
    padding: 0.5rem;
    border-radius: 0.4rem;
}

.comments > * > span {
    display: flex;
    flex-flow: row nowrap;
    align-items: center;
}

.comments > * > span > h3 {
    margin-left: 0.5rem;
}

.comments > * > p {
    margin-top: 0.5rem;
}
</style>