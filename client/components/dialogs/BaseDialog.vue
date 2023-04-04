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
                <button class="dialog-close" block @click="onClose">{{ getLanguage().close }}</button>
            </div>
        </div>
    </Dialog>
</template>

<script setup>
import { getLanguage } from '../../services/language';
const props = defineProps({
    isVisible: Boolean,
    onClose: Function,
});
</script>

<style></style>
