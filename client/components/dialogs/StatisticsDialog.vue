<template>
    <BaseDialog :isVisible="isVisible" :onClose="onClose">
        <div class="flex flex-col justify-center items-center relative">
            <h1 class="text-center font-pkmEmerald text-[32px] uppercase">STATISTICS</h1>
            <span class="w-fit text-center font-pkmEmerald text-[20px]"
                >Coming Soon.... <NotificationAnimation class="!mt-[36px] !mr-[165px]" />
            </span>
        </div>
    </BaseDialog>
</template>

<script setup>
import BaseDialog from './BaseDialog.vue';
import NotificationAnimation from '../animations/NotificationAnimation.vue';

const props = defineProps({
    isVisible: Boolean,
    onClose: Function,
});
</script>
