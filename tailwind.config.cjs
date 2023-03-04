/** @type {import('tailwindcss').Config} */
const colors = require('tailwindcss/colors');

module.exports = {
    darkMode: 'class',
    content: [
        "./index.html",
        "./client/**/*.{vue,js,ts,jsx,tsx}",
    ],
    theme: {
        extend: {
            colors: {
                    light: {
                        bg: colors.white,
                        border: colors.black,
                        text: colors.black
                    },
                    dark: {
                        bg: '#4a4a4a',
                        border: colors.gray[500],
                        text: colors.gray[300],
                    }
            },
            fontFamily: {
                pkmEmerald: ['pkmEmerald']
            },
        },
    },
    plugins: [],
}


