<template>
    <button
        class="absolute right-0 z-10 m-4 flex h-11 w-11 items-center justify-center rounded-full bg-black/30 text-[24px] text-white hover:bg-black/10 hover:!text-cyan-400 dark:bg-white/40 dark:text-[28px] dark:text-black hover:dark:bg-white/20 dark:hover:!text-amber-400"
        @click="toggleActive()"
    >
        <span v-if="isActive" class="pi pi-sun"></span>
        <span v-else class="pi pi-moon"></span>
    </button>
</template>

<script setup>
const props = defineProps({
    toggleActive: Function,
    isActive: Boolean,
});
</script>
