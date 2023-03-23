<template>
    <div class="single-result-container mb-1 items-center justify-center">
        <span
            class="border-b-2 border-light-border text-center font-pkmEmerald text-lg font-bold dark:border-dark-border"
            >{{ headerText }}
        </span>
    </div>
</template>

<script setup>
const props = defineProps({
    headerText: String,
});
</script>
