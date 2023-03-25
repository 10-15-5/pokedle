import { strict } from 'assert';
import mongoose from 'mongoose'


//TODO: REMOVE STRING
const connectToDB = async () => 
     mongoose.connect('');


const migrateUsers = async () => {
    const UserModel = mongoose.model('users', new mongoose.Schema({test: String},{ strict: false}));

    //Get all users TODO: replace this with all:
    const users = await UserModel.find({}).lean().exec()

    //map all users to have all new fields.
    const updatedUsers = users.map((user) => ({
        _id: user._id,
        createdAt: user.createdAt,
        updatedAt: user.updatedAt,
        classicCurrentStreak: user.currentStreak ? user.currentStreak : 0,
        classicMaxStreak: user.maxStreak ? user.maxStreak : 0,
        classicFirstTryWins: user.firstTryWins ? user.firstTryWins : 0,
        classicGamesWon: user.gamesWon?.length ? user.gamesWon : [],
        flavortextCurrentStreak: user.flavortextCurrentStreak ? user.flavortextCurrentStreak : 0,
        flavortextMaxStreak: user.flavortextMaxStreak ? user.flavortextMaxStreak : 0,
        flavortextFirstTryWins: user.flavortextFirstTryWins ? user.flavortextFirstTryWins : 0,
        flavortextGamesWon: user.flavortextGamesWon?.length ? user.flavortextGamesWon : [],
        silhouetteCurrentStreak: user.silhouetteCurrentStreak ? user.silhouetteCurrentStreak : 0,
        silhouetteMaxStreak: user.silhouetteMaxStreak ? user.silhouetteMaxStreak : 0,
        silhouetteFirstTryWins: user.silhouetteFirstTryWins ? user.silhouetteFirstTryWins : 0,
        silhouetteGamesWon: user.silhouetteGamesWon?.length ? user.silhouetteGamesWon : [],
    }))

    //Foreach -> findoneandreplace by _id
    const promiseArr = updatedUsers.map(user => UserModel.findOneAndReplace(user._id, user));

    await Promise.all(promiseArr);

    console.log("USERS SUCCESSFULLY MIGRATED");
}

const migrateDailyStats = async () => {
    const DailyStatsModel = mongoose.model('dailyStat', new mongoose.Schema({test: String},{ strict: false}),'dailyStats');

    //Get all dailyStats TODO: replace this with all:
    const dailyStats = await DailyStatsModel.find({}).lean().exec();

    //map all dailyStats to have all new fields.
    const updatedDailyStats = dailyStats.map((stats) => ({
        _id: stats._id,
        date: stats.date,

        classicFirstTryWins: stats.firstTryWins ? stats.firstTryWins : 0,
        classicGamesWon: stats.gamesWon ? stats.gamesWon : 0,
        classicNumberOfGuesses: stats.numberOfGuesses ? stats.numberOfGuesses : 0,

        flavortextFirstTryWins: stats.flavortextFirstTryWins ? stats.flavortextFirstTryWins : 0,
        flavortextGamesWon: stats.flavortextGamesWon ? stats.flavortextGamesWon : 0,
        flavortextNumberOfGuesses: stats.flavortextNumberOfGuesses ? stats.flavortextNumberOfGuesses : 0,

        silhouetteFirstTryWins: stats.silhouetteFirstTryWins ? stats.silhouetteFirstTryWins : 0,
        silhouetteGamesWon: stats.silhouetteGamesWon ? stats.silhouetteGamesWon : 0,
        silhouetteNumberOfGuesses: stats.silhouetteNumberOfGuesses ? stats.silhouetteNumberOfGuesses : 0,
    }))

    //Foreach -> findoneandreplace by _id
    const promiseArr = updatedDailyStats.map(stats => DailyStatsModel.findOneAndReplace(stats._id, stats));

    await Promise.all(promiseArr);

    console.log("DAILYSTATS SUCCESSFULLY MIGRATED");
}

const runMigration = async () => {
    await connectToDB();

    await migrateUsers();

    await migrateDailyStats();

    mongoose.connection.close()
    console.log("Migration finnished");
}

runMigration();