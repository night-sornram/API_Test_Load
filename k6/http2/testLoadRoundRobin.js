import http from "k6/http"
import {check, sleep} from "k6"

export const options = {
    vus: 1500,
    duration: '15s',
}


export default function () {
    const res = http.get("http://127.0.0.1:3000/http2/roundrobin/phone/0754952794")
    if(res.status != 200 && res.status != 408 && res.status != 0) console.log(res.status)
    check(res, { 'status was 200': (r) => r.status == 200 });
    check(res, { 'status was 408': (r) => r.status == 408 });
    check(res, { 'status was 0': (r) => r.status == 0 });
    sleep(1);
}