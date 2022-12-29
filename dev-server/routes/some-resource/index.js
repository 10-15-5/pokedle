import express from 'express';

const someResourceRouter = express.Router();

someResourceRouter.get('/some-resources', (req, res) => {
        res.status(200);
        res.send("Getting all of 'some resource'");
    }
);

export {
    someResourceRouter,
}