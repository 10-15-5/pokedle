<template>
    <button class="toggle-button group" @click="store.toggleShiny()">
        <span
            v-if="store.isShiny"
            class="pi pi-eye-slash group-hover:text-red-600 dark:group-hover:text-red-300"
        ></span>
        <span
            v-else
            class="pi pi-star group-hover:text-orange-400 dark:group-hover:text-orange-300"
        ></span>
    </button>
</template>

<script setup>
import { useStore } from '../../stores/store.js';
const store = useStore();
</script>
