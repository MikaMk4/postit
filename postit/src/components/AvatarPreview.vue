<template>
    <div class="avatar" :class="{ dim: isEditable, expandable: clickable, animated: appStore.animationsEnabled }">
        <router-link v-if="clickable" :to="{ name: 'user', params: { id: pId } }">
            <img :src="avatar" alt="Avatar" class="avatar-img">
        </router-link>
        <div v-else @click="editClicked">
            <img :src="avatar" alt="Avatar" class="avatar-img">
            <img src="../assets/edit.png" alt="Edit" class="edit-icon" v-if="isEditable">
        </div>
    </div>
</template>

<script setup>
import { defineProps, computed } from 'vue';
import { useAppStore } from '@/stores/AppStore';

const appStore = useAppStore();

const props = defineProps(
    {
        avatar: {
            type: String,
            required: true
        },
        size: {
            type: Number,
            default: 3.125
        },
        isEditable: {
            type: Boolean,
            default: false
        },
        pId: {
            type: String,
            default: ''
        }
    }
)

const emit = defineEmits(['changeAvatar']);

function editClicked() {
    if (props.isEditable) {
        emit('changeAvatar');
    }
}

const clickable = computed(
    () => props.pId !== ''
);

const aSize = computed(
    () => props.size + 'rem'
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
    height: auto;
    overflow: hidden;
    border-radius: 50%;
}

.avatar.expandable:hover {
    cursor: pointer;
}

.avatar.expandable.animated:hover {
    animation: expand-animation 0.25s forwards;
}

@keyframes expand-animation {
    0% {
        border-radius: 50%;
    }
    100% {
        border-radius: 25%;
    }
}

.dim:hover .avatar-img {
    filter: brightness(0.5);
    cursor: pointer;
}

.avatar-img {
    display: block;
    width: v-bind(aSize);
    height: auto;
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