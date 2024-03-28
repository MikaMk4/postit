<template>
    <div class="avatar">
        <img :src="avatar" alt="Avatar" class="avatar-img">
        <img src="../assets/edit.png" alt="Edit" class="edit-icon">
    </div>
</template>

<script setup>
import { defineProps, computed } from 'vue';

const props = defineProps(
    {
        avatar: {
            type: String,
            required: true
        },
        height: {
            type: Number,
            required: false,
            default: 100
        },
        width: {
            type: Number,
            required: false,
            default: 100
        }
    }
)

const aHeight = computed(() => props.height + 'px');
const aWidth = computed(() => props.width + 'px');

</script>

<style>
.avatar {
    position: relative;
    display: flex;
    flex-flow: column nowrap;
    justify-content: center;
    align-items: center;
    width: v-bind(aWidth);
    height: v-bind(aHeight);
    overflow: hidden;
    border-radius: 50%;
}

.avatar:hover .avatar-img {
    filter: brightness(0.5);
    cursor: pointer;
}

.avatar-img {
    display: block;
    width: 100%;
    height: 100%;
    object-fit: cover;
}

.avatar:hover .edit-icon {
    display: block;
    color: rgba(0, 0, 0, 0.5);
}

.edit-icon {
    display: none;
    position: absolute;
    top: 50%;
    right: 50%;
    width: 35%;
    height: 35%;
    transform: translate(50%, -50%);
    cursor: pointer;
}

</style>