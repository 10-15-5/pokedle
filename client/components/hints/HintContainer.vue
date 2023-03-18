<template>
    <div
        class="card min-w-[80px] max-w-[400px] flex-col font-pkmEmerald sm:max-w-[330px] sm:text-sm"
        :class="isShowHints ? 'w-[400px]' : ''"
    >
        <p v-if="!isHintFlavortextUnlocked" class="gap-2 p-2 px-3 sm:p-1 sm:px-2 sm:pt-2">
            Guess
            <b class="text-light-orange dark:text-dark-orange"> {{ guessesRemainingForHintFlavorText }} </b>
            more times for a HINT
        </p>
        <div
            v-else-if="!isShowHints"
            class="flex w-full flex-row items-center justify-between gap-2 p-2 px-3 pb-1 sm:pb-1 sm:pt-2"
        >
            <div class="flex flex-row text-base">
                <p>HINT{{ isHintCardOneUnlocked ? 'S' : '' }}</p>
            </div>
            <div>
                <button class="pi pi-plus" @click="isShowHints = true"></button>
            </div>
        </div>
        <div v-if="isShowHints" class="mr-3 flex flex-row justify-between">
            <div class="flex flex-row gap-2">
                <div
                    @click="selectHintFlavortext"
                    class="hint-tab-header"
                    :class="isShowHintFlavortext ? 'hint-tab-header-active' : ''"
                >
                    Flavortext
                </div>
                <div
                    @click="selectHintCardOne"
                    class="hint-tab-header flex gap-1"
                    :class="isShowHintCardOne ? 'hint-tab-header-active' : ''"
                >
                    Card Reveal 1
                    <div>
                        <span :class="isHintCardOneUnlocked ? 'pi pi-lock-open' : 'pi pi-lock'"></span>
                    </div>
                </div>
                <div
                    @click="selectHintCardTwo"
                    class="hint-tab-header flex gap-1"
                    :class="isShowHintCardTwo ? 'hint-tab-header-active' : ''"
                >
                    Card Reveal 2
                    <div>
                        <span :class="isHintCardTwoUnlocked ? 'pi pi-lock-open' : 'pi pi-lock'"></span>
                    </div>
                </div>
            </div>
            <div class="flex items-center justify-center">
                <button v-if="isShowHints" class="pi pi-minus" @click="isShowHints = false"></button>
            </div>
        </div>
        <div v-if="isShowHints" class="w-full bg-gray-100 p-4 pt-3 text-justify dark:bg-zinc-700">
            <div v-if="isShowHintFlavortext">
                {{ text }}
            </div>
            <div v-else-if="isShowHintCardOne">CARD REVEAL 1</div>
            <div v-else-if="isShowHintCardTwo">CARD REVEAL 2</div>
        </div>
    </div>
</template>

<script setup>
import { computed, ref } from 'vue';
const props = defineProps({
    text: String,
    numberOfGuesses: {
        type: Number,
        default: 0,
    },
    forceEnableHint: Boolean,
});

const hints = ['flavortext', 'card_reveal_1', 'card_reveal_2'];
const selectedHint = ref('flavortext');
const isShowHints = ref(false);

const guessesRemainingForHintFlavorText = computed(() => 4- props.numberOfGuesses);
const isHintFlavortextUnlocked = computed(() => (4 - props.numberOfGuesses) <= 0);
const isHintCardOneUnlocked = computed(() => (6 - props.numberOfGuesses) <= 0);
const isHintCardTwoUnlocked = computed(() => (8 - props.numberOfGuesses) <= 0);

const isShowHintFlavortext = computed(() => selectedHint.value === hints[0]);
const isShowHintCardOne = computed(() => selectedHint.value === hints[1]);
const isShowHintCardTwo = computed(() => selectedHint.value === hints[2]);

const selectHintFlavortext = () => {
    if(isHintFlavortextUnlocked.value) {
        selectedHint.value = hints[0]
    }
}
const selectHintCardOne = () => {
    if(isHintCardOneUnlocked.value) {
        selectedHint.value = hints[1]
    }
}
const selectHintCardTwo = () => {
    if(isHintCardTwoUnlocked.value) {
        selectedHint.value = hints[2]
    }
}
</script>
