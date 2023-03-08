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
                        text: colors.black,
                        orange: colors.orange[600],
                        green: colors.green[600],
                    },
                    dark: {
                        bg: '#4a4a4a',
                        border: colors.gray[500],
                        text: colors.gray[100],
                        orange: colors.orange[300],
                        green: colors.green[300],
                    }
            },
            fontFamily: {
                pkmEmerald: ['pkmEmerald']
            },
            screens: {
                'sm': { 'max': '576px'},
            },
        },
    },
    plugins: [],
}


