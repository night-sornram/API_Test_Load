import http from "k6/http"
import {sleep} from "k6"

export const options = {
    stages : [
        {duration: '5s', target: "7000"},
        {duration: '3s', target: "0"},
    ],
}

export default function () {
    const res = http.get("http://127.0.0.1:3000/phone/0754952794", null, {
        headers: {"Content-Type" : "application/json"}
    })
    sleep(1);
}