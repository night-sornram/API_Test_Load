import http from "k6/http"
import {sleep} from "k6"

export const options = {
    vus: 300,
    duration: '30s',
}


export default function () {
    const res = http.get("http://127.0.0.1:3000/default/phone/0754952794", null, {
        headers: {"Content-Type" : "application/json"}
    })
    sleep(1);
}