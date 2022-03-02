const express = require('express');
const app = express(); 
const PORT = 80;

app.get('/', function(req, res){
    res.send("Hello World");
})

app.listen(PORT, () => {
    console.log(`Server running in port ${PORT}`)
})