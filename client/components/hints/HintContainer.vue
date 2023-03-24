<template>
    <div
        class="card min-w-[80px] max-w-[400px] flex-col font-pkmEmerald sm:max-w-[330px] sm:text-sm"
        :class="isShowHints ? 'w-[400px]' : ''"
    >
        <p v-if="!isHintOneUnlocked" class="gap-2 p-2 px-3 sm:p-1 sm:px-2 sm:pt-2">
            Guess
            <b class="text-light-orange dark:text-dark-orange">
                {{ guessesRemainingForHintOne }}
            </b>
            more times for a HINT
        </p>
        <div
            v-else-if="!isShowHints"
            class="flex w-full flex-row items-center justify-between gap-2 p-2 px-3 pb-1 sm:pb-1 sm:pt-2"
        >
            <div class="flex flex-row text-base">
                <p>HINT{{ isHintTwoUnlocked ? 'S' : '' }}</p>
            </div>
            <div>
                <button class="pi pi-plus" @click="isShowHints = true"></button>
            </div>
        </div>
        <div v-if="isShowHints" class="mr-3 flex flex-row justify-between">
            <div class="flex flex-row gap-2">
                <div v-if="isHintOne"
                    @click="selectHintCardOne"
                    class="hint-tab-header ml-1 flex gap-1"
                    :class="isShowHintOne ? 'hint-tab-header-active' : ''"
                >
                    Hint 1
                    <div>
                        <span :class="isHintOneUnlocked ? 'pi pi-lock-open' : 'pi pi-lock'"></span>
                    </div>
                </div>
                <div v-if="isHintTwo"
                    @click="selectHintCardTwo"
                    class="hint-tab-header flex gap-1"
                    :class="isShowHintTwo ? 'hint-tab-header-active' : ''"
                >
                    Hint 2
                    <div>
                        <span :class="isHintTwoUnlocked ? 'pi pi-lock-open' : 'pi pi-lock'"></span>
                    </div>
                </div>
                <div v-if="isHintThree"
                    @click="selectHintCardThree"
                    class="hint-tab-header flex gap-1"
                    :class="isShowHintThree ? 'hint-tab-header-active' : ''"
                >
                    Hint 3
                    <div>
                        <span
                            :class="isHintThreeUnlocked ? 'pi pi-lock-open' : 'pi pi-lock'"
                        ></span>
                    </div>
                </div>
            </div>
            <div class="flex items-center justify-center">
                <button
                    v-if="isShowHints"
                    class="pi pi-minus"
                    @click="isShowHints = false"
                ></button>
            </div>
        </div>
        <div v-if="isShowHints" class="bg-gray-100 px-4 py-2 pt-3 text-justify dark:bg-zinc-700">
            <div v-if="isShowHintOne" :class="stylingHintOne">
                <slot name="hint1"></slot>
            </div>
            <div v-else-if="isShowHintTwo" :class="stylingHintTwo">
                <slot name="hint2"></slot>
            </div>
            <div v-else-if="isShowHintThree" :class="stylingHintThree">
                <slot name="hint3"></slot>
            </div>
        </div>
    </div>
</template>

<script setup>
import { computed, ref } from 'vue';

const props = defineProps({
    numberOfGuesses: {
        type: Number,
        default: 0,
    },
    stylingHintOne: String,
    stylingHintTwo: String,
    stylingHintThree: String,
    guessesRequiredForHintOne: Number,
    guessesRequiredForHintTwo: Number,
    guessesRequiredForHintThree: Number,

});

const hints = ['hint1', 'hint2', 'hint3'];
const selectedHint = ref('hint1');
const isShowHints = ref(false);

const isHintOne = computed(() => props.guessesRequiredForHintOne)
const isHintTwo = computed(() => isHintOne.value && props.guessesRequiredForHintTwo)
const isHintThree = computed(() => isHintTwo.value && props.guessesRequiredForHintThree)


const guessesRemainingForHintOne = computed(
    () => props.guessesRequiredForHintOne - props.numberOfGuesses
);
const isHintOneUnlocked = computed(
    () => props.guessesRequiredForHintOne - props.numberOfGuesses <= 0
);
const isHintTwoUnlocked = computed(
    () => props.guessesRequiredForHintTwo - props.numberOfGuesses <= 0
);
const isHintThreeUnlocked = computed(
    () => props.guessesRequiredForHintThree - props.numberOfGuesses <= 0
);

const isShowHintOne = computed(() => selectedHint.value === hints[0]);
const isShowHintTwo = computed(() => selectedHint.value === hints[1]);
const isShowHintThree = computed(() => selectedHint.value === hints[2]);

const selectHintCardOne = () => {
    if (isHintOneUnlocked.value) {
        selectedHint.value = hints[0];
    }
};
const selectHintCardTwo = () => {
    if (isHintTwoUnlocked.value) {
        selectedHint.value = hints[1];
    }
};
const selectHintCardThree = () => {
    if (isHintThreeUnlocked.value) {
        selectedHint.value = hints[2];
    }
};
</script>
