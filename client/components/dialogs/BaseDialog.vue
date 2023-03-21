<template>
    <Dialog
        class="!transition-none"
        :visible="isVisible"
        :closeOnEscape="true"
        :position="'center'"
        :showHeader="false"
        @update:visible="onClose"
    >
        <div
            class="flex !h-screen !w-screen items-center justify-center bg-black/40"
            @click="onClose"
        >
            <div class="dialog-card" v-on:click.stop>
                <slot></slot>
                <button class="dialog-close" block @click="onClose">Close</button>
            </div>
        </div>
    </Dialog>
</template>

<script setup>
const props = defineProps({
    isVisible: Boolean,
    onClose: Function,
});
</script>

<style></style>
