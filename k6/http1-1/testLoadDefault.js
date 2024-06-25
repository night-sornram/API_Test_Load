import http from "k6/http"
import {check, sleep} from "k6"

export const options = {
    stages: [
        { duration: '10s', target: 1000 },
        { duration: '10s', target: 1500 },
        { duration: '10s', target: 2200 },
        { duration: '10s', target: 2200 },
    ],
}


export default function () {
    const res = http.get("http://127.0.0.1:3000/http1-1/default/phone/0754952794")
    if(res.status != 200 && res.status != 408 && res.status != 0) console.log(res.status)
    check(res, { 'status was 200': (r) => r.status == 200 });
    check(res, { 'status was 408': (r) => r.status == 408 });
    check(res, { 'status was 0': (r) => r.status == 0 });
    sleep(1);
}


