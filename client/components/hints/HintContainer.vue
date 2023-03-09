<template>
    <div class="card text-justify sm:text-sm font-pkmEmerald max-w-[400px] sm:max-w-[330px] flex-col p-4 sm:p-2">
        {{ text }}
    </div>
</template>

<script setup>
const props = defineProps({
    text: String
})

</script>
