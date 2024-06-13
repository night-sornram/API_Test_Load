import http from "k6/http"
import {sleep} from "k6"

export const options = {
    vus: 500,
    duration: '15s',
}


export default function () {
    const res = http.get("http://127.0.0.1:3000/http1-1/fasthttpgoroutine/phone/0754952794", null, {
        headers: {"Content-Type" : "application/json"}
    })
    sleep(1);
}