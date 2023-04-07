<template>
    <div
        v-if="!isClicked && isShowNotification"
        class="absolute right-0 top-0 mt-[-33px] mr-[-13px] sm:mt-[-27px] sm:mr-[-8px]"
    >
        <SpriteAnimation :isShowNotification="isShowNotification" @click="disable" />
    </div>
</template>

<script setup>
import { ref } from 'vue';
import SpriteAnimation from './SpriteAnimation.vue';

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
