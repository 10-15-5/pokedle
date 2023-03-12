import { createRequire } from "module"; // Bring in the ability to create the 'require' method
const require = createRequire(import.meta.url); // construct the require method
const pokemonData = require('../../server/data/pokemonData-v4.json');
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

    const flavorTextArrWONames = [];

    for (let i = 0; i < flavorTextArr.length; i++) {
       const name = pokemonData[i].name;
       const uppercaseName = name.toUpperCase();

      const currFlavorText = flavorTextArr[i].flavorText;

      const currWOName = currFlavorText.replace(uppercaseName, '___');
      flavorTextArrWONames.push(currWOName);
    }

    //Many places in text are '\n', replace wiht space here
    const withoutBackslashN = flavorTextArrWONames.map(text => (text.replace(/\n/g, ' ')))

    const updatedPokemonData = pokemonData.map((pokemon, idx) => ({
        ...pokemon,
        flavorText: withoutBackslashN[idx],
    }));

    const jsonData = JSON.stringify(updatedPokemonData);

    fs.writeFile("./scripts/pokemonDataModifications/result.json", jsonData, (err) => {
        if (err) {
            console.log("Error");
            console.log(err);
        }
    })
}

await runScript();