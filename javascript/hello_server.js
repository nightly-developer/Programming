const express = require('express');
const app = express();
const cors = require('cors');
const port = 3000;


app.use(cors());

app.get('/greet',(req,res)=>{
    res.status(200).json({title:"Hello World!"});
})

app.listen(port, () => {
    console.log(`Server is running on port ${port}`);
});