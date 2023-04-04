<template>
    <BaseDialog :isVisible="isVisible" :onClose="onClose">
        <h1 class="text-center font-pkmEmerald text-[32px] uppercase">{{ getLanguage().statisticsDialog.statistics }}</h1>
        <div class="flex flex-row justify-center items-end relative">
            <span class="w-fit text-center font-pkmEmerald text-[20px] relatives"
                >{{ getLanguage().statisticsDialog.comingSoon }}
            </span><SpriteAnimation />
        </div>
    </BaseDialog>
</template>

<script setup>
import BaseDialog from './BaseDialog.vue';
import SpriteAnimation from '../animations/SpriteAnimation.vue';
import { getLanguage } from '../../services/language';
const props = defineProps({
    isVisible: Boolean,
    onClose: Function,
});
</script>
