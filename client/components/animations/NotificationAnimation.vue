<template>
    <div
        v-if="!isClicked && isShowNotification"
        class="absolute right-0 top-0 mt-[-7px] mr-[-8px] cursor-pointer sm:mt-[-6px] sm:mr-[-6px]"
        @click="disable"
    >
        <span class="relative flex h-4 w-4 sm:h-[14px] sm:w-[14px]">
            <span class="absolute inline-flex h-full w-full animate-ping rounded-full bg-sky-400 opacity-75"></span>
            <span class="relative inline-flex h-4 w-4 rounded-full bg-teal-300 sm:h-[14px] sm:w-[14px]"></span>
        </span>
    </div>
</template>

<script setup>
import { ref } from 'vue';

const props = defineProps({
    isShowNotification: {
        type: Boolean,
        default: true,
    },
});

const isClicked = ref(false);

const disable = () => {
    isClicked.value = true;
};
</script>
