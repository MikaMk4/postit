<template>
    <div class="mini-post" @click="miniPostClicked">
        <div class="mini-post__title">
        <h2>{{ post.title }}</h2>
        </div>
        <div class="mini-post__content">
        <p>{{ post.content }}</p>
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
            this.$router.push({ name: 'post', params: { id: this.post.id } });
        }
    }
}
</script>

<style>
.mini-post {
    border: 1px solid var(--accent-color);
    border-radius: 5px;
    margin: 10px;
    padding: 10px;
    width: 200px;
}

.mini-post:hover {
    background-color: var(--accent-color);
    cursor: pointer;
}

.mini-post > * {
    margin: 0.5rem;
}
</style>