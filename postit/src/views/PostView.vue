<template>
    <div class="post-view">
        <div class="post">
            <div class="post-header">
                <AvatarPreview :avatar="userStore.user.avatar" :size="2" :animationsEnabled="appStore.animationsEnabled" :pId="userStore.user.id"/>
                <h3>{{ userStore.user.name }}</h3>
                <h1>{{ post.title }}</h1>
            </div>
            <div class="options-container">
                <div class="like-container">
                    <i class="fa fa-heart" :class="{ liked: isLiked, animated: appStore.animationsEnabled }" @click="toggleLike"></i>
                    <p>{{ post.likeCount }}</p>
                </div>
                <div class="delete-post-container" v-if="canDelete">
                    <i class="fa fa-trash" :class="{ animated: appStore.animationsEnabled }" @click="deletePost"></i>
                </div>
            </div>
            <div class="post-content-container">
                <p>{{ post.content }}</p>
            </div>
        </div>
        <div class="comment-container">
            <h2>Comments</h2>
            <div class="comments">
                <div v-for="comment in comments" :key="comment.id">
                    <span>
                        <AvatarPreview :avatar="comment.author.avatar" :size="2" :animationsEnabled="appStore.animationsEnabled" :pId="comment.authorId"/>
                        <h3>{{ comment.author.name }}</h3>
                    </span>
                    <p>{{ comment.content }}</p>
                </div>
            </div>
        </div>
    </div>
</template>

<script setup>
import { computed, onMounted, ref } from 'vue'
import router from '@/router';
import { useMeta } from 'vue-meta'
import AvatarPreview from '@/components/AvatarPreview.vue';
import { useUserStore } from '@/stores/UserStore.js';
import { useAppStore } from '@/stores/AppStore';
import { usePostStore } from '@/stores/PostStore';

const userStore = useUserStore()
const appStore = useAppStore()

const post = ref({
    authorId: '',
    title: '',
    content: '',
    likeCount: 0
})

const boardId = router.currentRoute.value.params.bid
const postId = router.currentRoute.value.params.pid
const postStore = usePostStore()
onMounted(() => {
    postStore.fetchPost(boardId, postId).then((fetchedPost) => {
        post.value = fetchedPost
        console.log('Post fetched:', fetchedPost)
    }).catch((err) => {
        if (err === 'Post not found') {
            router.push({ name: 'not-found' })
        }
    })
})

const comments = ref([])

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
    post.value.likeCount += isLiked.value ? 1 : -1
    post.value.id = postId
    postStore.updatePost(post.value)
}

const canDelete = computed(() => {
    if (userStore.user === null) return false

    return userStore.user.id === post.value.authorId
})

function deletePost() {
    alert('Are you sure you want to delete this post?')
    postStore.deletePost(boardId, postId)
    router.go(-1)
}
</script>

<style>
.post-view {
    display: flex;
    flex-flow: column nowrap;
    align-items: center;
    justify-content: center;
    width: 50rem;
    max-width: 75%;
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

.post-header {
    display: flex;
    flex-flow: row nowrap;
    align-items: center;
    justify-content: flex-start;
    margin-bottom: .25rem;
}

.post-header > h3 {
    margin-left: 0.5rem;
}

.post-header > h1 {
    margin-left: 1rem;
}

.options-container {
    display: flex;
    flex-flow: row nowrap;
    justify-content: flex-start;
    align-items: center;
    margin-bottom: 2rem;
}

.options-container > * {
    margin-right: 1rem;
}

.like-container {
    display: flex;
    flex-flow: row nowrap;
    align-items: center;
    justify-content: flex-start;
    font-size: 1.5rem;
}

.like-container > * {
    margin-right: 0.5rem;
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
    cursor: pointer;
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

.post-content-container {
    text-align: left;
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