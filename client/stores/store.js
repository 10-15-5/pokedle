import { defineStore } from 'pinia'
import { computed, reactive, ref } from 'vue'

export const useStore = defineStore('store', () => {
    const user = reactive({});
    const isGameWon = ref(false);
    const isShiny = ref(false);
    const theme = ref('')
    const isDifficultyInsane = ref(true);

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
    
    const setDifficultyInsane = (isDifficultyInsaneParam) => {
        isDifficultyInsane.value = isDifficultyInsaneParam;
        localStorage.isDifficultyInsane = `${isDifficultyInsane.value}`
    }


    const toggleDifficultyInsane = () => {
        setDifficultyInsane(!isDifficultyInsane.value);
    }

    const setUser = (newUser) => {
        Object.assign(user, newUser);
    }

    const setIsGameWon = (isWon) => {
        isGameWon.value = isWon;
        localStorage.setItem('isGameWon', `${isWon}`);
    }

    return {
        user,
        isGameWon,
        isShiny,
        isDifficultyInsane,
        theme,
        isDark,
        setUser,
        setIsGameWon,
        toggleShiny,
        setShiny,
        setDifficultyInsane,
        toggleDifficultyInsane,
        setTheme,
        toggleTheme,
        loadTheme
    }
});