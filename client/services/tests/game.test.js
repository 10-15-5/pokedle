import * as game from '../game.js';
import { GameModes, DateOfFirstPokeldeGameClassic, DateOfFirstPokeldeGameFlavortext, DateOfFirstPokeldeGameSilhouette } from '../../constants';
import moment from 'moment-timezone';
moment.tz.setDefault('UTC');

describe('game', () => {
    describe('getCurrentPokemonNumber', () => {
        test('returns the correct number of days', () => {

            const now = moment.tz('2023-03-16T00:00:01', 'UTC');

            const result = game.getCurrentPokemonNumber(GameModes.Classic, now);
            expect(result).toEqual(0);
        
        });
        test('returns the correct number of days', () => {

            const now = moment.tz('2023-03-16T23:59:00', 'UTC');

            const result = game.getCurrentPokemonNumber(GameModes.Classic, now);
            expect(result).toEqual(0);
        
        });
        test('returns the correct number of days', () => {

            const now = moment.tz('2023-03-17T00:00:01', 'UTC');

            const result = game.getCurrentPokemonNumber(GameModes.Classic, now);
            expect(result).toEqual(1);
        
        });
        test('returns the correct number of days', () => {

            const now = moment.tz('2023-03-17T23:59:00', 'UTC');
            const result = game.getCurrentPokemonNumber(GameModes.Classic, now);
            expect(result).toEqual(1);
        
        });
        test('returns the correct number of days', () => {

            const now = moment.tz('2023-04-02T00:01:00', 'UTC');

            console.log(now.toString());

            const result = game.getCurrentPokemonNumber(GameModes.Classic, now);
            expect(result).toEqual(17);
        
        });

        test('throws an error if gamemode is not provided', () => {
            expect(() => game.getCurrentPokemonNumber()).toThrow('Gamemode Required');
        });
    });
});
