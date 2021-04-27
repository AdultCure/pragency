const express = require('express');

const PORT = process.env.PORT || 4020;

const app = express();

app.get('/', (req, res) => {
    res.send('hello postgres + nodejs')
})

app.listen(PORT, () => console.log(`server started on port ${PORT}`));
