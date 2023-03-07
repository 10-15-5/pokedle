<template>
    <BaseDialog :isVisible="isVisible" :onClose="onClose">
        <h1 class="font-pkmEmerald text-[32px] uppercase">SETTINGS</h1>
        <div class="flex w-full items-center justify-center gap-4">
            <p class="font-pkmEmerald text-[24px] w-1/7">Theme:</p>
            <ThemeButton class="m-2" />
        </div>
    </BaseDialog>
</template>

<script setup>
import ThemeButton from '../ThemeButton.vue';
import BaseDialog from './BaseDialog.vue';

const props = defineProps({
    isVisible: Boolean,
    onClose: Function,
});
</script>
