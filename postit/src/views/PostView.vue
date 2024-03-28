<template>
    <div class="post">
        <h1>{{ post.title }}</h1>
        <div like-buttons>
            <button @click="likePost">Like</button>
            <button @click="dislikePost">Dislike</button>
        </div>
        <p>{{ post.content }}</p>
    </div>
</template>

<script>
export default {
    name: 'PostView',
    data() {
        return {
            post: {
                id: this.$route.params.id,
                title: 'First Post',
                content: 'This is the first post on the site.'
            }
        }
    }
}
</script>