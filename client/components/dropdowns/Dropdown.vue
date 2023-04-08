<template>
    <div class="relative font-pkmEmerald">
        <button
            class="card flex items-center justify-center !h-10 !w-20 capitalize"
            type="button"
            @click="show = !show"
        >
            {{ selectedItem }}
            <svg
                class="ml-2 h-4 w-4"
                aria-hidden="true"
                fill="none"
                stroke="currentColor"
                viewBox="0 0 24 24"
                xmlns="http://www.w3.org/2000/svg"
            >
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7"></path>
            </svg>
        </button>
        <!-- Dropdown menu -->
        <div v-show="show" class="card absolute z-10 w-full flex-col cente">
            <span
                v-for="(item, k) in items"
                :key="k"
                :value="item"
                class="hover:bg-gray-100 dark:hover:bg-gray-600 dark:hover:text-white capitalize pl-1 py-1 cursor-pointer"
                @click="selectItem(item)"
                >{{ item }}</span
            >
        </div>
    </div>
</template>

<script setup>
import { onMounted, ref } from 'vue';
const emit = defineEmits(['onSelectItem']);

const show = ref(false);

const selectItem = (item) => {
    show.value = false;
    emit("onSelectItem", item)
}

const props = defineProps({
    selectedItem: String,
    items: [],
});
</script>
