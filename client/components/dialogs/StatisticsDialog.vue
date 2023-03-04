<template>
    <BaseDialog :isVisible="isVisible" :onClose="onClose">
        <h1 class="font-pkmEmerald text-[32px] uppercase">STATISTICS</h1>
                <p class="font-pkmEmerald text-[20px]">Coming Soon....</p>
    </BaseDialog>
</template>

<script setup>
import BaseDialog from './BaseDialog.vue';

const props = defineProps({
    isVisible: Boolean,
    onClose: Function,
});
</script>
