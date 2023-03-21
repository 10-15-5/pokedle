import { defineStore } from 'pinia'
import { computed, reactive, ref } from 'vue'

export const useStore = defineStore('store', () => {
    const user = reactive({});
    const isClassicGameWon = ref(false);
    const isShiny = ref(false);
    const theme = ref('')
    const isHintMode = ref(false);

    const setShiny = (isShinyParam) => {
        isShiny.value = isShinyParam;
        localStorage.setItem('isShiny', `${isShiny.value}`)
    }

    const toggleShiny = () => {
        setShiny(!isShiny.value);
    }

    const isDark = computed(() => theme.value === 'dark');

    const setTheme = (t) => {
        theme.value = t;
        localStorage.theme = t
        addThemeTagToHTML();
    }

    const addThemeTagToHTML = () => {
        if (localStorage.theme == 'dark') {
            document.documentElement.classList.add('dark')
          } else {
            document.documentElement.classList.remove('dark')
          }
    }

    const loadTheme = () => {
        theme.value = localStorage.theme;
        addThemeTagToHTML();
    }
    
    const toggleTheme = () => {
        setTheme(theme.value === 'dark' ? '' : 'dark');
    }
    
    const setHintMode = (isHintModeParam) => {
        isHintMode.value = isHintModeParam;
        localStorage.isHintMode = `${isHintMode.value}`
    }


    const toggleHintMode = () => {
        setHintMode(!isHintMode.value);
    }

    const setUser = (newUser) => {
        Object.assign(user, newUser);
    }

    const setIsClassicGameWon = (isWon) => {
        isClassicGameWon.value = isWon;
        localStorage.setItem('isClassicGameWon', `${isWon}`);
    }

    return {
        user,
        isClassicGameWon,
        isShiny,
        isHintMode,
        theme,
        isDark,
        setUser,
        setIsClassicGameWon,
        toggleShiny,
        setShiny,
        setHintMode,
        toggleHintMode,
        setTheme,
        toggleTheme,
        loadTheme
    }
});