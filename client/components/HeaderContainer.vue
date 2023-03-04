<template>
    <div class="flex items-center justify-center">
        <HeaderIcon>
            <p class="text-[30px]">🔥</p>
            <p class="streak-count mt-2 font-pkmEmerald text-light-text dark:text-gray-600">
                {{ store.user.currentStreak }}
            </p>
        </HeaderIcon>
        <HeaderIcon @click="isDialogSettings = true">
            <span class="pi pi-cog text-[20px] text-light-text dark:text-dark-text"></span>
        </HeaderIcon>
        <HeaderIcon @click="isDialogStatistics = true">
            <span class="pi pi-chart-bar text-[20px] text-light-text dark:text-dark-text"></span>
        </HeaderIcon>
        <HeaderIcon @click="isDialogInfo = true">
            <span class="pi pi-info-circle text-[20px] text-light-text dark:text-dark-text"></span>
        </HeaderIcon>
        <SettingsDialog
            :isVisible="isDialogSettings"
            :onClose="
                () => {
                    isDialogSettings = false;
                }
            "
        />
        <StatisticsDialog
            :isVisible="isDialogStatistics"
            :onClose="
                () => {
                    isDialogStatistics = false;
                }
            "
        />
        <InfoDialog
            :isVisible="isDialogInfo"
            :onClose="
                () => {
                    isDialogInfo = false;
                }
            "
        />
    </div>
</template>

<script setup>
import { useStore } from '../stores/store';
import { ref } from 'vue';
import HeaderIcon from './HeaderIcon.vue';
import InfoDialog from './dialogs/InfoDialog.vue';
import SettingsDialog from './dialogs/SettingsDialog.vue';
import StatisticsDialog from './dialogs/StatisticsDialog.vue';

const store = useStore();
let isDialogStatistics = ref(false);
let isDialogSettings = ref(false);
let isDialogInfo = ref(false);
</script>

<style scoped>
p {
    padding: 4px;
}

.streak-count {
    position: absolute;
    font-weight: bold;
    margin-left: 1px;
    text-shadow: -0.5px 0.5px 1px #fffcfc, 0.5px 0.5px 1px #ffffff, 0.5px -0.5px 0 #ffffff,
        -0.5px -0.5px 0 #ffffff;
}
</style>
