const express = require('express');
const app = express();
const port = 3000;

app.use(express.json()); 
app.use(express.urlencoded({
    extended: true
}));

app.get('/', (req, res) => {
    res.status(200).send("home page");
});

app.post('/post', (req, res) => {
    const myjson = req.body;
    res.status(200).send(myjson);
});

app.post('/postform', (req, res) => {
    res.status(200).send(JSON.stringify(req.body));
});

app.listen(port, () => {
    console.log(`Server is running on port ${port}`);
});
