<template>
    <Dialog
        class="w-[800px]"
        :visible="isVisible"
        :dismissable-mask="true"
        :closeOnEscape="true"
        :modal="true"
        :position="'center'"
        @update:visible="onClose"
    >
        <div class="dialog-card">
            <slot></slot>
            <button class="dialog-close" block @click="onClose">Close</button>
        </div>
    </Dialog>
</template>

<script setup>
const props = defineProps({
    isVisible: Boolean,
    onClose: Function,
});
</script>
