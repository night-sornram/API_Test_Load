import http from "k6/http"
import {check,sleep} from "k6"

export const options = {
    vus: 1000,
    duration: '10s',
}


export default function () {
    const res = http.get("http://127.0.0.1:3000/grpc/default/phone/0754952794", null, {
        headers: {"Content-Type" : "application/json"}
    })

    check(res, {
        'status is 200': (r) => r.status === 200,
    });


    sleep(1);
}