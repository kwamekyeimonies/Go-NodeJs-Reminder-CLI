require('dotenv').config()
const express = require('express')
const port = process.env.PORT
const app = express()
const notifier = require('node-notifier')
const path = require('path')

app.use(express.json())
app.get(
    "/health",
    (req,res)=>{
        res.status(200).send("Notifier api working.....")
    }
)
app.post(
    "/notify",
    (req,res)=>{
        notify(req.body,reply=> res.send(reply))
    }
)
app.listen(
    port,
    ()=>{
        console.log(`Server running on http://localhost:${port}`)
    }
)

const notify=({title,message},cb)=>{
    notifier.notify(
        {
            title:title || "Unknown title",
            message:message || "Unknown message",
            icon:path.join(__dirname,"dexoangle.jpg"),
            sound:true,
            wait:true,
            reply:true,
            closeLable:"Completed?",
            timeout:15
        },
        (err,response,reply)=>{
            cb(reply)
        }
    )
}
