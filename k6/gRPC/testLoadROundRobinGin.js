import http from "k6/http"
import {check, sleep} from "k6"

export const options = {
    stages: [
        {duration: '10s', target: 2000},
        {duration: '10s', target: 2000},
    ],
}


export default function () {
    const res = http.get("http://127.0.0.1:4000/grpc/roundrobin/phone/0754952794")
    if (res.status != 200 && res.status != 503 && res.status != 429 && res.status != 0) console.log(res.status)
    check(res, {'status was 200': (r) => r.status == 200});
    check(res, {'status was 429': (r) => r.status == 429});
    check(res, {'status was 503': (r) => r.status == 503});
    check(res, {'status was 0': (r) => r.status == 0});
    sleep(1);
}
