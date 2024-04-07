<template>
  <div class="home">
    <h1>Home</h1>
    <p>Welcome to PostIt!

Discover a community-driven platform where diverse voices converge, ideas flourish, and discussions thrive. Whether you're seeking insightful conversations, sharing your passions, or exploring new interests, PostIt is your digital hub for engaging content and meaningful connections.

Join our vibrant community and:

    Explore a variety of topics ranging from technology and science to arts and entertainment.
    Participate in discussions, ask questions, and share your perspectives with fellow enthusiasts.
    Stay updated with trending posts, latest news, and compelling stories curated by our passionate users.
    Customize your experience by subscribing to your favorite communities and tailoring your feed to your interests.
    Engage with like-minded individuals, forge new friendships, and broaden your horizons in a welcoming and inclusive environment.

At PostIt, we believe in the power of community, the value of diverse perspectives, and the joy of meaningful interactions. Come be a part of something special.

Start exploring now!</p>
  </div>
</template>

<script setup>
import { useMeta } from 'vue-meta'

useMeta({
  title: 'Home',
  htmlAttrs: {
    lang: 'en',
    amp: true
  }
})
</script>

<style>
.home {
  display: flex;
  flex-flow: column nowrap;
  justify-content: center;
}

.home * {
  margin: 0.5rem;
}

.home p {
  text-align: left;
  font-size: 2rem;
}
</style>