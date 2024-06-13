import http from "k6/http"
import {sleep} from "k6"

export const options = {
    vus: 2000,
    duration: '10s',
}


export default function () {
    const res = http.get("http://127.0.0.1:3000/grpc/roundrobin/phone/0754952794", null, {
        headers: {"Content-Type" : "application/json"}
    })
    sleep(1);
}