<template>
    <button
        class="toggle-button group"
        @click="store.toggleShiny()"
    >
        <span v-if="store.isShiny" class="pi pi-eye-slash
        group-hover:!text-red-600 group-hover:dark:!text-red-300 "></span>
        <span v-else class="pi pi-palette group-hover:!text-light-green group-hover:dark:!text-dark-green"></span>
    </button>
</template>

<script setup>
import { useStore } from '../../stores/store.js';
const store = useStore();
</script>
