import { defineStore } from 'pinia';
import { computed, reactive, ref } from 'vue';

export const useStore = defineStore('store', () => {
    const user = reactive({});
    const isShiny = ref(false);
    const theme = ref('');
    const isHintMode = ref(false);
    const isSound = ref(true)

    const isClassicGameWon = ref(false);
    const isFlavortextGameWon = ref(false);
    const isSilhouetteGameWon = ref(false);

    const currentStreak = ref(0);

    const setShiny = (isShinyParam) => {
        isShiny.value = isShinyParam;
        localStorage.isShiny = `${isShiny.value}`;
    };

    const toggleShiny = () => {
        setShiny(!isShiny.value);
    };

    const isDark = computed(() => theme.value === 'dark');

    const setTheme = (t) => {
        theme.value = t;
        localStorage.theme = t;
        addThemeTagToHTML();
    };

    const addThemeTagToHTML = () => {
        if (localStorage.theme == 'dark') {
            document.documentElement.classList.add('dark');
        } else {
            document.documentElement.classList.remove('dark');
        }
    };

    const loadTheme = () => {
        theme.value = localStorage.theme;
        addThemeTagToHTML();
    };

    const toggleTheme = () => {
        setTheme(theme.value === 'dark' ? '' : 'dark');
    };

    const setHintMode = (isHintModeParam) => {
        isHintMode.value = isHintModeParam;
        localStorage.isHintMode = `${isHintMode.value}`;
    };

    const toggleHintMode = () => {
        setHintMode(!isHintMode.value);
    };

    const setSound = (sound) => {
        isSound.value = sound;
        localStorage.isSound = `${isSound.value}`;
    };

    const toggleSound = () => {
        setSound(!isSound.value);
    };

    const setUser = (newUser) => {
        Object.assign(user, newUser);
    };

    const setIsClassicGameWon = (isWon) => {
        isClassicGameWon.value = isWon;
        localStorage.isClassicGameWon = `${isWon}`;
    };

    const setIsFlavortextGameWon = (isWon) => {
        isFlavortextGameWon.value = isWon;
        localStorage.isFlavortextGameWon = `${isWon}`;
    };

    const setIsSilhouetteGameWon = (isWon) => {
        isSilhouetteGameWon.value = isWon;
        localStorage.isSilhouetteGameWon = `${isWon}`;
    };

    const setCurrentStreak = (val) => {
       currentStreak.value = val;
    }

    return {
        user,
        isClassicGameWon,
        isShiny,
        isHintMode,
        theme,
        isDark,
        isSound,
        isFlavortextGameWon,
        isSilhouetteGameWon,
        currentStreak,
        setUser,
        setIsClassicGameWon,
        toggleShiny,
        setShiny,
        setHintMode,
        toggleHintMode,
        setTheme,
        toggleTheme,
        loadTheme,
        setIsFlavortextGameWon,
        setIsSilhouetteGameWon,
        setCurrentStreak,
        setSound,
        toggleSound,
    };
});
