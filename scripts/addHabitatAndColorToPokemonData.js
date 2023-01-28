import { createRequire } from "module"; // Bring in the ability to create the 'require' method
const require = createRequire(import.meta.url); // construct the require method
const pokemonSpeciesData = require('./pokemonSpecies.json');
const pokemonData = require('../server/database/pokemonData-v3.json');
import fs from 'fs';
import axios from "axios";

const runScript = async () => {

    const promiseArr = pokemonSpeciesData.map((pokemon) => axios.get(pokemon.url))

    const res = await Promise.all(promiseArr);

    const habitatAndColour = res.map((resp, number) => ({
        //number: number + 1,
        habitat: resp.data.habitat.name,
        color: resp.data.color.name,
    }));

    const updatedPokemonData = pokemonData.map((pokemon) => ({
        ...pokemon,
        habitat: habitatAndColour[pokemon.number-1].habitat,
        color: habitatAndColour[pokemon.number-1].color,
    })) 

    const jsonData = JSON.stringify(updatedPokemonData);

    fs.writeFile("./scripts/result.json", jsonData, (err) => {
        if (err) {
            console.log("Error");
            console.log(err);
        }
    })
}

await runScript();