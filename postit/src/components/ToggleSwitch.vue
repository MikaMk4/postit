<template>
    <label class="switch">
        <input type="checkbox" v-model="isActive" :class="{ animated: props.animationsEnabled }" />
        <span class="slider round" :class="{ animated: props.animationsEnabled }"></span>
    </label>
</template>

<script setup>
import { defineModel } from 'vue';

const props = defineProps({
    animationsEnabled: {
        type: Boolean,
        default: true
    }
})

const isActive = defineModel()
</script>

<style>
 /* The switch - the box around the slider */
 .switch {
  position: relative;
  display: inline-block;
  width: 3.75rem;
  height: 2.125rem;
}

/* Hide default HTML checkbox */
.switch input {
  opacity: 0;
  width: 0;
  height: 0;
}

/* The slider */
.slider {
  position: absolute;
  cursor: pointer;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background-color: #ccc;
}

.slider.animated {
  -webkit-transition: .2s;
  transition: .2s;
}

.slider:before {
  position: absolute;
  content: "";
  height: 1.625rem;
  width: 1.625rem;
  left: 0.25rem;
  bottom: 0.25rem;
  background-color: white;
}

.slider.animated:before {
  -webkit-transition: .2s;
  transition: .2s;
}

input:checked + .slider {
  background-color: #2196F3;
}

input:focus + .slider {
  box-shadow: 0 0 0.0625rem #2196F3;
}

input:checked + .slider:before {
  -webkit-transform: translateX(1.625rem);
  -ms-transform: translateX(1.625rem);
  transform: translateX(1.625rem);
}

/* Rounded sliders */
.slider.round {
  border-radius: 2.125rem;
}

.slider.round:before {
  border-radius: 50%;
}
</style>