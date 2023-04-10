import { createRequire } from 'module'; // Bring in the ability to create the 'require' method
const require = createRequire(import.meta.url); // construct the require method
const pokemonData = require('../../server/data/pokemonData-v4.json');
import fs from 'fs';
import axios from 'axios';

const LANGUAGES = ['de', 'fr', 'es', 'it', 'ja', 'ko', 'en'];
const languageMap = new Map([
    ['de', 'german'],
    ['fr', 'french'],
    ['es', 'spanish'],
    ['it', 'italian'],
    ['ja', 'japanese'],
    ['ko', 'korean'],
    ['en', 'english'],
]);

const saveAsResultFile = (updatedPokemonData) => {
    const jsonData = JSON.stringify(updatedPokemonData);

    fs.writeFile('./scripts/pokemonDataModifications/result.json', jsonData, (err) => {
        if (err) {
            console.log('Error');
            console.log(err);
        }
    });
};

const addNamesAndFlavortexts = async () => {
    const promiseArr = pokemonData.map((pokemon) =>
        axios.get(`https://pokeapi.co/api/v2/pokemon-species/${pokemon.number}/`)
    );

    //For single call:
    //const promiseArr = [axios.get(`https://pokeapi.co/api/v2/pokemon-species/${3}/`)];

    const res = await Promise.all(promiseArr);

    const arr = res.map((resp) => {
        const namesAndLanguages = resp.data.names.filter((name) => LANGUAGES.includes(name.language.name));

        const flavor_text_entries = resp.data.flavor_text_entries.filter(
            (flaverTextEntry) =>
                flaverTextEntry.version.name === 'omega-ruby' && LANGUAGES.includes(flaverTextEntry.language.name)
        );

        const flavortexts = flavor_text_entries.map((entry) => ({
            flavortext: entry.flavor_text,
            language: languageMap.get(entry.language.name),
        }));

        const names = namesAndLanguages.map((namesAndLanguage) => ({
            name: namesAndLanguage.name,
            language: languageMap.get(namesAndLanguage.language.name),
        }));

        return {
            flavortexts,
            names,
            name: resp.data.name,
        };
    });

    for (let i = 0; i < arr.length; i++) {
        const curr = arr[i];

        const languageAndNameMap = new Map();

        curr.names.forEach((element) => {
            languageAndNameMap.set(element.language, element.name);
        });

        console.log(languageAndNameMap);
        console.log(languageAndNameMap.get('german'));

        //Remove pokemon name from flavortext
        curr.flavortexts = curr.flavortexts.map((entry) => ({
            ...entry,
            flavortext: entry.flavortext.replace(languageAndNameMap.get(entry.language), '___'),
        }));

        //Remove "\n" from flavortext
        curr.flavortexts = curr.flavortexts.map((entry) => ({
            ...entry,
            flavortext: entry.flavortext.replace(/\n/g, ' '),
        }));
    }

    console.log(arr[0]);

    const updatedPokemonData = pokemonData.map((pokemon, idx) => {
        return {
            ...pokemon,
            flavortexts: arr[idx].flavortexts,
            names: arr[idx].names,
        };
    });

    updatedPokemonData.forEach((pokeData) => {
        delete pokeData.flavorText, delete pokeData.legendary;
    });

    console.log(updatedPokemonData[0]);

    saveAsResultFile(updatedPokemonData);
};

const addTypes = async () => {
    const promiseArr = [];

    for (let i = 1; i < 19; i++) {
        promiseArr.push(axios.get(`https://pokeapi.co/api/v2/type/${i}`));
    }

    const res = await Promise.all(promiseArr);

    const arr = res.map((resp) => {
        const namesAndLanguages = resp.data.names.filter((name) => LANGUAGES.includes(name.language.name));

        const names = namesAndLanguages.map((namesAndLanguage) => ({
            name: namesAndLanguage.name,
            language: languageMap.get(namesAndLanguage.language.name),
        }));

        return {
            name: resp.data.name,
            names,
        };
    });

    const jsonData = JSON.stringify(arr);

    fs.writeFile('./scripts/pokemonDataModifications/types.json', jsonData, (err) => {
        if (err) {
            console.log('Error');
            console.log(err);
        }
    });
};

//await addNamesAndFlavortexts();
await addTypes();
