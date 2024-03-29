<template>
    <div class="avatar" :class="{ dim: isEditable }">
        <img :src="avatar" alt="Avatar" class="avatar-img">
        <img src="../assets/edit.png" alt="Edit" class="edit-icon" v-if="isEditable">
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
        size: {
            type: Number,
            default: 50
        },
        isEditable: {
            type: Boolean,
            default: false
        }
    }
)

const aSize = computed(
    () => props.size + 'px'
);
</script>

<style>
.avatar {
    position: relative;
    display: flex;
    flex-flow: column nowrap;
    justify-content: center;
    align-items: center;
    width: v-bind(aSize);
    height: v-bind(aSize);
    overflow: hidden;
    border-radius: 50%;
}

.avatar:hover .avatar-img {
    cursor: pointer;
}

.dim:hover .avatar-img {
    filter: brightness(0.5);
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