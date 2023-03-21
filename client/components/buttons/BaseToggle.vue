<template>
    <div class="card relative flex !h-10 !w-20 flex-row dark:!bg-dark-border" @click="toggle()">
        <div
            class="flex h-full w-full cursor-pointer items-center justify-center rounded-l-sm"
            :class="isEnabled ? '' : ' bg-red-500 dark:bg-red-400'"
        >
            <slot v-if="isEnabled" name="leftIcon" class="absolute"></slot>
        </div>
        <div
            class="flex h-full w-full cursor-pointer items-center justify-center rounded-r-sm"
            :class="!isEnabled ? '' : 'bg-light-emerald dark:bg-green-400'"
        >
            <slot v-if="!isEnabled" name="rightIcon" class="absolute"></slot>
        </div>
    </div>
</template>

<script setup>
const props = defineProps({
    isEnabled: Boolean,
    toggle: Function,
});
</script>
