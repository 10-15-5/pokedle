import { createRequire } from "module"; // Bring in the ability to create the 'require' method
const require = createRequire(import.meta.url); // construct the require method
const pokemonData = require('../server/data/pokemonData-v4.json');
import fs from 'fs';
import axios from "axios";

const runScript = async () => {

    const promiseArr = pokemonData.map((pokemon) => axios.get(`https://pokeapi.co/api/v2/pokemon-species/${pokemon.number}/`))

    const res = await Promise.all(promiseArr);

    const flavorTextArr = res.map((resp) => {

        const flavor_text_entry = resp.data.flavor_text_entries.find(flaverTextEntry => flaverTextEntry.version.name === 'emerald');


        return {
            flavorText: flavor_text_entry.flavor_text
        }
    });

    //Remove uppercase pokemon name, will leave '___é___' in some places
    const withUnderscoreAndÈ = flavorTextArr.map(flavorTextEntry => (
        flavorTextEntry.flavorText.replace(/[A-Z]{2,}/g, '___')
    ))

    //Replace occurences of '___é___' with '___'
    const withoutÈ = withUnderscoreAndÈ.map(withUnderscoreAndÈ => (
        withUnderscoreAndÈ.replace('___é___', '___')
    ))

    //Many places in text are '\n', replace wiht space here
    const withoutBackslashN = withoutÈ.map(withoutÈ => (withoutÈ.replace(/\n/g, ' ')))

    const updatedPokemonData = pokemonData.map((pokemon, idx) => ({
        ...pokemon,
        flavorText: withoutBackslashN[idx],
    }));

    const jsonData = JSON.stringify(updatedPokemonData);

    fs.writeFile("./scripts/result.json", jsonData, (err) => {
        if (err) {
            console.log("Error");
            console.log(err);
        }
    })
}

await runScript();