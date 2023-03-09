<template>
    <div class="card max-w-[400px] flex-col p-4 gap-3">
        {{ text }}
    </div>
</template>

<script setup>
const props = defineProps({
    text: String
})

</script>
