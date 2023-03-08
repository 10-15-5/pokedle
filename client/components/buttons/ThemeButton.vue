<template>
    <button
        class="toggle-button hover:!text-cyan-400 dark:hover:!text-amber-400"
        @click="toggleDark()"
    >
        <span v-if="isDark" class="pi pi-sun"></span>
        <span v-else class="pi pi-moon"></span>
    </button>
</template>

<script setup>
import { useDark, useToggle } from '@vueuse/core';

const isDark = useDark();
const toggleDark = useToggle(isDark);
</script>
