<template>
    <BaseDialog :isVisible="isVisible" :onClose="onClose">
        <h1 class="font-pkmEmerald text-[32px] uppercase">SETTINGS</h1>
        <div class=" grid grid-cols-3 items-center justify-center gap-4">
            <p class="font-pkmEmerald text-[24px] col-span-2">Theme:</p>
            <ThemeButton class="m-2" />
            <p class="font-pkmEmerald text-[24px] col-span-2">Shiny Only:</p>
            <ShinyButton class="m-2" />
        </div>
    </BaseDialog>
</template>

<script setup>
import ThemeButton from '../buttons/ThemeButton.vue';
import ShinyButton from '../buttons/ShinyButton.vue';

import BaseDialog from './BaseDialog.vue';

const props = defineProps({
    isVisible: Boolean,
    onClose: Function,
});
</script>
