<template>
    <div class="mini-post" @click="miniPostClicked">
        <div class="mini-post__title">
            <h2>{{ post.title }}</h2>
        </div>
            <div class="mini-post__content">
            <p>{{ post.content }}</p>
        </div>
        <div class="mini-post-tn">
            <img :src="post.thumbnail" alt="thumbnail" />
        </div>
    </div>
</template>

<script>
export default {
    name: 'MiniPost',
    props: {
        post: {
            type: Object,
            required: true
        }
    },
    methods: {
        miniPostClicked() {
            this.$router.push({ name: 'post', params: { pid: this.post.id, bid: this.post.boardId } });
        }
    }
}
</script>

<style>
.mini-post {
    border: 0.1rem solid var(--accent-color);
    border-radius: 0.4rem;
    margin: 1rem;
    padding: 1rem;
    width: 20rem;
}

.mini-post:hover {
    background-color: var(--accent-color);
    cursor: pointer;
}

.mini-post > * {
    margin: 0.5rem;
}

.mini-post-tn > img {
    width: 100%;
    height: auto;
}
</style>